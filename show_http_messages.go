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
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"url"

	"github.com/google/uuid"
)

func (this Boolog) ShowHttpRequest(req http.Request, callback CallbackFunction) {
	timestamp := time.Now()
	reqUrl := req.URL
	var result *strings.Builder
	params, _ := url.ParseQuery(reqUrl.RawQuery)

	result.WriteString("<div class=\"outgoing implied_caution\">\r\n")

	textRendition := fmt.Sprintf("%s %s", req.Method, reqUrl.Path)

	result.WriteString(fmt.Sprintf("<center><h2>%sn</h2>", textRendition))
	result.WriteString(fmt.Sprintf("<small><b><i>%s</i></b></small>", reqUrl.Host))

	// Hide the full URL
	identifier := uuid.NewString()
	result.WriteString(fmt.Sprintf("<br><br><label for=\"%s\">\r\n<input id=\"%s\" type=\"checkbox\"><small><i>(show complete URL)</i></small>\r\n<div class=\"%s\">\r\n", identifier, identifier, encapsulationTag()))

	result.WriteString(fmt.Sprintf("<br>\r\n%s\r\n", strings.ReplaceAll(reqUrl.RawPath, "&", "&amp;")))
	result.WriteString("</div>\r\n</label>")

	if len(params) < 1 {
		result.WriteString("<br><br><small><i>(no query)</i></small>")
	} else {
		result.WriteString("<br><br><b>Queries</b><br><table class=\"gridlines\">\r\n")

		for key, values := range params {
			for _, value := range values {
				result.WriteString("<tr><td>")
				result.WriteString(key)
				result.WriteString("</td><td>")

				if value == "" {
					result.WriteString("(unset)")
				} else {
					processString(key, value, callback)
				}
			}
			result.WriteString("</td></tr>")
		}
		result.WriteString("\r\n</table>")
	}

	payload, _ := req.GetBody()
	defer payload.Close()

	result.WriteString("<br>")
	result.WriteString(this.renderHeadersAndBody(req.Header, payload, callback))

	this.writeToHTML(result.String(), EMOJI_OUTGOING, timestamp)
	this.echoPlainText(textRendition, EMOJI_OUTGOING, timestamp)
}

// TODO: Will need to figure out how to properly parse/render the headers & body
func (this Boolog) renderHeadersAndBody(header http.Header, io.ReadCloser, timestamp time.Time) string {

}
