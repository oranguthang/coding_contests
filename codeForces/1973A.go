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

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var p1, p2, p3 int
		fmt.Fscan(in, &p1, &p2, &p3)

		if isValid(p1, p2, p3) {
			fmt.Fprintln(out, maxDraws(p1, p2, p3))
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}

func isValid(p1, p2, p3 int) bool {
	// Input validation
	totalPoints := p1 + p2 + p3
	return (totalPoints%2 == 0) && (2*p3 <= totalPoints)
}

func maxDraws(p1, p2, p3 int) int {
	// Max number of drafts
	totalPoints := p1 + p2 + p3
	games := totalPoints / 2
	worstPlayerWins := games - p3
	return games - worstPlayerWins
}
