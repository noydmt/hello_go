package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		i, i2 int = 100, 200
		i3    int
		s, s2 string = "Hello Go", "Hello World"
		s3    string
		t, f  bool = true, false
	)

	fmt.Println(time.Now())
	fmt.Println(i, s)   // 100 Hello Go
	fmt.Println(i2, s2) // 200 Hello World
	fmt.Println(i3, s3) // 0 ※初期化せずに変数宣言するとそれぞれの型のデフォルト値が出力。string は空文字
	fmt.Println(t, f)

	if t {
		fmt.Println("if文を評価")
	}
}
