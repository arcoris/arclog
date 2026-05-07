/*
   Copyright 2026 The ARCORIS Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package field

import (
	"errors"
	"fmt"
	"math"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	encoderconvert "arcoris.dev/arclog/api/encoder/convert"
)

var (
	// ErrUnsupportedType is returned when AddTo receives a Field with a Type
	// value that is not defined by this package.
	//
	// The error is intended for fields assembled manually or by incompatible
	// extension code. Fields produced by this package's constructors always use
	// a supported type tag.
	ErrUnsupportedType = errors.New("field: unsupported field type")

	// ErrNilEncoder is returned when AddTo needs to emit a field but no
	// encoder.ObjectEncoder was provided.
	//
	// Skip fields do not require an encoder and remain a no-op even when enc is
	// nil.
	ErrNilEncoder = errors.New("field: nil object encoder")
)

// AddTo appends f to enc and returns the authoritative buffer for subsequent
// writes.
//
// AddTo is API-level dispatch, not concrete encoding. It maps Field.Type to the
// corresponding encoder.ObjectEncoder method and preserves the returned-buffer
// contract of the encoder package. Concrete encoders own formatting details.
//
// Skip fields are ignored. Inline fields call the stored object marshaler in the
// current encoder namespace. Namespace fields delegate to ObjectEncoder.
// OpenNamespace. Object, array, and reflection fields preserve the error
// semantics of the encoder method they call. Error and fmt.Stringer fields use
// strict encoder-bound conversion: nil-like values have already been normalized
// by constructors, and panics from Error or String propagate unchanged.
//
// AddTo expects fields to be built by this package's constructors. Manually
// assembled fields with mismatched Type and payload storage may panic through a
// failed type assertion; unknown Type values return ErrUnsupportedType.
func (f Field) AddTo(dst *buffer.Buffer, enc encoder.ObjectEncoder) (*buffer.Buffer, error) {
	switch f.Type {
	case SkipType:
		return dst, nil
	}

	if enc == nil {
		return dst, ErrNilEncoder
	}

	switch f.Type {
	case BoolType:
		return enc.AddBool(dst, f.Key, f.Integer == 1), nil
	case ByteStringType:
		return enc.AddByteString(dst, f.Key, f.Interface.([]byte)), nil
	case Complex128Type:
		return enc.AddComplex128(dst, f.Key, f.Interface.(complex128)), nil
	case Complex64Type:
		return enc.AddComplex64(dst, f.Key, f.Interface.(complex64)), nil
	case DurationType:
		return enc.AddDuration(dst, f.Key, time.Duration(f.Integer)), nil
	case Float64Type:
		return enc.AddFloat64(dst, f.Key, math.Float64frombits(uint64(f.Integer))), nil
	case Float32Type:
		return enc.AddFloat32(dst, f.Key, math.Float32frombits(uint32(f.Integer))), nil
	case Int64Type:
		return enc.AddInt64(dst, f.Key, f.Integer), nil
	case Int32Type:
		return enc.AddInt32(dst, f.Key, int32(f.Integer)), nil
	case Int16Type:
		return enc.AddInt16(dst, f.Key, int16(f.Integer)), nil
	case Int8Type:
		return enc.AddInt8(dst, f.Key, int8(f.Integer)), nil
	case IntType:
		return enc.AddInt(dst, f.Key, int(f.Integer)), nil
	case StringType:
		return enc.AddString(dst, f.Key, f.String), nil
	case TimeType:
		return enc.AddTime(dst, f.Key, f.timeValue()), nil
	case TimeFullType:
		return enc.AddTime(dst, f.Key, f.Interface.(time.Time)), nil
	case Uint64Type:
		return enc.AddUint64(dst, f.Key, uint64(f.Integer)), nil
	case Uint32Type:
		return enc.AddUint32(dst, f.Key, uint32(f.Integer)), nil
	case Uint16Type:
		return enc.AddUint16(dst, f.Key, uint16(f.Integer)), nil
	case Uint8Type:
		return enc.AddUint8(dst, f.Key, uint8(f.Integer)), nil
	case UintType:
		return enc.AddUint(dst, f.Key, uint(f.Integer)), nil
	case UintptrType:
		return enc.AddUintptr(dst, f.Key, uintptr(f.Integer)), nil
	case ObjectMarshalerType:
		return enc.AddObject(dst, f.Key, f.Interface.(encoder.ObjectMarshaler))
	case ArrayMarshalerType:
		return enc.AddArray(dst, f.Key, f.Interface.(encoder.ArrayMarshaler))
	case InlineMarshalerType:
		return f.Interface.(encoder.ObjectMarshaler).MarshalLogObject(dst, enc)
	case ReflectType:
		return enc.AddReflected(dst, f.Key, f.Interface)
	case NamespaceType:
		return enc.OpenNamespace(dst, f.Key), nil
	case StringerType:
		return encoderconvert.AddStringer(dst, enc, f.Key, f.Interface.(fmt.Stringer)), nil
	case ErrorType:
		return encoderconvert.AddError(dst, enc, f.Key, f.Interface.(error)), nil
	default:
		return dst, fmt.Errorf("%w: %s", ErrUnsupportedType, f.Type)
	}
}

// timeValue reconstructs the compact time representation used by Time.
//
// A nil or non-location payload is treated as UTC so malformed values cannot
// panic during ordinary time reconstruction. Constructor-produced Time fields
// store either a *time.Location here or a full time.Time in Interface.
func (f Field) timeValue() time.Time {
	loc, _ := f.Interface.(*time.Location)
	if loc == nil {
		return time.Unix(0, f.Integer)
	}
	return time.Unix(0, f.Integer).In(loc)
}
