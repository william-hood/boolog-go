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

const THEME_CLASSIC = `
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
            background-image: linear-gradient(to bottom right, white, WhiteSmoke);
        }

        .failing_test_result {
            background-image: linear-gradient(to bottom right, MistyRose, salmon);
        }

        .inconclusive_test_result {
            background-image: linear-gradient(to bottom right, LemonChiffon, Moccasin);
        }

        .passing_test_result {
            background-image: linear-gradient(to bottom right, honeydew, palegreen);
        }

        .implied_good {
            background-image: linear-gradient(to bottom right, mintcream, honeydew);
        }

        .implied_caution {
            background-image: linear-gradient(to bottom right, LemonChiffon, oldlace);
        }

        .implied_bad {
            background-image: linear-gradient(to bottom right, Seashell, LavenderBlush);
        }

        .neutral {
            background-image: linear-gradient(to bottom right, White, LightGrey);
        }

        .old_parchment {
            background-image: radial-gradient(LightGoldenrodYellow, Cornsilk, Wheat);
        }

        .plate {
            background-image: radial-gradient(GhostWhite, LightSteelBlue);
        }

        .exception {
            background-image: linear-gradient(to bottom right, yellow, salmon);
        }

        .decaf_green {
            background-image: linear-gradient(to bottom right, #B0C6B2, #83A787);
        }

        .decaf_orange {
            background-image: linear-gradient(to bottom right, #EED886, #D0A403);
        }

        .decaf_green_light_roast {
            background-image: linear-gradient(to bottom right, #DCE9DD, #B0C6B2);
        }

        .decaf_orange_light_roast {
            background-image: linear-gradient(to bottom right, #F2E5B4, #EED886);
        }

        .desert_horizon {
            background-image: linear-gradient(to bottom, #127FCF, #52ACEE, #7EC7FD, #F7EFCA, #F6EDC2, #F5EBBA);
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
            background-color: black;
        }
        
        .centered {
            text-align: center;
        }

        .highlighted {
            background-image: linear-gradient(to bottom right, yellow, gold);
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
