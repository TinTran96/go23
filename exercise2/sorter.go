package main

import (
	"flag"
	"fmt"

	ultil "github.com/TinTran96/go23/exercise2/ultility"
)

func main() {
	mixFlags := flag.String("mix", "foo", "a mix string with number as string")
	intFlags := flag.String("int", "foo", "a number")
	strFlags := flag.String("string", "foo", "a string")

	flag.Parse()

	var input []string
	var isValid bool = false
	if *mixFlags != "foo" {
		input = append(flag.Args(), *mixFlags)
		ultil.Sort("mix", input)
	} else {
		if *intFlags != "foo" {
			input = append(flag.Args(), *intFlags)
			isValid = ultil.ValidateInput("int", input)
			if isValid {
				ultil.Sort("int", input)
			} else {
				fmt.Println("Invalid Input")
			}
		} else {
			if *strFlags != "foo" {
				input = append(flag.Args(), *strFlags)
				isValid = ultil.ValidateInput("string", input)
				if isValid {
					ultil.Sort("string", input)
				} else {
					fmt.Println("Invalid Input")
				}
			} else {
				fmt.Println("Invalid Flag")
			}
		}
	}

}
