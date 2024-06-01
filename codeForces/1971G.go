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

	var testsCount, n int
	fmt.Fscan(in, &testsCount)
	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &n)
		arr := make([]int, n, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		//for i := 0; i < len(arr); i++ {
		//	for j := i + 1; j < len(arr); j++ {
		//		if (arr[j] < arr[i]) && (arr[j]^arr[i] < 4) {
		//			arr[j], arr[i] = arr[i], arr[j]
		//		}
		//	}
		//}
		sort.Slice(arr, func(i, j int) bool {
			return (arr[j] > arr[i]) && (arr[j]^arr[i] < 4)
		})
		for i := 0; i < len(arr); i++ {
			fmt.Fprint(out, arr[i], " ")
		}
		fmt.Fprintln(out)
	}
}
