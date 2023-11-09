package reverse

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Process takes a file path and converts the ASCII art in the file back into text
func Process(filePath string) error {
	if filePath == "" {
		return errors.New("no file provided")
	}

	template, err := openFileAndReadLineByLine("standard.txt", "fonts")
	if err != nil {
		return fmt.Errorf("error reading template file: %v", err)
	}
	segmentedTemplate := segment(template)
	targetTemplate, err := openFileAndReadLineByLine(filePath, "examples")
	if err != nil {
		return fmt.Errorf("error reading target file: %v", err)
	}
	result := ""

	for len(targetTemplate[0]) > 0 {
		for i, v := range segmentedTemplate {
			if checkIfLetterIsPresent(v, targetTemplate) {
				result = result + string(rune(i+32))
				targetTemplate = removeLetter(len(v[0]), targetTemplate)
			}
		}
	}
	fmt.Println(result)
	return nil
}

func openFileAndReadLineByLine(input string, folder string) ([]string, error) {
	// Prepend the folder name to the input file name
	bytes, err := os.ReadFile(folder + "/" + input)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	lines := strings.Split(string(bytes), "\n")
	return lines, nil
}

func segment(template []string) [][]string {
	var result [][]string

	for i := 0; i < len(template)-1; i = i + 9 {
		temp := template[i+1 : i+9]
		result = append(result, temp)
	}
	return result
}

func removeLetter(length int, word []string) []string {
	for i, v := range word[0 : len(word)-1] {
		word[i] = v[length:]
	}
	return word
}

func checkIfLetterIsPresent(letter, word []string) bool {
	found := true
	if len(letter[0]) > len(word[0]) {
		return false
	}
	for i, v := range word[0 : len(word)-1] {

		if letter[i] != v[:len(letter[i])] {
			found = false
		}
	}
	return found
}
