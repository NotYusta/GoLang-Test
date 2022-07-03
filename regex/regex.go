package regex

import (
	"fmt"
	"regexp"
	"strings"
)

func GenerateRegexWord(value string) string {
	var words []string 
	for _, word := range strings.Split(value, "") {
		if word == "a" {
			word = "(" + "a|@" + ")"
		}

		if word == "1" || word == "i" {
			word = "(" +  "1|i" +")"
		}

		if word == "u" || word == "0" || word == "o" {
			word = "(" + "u|0|o" + ")"
		}

		if word == "3" || word == "e" {
			word = "(" + "3|e" + ")"
		}

		if word == "9" || word == "6" || word == "g" {
			word = "(" + "9|6|g" + ")"
		}

		if word == " " || word == "_" || word == "-" || word == "=" {
			word = "(" + "_|-|=| |" + ")"
		}

		words = append(words, word)
	}

	return strings.Join(words, "")
}

func Test() {
	word := GenerateRegexWord("anjay_2")
	rwordgex, _ := regexp.Compile(word)
	// word result
	fmt.Println(word)

	// regex match result

	regexTestWords := []string{
		"anjay 2",
		"anjay-2",
		"anjay_2",
		"anjay2",
		"anjay@2",
	}

	for _, regexTestWord := range regexTestWords {
		fmt.Println("word: " + regexTestWord)
		fmt.Println("value: " + fmt.Sprintf("%v", rwordgex.MatchString(regexTestWord)))
	}
}