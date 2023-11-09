package output

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var SelectedFont string

func Process(args []string, outputFileName string) {
	if len(args) < 1 {
		return
	}

	// Select the font based on provided arguments
	for _, arg := range args {
		switch arg {
		case "standard":
			SelectedFont = "fonts/standard.txt"
		case "shadow":
			SelectedFont = "fonts/shadow.txt"
		case "thinkertoy":
			SelectedFont = "fonts/thinkertoy.txt"
		}
	}

	bannerStr := args[0]
	for _, arg := range args[1:] {
		if arg != "standard" && arg != "thinkertoy" && arg != "shadow" {
			bannerStr += " " + arg
		}
	}

	var multilineBanner bool

	var bannerStrArr []string
	for _, char := range bannerStr {
		bannerStrArr = append(bannerStrArr, string(char))
	}

	end := len(bannerStrArr) - 1

	for i := range bannerStrArr {
		if i != 0 {
			if i == end && bannerStrArr[i] == "!" && bannerStrArr[i-1] == "\\" {
				bannerStrArr = remove(bannerStrArr, i-1)
				i = i - 1
			}
			if bannerStrArr[i] == "n" && bannerStrArr[i-1] == "\\" {
				multilineBanner = true
			}
		}
	}
	bannerStr = strings.Join(bannerStrArr, "")

	// Specify directory where you want to store files
	outputDirectory := "tests"

	// Create the directory if it doesn't exist
	err := os.MkdirAll(outputDirectory, 0755) // 0755 sets read/write permissions for owner and read for others
	if err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}

	// Join directory and file name to create full path
	fullPath := filepath.Join(outputDirectory, outputFileName)

	// Create output file in the specified directory
	file, err := os.Create(fullPath)

	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	defer file.Close()

	banner := ""
	if multilineBanner {
		lines := strings.Split(bannerStr, "\\n")
		for _, line := range lines {
			for i := 0; i < 8; i++ {
				for _, char := range line {
					banner += GetLine(SelectedFont, 1+int(char-' ')*9+i)
				}
				fmt.Fprintln(file, banner)
				banner = ""
			}
		}
	} else {
		for i := 0; i < 8; i++ {
			for _, char := range bannerStr {
				banner += GetLine(SelectedFont, 1+int(char-' ')*9+i)
			}

			fmt.Fprintln(file, banner)
			banner = ""
		}
	}
}

func GetLine(fontFileName string, lineNumber int) string {
	file, err := os.Open(fontFileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	currentLine := 0
	line := ""
	for scanner.Scan() {
		if currentLine == lineNumber {
			line = scanner.Text()
		}
		currentLine++
	}
	return line
}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
