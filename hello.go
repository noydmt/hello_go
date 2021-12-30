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

	var (
		i2 int    = 200
		s2 string = "Hello Go2"
	)
	fmt.Println(i2, s2)
}
