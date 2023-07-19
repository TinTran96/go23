package main

import (
	"os"
)

func reorderNameVN(inputName []string, size int) string {
	var firstName string
	var lastName string

	var result string

	firstName = inputName[0]
	lastName = inputName[1]

	if size > 3 {
		var middleName string
		i := 2
		for i < size-1 {
			middleName += inputName[i]
			i = i + 1
			if i < size-1 {
				middleName += " "
			}
		}
		result = lastName + " " + middleName + " " + firstName
	} else {
		result = lastName + " " + firstName
	}

	return result
}

func reorderNameUS(inputName []string, size int) string {
	var firstName string
	var lastName string

	var result string

	firstName = inputName[0]
	lastName = inputName[1]

	if size > 3 {
		var middleName string
		i := 2
		for i < size-1 {
			middleName += inputName[i]
			i = i + 1
			if i < size-1 {
				middleName += " "
			}
		}
		result = firstName + " " + middleName + " " + lastName
	} else {
		result = firstName + " " + lastName
	}

	return result
}

func main() {

	args := os.Args[1:]

	size := len(args)
	countryCode := args[size-1]

	switch countryCode {
	case "US":
		println("Output: " + reorderNameUS(args, size))
	case "VN":
		println("Output: " + reorderNameVN(args, size))
	default:
		println("Country Code " + countryCode + " Invalid")
	}
}
