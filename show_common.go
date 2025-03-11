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
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

var nonRecurseTypes = []string{"string", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "bool", "byte", "rune", "float32", "float64", "complex64", "complex128"}

func shouldRecurse(candidate any) bool {
	if candidate == nil {
		return false
	}

	if slices.Contains(nonRecurseTypes, reflect.TypeOf(candidate).String()) {
		return false
	}

	return true
}

// shouldRender() from the Kotlin version probably doesn't apply here as unexported fields are not visible to reflection in Go.

func (this *Boolog) beginShow(timestamp time.Time, typeName string, variableName string, style string, recurseLevel int) *strings.Builder {
	result := new(strings.Builder)
	result.WriteString(fmt.Sprintf("\r\n<div class=\"object %s\">\r\n", style))

	if recurseLevel > 0 {
		identifier := uuid.NewString()
		result.WriteString(fmt.Sprintf("<label for=\"%s\">\r\n<input id=\"%s\" class=\"gone\" type=\"checkbox\">\r\n", identifier, identifier))
	}
	result.WriteString(fmt.Sprintf("<center><h2>%s</h2>\r\n<small>", typeName))
	if variableName != nAMELESS {
		result.WriteString("<b>\"")
	}
	result.WriteString(variableName)
	if variableName != nAMELESS {
		result.WriteString("\"</b>")
	}
	result.WriteString("</small></center>\r\n")

	if recurseLevel > 0 {
		result.WriteString(fmt.Sprintf("<div class=\"%s\">\r\n", encapsulationTag()))
	}

	this.echoPlainText(fmt.Sprintf("Showing %s: %s (details in HTML log)", typeName, variableName), EMOJI_OBJECT, timestamp)

	return result
}
