package go_algo

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	s := "12345654321"
	ret := Palindrome(s)
	fmt.Println(ret)
}

func TestMoveZeros(t *testing.T) {
	s := []int{1, 2, 3, 0, 0, 4, 5, 6, 0, 7, 8, 9}
	ret := MoveZeros(s)
	fmt.Println(ret)
}


