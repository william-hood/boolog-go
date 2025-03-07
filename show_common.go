// Copyright (c) 2025 William Arthur Hood
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package boolog

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func shouldRecurse(candidate any) bool {
	if candidate == nil {
		return false
	}

	if reflect.TypeOf(candidate).String() == "string" {
		return false
	}

	// TODO: Probably need to exhaust all "primitive" types with the check above.
	return true
}

// shouldRender() from the Kotlin version probably doesn't apply here.

func (this Boolog) beginShow(timestamp time.Time, typeName string, variableName string, style string, recurseLevel int) *strings.Builder {
	result := new(strings.Builder)
	result.WriteString(fmt.Sprintf("\r\n<div class=\"object %s centered\">\r\n", style))

	// TODO

	return result
}
