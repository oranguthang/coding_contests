package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type SecondsRecord struct {
	seconds, idx int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount, participantsCount int
	fmt.Fscan(in, &testsCount)

	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &participantsCount)
		times := make([]SecondsRecord, participantsCount)
		for j := 0; j < participantsCount; j++ {
			fmt.Fscan(in, &times[j].seconds)
			times[j].idx = j
		}

		sort.Slice(times, func(i, j int) bool {
			return times[i].seconds < times[j].seconds
		})

		places := make([]int, participantsCount)
		currentPlace := 1
		for j := 0; j < participantsCount; j++ {
			if j != 0 && (times[j].seconds == times[j-1].seconds || times[j].seconds == times[j-1].seconds+1) {
				places[times[j].idx] = places[times[j-1].idx]
			} else {
				places[times[j].idx] = currentPlace
			}
			currentPlace++
		}

		for _, value := range places {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}
