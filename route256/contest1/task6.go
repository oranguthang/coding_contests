package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type CardRecord struct {
	maxCard, idx int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var friendsCount, cardsCount int
	fmt.Fscan(in, &friendsCount, &cardsCount)

	friendsCards := make([]CardRecord, friendsCount)
	for i := 0; i < friendsCount; i++ {
		fmt.Fscan(in, &friendsCards[i].maxCard)
		friendsCards[i].idx = i
	}

	sort.Slice(friendsCards, func(i, j int) bool {
		return friendsCards[i].maxCard < friendsCards[j].maxCard
	})

	places := make([]int, friendsCount)
	currentCard := friendsCards[0].maxCard + 1
	for i := 0; i < friendsCount; {
		if currentCard > friendsCards[i].maxCard {
			places[friendsCards[i].idx] = currentCard
			i++
		}
		currentCard++
	}

	if currentCard > cardsCount+1 {
		fmt.Println(-1)
	} else {
		for _, value := range places {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}
