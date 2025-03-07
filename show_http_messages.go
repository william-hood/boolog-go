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
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

const HTTP_MESSAGE_BODY = "HTTP Req/Resp Body/Payload"

func (this Boolog) ShowHttpRequest(req http.Request, callback CallbackFunction) {
	timestamp := time.Now()
	reqUrl := req.URL
	var result strings.Builder
	params, _ := url.ParseQuery(reqUrl.RawQuery)

	result.WriteString("<div class=\"outgoing implied_caution\">\r\n")

	textRendition := fmt.Sprintf("%s %s", req.Method, reqUrl.Path)

	result.WriteString(fmt.Sprintf("<center><h2>%sn</h2>", textRendition))
	result.WriteString(fmt.Sprintf("<small><b><i>%s</i></b></small>", reqUrl.Host))

	// Hide the full URL
	identifier := uuid.NewString()
	result.WriteString(fmt.Sprintf("<br><br><label for=\"%s\">\r\n<input id=\"%s\" type=\"checkbox\"><small><i>(show complete URL)</i></small>\r\n<div class=\"%s\">\r\n", identifier, identifier, encapsulationTag()))

	result.WriteString(fmt.Sprintf("<br>\r\n%s\r\n", strings.ReplaceAll(reqUrl.String(), "&", "&amp;")))
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
					result.WriteString(processString(key, value, callback))
				}
			}
			result.WriteString("</td></tr>")
		}
		result.WriteString("\r\n</table>")
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewReader((bodyBytes)))
	}

	result.WriteString("<br>")
	result.WriteString(this.renderHeadersAndBody(req.Header, bodyBytes, callback))

	this.writeToHTML(result.String(), EMOJI_OUTGOING, timestamp)
	this.echoPlainText(textRendition, EMOJI_OUTGOING, timestamp)
}

func (this Boolog) ShowHttpResponse(resp http.Response, callback CallbackFunction) {
	var result strings.Builder

	timestamp := time.Now()
	statusCode := resp.StatusCode
	style := "implied_bad"
	if (statusCode > 199) && (statusCode < 300) {
		style = "implied_good"
	}

	result.WriteString(fmt.Sprintf("<div class=\"incoming %s\">\r\n", style))

	textRendition := resp.Status

	result.WriteString(fmt.Sprintf("<center><h2>%s</h2>", textRendition))

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewReader((bodyBytes)))
	}

	result.WriteString(this.renderHeadersAndBody(resp.Header, bodyBytes, callback))

	this.writeToHTML(result.String(), EMOJI_INCOMING, timestamp)
	this.echoPlainText(textRendition, EMOJI_INCOMING, timestamp)
}

func (this Boolog) renderHeadersAndBody(headerMap http.Header, bodyBytes []byte, callback CallbackFunction) string {
	var result strings.Builder

	// Headers
	if len(headerMap) > 0 {
		result.WriteString("<br><b>Headers</b><br>")
		var renderedHeaders strings.Builder
		renderedHeaders.WriteString("<table class=\"gridlines\">\r\n")

		for key, values := range headerMap {
			renderedHeaders.WriteString("<tr><td>")
			renderedHeaders.WriteString(key)
			renderedHeaders.WriteString("</td><td>")

			if len(values) < 1 {
				renderedHeaders.WriteString("<small><i>(empty)</i></small>")
			} else if len(values) == 1 {
				// Attempt Base64 Decode and JSON pretty-print here.
				renderedHeaders.WriteString(processString(key, values[0], callback))
			} else {
				renderedHeaders.WriteString("<table class=\"gridlines neutral\">\r\n")
				for _, thisValue := range values {
					renderedHeaders.WriteString("<tr><td>")
					// Attempt Base64 Decode and JSON pretty-print here.
					renderedHeaders.WriteString(processString(key, thisValue, callback))
					renderedHeaders.WriteString("</td></tr>")
				}
				renderedHeaders.WriteString("\r\n</table>")
			}

			renderedHeaders.WriteString("</td></tr>")
		}
		renderedHeaders.WriteString("\r\n</table><br>")

		if len(headerMap) > MAX_HEADERS_TO_DISPLAY {
			identifier1 := uuid.NewString()
			result.WriteString(fmt.Sprintf("<label for=\"%s\">\r\n<input id=\"%s\" type=\"checkbox\">\r\n(show %d headers)\r\n<div class=\"%s\">\r\n", identifier1, identifier1, len(headerMap), encapsulationTag()))
			result.WriteString(renderedHeaders.String())
			result.WriteString("</div></label>")
		} else {
			result.WriteString(renderedHeaders.String())
		}
	} else {
		result.WriteString("<br><br><small><i>(no headers)</i></small><br>\r\n")
	}

	// Body
	if len(bodyBytes) < 1 {
		result.WriteString("<br><br><small><i>(no payload)</i></small></center>")
	} else {
		// Attempt Base64 Decode and JSON pretty-print here.
		stringPayload := string(bodyBytes)
		payloadSize := len(stringPayload)
		result.WriteString("<br><b>Payload</b><br></center>\r\n")
		renderedBody := TreatAsCode(processString(HTTP_MESSAGE_BODY, stringPayload, callback))

		if payloadSize > MAX_BODY_LENGTH_TO_DISPLAY {
			identifier2 := uuid.NewString()
			result.WriteString(fmt.Sprintf("<label for=\"%s\">\r\n<input id=\"%s\" type=\"checkbox\">\r\n(show large payload)\r\n<div class=\"%s\">\r\n", identifier2, identifier2, encapsulationTag()))
			result.WriteString(renderedBody)
			result.WriteString("</div></label>")
		} else {
			result.WriteString(renderedBody)
		}
	}

	return result.String()
}

func (this Boolog) ShowHttpTransaction(req http.Request, callback CallbackFunction) *http.Response {
	this.ShowHttpRequest(req, callback)
	resp, _ := http.DefaultClient.Do(&req)
	this.ShowHttpResponse(*resp, callback)
	return resp
}
