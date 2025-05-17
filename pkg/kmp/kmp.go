package main

import (
	"fmt"
	"os"
	"time"
)

func computeLPSArray(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)

	length := 0

	lps[0] = 0

	i := 1
	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func KMPSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)

	if m == 0 {
		return []int{}
	}

	lps := computeLPSArray(pattern)

	positions := []int{}

	i := 0
	j := 0

	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			positions = append(positions, i-j)
			j = lps[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return positions
}

func displayResults(text, pattern string, positions []int) {
	fmt.Println("Text length:", len(text))
	fmt.Println("Pattern length:", len(pattern))
	fmt.Println("Number of occurrences:", len(positions))

	if len(positions) > 0 {
		fmt.Println("\nOccurrences at positions:")
		for i, pos := range positions {
			fmt.Printf("%d. Position %d: %s\n", i+1, pos, highlightMatch(text, pos, len(pattern)))
		}
	} else {
		fmt.Println("\nPattern not found in the text.")
	}
}

func highlightMatch(text string, pos, patternLen int) string {
	contextSize := 10
	startPos := max(0, pos-contextSize)
	endPos := min(len(text), pos+patternLen+contextSize)

	substring := text[startPos:endPos]

	relativePos := pos - startPos
	matchedPart := substring[relativePos : relativePos+patternLen]
	highlightedSubstring := substring[:relativePos] + "<<" + matchedPart + ">>" + substring[relativePos+patternLen:]

	return highlightedSubstring
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sampleText := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum
dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident,
sunt in culpa qui officia deserunt mollit anim id est laborum.`

	samplePattern := "dolor"

	if len(os.Args) > 2 {
		sampleText = os.Args[1]
		samplePattern = os.Args[2]
	} else {
		fmt.Println("Using default sample text and pattern.")
		fmt.Println("To use custom text and pattern: go run main.go \"text\" \"pattern\"")
		fmt.Println()
	}

	text := sampleText
	pattern := samplePattern

	startTime := time.Now()

	positions := KMPSearch(text, pattern)

	executionTime := time.Since(startTime)

	fmt.Println("String Matching Results (KMP Algorithm)")
	fmt.Println("======================================")
	fmt.Println("Pattern to find:", pattern)
	fmt.Println()

	displayResults(text, pattern, positions)

	fmt.Printf("\nExecution time: %v\n", executionTime)
	fmt.Println("\nTime Complexity: O(n+m) where n is text length and m is pattern length")
	fmt.Println("Space Complexity: O(m) for the LPS array of pattern")
}
