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
		i3 int
		s3 string
	)
	fmt.Println(i2, s2) // 200 Hello Go2
	fmt.Println(i3, s3) // 0 ※初期化せずに変数宣言するとそれぞれの型のデフォルト値が出力。string は空文字

}
