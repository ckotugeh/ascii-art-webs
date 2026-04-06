package main

import (
	"bufio"
	"os"
	"strings"
)

// generateASCIIArt generates ASCII art from the provided text and banner.
// It assumes banner files follow the standard format: 9 lines per character (1 spacer + 8 content).
func generateASCIIArt(text, banner string) (string, error) {
	// 1. Open the banner file
	file, err := os.Open("banners/" + banner + ".txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 2. Read all lines into a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// 3. Prepare output using strings.Builder for efficiency
	var result strings.Builder

	// 4. Split input text into lines (handling literal \n strings if needed)
	inputLines := strings.Split(text, "\n")

	for _, line := range inputLines {
		// Handle empty lines in input (e.g., if the user hits Enter twice)
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// 5. Each ASCII character is 8 rows high
		// We start at 1 to skip the empty spacer line at the top of each block
		for i := 1; i <= 8; i++ {
			for _, char := range line {
				// Ignore non-printable characters outside standard ASCII range
				if char < 32 || char > 126 {
					continue
				}

				// Each character block is 9 lines total (8 rows + 1 spacer)
				// Index 0 in banner files is usually the spacer for space (char 32)
				index := (int(char) - 32) * 9

				// Check bounds to ensure the banner file actually contains this character
				if index+i < len(lines) {
					result.WriteString(lines[index+i])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
