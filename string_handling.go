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
	"strings"
)

// Type CallbackFunction can be used to allow additional processing for strings passed through Boolog. Example uses would be JSON Pretty-Printing or Base64 decoding.
type CallbackFunction func(string, string) string

// Pass the result of this to .Info() (or any other method that outputs to HTML) and the text will show as a monospace font, keeping any line breaks or other formatting used.
func TreatAsCode(value string) string {
	return fmt.Sprintf("<pre><code><xmp>%s</xmp></code></pre>", value)
}

func processString(fieldName string, fieldValue string, callback CallbackFunction) string {
	result := strings.Clone(fieldValue)

	if callback != nil {
		result = callback(fieldName, result)
	}

	//result = result.replace("<", "&lt;").replace(">", "&gt;")
	return result
}
