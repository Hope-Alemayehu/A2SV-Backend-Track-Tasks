package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// wordCountCleaner removes unwanted characters and preserves whitespace.
func wordCountCleaner(s string) string {
	var cleanedWord strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			cleanedWord.WriteRune(r)
		}
	}
	return cleanedWord.String()
}

// palindromStringCleaner removes unwanted characters for palindrome checking.
func palindromStringCleaner(s string) string {
	var cleanString strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleanString.WriteRune(unicode.ToLower(r))
		}
	}
	return cleanString.String()
}

// wordFrequencyCounter returns the frequency of each word in a string.
func wordFrequencyCounter(s string) map[string]int {
	words := wordCountCleaner(s)
	wordsArray := strings.Fields(words)
	freqCounter := make(map[string]int)
	for _, word := range wordsArray {
		freqCounter[strings.ToLower(word)]++
	}
	return freqCounter
}

// checkPalindrome checks if a string is a palindrome.
func checkPalindrome(s string) bool {
	word := palindromStringCleaner(s)
	left, right := 0, len(word)-1
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {

	var choice int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press 1 To count your word frequency")
	fmt.Println("Press 2 To check if a string is a Palindrome")
	fmt.Println("Press 3 To Exit")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter a String: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)

		result := wordFrequencyCounter(s)
		fmt.Println("Word Frequency:")
		for word, frequency := range result {
			fmt.Printf("%s: %d\n", word, frequency)
		}

	case 2:
		fmt.Print("Enter a String: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)

		isPalindrome := checkPalindrome(s)
		if isPalindrome {
			fmt.Printf("%s is a palindrome\n", s)
		} else {
			fmt.Printf("%s is not a palindrome\n", s)
		}

	case 3:
		fmt.Println("Exiting program...")
		os.Exit(0)
	default:
		fmt.Println("Enter a valid choice")
	}
}
