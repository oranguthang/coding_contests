package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//f, err := os.Open("D:\\Downloads\\282\\84")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//oldStdin := os.Stdin
	//defer func() { os.Stdin = oldStdin }()
	//os.Stdin = f

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var lettersCount, testsCount int
	fmt.Fscan(in, &lettersCount, &testsCount)
	lettersMap := make(map[string]int, 0)
	for i := 0; i < lettersCount; i++ {
		var letter string
		fmt.Fscan(in, &letter)
		lettersMap[letter]++
	}
	for i := 0; i < testsCount; i++ {
		pinMap := make(map[string]int, 0)
		var pin string
		fmt.Fscan(in, &pin)
		for _, char := range pin {
			pinMap[string(char)]++
		}
		if fmt.Sprint(lettersMap) == fmt.Sprint(pinMap) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
