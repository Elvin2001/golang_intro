package main

import "fmt"

const name string = "Elvin"
const age int = 22
const salary int = 20000

func main() {
	//fmt.Println("Hello, Elvin!")

	//var a int8 = -1            // -128 <> 127, 1 byte, 8 bit
	//var b int16 = 32767        // -32768 <> 32767, 2 byte, 16 bit
	//var c int32 = 30           // -2bil <> 2bil, 4 byte
	//var d int64 = -192929292   // -9pent <> 9pent, 8 byte
	//var e uint8 = 23           // 0 <> 255, 1 byte
	//var f uint16 = 24303       // 0 <> 65535, 2 byte
	//var g uint32 = 12000000    // 0 <> 4 bil, 4 byte
	//var h uint64               // 0 <> 18pent
	//var i byte                 // synonym uint8
	//var j rune                 // synonym int32
	//var k int                  //
	//var m uint                 //
	//var a1 float32 = 1.8       // // 1.4 * 10^ - 45 <> 3.4 * 10^35, 4 byte
	//var b1 float64 = 13.248838 // 4.8 * 10^ - 320 <> 1.8 * 1^305, 8 byte
	//var c1 complex64 = 1 + 2i
	//var d1 complex128 = 4 + 90i
	//
	//var b23 bool = true
	//var b13 bool = false
	//
	//var name string = "Elvin"

	//var (
	//	name = "Elvin"
	//	age  = 22
	//)

	//var name1, age1 = "Arthur", 23
	//
	//fmt.Println(fmt.Sprintf("My name is %s, and age: %d", name1, age1))

	getName()

}

func getName() {
	fmt.Println("My name: " + name)
}

func test() (string, string) {
	a := "hello"
	b := "world"
	return a, b
}

func test1() string {
	a := "hee"
	return a
}
