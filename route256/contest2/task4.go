package main

import (
	"bufio"
	"fmt"
	"os"
)

//type BankCurrency struct {
//	rubleDollar, rubleEuro, dollarRuble, dollarEuro, euroRuble, euroDollar float64
//}

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
	var fromVal, toVal int
	for i := 0; i < testsCount; i++ {
		banksCurrency := make([][]float64, 3)
		for j := range banksCurrency {
			banksCurrency[j] = make([]float64, 6)
		}
		for j := 0; j < 3; j++ {
			for k := 0; k < 6; k++ {
				fmt.Fscan(in, &fromVal, &toVal)
				banksCurrency[j][k] = float64(toVal) / float64(fromVal)
			}
		}
		results := []float64{
			// ruble - dollar
			banksCurrency[0][0],
			banksCurrency[1][0],
			banksCurrency[2][0],

			// ruble - dollar, dollar - ruble, ruble - dollar
			banksCurrency[0][0] * banksCurrency[1][2] * banksCurrency[2][0],
			banksCurrency[0][0] * banksCurrency[2][2] * banksCurrency[1][0],
			banksCurrency[1][0] * banksCurrency[0][2] * banksCurrency[2][0],
			banksCurrency[1][0] * banksCurrency[2][2] * banksCurrency[0][0],
			banksCurrency[2][0] * banksCurrency[0][2] * banksCurrency[1][0],
			banksCurrency[2][0] * banksCurrency[1][2] * banksCurrency[0][0],

			// ruble - dollar, dollar - euro, euro - dollar
			banksCurrency[0][0] * banksCurrency[1][3] * banksCurrency[2][5],
			banksCurrency[0][0] * banksCurrency[2][3] * banksCurrency[1][5],
			banksCurrency[1][0] * banksCurrency[0][3] * banksCurrency[2][5],
			banksCurrency[1][0] * banksCurrency[2][3] * banksCurrency[0][5],
			banksCurrency[2][0] * banksCurrency[0][3] * banksCurrency[1][5],
			banksCurrency[2][0] * banksCurrency[1][3] * banksCurrency[0][5],

			// ruble - euro, euro - ruble, ruble - dollar
			banksCurrency[0][1] * banksCurrency[1][4] * banksCurrency[2][0],
			banksCurrency[0][1] * banksCurrency[2][4] * banksCurrency[1][0],
			banksCurrency[1][1] * banksCurrency[0][4] * banksCurrency[2][0],
			banksCurrency[1][1] * banksCurrency[2][4] * banksCurrency[0][0],
			banksCurrency[2][1] * banksCurrency[0][4] * banksCurrency[1][0],
			banksCurrency[2][1] * banksCurrency[1][4] * banksCurrency[0][0],

			// ruble - euro, euro - dollar
			banksCurrency[0][1] * banksCurrency[1][5],
			banksCurrency[0][1] * banksCurrency[2][5],
			banksCurrency[1][1] * banksCurrency[2][5],
			banksCurrency[1][1] * banksCurrency[0][5],
			banksCurrency[2][1] * banksCurrency[0][5],
			banksCurrency[2][1] * banksCurrency[1][5],
		}

		//fmt.Println(banksCurrency)
		//fmt.Println(results)
		maxRes := results[0]
		for _, val := range results {
			if maxRes < val {
				maxRes = val
			}
		}
		fmt.Println(maxRes)
	}
}
