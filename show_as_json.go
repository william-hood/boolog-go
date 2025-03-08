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
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func (this Boolog) ShowAsJson(target any, targetVariableName string) string {
	return this.showAsJsonDetailed(target, targetVariableName, EMOJI_OBJECT, "plate")
}

func (this Boolog) showAsJsonDetailed(target any, targetVariableName string, emoji string, style string) string {
	timestamp := time.Now()

	jsonTarget, _ := json.MarshalIndent(target, "", "   ")
	renderedTarget := TreatAsCode(string(jsonTarget))

	targetType := reflect.TypeOf(target)
	targetTypeName := targetType.String()

	this.Debug(fmt.Sprintf("Target Type Name: %s", targetTypeName))

	result := this.beginShow(timestamp, targetTypeName, targetVariableName, fmt.Sprintf("%s left_justified", style), 0)

	if len(renderedTarget) > MAX_BODY_LENGTH_TO_DISPLAY {
		identifier2 := uuid.NewString()
		result.WriteString(fmt.Sprintf("<label for=\"%s\">\r\n<input id=\"%s\" type=\"checkbox\">\r\n(show large object)\r\n<div class=\"%s\">\r\n", identifier2, identifier2, encapsulationTag()))
		result.WriteString(renderedTarget)
		result.WriteString("</div></label>")
	} else {
		result.WriteString(renderedTarget)
	}

	result.WriteString("\r\n</div>")

	rendition := result.String()

	this.writeToHTML(rendition, emoji, timestamp)

	return rendition
}
