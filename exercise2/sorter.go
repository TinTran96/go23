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

	if *mixFlags != "foo" {
		input = append(flag.Args(), *mixFlags)
	} else {
		if *intFlags != "foo" {
			input = append(flag.Args(), *intFlags)
		} else {
			if *strFlags != "foo" {
				input = append(flag.Args(), *strFlags)
			} else {
				fmt.Println("Invalid Flag")
			}
		}
	}

	ultil.Sort(input)

}
