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

	var a, b int
	fmt.Fscan(in, &a, &b)

	fmt.Println(a - b)
}
