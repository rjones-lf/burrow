// MIT License
//
// Copyright (c) 2018 Finterra
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlterPower(t *testing.T) {
	NewConcreteAccountFromSecret("seeeeecret")
	val := AsMutableValidator(NewValidator(PrivateKeyFromSecret("seeeeecret").PublicKey(), 100, 1))
	val.AddStake(100)
	assert.Equal(t, int64(1), val.Power())
	assert.Equal(t, uint64(200), val.Stake())
}

func TestEncoding(t *testing.T) {
	val1 := NewValidator(PrivateKeyFromSecret("seeeeecret").PublicKey(), 100, 1)
	bytes, _ := val1.Bytes()
	val2 := LoadValidator(bytes)
	assert.Equal(t, val1, val2)
}
