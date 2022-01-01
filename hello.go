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

	var interface1 interface{}
	fmt.Println(interface1) // <nil>
	interface1 = 2
	fmt.Println(interface1) // 2
	interface1 = [3]int{1, 2, 3}
	fmt.Println(interface1) // [1,2,3]
	interface1 = "aaa"
	fmt.Println(interface1) // aaa
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

	var byteA []byte = []byte{72, 73}
	fmt.Println(byteA)         // [72,73]
	fmt.Println(string(byteA)) // HI

	byteB := []byte("HI")
	fmt.Println(byteB) // [72,73]

	var array1 [3]int
	fmt.Println(array1)        // [0,0,0]
	fmt.Printf("%T\n", array1) // [3]int

	var array2 [3]string = [3]string{"A", "B"}
	fmt.Println(array2) // [A B ]

	array3 := [3]int{}
	fmt.Println(array3) // [0 0 0]

	array4 := [...]string{"A", "B", "C", "D"}
	fmt.Println(array4)        // [A B C D]
	fmt.Printf("%T\n", array4) // [4]string
	array5 := [...]string{"A"}
	fmt.Printf("%T\n", array5) // [1]string

	fmt.Println(array4[2]) // C
	array4[3] = "E"
	fmt.Println(array4)      // [A B C E]
	fmt.Println(len(array4)) // 4
}
