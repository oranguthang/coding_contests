package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type PatientRecord struct {
	orderNum, idx int
	isMoved       bool
}

func moveDown(idx, windowsCount, patientsCount int, patientsVisits []PatientRecord, places []rune) bool {
	tmpPatientsVisits := make([]PatientRecord, patientsCount)
	tmpPlaces := make([]rune, patientsCount)
	copy(tmpPatientsVisits, patientsVisits)
	copy(tmpPlaces, places)
	for i := idx; i >= 0; i-- {
		if tmpPatientsVisits[i].isMoved {
			return false
		}
		if tmpPatientsVisits[i].orderNum == 1 {
			return false
		}
		tmpPatientsVisits[i].orderNum--
		tmpPatientsVisits[i].isMoved = true
		tmpPlaces[tmpPatientsVisits[i].idx] = '-'
		if i-1 >= 0 && tmpPatientsVisits[i].orderNum > tmpPatientsVisits[i-1].orderNum {
			break
		} else if i-1 >= 0 && tmpPatientsVisits[i].orderNum < tmpPatientsVisits[i-1].orderNum {
			return false
		}
	}
	copy(patientsVisits, tmpPatientsVisits)
	copy(places, tmpPlaces)
	return true
}

func moveUp(idx, windowsCount, patientsCount int, patientsVisits []PatientRecord, places []rune) bool {
	tmpPatientsVisits := make([]PatientRecord, patientsCount)
	tmpPlaces := make([]rune, patientsCount)
	copy(tmpPatientsVisits, patientsVisits)
	copy(tmpPlaces, places)
	for i := idx; i < patientsCount; i++ {
		if tmpPatientsVisits[i].isMoved {
			return false
		}
		if tmpPatientsVisits[i].orderNum == windowsCount {
			return false
		}
		tmpPatientsVisits[i].orderNum++
		tmpPatientsVisits[i].isMoved = true
		tmpPlaces[tmpPatientsVisits[i].idx] = '+'
		if i+1 < patientsCount && tmpPatientsVisits[i].orderNum < tmpPatientsVisits[i+1].orderNum {
			break
		} else if i+1 < patientsCount && tmpPatientsVisits[i].orderNum > tmpPatientsVisits[i+1].orderNum {
			return false
		}
	}
	copy(patientsVisits, tmpPatientsVisits)
	copy(places, tmpPlaces)
	return true
}

func main() {
	f, err := os.Open("D:\\Downloads\\282\\84")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = f

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount, windowsCount, patientsCount int
	fmt.Fscan(in, &testsCount)
	//moveDowns := 0
	//moveUps := 0

	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &windowsCount, &patientsCount)
		patientsVisits := make([]PatientRecord, patientsCount)
		for i := 0; i < patientsCount; i++ {
			fmt.Fscan(in, &patientsVisits[i].orderNum)
			patientsVisits[i].idx = i
			patientsVisits[i].isMoved = false
		}

		sort.Slice(patientsVisits, func(i, j int) bool {
			return patientsVisits[i].orderNum < patientsVisits[j].orderNum
		})

		canReschedule := true

		places := []rune(strings.Repeat("0", patientsCount))
		for j := 0; j < patientsCount-1; j++ {
			if patientsVisits[j].orderNum == patientsVisits[j+1].orderNum {
				//if j == 0 {
				//	canReschedule = moveDown(j, windowsCount, patientsCount, patientsVisits, places)
				//	if canReschedule {
				//		continue
				//	}
				//}
				canReschedule = moveDown(j, windowsCount, patientsCount, patientsVisits, places)
				if canReschedule {
					continue
				}
				canReschedule = moveUp(j+1, windowsCount, patientsCount, patientsVisits, places)
				if !canReschedule {
					break
				}
				//canReschedule = moveDown(j, windowsCount, patientsCount, patientsVisits, places)
				//moveDowns++
				//if !canReschedule {
				//	canReschedule = moveDown(j, windowsCount, patientsCount, patientsVisits, places)
				//	//canReschedule = moveUp(j+1, windowsCount, patientsCount, patientsVisits, places)
				//	//moveUps++
				//	if !canReschedule {
				//		break
				//	}
				//}
			}
		}

		if !canReschedule {
			fmt.Println("x")
		} else {
			fmt.Println(string(places))
		}
		//fmt.Println(moveDowns, moveUps)
	}
}
