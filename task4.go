package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, substring := parseInput()

	findPalindrome(text, substring)

}

func parseInput() (string, string) {
	var text, substring string
	reader := bufio.NewReader(os.Stdin)

	println("Enter text")
	text, _ = reader.ReadString('\n')
	println("Enter substring")
	substring, _ = reader.ReadString('\n')
	// get rid of '\n' in the end
	text = text[:len(text)-1]
	substring = substring[:len(substring)-1]

	return text, substring
}

func findPalindrome(text string, substring string) {
	var word string
	var foundPalindrome = false
	endIndex := len(text) - 1
	var startIndex int

	for endIndex > 0 {
		skipSpaces(text, &endIndex)
		getWord(text, &word, &startIndex, &endIndex)
		if strings.Contains(word, substring) {
			foundPalindrome = isPalindrome(word)
			if foundPalindrome {
				break
			}
		}
		endIndex = startIndex
		if endIndex > 0 {
			endIndex--
		}
	}

	printResults(text, substring, word, startIndex, foundPalindrome)

}

func getWord(text string, word *string, pstartIndex *int, pendIndex *int) {
	*pstartIndex = *pendIndex
	for *pstartIndex-1 >= 0 && unicode.IsLetter(rune(text[*pstartIndex-1])) {
		*pstartIndex--
	}

	*word = text[*pstartIndex : *pendIndex+1]

}

func skipSpaces(text string, pendIndex *int) {
	for unicode.IsSpace(rune(text[*pendIndex])) {
		*pendIndex--
		if *pendIndex == 0 {
			break
		}
	}
}

func isPalindrome(word string) bool {
	startIndex := 0
	endIndex := len(word) - 1
	result := true
	for endIndex >= startIndex {
		if word[startIndex] != word[endIndex] {
			result = false
			break
		}
		endIndex--
		startIndex++
	}

	return result
}

func printResults(text string, substring string, word string, startIndex int, foundPalindrome bool) {
	if foundPalindrome {
		// fmt.Printf("%s\n", word)
		lastOccurrence := startIndex + strings.LastIndex(word, substring)
		fmt.Printf("%s%s", text[:lastOccurrence], text[lastOccurrence+len(substring):len(text)])
	} else {
		println("There is no palindrome")
	}
}