package regex

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	AlphaBetOnlyRegex = regexp.MustCompile("[a-zA-Z]+")
	NumberOnlyRegex   = regexp.MustCompile("[0-9]+")
)

func GenerateTypoRegexWord(value string) string {
	var words []string
	splittedWords := strings.Split(value, "")
	for wi, word := range splittedWords {
		var wordValue string
		if word == "a" || word == "A" {
			wordValue = "(" + "a|A|@?" + ")"
		} else if word == "1" || word == "I" || word == "i" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "1|i|I?" + ")"
			} else {
				wordValue = "(" + "1|i|I" + ")"
			}
		} else if word == "u" || word == "U" || word == "0" || word == "o" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "u|U|0|o?" + ")"
			} else {
				wordValue = "(" + "u|U|0|o" + ")"
			}
		} else if word == "3" || word == "e" || word == "E" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "3|e|E?" + "?)"
			} else {
				wordValue = "(" + "3|e|E" + ")"
			}
		} else if word == "9" || word == "6" || word == "g" || word == "G" {
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "9|6|g|G?" + ")"
			} else {
				wordValue = "(" + "9|6|g|G" + ")"
			}
		} else if word == " " || word == "_" || word == "-" || word == "=" {
			wordValue = "(.?)"
		} else {
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + word + "?)"
			} else {
				wordValue = word
			}
		}

		words = append(words, wordValue)
	}


	return strings.Join(words, "")
}

func GenerateTypoRegexWordLowercase(value string) string {
	var words []string
	splittedWords := strings.Split(value, "")
	for wi, word := range splittedWords {
		var wordValue string

		if word == "a" {
			wordValue = "(" + "a|@?" + ")"
		} else if word == "1" || word == "i" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				if splittedWords[wi-1] == word {
					words[wi-1] = "(" + "1|i?" + ")"
				}

				wordValue = "(" + "1|i?" + ")"
			} else {
				wordValue = "(" + "1|i" + ")"
			}
		} else if word == "u" || word == "0" || word == "o" {
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "u|0|o?" + ")"
			} else {
				wordValue = "(" + "u|0|o" + ")"
			}
		} else if word == "3" || word == "e" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "3|e?" + "?)"
			} else {
				wordValue = "(" + "3|e" + ")"
			}
		} else if word == "9" || word == "6" || word == "g" {
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "9|6|g?" + ")"
			} else {
				wordValue = "(" + "9|6|g" + ")"
			}
		} else if word == " " || word == "_" || word == "-" || word == "=" {
			wordValue = "(.?)"
		} else {
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + word + "?)"
			} else {
				wordValue = word
			}
		}

		words = append(words, wordValue)
	}

	return strings.Join(words, "")
}


func GenerateTypoRegexWordUppercase(value string) string {
	var words []string
	splittedWords := strings.Split(value, "")
	for wi, word := range splittedWords {
		var wordValue string

		if word == "A" {
			wordValue = "(" + "A|@?" + ")"
		} else if word == "1" || word == "I" {
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "1|I?" + ")"
			} else {
				wordValue = "(" + "1|I" + ")"
			}
		} else if word == "U" || word == "0" || word == "O" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "U|0|O?" + ")"
			} else {
				wordValue = "(" + "U|0|O" + ")"
			}
		} else if word == "3" || word == "E" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "3|E?" + "?)"
			} else {
				wordValue = "(" + "3|E" + ")"
			}
		} else if word == "9" || word == "6" || word == "G" {
			
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + "9|6|G?" + ")"
			} else {
				wordValue = "(" + "9|6|G" + ")"
			}
		} else if word == " " || word == "_" || word == "-" || word == "=" {
			wordValue = "(.?)"
		} else {
			if IsWordOptionableRegex(wi, len(splittedWords) -1, splittedWords, word) {
				wordValue = "(" + word + "?)"
			} else {
				wordValue = word
			}
		}

		words = append(words, wordValue)
	}

	return strings.Join(words, "")
}

// pWord : previous word, cWord: current word, nWord: next word
func IsWordOptionableRegex(wordIndex int, splittedWordsLength int, splittedWords []string, cWord string) bool {
	// (If index equals 0 or equals splitted words length) or (word index not equals splitted words length and next word is current word) or (word index not equals 0 and (word is number only or previous word is equals current word)) 
	
	// search if the word index is 0 or equals splitted words length
	return (wordIndex == 0 || wordIndex == splittedWordsLength) || 
	// search if the index is not equal with the splitted words length and next word is equal with the current word 
	(wordIndex != splittedWordsLength && splittedWords[wordIndex + 1] == cWord) ||
	// search if the word is a number if the index is not 0 (after a word)
	(wordIndex != 0 && (NumberOnlyRegex.MatchString(cWord) || splittedWords[wordIndex - 1] == cWord))
}

func TotalRegexGroup(reg *regexp.Regexp, word string) (totalGroup int) {
	for _, w := range reg.FindStringSubmatch(word) {
		if w == " " {
			continue
		} else {
			totalGroup++
		}
	}

	return totalGroup
}

func Test() {
	anjayTypoWordTest()
}

func anjayTypoWordTest() {
	println()
	
	originalWord := "anjay_2"
	word := GenerateTypoRegexWord(originalWord)
	rwordgex, _ := regexp.Compile(".*" + word + ".*")
	// word result
	fmt.Println("original word: " + originalWord)
	fmt.Println("regex pattern: " + word)
	println()

	// regex match result

	regexTestWords := []string{
		"anjay 2",
		"njay 2",
		"NJAY",
		"anj@y-2",
		"@njay2",
		"anjay@2",
		"   anjay@2   ",
		"anjay_2aaaaaaaaaaaaaa",
	}

	for _, regexTestWord := range regexTestWords {
		fmt.Println("word: " + regexTestWord)
		fmt.Println("group: " + fmt.Sprintf("%v", TotalRegexGroup(rwordgex, regexTestWord)))
		fmt.Println("value: " + fmt.Sprintf("%v", rwordgex.MatchString(regexTestWord)))
	}
	
	println()
}