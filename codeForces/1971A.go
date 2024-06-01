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

	var testsCount, a, b int
	fmt.Fscan(in, &testsCount)
	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &a, &b)
		if a > b {
			fmt.Fprintln(out, b, a)
		} else {
			fmt.Fprintln(out, a, b)
		}
	}
}
