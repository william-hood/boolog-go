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

const THEME_DARK_GRADIENT = `
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
            border: 0.1em solid black;
            display: inline-block;
            background-image: linear-gradient(to bottom right, #3A3B3C, #34282C);
        }

        .failing_test_result {
            background-image: linear-gradient(to bottom right, #800000, #2B1B17);
        }

        .inconclusive_test_result {
            background-image: linear-gradient(to bottom right, #AF9B60, #493D26);
        }

        .passing_test_result {
            background-image: linear-gradient(to bottom right, #4E9258, #228B22);
        }

        .implied_good {
            background-image: linear-gradient(to bottom right, #228B22, #254117);
        }

        .implied_caution {
            background-image: linear-gradient(to bottom right, #AF9B60, #966F33);
        }

        .implied_bad {
            background-image: linear-gradient(to bottom right, #B21807, #660000);
        }

        .neutral {
            background-image: linear-gradient(to bottom right, #4D4D4F, #040720);
        }

        .old_parchment {
            background-image: radial-gradient(#E6BF83, #C8AD7F, #C19A6B);
        }

        .plate {
            background-image: radial-gradient( #838996, #2B3856);
        }

        .exception {
            background-image: linear-gradient(to bottom right, #E8A317, #660000);
        }


        body {
            background-color: #3A3B3C;
            color: #FFFFFF;
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
            background-color: #FFFFFF;
        }
        
        .centered {
            text-align: center;
        }

        .highlighted {
            background-image: linear-gradient(to bottom right, #E8A317, #CA762B);
        }

        .outlined {
            display: inline-block;
            border-radius: 0.5em;
            border: 0.05em solid black;
            padding: 0.2em 0.2em;
        }

        .object {
            border-radius: 1.5em;
            border: 0.3em solid black;
            display: inline-block;
            padding: 0.4em 0.4em;
        }

        .incoming {
            border-radius: 3em 0.5em 0.5em 3em;
            border: 0.3em solid black;
            display: inline-block;
            padding: 1em 1em;
        }

        .outgoing {
            border-radius: 0.5em 3em 3em 0.5em;
            border: 0.3em solid black;
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
        }
        
        label {
            cursor: pointer;
        }
    </style>
`
