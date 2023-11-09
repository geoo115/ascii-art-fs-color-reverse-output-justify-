package color

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	colorReset = "\033[0m"
)

func Process(input string, colors []string, banner string, colorWords string) error {
	var selectedColors []string

	for _, color := range colors {
		var selectedColor string
		if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
			rgbValues := strings.TrimPrefix(color, "rgb(")
			rgbValues = strings.TrimSuffix(rgbValues, ")")
			rgbValuesSplit := strings.Split(rgbValues, ",")
			if len(rgbValuesSplit) != 3 {
				return fmt.Errorf("invalid RGB color value")
			}
			r, err := strconv.Atoi(strings.TrimSpace(rgbValuesSplit[0]))
			if err != nil {
				return fmt.Errorf("invalid RGB color value: %v", err)
			}
			g, err := strconv.Atoi(strings.TrimSpace(rgbValuesSplit[1]))
			if err != nil {
				return fmt.Errorf("invalid RGB color value: %v", err)
			}
			b, err := strconv.Atoi(strings.TrimSpace(rgbValuesSplit[2]))
			if err != nil {
				return fmt.Errorf("invalid RGB color value: %v", err)
			}
			selectedColor = fmt.Sprintf("\u001b[38;2;%d;%d;%dm", r, g, b)
		} else {
			switch color {
			case "red":
				selectedColor = "\u001b[38;2;255;0;0m"
			case "green":
				selectedColor = "\u001b[38;2;0;255;0m"
			case "yellow":
				selectedColor = "\u001b[38;2;255;255;0m"
			case "blue":
				selectedColor = "\u001b[38;2;0;0;255m"
			case "purple":
				selectedColor = "\u001b[38;2;161;32;255m"
			case "cyan":
				selectedColor = "\u001b[38;2;0;183;235m"
			case "white":
				selectedColor = "\u001b[38;2;255;255;255m"
			case "pink":
				selectedColor = "\u001b[38;2;255;0;255m"
			case "grey":
				selectedColor = "\u001b[38;2;128;128;128m"
			case "black":
				selectedColor = "\u001b[38;2;0;0;0m"
			case "brown":
				selectedColor = "\u001b[38;2;160;128;96m"
			case "orange":
				selectedColor = "\u001b[38;2;255;160;16m"
			default:
				selectedColor = colorReset
			}
		}

		// selectedColor = ... whatever it is ...
		selectedColors = append(selectedColors, selectedColor)
	}
	// Create color queue
	colorQueue := NewQueue()

	// Push all colors into queue
	for _, color := range selectedColors {
		colorQueue.Push(color)
	}

	// Default banner
	if banner == "" {
		banner = "standard"
	}

	// Read the art template file
	bytes, err := os.ReadFile(fmt.Sprintf("fonts/%s.txt", banner))
	if err != nil {
		return err
	}

	// Split the lines based on banner type
	var lines []string
	if banner == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	colorWordSlice := strings.Split(colorWords, ",")

	// Create ASCII Art
	createArt(input, colorQueue, colorWordSlice, lines)
	return nil
}

func createArt(input string, colorQueue *Queue, colorWords []string, template []string) {
	for line := 1; line <= 8; line++ {
		remainingInput := input
		for remainingInput != "" {
			colorMatch := ""
			color := colorReset
			for _, word := range colorWords {
				if strings.HasPrefix(remainingInput, word) {
					// Found a match. Apply color and move on.
					colorMatch = word
					color = colorQueue.Pop().(string)
					colorQueue.Push(color) // Requeue the color
					break
				}
			}
			if colorMatch == "" {
				// No match found. Print without color and move on.
				fmt.Print(template[line+int(remainingInput[0]-32)*9], colorReset)
				remainingInput = remainingInput[1:]
			} else {
				for _, r := range colorMatch {
					fmt.Print(color, template[line+int(r-32)*9], colorReset)
				}
				remainingInput = strings.TrimPrefix(remainingInput, colorMatch)
			}
		}
		fmt.Println()
	}
}

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}
