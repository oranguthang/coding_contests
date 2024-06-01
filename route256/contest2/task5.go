package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func isEmpty(obj interface{}) bool {

	if obj == nil {
		return true
	}

	objValue := reflect.ValueOf(obj)

	switch objValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		ref := objValue.Elem().Interface()
		return isEmpty(ref)
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(obj, zero.Interface())
	}
}

func checkEmptySubSlices(data *interface{}) bool {
	f := true
	switch obj := (*data).(type) {
	case []interface{}:
		if len(obj) == 0 {
			return true
		}
		for _, val := range obj {
			f = f && checkEmptySubSlices(&val)
		}
		return true
	}
	return false
}

// {"a":"f","b":{"c":{"d":["string"],"e":["ababa"]}}}
func deleteEmptyMapsAndSlices(data *interface{}) {
	switch obj := (*data).(type) {
	case map[string]interface{}:
		//if len(obj) == 0 {
		//	fmt.Println("empty map")
		//}
		for k, val := range obj {
			deleteEmptyMapsAndSlices(&val)
			if isEmpty(val) {
				delete(obj, k)
			}
		}
	case []interface{}:
		//if len(obj) == 0 {
		//	fmt.Println("empty slice")
		//}
		var lst []interface{}
		for j := len(obj) - 1; j >= 0; j-- {
			//for k, val := range obj {
			deleteEmptyMapsAndSlices(&obj[j])
			if isEmpty(obj[j]) || checkEmptySubSlices(&obj[j]) {
				//obj = append(obj[:j], obj[j+1:]...)
				lst = append(lst, obj[j])
			}
		}
		//if isEmpty(obj) {
		//	print("w")
		//	obj = nil
		//}
		obj = lst
		//case string, float64, bool, int:
		//default:
		//	*data = nil
		//}
	}
}

func main() {
	f, err := os.Open("D:\\Downloads\\5_1")
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

	var testsCount, testLinesCount int
	fmt.Fscan(in, &testsCount)

	var jsonList []interface{}
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

		var jsonObj interface{}
		json.Unmarshal([]byte(jsonString), &jsonObj)
		deleteEmptyMapsAndSlices(&jsonObj)

		jsonList = append(jsonList, jsonObj)
		//jsonList = append(jsonList, result)
		//results = append(results, string(result))
	}
	result, _ := json.Marshal(jsonList)
	fmt.Println(string(result))
}
