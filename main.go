package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

func main() {

	// badly formatted code #gofmt

	// declare variable here
	var a = 1
		// declare another variable here
			var b = 16
			//calculate the sum
		sum := a + b
		color.Red("Hello red planet")
		fmt.Println("----------------------------------------------")
	// end badly formatted code

	// variables decalared but not used add noise
	res1, _ := strconv.ParseBool("this will fail")
	res2, err := strconv.ParseBool("this will fail")
	res3, _ := 			strconv.ParseBool("this will fail")
	res4, _ = strconv.ParseBool("this will fail")
}