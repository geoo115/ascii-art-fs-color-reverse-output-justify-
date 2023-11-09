package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Arg := os.Args[1:]

	if len(Arg) != 2 {
		return
	}

	for _, r := range Arg[0] {
		if r < ' ' || r > '~' {
			return
		}
	}

	bytes, err := os.ReadFile("standard.txt")
	if err != nil {
		os.Exit(1)
	}
	lines := strings.Split(string(bytes), "\n")

	var arr []rune
	Newline := false

	for i, r := range Arg[0] {
		if Newline {
			Newline = false
			printArt(arr, lines, Arg[1])
			arr = []rune{}
			continue
		}

		if r == '\\' && len(Arg[0]) != i+1 {
			if Arg[0][i+1] == 'n' {
				Newline = true
				continue
			}
		}
		arr = append(arr, r)
	}
	printArt(arr, lines, Arg[1])
}

func printArt(arr []rune, lines []string, outputFile string) {
	if len(arr) != 0 {
		// Open the output file in append mode
		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		for line := 1; line <= 8; line++ {
			var artLine string
			for _, r := range arr {
				skip := (r - 32) * 9
				artLine += lines[line+int(skip)]
			}
			artLine += "\n"
			_, err = f.WriteString(artLine)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		fmt.Println()
	}
}
