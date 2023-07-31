package ex4

import (
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

func NumDifferentIntegers(word string) []int {
	var arrInt []int
	var num string = ""
	for i := 0; i < len(word); i++ {

		_, err := strconv.Atoi(string(word[i]))
		if num != "" && err != nil {
			s, err1 := strconv.Atoi(string(num))
			if checkIntExist(arrInt, s) {
				num = ""
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
