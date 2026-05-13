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

// Bytes constructs a bytes field descriptor.
//
// The slice is borrowed and the constructor does not copy it. Callers must not
// mutate value until the log call or context binding has consumed it. Runtime
// encoders that retain fields must copy, encode, or bind bytes as needed.
// Bytes stores value in Field.Bytes and leaves Field.Interface nil.
func Bytes(key string, value []byte) Field {
	return Field{Key: key, Type: BytesType, Bytes: value}
}
