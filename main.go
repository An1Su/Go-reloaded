package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"go-reloaded/piscine"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide file name.")
		return
	}
	if len(os.Args) != 3 {
		fmt.Println("Both input and output filenames should be provided.")
		return
	}

	fileAdr := os.Args[1]
	newfileAdr := os.Args[2]
	originalText := readText(fileAdr)

	modifiedText := originalText

	for {
		keyWord, nrOfWords, keywordPosition, fullKeyword := checkExpressions(modifiedText)
		if keyWord == "" {
			break
		}
		modifiedText = textModifier(modifiedText, keyWord, nrOfWords, keywordPosition, fullKeyword)
	}

	modifiedText = rightArticle(modifiedText)
	modifiedText = punctuationFormat(modifiedText)
	modifiedText = quoteFormat(modifiedText)
	err := writeText(newfileAdr, modifiedText)
	if err != nil {
		fmt.Println("Error writting to the file:", err)
		return
	}
}

func readText(fileAdr string) (string) {

	content, err := os.ReadFile(fileAdr)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func writeText(newfileAdr string, modifiedText string) error {
	file, err := os.Create(newfileAdr)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(modifiedText)
	return err
}

func textModifier(originalText string, keyWord string, nrOfWords int, keywordPosition int, fullKeyword string) string {
	modifiedText := string(originalText[:keywordPosition]) + string(originalText[keywordPosition+len(fullKeyword):])

	for i := 0; i < nrOfWords; i++ {

		newWordPosition, theWord := findWord(modifiedText, keywordPosition)

		if newWordPosition == -1 {
			break
		}

		var newWord string

		switch keyWord {

		case "hex":
			digit, _ := strconv.ParseInt(theWord, 16, 64)
			newWord = strconv.FormatInt(digit, 10)
		case "bin":
			digit, _ := strconv.ParseInt(theWord, 2, 64)
			newWord = strconv.FormatInt(digit, 10)
		case "up":
			newWord = piscine.ToUpper(theWord)
		case "low":
			newWord = piscine.ToLower(theWord)
		case "cap":
			newWord = piscine.Capitalize(theWord)

		}

		modifiedText = string(modifiedText[:newWordPosition]) + newWord + string(modifiedText[newWordPosition+len(theWord):])

		keywordPosition = newWordPosition

		if keywordPosition == 0 {
			break
		}

	}

	return modifiedText
}

func checkExpressions(originalText string) (string, int, int, string) { // Returns keyWord, nrOfWords, keywordPosition, fullKeyword
	keyWord := ""
	fullKeyword := ""
	keywordPosition := -1
	nrOfWords := -1

	regCheck := regexp.MustCompile(`(?:\s*)?\((hex|bin|up|low|cap)(?:\,\s*)?(\d+)?\)`)

	foundKeyWord := regCheck.FindStringSubmatch(originalText)

	if len(foundKeyWord) == 3 {

		keywordPosition = regCheck.FindStringIndex(originalText)[0]
		keyWord = foundKeyWord[1]
		nrOfWords = 1

		if foundKeyWord[2] != "" {
			nrOfWords, _ = strconv.Atoi(foundKeyWord[2])
		}

		fullKeyword = foundKeyWord[0]

	}

	return keyWord, nrOfWords, keywordPosition, fullKeyword
}

func findWord(originalText string, keywordPosition int) (int, string) {
	var theWord string
	wordFound := false
	sliceOfRunes := []rune(originalText)

	for i := keywordPosition - 1; i >= 0; i-- {
		if sliceOfRunes[i] == ' ' && wordFound {
			wordStartposition := i + 1
			theWord = string(sliceOfRunes[wordStartposition:keywordPosition])
			return wordStartposition, theWord

		} else if sliceOfRunes[i] != ' ' {
			wordFound = true
		}
	}

	theWord = string(sliceOfRunes[:keywordPosition])
	return 0, theWord
}

func rightArticle(text string) string {
	runeText := []rune(text)
	var result []rune

	for i := 0; i < len(runeText); i++ {
		char := runeText[i]

		if (char == 'a' || char == 'A') && i < len(runeText)-1 {
			nextChar := runeText[i+1]

			if nextChar == ' ' {
				if i+2 < len(runeText) && anArticle(runeText[i+2]) {
					result = append(result, char, 'n')
					continue
				}
			}

			if i+1 < len(runeText) && nextChar == 'n' {
				if i > 0 && i+2 < len(runeText) && runeText[i+2] == ' ' && runeText[i-1] == ' ' {
					if i+3 < len(runeText) && !anArticle(runeText[i+3]) {
						result = append(result, 'a')
						i++
						continue
					}
				}
			}
		}

		result = append(result, char)
	}

	return string(result)
}

func anArticle(c rune) bool {
	s := "aeiouhAEIOUH"
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func punctuationFormat(text string) string {
	rText := []rune(text)
	var result []rune

	// check the space number after the punctuation!!!

	for i, char := range rText {

		if (i == 0 || i == len(rText)-1) && char == ' ' {
			continue
		}

		switch char {

		case ' ':

			if i+1 < len(rText) && (keyPunctuation(rText[i+1]) || rText[i+1] == ' ') {
				continue
			}

		default:
			if keyPunctuation(char) {

				result = append(result, char)
				if i+1 < len(rText) && !keyPunctuation(rText[i+1]) && rText[i+1] != ' ' {
					result = append(result, ' ')
				}
				continue
			}
		}

		result = append(result, char)
	}

	return string(result)
}

func keyPunctuation(char rune) bool {
	switch char {
	case '.', ',', '!', '?', ':', ';':
		return true
	}
	return false
}

func quoteFormat(s string) string {
	var result string
	quoteMet := false

	for i := 0; i < len(s); i++ {

		if !(quoteMet) {
			if s[i] == '\'' {
				quoteMet = true
				result += string(s[i])
				continue
			}
		} else {

			if s[i] == ' ' && s[i-1] == '\'' {
				continue
			}

			if i < len(s)-1 && s[i] == ' ' && s[i+1] == '\'' {
				result += string(s[i+1])
				i++
				quoteMet = false
				continue
			}

		}

		result += string(s[i])
	}

	return result
}
