package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type directoryStruct struct {
	Dir        string
	Files      []string           `json:",omitempty"`
	Folders    []*directoryStruct `json:",omitempty"`
	IsInfested bool               `json:",omitempty"`
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount, testLinesCount int
	fmt.Fscan(in, &testsCount)

	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &testLinesCount)
		testLinesCount++ // include \n sign from previous line
		lines := make([]string, testLinesCount, testLinesCount)
		for j := 0; j < testLinesCount; j++ {
			if line, err := in.ReadString('\n'); err == nil {
				lines[j] = strings.TrimSpace(line)
			}
		}

		jsonString := strings.Join(lines, "")
		dir := new(directoryStruct)
		_ = json.Unmarshal([]byte(jsonString), &dir)

		dirsQueue := make([]*directoryStruct, 0, 10)
		dirsQueue = append(dirsQueue, dir)
		infestedFilesCount := 0
		//infestedFoldersSet := make(map[string]struct{}, 10)

		for len(dirsQueue) > 0 {
			currDir := dirsQueue[0]
			dirsQueue = dirsQueue[1:]
			//_, isInfested := infestedFoldersSet[currDir.Dir]
			if currDir.Files != nil && !currDir.IsInfested {
				for fileIdx := range currDir.Files {
					if strings.HasSuffix(currDir.Files[fileIdx], ".hack") {
						currDir.IsInfested = true
						break
					}
				}
			}
			if currDir.IsInfested {
				infestedFilesCount += len(currDir.Files)
				for folderIdx := range currDir.Folders {
					currDir.Folders[folderIdx].IsInfested = true
				}
			}
			for folderIdx := range currDir.Folders {
				dirsQueue = append(dirsQueue, currDir.Folders[folderIdx])
			}
		}
		fmt.Println(infestedFilesCount)
	}
}
