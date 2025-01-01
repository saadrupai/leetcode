package main

import (
	"fmt"
	"strings"
)

func encode(arr []string) {
	encodedStr := ""
	for _, str := range arr {
		encodedStr = encodedStr + "saad" + str
	}
	fmt.Println("encoded_string : ", encodedStr)

	decode(encodedStr)
}

func decode(s string) {
	fmt.Println("decoded_list : ", strings.Split(s, "saad"))

}

func main() {
	arr := []string{"hello", "world", "foo|bar", ""}
	fmt.Println(arr)
	encode(arr)
}

// You are given a list of strings. Your task is to implement:

// encode(strs): Encodes the list of strings into a single string.
// decode(s): Decodes the encoded string back into the original list of strings.
// You need to ensure that:

// The encoding method handles edge cases, such as empty strings or special characters (e.g., |, #, \).
// The decoding method can correctly reconstruct the original list.
