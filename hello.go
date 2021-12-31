package main

import (
	"fmt"
	"time"
)

// i5 := 500 main関数の外で型推論で宣言するとエラー
var i5 int = 500 // main関数の外でも型を明示的にすれば変数宣言できる

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

	i3 = 300
	s3 = "Hello Go World"
	fmt.Println(i3, s3) // 300 Hello Go World

	i4 := 400
	fmt.Println(i4)

	fmt.Println(i5)

	outer()
}

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)

	var i int = 300
	fmt.Printf("%T\n", i)        // int %T が型を表示するフォーマット
	fmt.Printf("%T\n", int64(i)) // int64
	fmt.Printf("%T\n", int32(i)) // int32
	fmt.Printf("%T\n", int16(i)) // int16
	fmt.Printf("%T\n", int8(i))  // int8

	var str_hello string = "Hello"
	fmt.Println(str_hello[0])         // 72
	fmt.Println(string(str_hello[0])) // H

	byteA := []byte{72, 73}
	fmt.Println(byteA)         // [72,73]
	fmt.Println(string(byteA)) // HI
}
