package main

import "fmt"

func bytes() {
	fmt.Println("-----learning about bytes -------------")
	numTochar()
	seqOfBytes()
	multiByteChar()
	bitOps()
	decoder()
}

func numTochar() {
	var b byte = 78
	c := byte(0b01010101)
	var d byte
	d = 77

	fmt.Println(b, c, d)
	fmt.Println(string(c))
	fmt.Printf("%c, %c \n", b, d)

}

func seqOfBytes() {
	s := "Go"

	fmt.Println([]byte(s)) // [71 111]

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // accessing raw bytes
	}
}

func multiByteChar() {
	s := "👉"
	c := "a"

	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println([]rune(s))
	//fmt.Println(byte(c)) doesn't work
	fmt.Println([]byte(c)[0])
	b := c[0]
	fmt.Println(b)

}

func bitOps() {
	var a byte = 5 // 00000101
	var b byte = 3 // 00000011

	fmt.Println(a & b)  // 1 (AND)
	fmt.Println(a | b)  // 7 (OR)
	fmt.Println(a ^ b)  // 6 (XOR)
	fmt.Println(a << 1) // 10 (shift left)
}

func decoder() {
	data := []byte{72, 105, 32, 226, 130, 185}

	fmt.Println(string(data)) // "Hi ₹"

	for i, b := range data {
		fmt.Printf("Index %d: %08b (%d)\n", i, b, b)
	}
}

/*
UTF-8 DECODING – HOW BYTES BECOME A UNICODE CODE POINT (RUNE)

Key idea:
A string in Go is a sequence of bytes (UTF-8 encoded).
A rune is a Unicode code point (actual character).

Example:
"₹" → bytes: [226 130 185] → rune: 8377

--------------------------------------------------
STEP 1: FIRST BYTE TELLS THE LENGTH

UTF-8 uses prefix bits in the first byte as a signal:

0xxxxxxx        → 1 byte (ASCII)
110xxxxx        → 2 bytes
1110xxxx        → 3 bytes
11110xxx        → 4 bytes

Example:
226 → 11100010 → starts with 1110 → means 3-byte character

--------------------------------------------------
STEP 2: CONTINUATION BYTES

Remaining bytes must start with:
10xxxxxx

Example:
130 → 10000010
185 → 10111001

These are NOT new characters — they belong to the first byte.

--------------------------------------------------
STEP 3: REMOVE PREFIX (HEADER) BITS

We remove the structural bits:

First byte:
11100010 → remove "1110" → 0010

Second byte:
10000010 → remove "10" → 000010

Third byte:
10111001 → remove "10" → 111001

--------------------------------------------------
STEP 4: JOIN THE DATA BITS

0010   000010   111001

Combine into one binary number:

0010000010111001

--------------------------------------------------
STEP 5: CONVERT TO DECIMAL

0010000010111001 → 8377

This is the Unicode code point:
U+20B9 → ₹

--------------------------------------------------
FINAL SUMMARY:

Bytes → remove UTF-8 prefixes → join bits → Unicode number

[226 130 185] → 8377 → "₹"

--------------------------------------------------
IMPORTANT NOTES:

- string = sequence of UTF-8 bytes
- rune   = Unicode code point (int32)

- len("₹") = 3  (bytes)
- len([]rune("₹")) = 1 (character)

- s[0] gives a byte, NOT a character
- use []rune(s) to correctly access characters

--------------------------------------------------
MENTAL MODEL:

UTF-8 splits a Unicode number into multiple bytes.
Decoding joins those bytes back into the original number.
*/
