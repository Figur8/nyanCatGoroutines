// Copyright 2017 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build example

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

func run() error {
	f, err := os.Open("Nyanyanyanyanyanyanya.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer p.Close()

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

/**
  Unicode / UTF8 Characters
  https://en.wikipedia.org/wiki/UTF-8
  https://en.wikipedia.org/wiki/Unicode#UTF
  https://en.wikibooks.org/wiki/Unicode/Character_reference/2000-2FFF
*/
const pentagon = "\u2B1F"
const hexagon = "\u2B22"
const circle = "\u2B24"
const elipse = "\u2B2E"
const black = "30"
const red = "31"
const green = "32"
const yellow = "33"
const blue = "34"
const purple = "35"
const lightblue = "36"
const gray = "37"

func printExpecificColoredTextForEachColorInList(text string, colorList []string) {
	for _, color := range colorList {
		if text == pentagon {
			if color != yellow {
				continue
			}
			printColoredText(text+" ", color)
		} else if color == purple {
			if text == elipse {
				continue
			}
			printColoredText(text+" ", color)
		} else {
			printColoredText(text+" ", color)
			if text == circle {
				return
			}
		}
	}
}

func printColoredText(text string, colorCode string) {
	print(mountColoredText(text, colorCode))
}

/**
  ANSI escape is what can help us make visual changes
  https://en.wikipedia.org/wiki/ANSI_escape_code
*/
func mountColoredText(text, colorCode string) string {
	CSIColorClear := "\033[0m"
	CSIColorStart := "\033[1;" + colorCode + "m"
	return CSIColorStart + text + CSIColorClear
}

func main() {
	go run()

	colorList := []string{
		yellow,
		blue,
		purple,
		gray,
	}
	textList := []string{
		pentagon,
		elipse,
		hexagon,
		circle,
	}
	defer println("")
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		loopTimes, err := strconv.Atoi(argsWithoutProg[0])
		if err != nil {
			println("Fail to undestand Loop Times Value:")
			panic(err)
		}
		i := 0
		for i < loopTimes {
			printAllColoredTextInList(textList, colorList)
			i++
		}
		return
	}
	for {
		printAllColoredTextInList(textList, colorList)
	}
}

func printAllColoredTextInList(textList, colorList []string) {
	for index := 0; index < len(textList); index++ {
		printExpecificColoredTextForEachColorInList(textList[index], colorList)
	}
}
