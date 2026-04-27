package main

import "fmt"

func integers() {
	fmt.Println("****learning about the integer data type****")
	integer8()
	integer16()
	integer32()
	integer64()
	basicUnsigned()
}

func integer8() {
	//min value -128, max value 127, includes 0
	// 2**8(2power8) values
	// bit value start with 0 means positive, if they start with 1 it means negative.
	fmt.Println("****int8 ****")
	var b int8 = -128
	a := int8(127)
	a++
	fmt.Println(a, b) // -128
}

func integer16() {
	fmt.Println("****int16****")
	a := int16(234)
	var b int16 = -12345
	var c int16
	fmt.Println(a, b, c)
}

func integer32() {
	fmt.Println("****int32****")
	a := int32(123456)
	var b int32 = -1929347
	var c int32
	fmt.Println(a, b, c)
}

func integer64() {
	fmt.Println("***int64***")
	a := int64(123456689093737)
	var b int64 = -12938383838838833
	var c int64
	fmt.Println(a, b, c)
}

func basicUnsigned() {
	fmt.Println("\n-- Basic Unsigned Types --")

	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 100

	fmt.Println("uint8 :", u8)
	fmt.Println("uint16:", u16)
	fmt.Println("uint32:", u32)
	fmt.Println("uint64:", u64)
	fmt.Println("uint  :", u)
}
