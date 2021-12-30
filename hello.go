package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	if t {
		fmt.Println("if文を評価")
	}
}
