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
	fmt.Print("Unpacked string:", unpstring)
}

func Reformat(str string) (out string) {
	var repeat string
	var prevLetter rune
	var backslash bool
	for ind, currRune := range str {
		if unicode.IsDigit(currRune) {
			if backslash {
				out = out + string(currRune)
				prevLetter = currRune
				backslash = false
			} else {
				repeat = repeat + string(currRune)
				if ind+1 != len(str) {
					if unicode.IsDigit(rune(str[ind+1])) {
						continue
					} else {
						out = out + RepeatLetters(repeat, prevLetter)
					}
				} else {
					out = out + RepeatLetters(repeat, prevLetter)
				}
			}
		} else {
			fmt.Print(string(currRune))
			repeat = ""
			if currRune == '\\' {
				if backslash {
					backslash = false
					prevLetter = currRune
				} else {
					backslash = true
				}
			} else {
				if ind+1 != len(str) && !unicode.IsDigit(rune(str[ind+1])) {
					out = out + string(currRune)
				} else {
					prevLetter = currRune
				}
				if ind == len(str)-1 {
					out = out + string(currRune)
				}
			}

		}
	}
	return
}

func RepeatLetters(repeat string, prevLetter rune) string {
	repeatnum, _ := strconv.Atoi(repeat)
	return strings.Repeat(string(prevLetter), repeatnum)
}
