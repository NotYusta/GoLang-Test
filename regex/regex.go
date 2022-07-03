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
			word = "(" + "a|@?" + ")"
		} else if word == "1" || word == "i" {
			word = "(" +  "1|i" +")"
		} else if word == "u" || word == "0" || word == "o" {
			word = "(" + "u|0|o" + ")"
		} else if word == "3" || word == "e" {
			word = "(" + "3|e" + ")"
		} else if word == "9" || word == "6" || word == "g" {
			word = "(" + "9|6|g" + ")"
		} else if word == " " || word == "_" || word == "-" || word == "=" {
			word = "(.?)"
		}

		words = append(words, word)
	}

	return strings.Join(words, "")
}

func Test() {
	anjayWordTest()
}

func anjayWordTest() {
	println()
	
	originalWord := "anjay_2"
	word := GenerateRegexWord(originalWord)
	rwordgex, _ := regexp.Compile(word)
	// word result
	fmt.Println("original word: " + originalWord)
	fmt.Println("regex pattern: " + word)
	println()

	// regex match result

	regexTestWords := []string{
		"anjay 2",
		"njay 2",
		"anj@y-2",
		"anjay_2",
		"@njay2",
		"anjay@2",
	}

	for _, regexTestWord := range regexTestWords {
		fmt.Println("word: " + regexTestWord)
		fmt.Println("value: " + fmt.Sprintf("%v", rwordgex.MatchString(regexTestWord)))
	}
	
	println()
}