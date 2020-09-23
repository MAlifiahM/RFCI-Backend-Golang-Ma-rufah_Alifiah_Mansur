package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

func main() {
	collection := []int{3, 5, 2, -4, 8, 11}
	fmt.Println(foo(7, collection))
}

func foo(input int, numbers []int) (int, [][]string) {
	var result []int
	var result2 []int
	var nextIndex int
	var finalResultR1 []string
	var finalResultR2 []string
	var finalResult [][]string

	for idx, number := range numbers {
		nextIndex = idx + 1
		idxnext3 := idx - 2

		if nextIndex >= len(numbers) {
			result2 = append(result2, numbers[idxnext3])
			result2 = append(result2, number)
			r2, _ := json.Marshal(result2)
			resultR2 := strings.Trim(string(r2), "[]")
			finalResultR2 = append(finalResultR2, resultR2)
			break;
		}
		nextNumber := numbers[nextIndex]
		if number + nextNumber == input {
			result = append(result, nextNumber)
			result = append(result, number)
			r1, _ := json.Marshal(result)
			resultR1 := strings.Trim(string(r1), "[]")
			finalResultR1 = append(finalResultR1, resultR1)
		}
	}
	finalResult = append(finalResult, finalResultR2)
	finalResult = append(finalResult, finalResultR1)
	return input, finalResult
}
