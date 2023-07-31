package ex4

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

func ReorderName(input []string) string {
	size := len(input)
	countryCode := input[size-1]

	switch countryCode {
	case "US":
		return reorderNameUS(input, size)
	case "VN":
		return reorderNameVN(input, size)
	default:
		return "Country Code " + countryCode + " Invalid"
	}

}
