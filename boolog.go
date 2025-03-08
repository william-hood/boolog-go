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
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type HeaderFunction func(string) string

type Boolog struct {
	Title          string
	forPlainText   *os.File
	forHTML        *os.File
	ShowTimeStamps bool
	ShowEmojis     bool
	Content        *strings.Builder
	isConcluded    bool
	firstEcho      bool
}

func (this Boolog) WasUsed() bool {
	return (this.Content.Len() - len(STARTING_CONTENT)) > 0
}

func (this Boolog) echoPlainText(message string, emoji string, timestamp time.Time) error {
	if this.forPlainText == nil {
		// Silently decline
		return nil
	}

	if this.isConcluded {
		return errors.New(BOOLOG_CONCLUDED)
	}

	if this.firstEcho {
		this.firstEcho = false
		this.echoPlainText("", EMOJI_TEXT_BLANK_LINE, timestamp)
		this.echoPlainText(this.Title, EMOJI_BOOLOG, timestamp)
	}

	if this.ShowTimeStamps {
		this.forPlainText.WriteString(timestamp.Format(TIME_FMT_TEXT))
		this.forPlainText.WriteString("\t")
	}

	if this.ShowEmojis {
		this.forPlainText.WriteString(fmt.Sprintf("%s\t", emoji))
	}

	this.forPlainText.WriteString(fmt.Sprintf("%s\r\n", message))

	return nil
}

func (this Boolog) writeToHTML(message string, emoji string, timestamp time.Time) error {
	if this.isConcluded {
		return errors.New(BOOLOG_CONCLUDED)
	}

	this.Content.WriteString("<tr>")

	if this.ShowTimeStamps {
		this.Content.WriteString(fmt.Sprintf("<td class=\"min\"><small>%s</small></td><td>&nbsp;</td><td class=\"min\"><small>%s</small></td><td>&nbsp;</td>", timestamp.Format(TIME_FMT_DATE), timestamp.Format(TIME_FMT_TIME)))
	}

	if this.ShowEmojis {
		this.Content.WriteString(fmt.Sprintf("<td><h2>%s</h2></td>", emoji))
	}

	this.Content.WriteString(fmt.Sprintf("<td>%s</td></tr>\r\n", message))

	return nil
}

func (this Boolog) Conclude() string {
	if !this.isConcluded {
		timestamp := time.Now()
		this.echoPlainText("", EMOJI_TEXT_BOOLOG_CONCLUDE, timestamp)
		this.echoPlainText("", EMOJI_TEXT_BLANK_LINE, timestamp)

		if this.forPlainText != os.Stdin {
			this.forPlainText.Sync()
			this.forPlainText.Close()
		}

		this.isConcluded = true

		this.Content.WriteString("\r\n</table>")

		if this.forHTML != nil {
			this.forHTML.WriteString(this.Content.String())
			this.forHTML.WriteString("\r\n</body>\r\n</html>")
			this.forHTML.Sync()
			this.forHTML.Close()
		}
	}

	return this.Content.String()
}

func (this Boolog) Info(message string) error {
	return this.InfoDetailed(message, EMOJI_TEXT_BLANK_LINE)
}

func (this Boolog) InfoDetailed(message string, emoji string) error {
	timestamp := time.Now()
	htmlErr := this.writeToHTML(message, emoji, timestamp)
	if htmlErr != nil {
		return htmlErr
	}
	textErr := this.echoPlainText(message, emoji, timestamp)
	return textErr
}

func (this Boolog) Debug(message string) error {
	timestamp := time.Now()
	htmlErr := this.writeToHTML(highlight(message), EMOJI_DEBUG, timestamp)
	if htmlErr != nil {
		return htmlErr
	}
	textErr := this.echoPlainText(message, EMOJI_DEBUG, timestamp)
	return textErr
}

func (this Boolog) Error(message string) error {
	timestamp := time.Now()
	htmlErr := this.writeToHTML(highlight(message), EMOJI_ERROR, timestamp)
	if htmlErr != nil {
		return htmlErr
	}
	textErr := this.echoPlainText(message, EMOJI_ERROR, timestamp)
	return textErr
}

func (this Boolog) SkipLine() error {
	timestamp := time.Now()
	htmlErr := this.writeToHTML("", EMOJI_TEXT_BLANK_LINE, timestamp)
	if htmlErr != nil {
		return htmlErr
	}
	textErr := this.echoPlainText("", EMOJI_TEXT_BLANK_LINE, timestamp)
	return textErr
}

func (this Boolog) ShowBoolog(subordinate Boolog) (string, error) {
	return this.ShowBoologDetailed(subordinate, EMOJI_BOOLOG, "neutral", 0)
}

func (this Boolog) ShowBoologDetailed(subordinate Boolog, emoji string, style string, recurseLevel int) (string, error) {
	var err error = nil
	timestamp := time.Now()
	subordinateContent := subordinate.Conclude()
	result := wrapAsSubordinate(subordinate.Title, subordinateContent, style)

	if recurseLevel < 1 {
		err = this.writeToHTML(result, emoji, timestamp)
	}

	return result, err
}

func NewBoolog(logTitle string, htmlOutputFileName string) Boolog {
	return NewBoologDetailed(logTitle, htmlOutputFileName, nil, "", "")
}

func NewBoologDetailed(logTitle string, htmlOutputFileName string, htmlHeaderFunction HeaderFunction, htmlStyling string, textOutputFileName string) Boolog {
	result := new(Boolog)

	if logTitle == "" {
		logTitle = UNKNOWN
	}

	result.Title = logTitle
	result.ShowTimeStamps = true
	result.ShowEmojis = true
	result.isConcluded = false
	result.firstEcho = true
	result.Content = &strings.Builder{}
	result.Content.WriteString(STARTING_CONTENT)

	if textOutputFileName == "" {
		result.forPlainText = os.Stdout
	} else {
		plainText, err := os.Create(textOutputFileName)
		if err != nil {
			panic(err)
		}
		result.forPlainText = plainText
	}

	if htmlOutputFileName != "" {
		hTML, err := os.Create(htmlOutputFileName)
		if err != nil {
			panic(err)
		}
		result.forHTML = hTML

		if htmlStyling == "" {
			htmlStyling = CLASSIC_STYLING
		}

		result.forHTML.WriteString(fmt.Sprintf("<html>\r\n<meta charset=\"UTF-8\">\r\n<head>\r\n<title>%s</title>\r\n", result.Title))
		result.forHTML.WriteString(htmlStyling)
		result.forHTML.WriteString("</head>\r\n<body>\r\n")

		if htmlHeaderFunction == nil {
			htmlHeaderFunction = defaultHeader
		}

		result.forHTML.WriteString(htmlHeaderFunction(result.Title))
	}

	return *result
}

func defaultHeader(title string) string {
	var builder strings.Builder
	builder.WriteString("<h1>")
	builder.WriteString("<img src=\"data:image/png;base64,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\" alt=\"Boolog Logo\" />")
	builder.WriteString("&nbsp;")
	builder.WriteString(title)
	builder.WriteString("</h1>\r\n<hr>\r\n<small><i>Powered by Boolog...</i></small>\r\n\r\n")

	return builder.String()
}

// In Kotlin there's an optional style variable that defaults to "highlighted". For this I just hardcoded it.
func highlight(message string) string {
	var builder strings.Builder

	builder.WriteString("<p class=\"highlighted outlined\">&nbsp;")
	builder.WriteString(message)
	builder.WriteString("&nbsp;</p>")

	return builder.String()
}

func encapsulationTag() string {
	return "lvl-" + uuid.NewString()
}

func wrapAsSubordinate(boologTitle string, boologContent string, style string) string {
	identifier := uuid.NewString()
	return fmt.Sprintf("\r\n\r\n<div class=\"boolog %s\">\r\n<label for=\"%s\">\r\n<input id=\"%s\" class=\"gone\" type=\"checkbox\">\r\n<h2>%s</h2>\r\n<div class=\"%s\">\r\n%s\r\n</div></label></div>", style, identifier, identifier, boologTitle, encapsulationTag(), boologContent)
}
