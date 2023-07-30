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

func numDifferentIntegers(word string) {
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

	for i := 0; i < len(arrInt); i++ {
		fmt.Printf("%d \t", arrInt[i])
	}
}

func main() {
	word := "A1b01c001"

	numDifferentIntegers(word)
}
