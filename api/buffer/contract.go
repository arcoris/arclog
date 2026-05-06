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

// The contracts in this file are internal compile-time shape checks for
// *Buffer. External packages should depend on the concrete *buffer.Buffer type,
// not on broad interfaces that would freeze more API surface than needed.
//
// The contract is split by responsibility so the required method groups remain
// readable: writes, primitive appends, borrowed views, and lifecycle. Keeping
// the shape unexported also makes accidental broad public interface exposure
// less likely.
type bufferContract interface {
	writerContract
	appendContract
	viewContract
	lifecycleContract
}

var _ bufferContract = (*Buffer)(nil)

// writerContract verifies the io-style write operations on *Buffer.
type writerContract interface {
	io.Writer

	WriteByte(byte) error
	WriteString(string) (int, error)
}

// appendContract verifies the primitive append operations on *Buffer.
type appendContract interface {
	AppendTime(time.Time, string)
	AppendDuration(time.Duration)

	AppendByte(byte)
	AppendBytes([]byte)
	AppendString(string)
	AppendBool(bool)

	AppendComplex128(complex128)
	AppendComplex64(complex64)

	AppendFloat64(float64)
	AppendFloat32(float32)

	AppendInt(int)
	AppendInt64(int64)
	AppendInt32(int32)
	AppendInt16(int16)
	AppendInt8(int8)

	AppendUint(uint)
	AppendUint64(uint64)
	AppendUint32(uint32)
	AppendUint16(uint16)
	AppendUint8(uint8)
	AppendUintptr(uintptr)
}

// viewContract verifies borrowed read access to the current buffer contents.
type viewContract interface {
	Bytes() []byte
	String() string
	Len() int
	Cap() int
}

// lifecycleContract verifies explicit buffer reuse and release operations.
type lifecycleContract interface {
	Reset()
	Free()
}
