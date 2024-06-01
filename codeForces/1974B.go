package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount, sLen int
	var s string
	fmt.Fscan(in, &testsCount)
	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &sLen)
		fmt.Fscan(in, &s)
		sRune := []rune(s)
		sort.Slice(sRune, func(i int, j int) bool {
			return sRune[i] < sRune[j]
		})
		rList := make([]rune, 0, 26)

		rList = append(rList, sRune[0])
		for j, k := 0, 0; j < len(sRune); j++ {
			if rList[k] != sRune[j] {
				k++
				rList = append(rList, sRune[j])
			}
		}

		rMap := make(map[rune]rune, 26)
		for j := 0; j < len(rList); j++ {
			rMap[rList[j]] = rList[len(rList)-j-1]
		}
		resRune := []rune(s)
		for j := 0; j < len(resRune); j++ {
			resRune[j] = rMap[resRune[j]]
			fmt.Fprint(out, string(resRune[j]))
		}
		fmt.Fprintln(out)
	}
}
