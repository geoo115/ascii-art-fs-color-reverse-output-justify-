package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func Process(input, banner string) error {
	for _, r := range input {
		if r < ' ' || r > '~' {
			return fmt.Errorf("invalid character: %c", r)
		}
	}

	// read the file from the fonts folder
	bytes, err := os.ReadFile(fmt.Sprintf("fonts/%s.txt", banner))
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var lines []string
	if banner == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	var arr []rune
	Newline := false
	var lineCount int
	var offset int

	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		for i, r := range input {
			if Newline {
				Newline = false
				art(arr, lines)
				arr = []rune{}
				continue
			}
			if r == '\\' && len(input) != i+1 {
				if input[i+1] == 'n' {
					Newline = true
					continue
				}
			}
			arr = append(arr, r)
		}
		art(arr, lines)

	} else {
		switch banner {
		case "card":
			lineCount = 7
			offset = 223

		case "colossal":
			lineCount = 9
			offset = 289
		case "metric":
			lineCount = 11
			offset = 705
		case "graffiti":
			lineCount = 7
			offset = 222
		case "matrix":
			lineCount = 10
			offset = 320
		case "rev":
			lineCount = 11
			offset = 353
		}

		for i, r := range input {
			if Newline {
				Newline = false
				printArt(arr, lines, lineCount, offset)
				arr = []rune{}
				continue
			}

			if r == '\\' && i != len(input)-1 {
				if input[i+1] == 'n' {
					Newline = true
					continue
				}
			}
			arr = append(arr, r)
		}
		printArt(arr, lines, lineCount, offset)
	}
	return nil
}

func printArt(arr []rune, lines []string, lineCount int, offset int) {
	if len(arr) != 0 {
		for line := 1; line <= lineCount; line++ {
			for _, r := range arr {
				skip := (r * rune(lineCount)) - rune(offset)
				fmt.Print(lines[line+int(skip)])
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}

func art(arr []rune, lines []string) {
	if len(arr) != 0 {
		for line := 1; line <= 8; line++ {
			for _, r := range arr {
				skip := (r - 32) * 9
				fmt.Print(lines[line+int(skip)])
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}
