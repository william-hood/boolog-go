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

	"github.com/google/uuid"
)

func (this Boolog) ShowArray(target []any, targetVariableName string, recurseLevel int) string {
	if recurseLevel > MAX_SHOW_OBJECT_RECURSION {
		return fmt.Sprintf("<div class=\"outlined\">%s Too Many Levels In %s</div>", EMOJI_INCONCLUSIVE_TEST, EMOJI_INCONCLUSIVE_TEST)
	}

	targetValue := reflect.ValueOf(&target).Elem()
	targetType := targetValue.Type()
	targetTypeName := targetType.Name()

	if len(target) < 1 {
		return fmt.Sprintf("<div class=\"outlined\">(%s \"%s\" is empty)</div>", targetTypeName, targetVariableName)
	}

	timestamp := time.Now()

	result := this.beginShow(timestamp, targetTypeName, targetVariableName, "neutral", recurseLevel)
	content := new(strings.Builder)
	content.WriteString("<br><table class=\"gridlines\">\r\n")

	for index, value := range target {
		content.WriteString("<tr><td>")
		content.WriteString(fmt.Sprintf("%d", index))
		content.WriteString("</td><td>")
		content.WriteString(this.show(value, fmt.Sprintf("Array slot #%d", index), recurseLevel+1))
		content.WriteString("</td></tr>\r\n")
	}

	content.WriteString("\r\n</table><br>")

	if len(target) > MAX_OBJECT_FIELDS_TO_DISPLAY {
		identifier2 := uuid.NewString()
		result.WriteString(fmt.Sprintf("<label for=\"%s\">\r\n<input id=\"%s\" type=\"checkbox\">\r\n(show %d items)\r\n<div class=\"%s\">\r\n", identifier2, identifier2, len(target), encapsulationTag()))
		result.WriteString(content.String())
		result.WriteString("</div></label>")
	} else {
		result.WriteString(content.String())
	}

	if recurseLevel > 0 {
		result.WriteString("\r\n</div></label>")
	}

	result.WriteString("\r\n</div>")

	rendition := result.String()

	if recurseLevel < 1 {
		this.writeToHTML(rendition, EMOJI_OBJECT, timestamp)
	}

	return rendition
}
