package main

import (
	"fmt"
	"unicode"
)

// Count words in a camelCase string
func camelcase(s string) int32 {
	var n int32
	for _, c := range s {
		if unicode.IsUpper(c) {
			n += 1
		}
	}
	return n + 1
}

// Cipher a string using caesar algorithm
func caesarCipher(s string, k int32) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	encrypted := ""

	for _, c := range s {
		if unicode.IsLower(c) {
			encrypted += string(lower[(int32(c)-int32('a')+k)%26])
		} else if unicode.IsUpper(c) {
			encrypted += string(upper[(int32(c)-int32('A')+k)%26])
		} else {
            encrypted += string(c)
        }
	}

	return encrypted
}

func main() {
	fmt.Println(caesarCipher("Hello World zzz", 1))
}
