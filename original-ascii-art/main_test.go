package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	testCases := []struct {
		input    string
		fileName string
	}{
		{"HELLO", "HELLO.txt"},
		{"HeLlo HuMaN", "HeLloHuMaN.txt"},
		{"1Hello 2There", "1Hello2There.txt"},
		{"Hello\\nThere", "hellothere.txt"},
		{"{Hello & There #}", "1234.txt"},
		{"hello There 1 to 2!", "helloThere1to2.txt"},
		{"MaD3IrA&LiSboN", "MaD3IrA&LiSboN.txt"},
		{"1a\"#FdwHywR&/()=", "rand.txt"},
		{"{|}~", "random.txt"},
		{"[\\]^_ 'a", "any.txt"},
		{"RGB", "RGB.txt"},
		{":;<=>?@", "messy.txt"},
		{"\\!\" #$%&'()*+,-./", "lala.txt"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOQPRSTUVWXYZ.txt"},
		{"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz.txt"},
	}

	for _, tc := range testCases {
		t.Run(tc.fileName, func(innerT *testing.T) {
			check(innerT, tc.input, tc.fileName)
		})
	}
}

func check(t *testing.T, input string, fileName string) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", input}

	main()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	output, _ := os.ReadFile(fileName)

	t.Logf("Testing for argument %q", input)
	if out != string(output) {
		for i := range output {
			if out[i] != output[i] {
				if output[i] != 10 {
					output[i] = 42
				}
			}
		}
		t.Errorf("Ascii art failed, expected\n %v, got\n %v", string(output), out)
	} else {
		t.Log("PASS")
	}
}
