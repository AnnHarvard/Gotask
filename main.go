package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var pstring, unpstring string
	fmt.Print("Packed string:")
	fmt.Scan(&pstring)
	unpstring = Reformat(pstring)
	if unpstring == "" {
		fmt.Print("Некорректная строка!")
	} else {
		fmt.Print("Unpacked string:", unpstring)
	}

}

func Reformat(str string) (out string) {
	var repeat string
	var prevLetter rune
	var backslash bool
	if _, err := strconv.Atoi(str); err == nil {
		return ""
	}
	strrune := []rune(str)
	for ind, currRune := range strrune {
		if unicode.IsDigit(currRune) {
			if backslash {
				out = out + string(currRune)
				prevLetter = currRune
				backslash = false
			} else {

				repeat = repeat + string(currRune)
				if ind+1 != len(strrune) {
					if unicode.IsDigit(rune(strrune[ind+1])) {
						continue
					} else {
						out = out + RepeatLetters(repeat, prevLetter)
					}
				} else {
					out = out + RepeatLetters(repeat, prevLetter)
				}
			}
		} else {
			repeat = ""
			if currRune == '\\' {
				if backslash {
					backslash = false
					prevLetter = currRune
				} else {
					backslash = true
				}
			} else {
				if ind+1 != len(strrune) && !unicode.IsDigit(rune(strrune[ind+1])) {
					out = out + string(currRune)
					fmt.Println("tr", currRune, " ", string(currRune))
					fmt.Println("r", string(strrune[ind+1]))
				} else {
					prevLetter = currRune
				}
				if ind == len(strrune)-1 {
					out = out + string(currRune)
				}
			}

		}
	}
	return
}

func RepeatLetters(repeat string, prevLetter rune) string {
	repeatnum, _ := strconv.Atoi(repeat)
	if unicode.IsDigit(prevLetter) {
		repeatnum--
	}
	return strings.Repeat(string(prevLetter), repeatnum)
}
