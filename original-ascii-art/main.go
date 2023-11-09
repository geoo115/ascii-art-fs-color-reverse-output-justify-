package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Args = []string{"cmd", "Hello\\n\\nThere"}
	// Get command-line arguments and store them in the 'Arg' variable, excluding the program name
	Arg := os.Args[1:]

	// Ensure that only one command-line argument is provided
	if len(Arg) != 1 {
		return
	}

	// Check that the provided argument only contains valid characters (' ' to '-')
	// This limits the input to printable ASCII characters, which have ASCII art representations in the "standard.txt" file
	for _, r := range Arg[0] {
		if r < ' ' || r > '~' {
			return
		}
	}

	// Read the contents of the "standard.txt" file, which contains the ASCII art representations for each character
	bytes, err := os.ReadFile("standard.txt")
	// If there's an error reading the file, print the error and exit the program
	if err != nil {
		os.Exit(1)
	}
	// Split the file content into an array of strings, with each line as an element
	lines := strings.Split(string(bytes), "\n")

	// Initialize a rune slice 'arr' to store the characters from the input string
	var arr []rune
	// Declare a boolean variable 'Newline' to control the processing of newline characters
	Newline := false

	// Iterate over each character in the input string
	for i, r := range Arg[0] {
		/* If 'Newline' is true, print the current set of characters as ASCII art,
		reset the 'arr' slice, and set 'Newline' to false */
		if Newline {
			Newline = false
			printArt(arr, lines)
			arr = []rune{}
			continue
		}

		// Check if the current character is a backslash ('\\') and not the last character in the input string
		if r == '\\' && len(Arg[0]) != i+1 {
			/* If the character following the backslash character is an ('n'),
			set 'Newline' to true and skip the current iteration */
			if Arg[0][i+1] == 'n' {
				Newline = true
				continue
			}
		}
		// Add the current character to the 'arr' slice
		arr = append(arr, r)
	}
	// Print the final set of characters as ASCII art
	printArt(arr, lines)
}

// printArt is a helper function that takes a rune array and
// an array of lines containing the art data, and prints the
// ASCII art representation of the input text
func printArt(arr []rune, lines []string) {
	// Iterate through the 8 lines of ASCII art for each character
	if len(arr) != 0 {
		for line := 1; line <= 8; line++ {
			// Loop through each character (rune) in the 'arr' slice
			for _, r := range arr {
				// Calculate the index offset for the current character's ASCII art
				// Each character has 9 lines of ASCII art (including a blank line for separation)
				skip := (r - 32) * 9

				// Print the corresponding line of ASCII art for the current character
				/* Use the 'skip' value to find the correct starting index for the character's art
				in the 'lines' slice and add the current line number to print the correct line of the art */
				// time.Sleep(300 * time.Millisecond)
				fmt.Print(lines[line+int(skip)])
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}
