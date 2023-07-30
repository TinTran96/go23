package main

import (
	"fmt"
	"strconv"
)

func checkIntExist(arr []int, key int) bool {
	for i := 0; i < len(arr); i++ {
		if key == arr[i] {
			return true
		}
	}
	return false
}

func numDifferentIntegers(word string) []int {
	var arrInt []int
	var num string = ""
	for i := 0; i < len(word); i++ {

		_, err := strconv.Atoi(string(word[i]))
		if num != "" && err != nil {
			s, err1 := strconv.Atoi(string(num))
			if checkIntExist(arrInt, s) {
				continue
			} else {
				if err1 == nil {
					arrInt = append(arrInt, s)
				}
			}
			num = ""
		} else {
			if err == nil {
				num = num + string(word[i])
			} else {
				continue
			}
		}
	}
	return arrInt

}

func printResult(arrInt []int) {
	fmt.Printf("%d(", len(arrInt))
	sizeArrInt := len(arrInt)
	for i := 0; i < sizeArrInt; i++ {
		fmt.Printf("%d", arrInt[i])
		if i != (sizeArrInt - 1) {
			fmt.Printf(", ")
		}
	}
	fmt.Printf(")")
}

func main() {
	word := "a123bc34d8ef34"

	printResult(numDifferentIntegers(word))
}
