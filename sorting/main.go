package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"strconv"
)

var toBeSorted [7]int = [7]int{4, 9, 7, 5, 8, 9, 3}

func bubbleSort(input [7]int) {
	swapping := 0
	number := 0
	var swap []int
	var result []string
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				swapping++
				swap = append(swap, input[j])
				swap = append(swap, input[j+1])
				s, _ := json.Marshal(swap)
				resultSwaps := strings.Trim(string(s), "[]")
				result = append(result, resultSwaps)
				r, _ := json.Marshal(input)

				resultInput := strings.Trim(string(r), "[]")
				finalResultInput := strings.Replace(resultInput, ",", " ", -1)
				number++
				numbers := strconv.FormatInt(int64(number), 10)
				finalNumbers := numbers+"."
				fmt.Println(finalNumbers, result, "->" ,finalResultInput)
				
				swap = swap[:0]
				result = result[:0]
			}
		}
	}
	var stringSwap = "Jumlah swap : "
	var concratenated = fmt.Sprint(stringSwap, swapping)
	fmt.Println(concratenated)
}

func main() {
	bubbleSort(toBeSorted)
}