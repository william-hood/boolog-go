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

const THEME_DARK_FLAT = `
    <style>
        html {
            font-family: sans-serif
        }

        [class*='lvl-'] {
            display: none;
            cursor: auto;
        }

        input:checked~[class*='lvl-'] {
            display: block;
        }

        .gone {
            display: none;
        }

        .boolog {
            font-family: sans-serif;
            border-radius: 0.25em;
            display: inline-block;
            background-color: #2E1A47;
            color: #F3C331;
        }

        .failing_test_result {
            background-color: #970005;
            color: #000000;
        }

        .inconclusive_test_result {
            background-color: #8A7839;
            color: #000000;
        }

        .passing_test_result {
            background-color: #007122;
            color: #000000;
        }

        .implied_good {
            background-color: #A0D6B4;
            color: #000000;
        }

        .implied_caution {
            background-color: #EDE7C1;
            color: #000000;
        }

        .implied_bad {
            background-color: #FA8072;
            color: #000000;
        }

        .neutral {
            background-color: #893AA0;
            color: #000000;
        }

        .old_parchment {
            background-color: #8E844C;
            color: #000000;
        }

        .plate {
            background-color: #909DFD;
            color: #000000;
        }

        .exception {
            background-color: #C11B17;
            color: #000000;
        }


        body {
            background-color: #0D0000;
            color: #F3C331;
        }
        table,
        th,
        td {
            padding: 0.1em 0em;
            margin-left: auto;
            margin-right: auto;
        }

        td.min {
            width: 1%;
            white-space: nowrap;
        }

        h1 {
            font-size: 3em;
            margin: 0em
        }

        h2 {
            font-size: 1.75em;
            margin: 0.2em
        }

        hr {
            border: none;
            height: 0.3em;
            background-color: #F3C331;
        }
        
        .centered {
            text-align: center;
        }

        .highlighted {
            background-color: #F3C331;
            color: #000000;
        }

        .outlined {
            display: inline-block;
            border-radius: 0.5em;
            padding: 0.2em 0.2em;
        }

        .object {
            border-radius: 1.5em;
            display: inline-block;
            padding: 0.4em 0.4em;
        }

        .incoming {
            border-radius: 3em 0.5em 0.5em 3em;
            display: inline-block;
            padding: 1em 1em;
        }

        .outgoing {
            border-radius: 0.5em 3em 3em 0.5em;
            display: inline-block;
            padding: 1em 1em;
        }

        .left_justified {
	        float: left;
        }

        table.gridlines,
        table.gridlines th,
        table.gridlines td {
            padding: 0.4em 0.4em;
            border-collapse: collapse;
            border: 0.02em solid black;
            color: #000000;
        }
        
        label {
            cursor: pointer;
        }
    </style>
`
