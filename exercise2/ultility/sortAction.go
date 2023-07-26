package ultility

import (
	"fmt"
	"sort"
	"strconv"
)

func SortString(input []string) []string {
	sort.Strings(input)
	return input
}

func SortInt(input []string) []int {
	var arr []int
	size := len(input)
	for i := 0; i < size; i++ {
		item := input[i]
		x, err := strconv.Atoi(item)
		if err == nil {
			arr = append(arr, x)
		}
	}
	sort.Ints(arr)
	return arr
}

func SortNum(input []string) []float64 {
	var arr []float64
	size := len(input)
	for i := 0; i < size; i++ {
		item := input[i]
		x, err := strconv.ParseFloat(item, 10)
		if err == nil {
			arr = append(arr, x)
		}
	}
	sort.Float64s(arr)
	return arr
}

func SortMix(input []string) {
	var numArr []string
	var numArrFloat []float64
	var stringArr []string
	size := len(input)

	for i := 0; i < size; i++ {
		item := input[i]
		if _, err := strconv.ParseFloat(item, 10); err == nil {
			numArr = append(numArr, item)
		} else {
			stringArr = append(stringArr, item)
		}
	}

	fmt.Print("Output:")
	numArrFloat = SortNum(numArr)
	stringArr = SortString(stringArr)
	for i := 0; i < len(numArrFloat); i++ {
		fmt.Print(" ", numArrFloat[i])
	}
	for i := 0; i < len(stringArr); i++ {
		fmt.Print(" ", stringArr[i])
	}
}

func ValidateInput(flag string, input []string) bool {
	size := len(input)
	if flag == "int" {
		for i := 0; i < size; i++ {
			item := input[i]
			if _, err := strconv.Atoi(item); err == nil {
				continue
			} else {
				return false
			}
		}
		return true
	}
	if flag == "string" {
		for i := 0; i < size; i++ {
			item := input[i]
			if _, err := strconv.Atoi(item); err == nil {
				return false
			} else {
				if _, err := strconv.ParseFloat(item, 10); err == nil {
					return false
				}
			}
		}
		return true
	}
	return false
}

func Sort(flag string, input []string) {
	var size int = len(input)
	var strArr []string
	var intArr []int

	if flag == "string" {
		strArr = SortString(input)
		fmt.Print("Output:")
		for i := 0; i < size; i++ {
			fmt.Print(" ", strArr[i])
		}
	}
	if flag == "int" {
		intArr = SortInt(input)
		fmt.Print("Output:")
		for i := 0; i < size; i++ {
			fmt.Print(" ", intArr[i])
		}
	}
	if flag == "mix" {
		SortMix(input)
	}
}
