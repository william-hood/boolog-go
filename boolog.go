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
		this.forPlainText.WriteString(fmt.Sprintf("<td><h2>%s</h2></td>", emoji))
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

func (this Boolog) Info(message string, emoji string) error {
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

func (this Boolog) Error(message string, emoji string) error {
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

func (this Boolog) showBoolog(subordinate Boolog) (string, error) {
	return this.showBoologDetailed(subordinate, EMOJI_BOOLOG, "neutral", 0)
}

func (this Boolog) showBoologDetailed(subordinate Boolog, emoji string, style string, recurseLevel int) (string, error) {
	var err error = nil
	timestamp := time.Now()
	subordinateContent := subordinate.Conclude()
	result := wrapAsSubordinate(subordinate.Title, subordinateContent, style)

	if recurseLevel < 1 {
		err = this.writeToHTML(result, emoji, timestamp)
	}

	return result, err
}

func NewBoologSimple(logTitle string, htmlOutputFileName string) Boolog {
	return NewBoolog(logTitle, htmlOutputFileName, "", "", "")
}

func NewBoolog(logTitle string, htmlOutputFileName string, htmlHeader string, htmlStyling string, textOutputFileName string) Boolog {
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

	{
		hTML, err := os.Create(htmlOutputFileName)
		if err != nil {
			panic(err)
		}
		result.forHTML = hTML

		if htmlStyling == "" {
			htmlStyling = CLASSIC_STYLING
		}

		result.forHTML.WriteString("<html>\r\n<meta charset=\"UTF-8\">\r\n<head>\r\n<title>$title</title>\r\n")
		result.forHTML.WriteString(htmlStyling)
		result.forHTML.WriteString("</head>\r\n<body>\r\n")

		if htmlHeader == "" {
			htmlHeader = defaultHeader(result.Title)
		}

		result.forHTML.WriteString(htmlHeader)
	}

	return *result
}

func defaultHeader(title string) string {
	var builder strings.Builder
	builder.WriteString("<h1>")
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
	return fmt.Sprintf("\r\n\r\n<div class=\"boolog %s\">\r\n<label for=\"%s\">\r\n<input id=\"$identifier\" class=\"gone\" type=\"checkbox\">\r\n<h2>%s</h2>\r\n<div class=\"%s\">\r\n%s\r\n</div></label></div>", style, identifier, boologTitle, encapsulationTag(), boologContent)
}
