package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount, x, y int
	var screensCount int
	fmt.Fscan(in, &testsCount)

	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &x, &y)

		if x == 0 {
			screensCount = int(math.Ceil(float64(y) / 2))
			fmt.Fprintln(out, screensCount)
		} else if y == 0 {
			screensCount = int(math.Ceil(float64(x) / 15))
			fmt.Fprintln(out, screensCount)
		} else {
			screensCount = int(math.Ceil(float64(y) / 2))
			// if we have an odd Y, we need to add 4 free spots to last screen
			freeSpots := screensCount*7 + y%2*4
			if x < freeSpots {
				fmt.Fprintln(out, screensCount)
			} else {
				fmt.Fprintln(out, screensCount+int(math.Ceil(float64(x-freeSpots)/15)))
			}
		}
	}
}
