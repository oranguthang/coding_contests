package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount int
	var s string
	fmt.Fscan(in, &testsCount)
	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &s)
		canSwapChars := false
		sRunes := []rune(s)
		for i := 0; i < len(sRunes)-1; i++ {
			if sRunes[i] != sRunes[i+1] {
				sRunes[i], sRunes[i+1] = sRunes[i+1], sRunes[i]
				canSwapChars = true
				break
			}
		}
		if canSwapChars {
			fmt.Fprintln(out, "YES")
			fmt.Fprintln(out, string(sRunes))
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
