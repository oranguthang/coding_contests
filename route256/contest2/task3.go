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

	var testsCount int
	fmt.Fscan(in, &testsCount)
	for i := 0; i < testsCount; i++ {
		var testString string
		isCorrestString := true
		fmt.Fscan(in, &testString)
		if len(testString) > 1 {
			baseLetter := testString[0]
			if baseLetter != testString[len(testString)-1] {
				isCorrestString = false
			} else {
				for j := 1; j < len(testString)-1; j++ {
					if testString[j] != baseLetter && testString[j-1] != testString[j+1] {
						isCorrestString = false
						break
					}
				}
			}

			//if len(testString) == 2 {
			//	if testString[0] == testString[1] {
			//		fmt.Println("YES")
			//	} else {
			//		fmt.Println("NO")
			//	}
			//}
		}
		if isCorrestString {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
