package ultility

import (
	"fmt"
	"sort"
)

func SortString(input []string) []string {
	sort.Strings(input)
	return input
}

func Sort(input []string) {
	var s []string

	s = SortString(input)
	fmt.Print("Output:")
	for i := 0; i < len(s); i++ {
		fmt.Print(" ", s[i])
	}

}
