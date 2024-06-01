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

	var testsCount, a, b, c, d int
	fmt.Fscan(in, &testsCount)
	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		if a > b {
			a, b = b, a
		}
		if c > d {
			c, d = d, c
		}
		if (a < c && b > d) || (c < a && d > b) || b < c || d < a {
			fmt.Fprintln(out, "NO")
		} else {

			fmt.Fprintln(out, "YES")
		}
	}
}
