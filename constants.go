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

const EMOJI_SETUP = "🛠"
const EMOJI_CLEANUP = "🧹"
const EMOJI_PASSING_TEST = "✅"
const EMOJI_SUBJECTIVE_TEST = "🤔"
const EMOJI_INCONCLUSIVE_TEST = "🛑"
const EMOJI_FAILING_TEST = "❌"
const EMOJI_DEBUG = "🐞"
const EMOJI_ERROR = "😱"
const EMOJI_BOOLOG = "📝"
const EMOJI_TEXT_BOOLOG_CONCLUDE = "⤴️"
const EMOJI_TEXT_BLANK_LINE = ""
const EMOJI_OBJECT = "🔲"
const EMOJI_CAUSED_BY = "→"
const EMOJI_OUTGOING = "↗️"
const EMOJI_INCOMING = "↩️"

const uNKNOWN = "(unknown)"
const nAMELESS = "(name not given)"
const aLREADY_CONCLUDED_MESSAGE = "An attempt was made to write to a Boolog that was already concluded.\r\n<li>Once a Boolog has been concluded it can no longer be written to.\r\n<li>Passing a Boolog to the ShowBoolog() method will automatically conclude it."
const mAX_OBJECT_FIELDS_TO_DISPLAY = 10
const mAX_SHOW_OBJECT_RECURSION = 10
const mAX_HEADERS_TO_DISPLAY = 10
const mAX_BODY_LENGTH_TO_DISPLAY = 500
const sTARTING_CONTENT = "<table class=\"left_justified\">\r\n"
const bOOLOG_CONCLUDED = "Boolog Concluded"
const TIME_FMT_DATE = "2006-01-02"
const TIME_FMT_TIME = "15:04:05.0000"
const TIME_FMT_TEXT = TIME_FMT_DATE + " " + TIME_FMT_TIME
