package main

import (
	"fmt"
	"errors"
	"unicode/utf8"
)


func main(){
	input := "Some days are great"
	reversed, revErr := Reverse(input)
	doubleReversed, doubleRevErr := Reverse(reversed)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("Returned string: %q\n, err: %v\n", reversed, revErr)
	fmt.Printf("Returned string: %q\n, err: %v\n", doubleReversed, doubleRevErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid for UTF8")
	}
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	fmt.Printf("input: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1{
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
