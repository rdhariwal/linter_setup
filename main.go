package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

func main() {

	// #gofmt begin
	var a = 1
		// declare another variable here
			var b = 16
			//calculate the sum
		sum := a + b
			fmt.Println(sum)
			color.Red("Hello red planet")
		fmt.Println("----------------------------------------------")
	//#gofmt end
	// #gofmt begin
	var a = 1
		// declare another variable here
			var b = 16
			//calculate the sum
		sum := a + b
			fmt.Println(sum)
			color.Red("Hello red planet")
		fmt.Println("----------------------------------------------")
	//#gofmt end
	//#typecheck begin
	// eval1 declared but not used
	eval1, _ := strconv.ParseBool("false")

	// err declared but not used
	eval2, err := strconv.ParseBool("true")
	fmt.Println(eval2)
	//#typecheck end

}