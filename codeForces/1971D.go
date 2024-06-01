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
		swapsCount := 0
		sRunes := []rune(s)
		if sRunes[0] == '1' {
			for i := 1; i < len(sRunes); i++ {
				if sRunes[i-1] != sRunes[i] {
					swapsCount++
				}
			}
			// for strings like 11110000 we have only one answer - 2
			if swapsCount == 1 {
				fmt.Fprintln(out, 2)
				continue
			} else {
				swapsCount = 0
			}
		}
		for i := 1; i < len(sRunes); i++ {
			if sRunes[i-1] != sRunes[i] {
				swapsCount++
			}
		}
		// we have at least 1 part
		if swapsCount == 0 {
			swapsCount++
		}
		fmt.Fprintln(out, swapsCount)
	}
}
