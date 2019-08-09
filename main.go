package main

import (
	"fmt"
	//	"github.com/fatih/color"
	//	"os"
	//	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	////// #gofmt begin
	//var a = 1
	//	// declare another variable here
	//		var b = 16
	//		//calculate the sum
	//	sum := a + b
	//		fmt.Println(sum)
	//		color.Red("Hello red planet")
	//	fmt.Println("----------------------------------------------")
	////#gofmt end
	//--------------------
	////#typecheck begin
	//// eval1 declared but not used
	//eval1, _ := strconv.ParseBool("false")
	//
	//// err declared but not used
	//eval2, err := strconv.ParseBool("true")
	//fmt.Println(eval2)
	////#typecheck end
	//
	////gosec start
	//username := "admin"
	//var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"
	//fmt.Println("Doing something with: ", username, password)
	////gosec end
}

//func testsql() {
////gosec start
//	db, err := sql.Open("sqlite3", ":memory:")
//	if err != nil {
//		panic(err)
//	}
//	q := fmt.Sprintf("SELECT * FROM foo where name = '%s'", os.Args[1])
//	rows, err := db.Query(q)
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
////gosec end
//}
