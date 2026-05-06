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

package buffer

import (
	"io"
	"time"
)

// Interface describes the append-oriented byte buffer operations required by
// ARCORIS logging encoders and logging cores.
//
// Implementations SHOULD be backed by mutable byte storage and SHOULD preserve
// Buffer's single-owner lifecycle semantics: bytes returned by Bytes are
// borrowed, writes append to the current logical contents, Reset clears the
// logical contents without promising to release capacity, and Free releases the
// value to its owner if it has one. Unless an implementation explicitly
// documents stronger guarantees, Interface values are not safe for concurrent
// mutation.
type Interface interface {
	io.Writer

	// WriteByte appends a single byte and reports success or failure.
	WriteByte(v byte) error

	// WriteString appends s and reports the number of bytes accepted.
	WriteString(s string) (int, error)

	// AppendTime appends t formatted with layout.
	AppendTime(t time.Time, layout string)

	// AppendDuration appends d as a base-10 number of nanoseconds.
	AppendDuration(d time.Duration)

	// AppendByte appends a single byte.
	AppendByte(v byte)

	// AppendBytes appends v exactly as provided.
	AppendBytes(v []byte)

	// AppendString appends s as bytes without escaping.
	AppendString(s string)

	// AppendBool appends v as "true" or "false".
	AppendBool(v bool)

	// AppendComplex128 appends v in the form "<real><+|-><imag>i".
	AppendComplex128(v complex128)

	// AppendComplex64 appends v in the form "<real><+|-><imag>i".
	AppendComplex64(v complex64)

	// AppendFloat64 appends v using compact 64-bit floating-point formatting.
	AppendFloat64(v float64)

	// AppendFloat32 appends v using compact 32-bit floating-point formatting.
	AppendFloat32(v float32)

	// AppendInt appends v as a base-10 signed integer.
	AppendInt(v int)

	// AppendInt64 appends v as a base-10 signed integer.
	AppendInt64(v int64)

	// AppendInt32 appends v as a base-10 signed integer.
	AppendInt32(v int32)

	// AppendInt16 appends v as a base-10 signed integer.
	AppendInt16(v int16)

	// AppendInt8 appends v as a base-10 signed integer.
	AppendInt8(v int8)

	// AppendUint appends v as a base-10 unsigned integer.
	AppendUint(v uint)

	// AppendUint64 appends v as a base-10 unsigned integer.
	AppendUint64(v uint64)

	// AppendUint32 appends v as a base-10 unsigned integer.
	AppendUint32(v uint32)

	// AppendUint16 appends v as a base-10 unsigned integer.
	AppendUint16(v uint16)

	// AppendUint8 appends v as a base-10 unsigned integer.
	AppendUint8(v uint8)

	// AppendUintptr appends v as a lowercase hexadecimal pointer-sized value.
	AppendUintptr(v uintptr)

	// Bytes returns a borrowed view of the current contents.
	//
	// Callers MUST treat the returned slice as read-only and MUST NOT retain it
	// after the next mutation, Reset, Free, or pool return.
	Bytes() []byte

	// Len reports the current number of bytes in the buffer.
	Len() int

	// Cap reports the current capacity of the underlying storage.
	Cap() int

	// Reset clears the logical contents while retaining reusable capacity.
	Reset()

	// Free releases the buffer to its originating pool, if any.
	Free()
}
