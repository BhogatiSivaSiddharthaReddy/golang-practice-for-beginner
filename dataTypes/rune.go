package main

import (
	"fmt"
	"unicode/utf8"
)

func runes() {
	fmt.Println("=== Rune (Unicode) Demo ===")

	basicRune()
	runeVsByte()
	stringIteration()
	unicodeLength()
	runeConversion()
	runeArithmetic()
}

func basicRune() {
	fmt.Println("\n-- Basic Rune --")

	var r rune = 'A'
	r1 := rune(76)
	var r2 rune = '你'

	fmt.Println("Rune value (A):", r)
	fmt.Printf("Rune as char: %c\n", r)

	fmt.Println(r1)

	fmt.Println("Rune value (你):", r2)
	fmt.Printf("Rune as char: %c\n", r2)
}

func runeVsByte() {
	fmt.Println("\n-- Rune vs Byte --")

	s := "A你"

	fmt.Println("String:", s)
	fmt.Println("Bytes length:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("byte[%d]=%d\n", i, s[i])
	}
}

func stringIteration() {
	fmt.Println("\n-- String Iteration (rune safe) --")

	s := "Go你"

	for i, r := range s {
		fmt.Printf("index=%d rune=%c value=%d\n", i, r, r)
	}
}

func unicodeLength() {
	fmt.Println("\n-- Length Difference --")

	s := "你好"

	fmt.Println("Byte length:", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
}

func runeConversion() {
	fmt.Println("\n-- Conversion --")

	s := "hello"

	runes := []rune(s)
	runes[0] = 'H'

	newStr := string(runes)

	fmt.Println("Modified string:", newStr)
}

func runeArithmetic() {
	fmt.Println("\n-- Rune Arithmetic --")

	r := 'A'

	fmt.Printf("%c + 1 = %c\n", r, r+1)
	fmt.Printf("%c + 32 = %c\n", r, r+32) // lowercase
}

/*
UTF-8 ENCODING – HOW A RUNE (UNICODE) BECOMES BYTES

Key idea:
A rune (Unicode code point) is converted into 1–4 bytes using UTF-8 rules.

Example:
'₹' → Unicode: U+20B9 (decimal 8377)
→ encoded into bytes: [226 130 185]

--------------------------------------------------
STEP 1: DECIDE NUMBER OF BYTES

Based on Unicode value:

U+0000 – U+007F     → 1 byte  (ASCII)
U+0080 – U+07FF     → 2 bytes
U+0800 – U+FFFF     → 3 bytes
U+10000 – U+10FFFF  → 4 bytes

Example:
₹ = U+20B9 → falls in U+0800–U+FFFF → needs 3 bytes

--------------------------------------------------
STEP 2: UTF-8 BYTE PATTERN

1 byte:
0xxxxxxx

2 bytes:
110xxxxx 10xxxxxx

3 bytes:
1110xxxx 10xxxxxx 10xxxxxx

4 bytes:
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

--------------------------------------------------
STEP 3: CONVERT RUNE TO BINARY

₹ = 8377

Binary:
0010000010111001

--------------------------------------------------
STEP 4: SPLIT BITS INTO GROUPS

For 3-byte encoding:

xxxx      xxxxxx      xxxxxx
0010      000010      111001

--------------------------------------------------
STEP 5: FILL INTO PATTERN

1110xxxx → 1110 0010 → 226
10xxxxxx → 10 000010 → 130
10xxxxxx → 10 111001 → 185

--------------------------------------------------
FINAL RESULT:

8377 → [226 130 185] → "₹"

--------------------------------------------------
SUMMARY:

Unicode number → split bits → add UTF-8 prefixes → bytes

--------------------------------------------------
MENTAL MODEL:

Encoding = break a number into pieces + add headers
Decoding = remove headers + join pieces

*/
