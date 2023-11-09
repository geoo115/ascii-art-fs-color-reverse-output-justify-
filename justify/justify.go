package justify

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Process() error {
	text, fonttype, alignment, err := readCmd()
	if err != nil {
		return err
	}
	template, err := openFileAndReadLineByLine(fonttype)
	if err != nil {
		return err
	}
	segmentedTemplate := segment(template)
	createArt(text, segmentedTemplate, alignment)
	return nil
}

func createArt(text string, template [][]string, alignment string) {
	var arr [][]string
	wordArr := []string{"", "", "", "", "", "", "", ""}
	// New boolean flag to skip next character if it's part of \n
	skipNext := false
	for i, r := range text {
		if skipNext {
			// Skip this character, it was part of the \n
			skipNext = false
			continue
		}
		if r == '\\' && len(text) > i+1 && text[i+1] == 'n' {
			// Found \n in text, print current line and start a new one
			arr = append(arr, wordArr)
			if alignment == "justify" {
				printJustify(arr)
			} else {
				for _, wordArr := range arr {
					print(wordArr, alignment)
				}
			}
			// Start a new line
			arr = [][]string{}
			wordArr = []string{"", "", "", "", "", "", "", ""}
			// Skip the next character
			skipNext = true
			continue
		}
		if r == ' ' {
			arr = append(arr, wordArr)
			// Add an ASCII "space" (a block of actual space characters)
			wordArr = []string{"          ", "          ", "          ", "          ", "          ", "          ", "          ", "          "}
			arr = append(arr, wordArr)
			// Prepare for the next word
			wordArr = []string{"", "", "", "", "", "", "", ""}
			continue
		}
		wordArr = customAppend(wordArr, template[int(r)-32])
	}
	arr = append(arr, wordArr) // Add the last word

	lineArr := make([]string, 8)
	if alignment == "justify" {
		printJustify(arr)
	} else {
		for _, wordArr := range arr {
			for i := range lineArr {
				lineArr[i] += wordArr[i] // append a space after each word
			}
		}
		print(lineArr, alignment) // print the line once all words are processed
	}
}

func print(str []string, alignment string) {
	n := len(str)
	artlen := len(str[0])
	width := getTerminalWidth()

	switch alignment {
	case "center":
		diff := width/2 - artlen/2
		for i := 0; i < n; i++ {
			for j := 0; j < diff; j++ {
				fmt.Print(" ")
			}
			fmt.Println(str[i])
		}
	case "right":
		diff := width - artlen
		for i := 0; i < n; i++ {
			for j := 0; j < diff; j++ {
				fmt.Print(" ")
			}
			fmt.Println(str[i])
		}
	case "left":
		for i := 0; i < n; i++ {
			fmt.Println(str[i])
		}
	default:
		fmt.Println("You have entered incorrect alignment option!")
		return
	}
}

func customAppend(str, item []string) []string {
	if str == nil {
		str = make([]string, 8)
	}
	for i := 0; i < 8; i++ {
		str[i] = str[i] + item[i]
	}
	return str
}
func segment(template []string) [][]string {
	var result [][]string
	for i := 0; i < len(template)-1; i = i + 9 {
		temp := template[i+1 : i+9]
		result = append(result, temp)
	}
	return result
}

func readCmd() (string, string, string, error) {
	args := os.Args[1:]
	if len(args) != 3 {
		return "", "", "", fmt.Errorf("for this exercise you have to enter three input parameters. Try: \"$./ascii-art-fs --align=<option> <text> <font type>\"")
	}

	if args[0][:8] != "--align=" {
		return "", "", "", fmt.Errorf("you have made a mistake in --align=")
	}

	if hasError(args[1]) {
		return "", "", "", fmt.Errorf("input has non-readable ASCII characters")
	}

	option := args[0][8:]
	return args[1], args[2], option, nil
}

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	width, _ := strconv.Atoi(strings.Fields(string(out))[1])
	return width
}

func hasError(str string) bool {
	for _, r := range str {
		if r < 32 || r > 126 {
			return true
		}
	}
	return false
}

func openFileAndReadLineByLine(banner string) ([]string, error) {
	file := fmt.Sprintf("fonts/%s.txt", banner)
	// file deepcode ignore PT: <please specify a reason of ignoring this>
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var lines []string
	if banner == "standard" || banner == "shadow" {
		lines = strings.Split(string(bytes), "\n")
	} else {
		lines = strings.Split(string(bytes), "\r\n")
	}
	return lines, nil
}

func printJustify(arr [][]string) {
	n := len(arr)
	if n == 0 {
		return
	}
	m := len(arr[0])
	totalWidth := 0
	for _, wordArr := range arr {
		totalWidth += len(wordArr[0])
	}
	totalSpaces := getTerminalWidth() - totalWidth
	spacePerWord := totalSpaces / (n - 1)
	for i := 0; i < m; i++ {
		for j, wordArr := range arr {
			fmt.Print(wordArr[i])
			if j < n-1 {
				fmt.Print(strings.Repeat(" ", spacePerWord))
			}
		}
		fmt.Println()
	}
}
