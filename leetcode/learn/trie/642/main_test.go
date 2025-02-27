package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	autocomplete := Constructor(
		[]string{"i love you", "island", "ironman", "i love leetcode"},
		[]int{5, 3, 2, 2},
	)

	r := autocomplete.Input('i')
	if !reflect.DeepEqual(r, []string{"i love you", "island", "i love leetcode"}) {
		t.Errorf("Wrong answer, test 1")
	}

	r = autocomplete.Input(' ')
	if !reflect.DeepEqual(r, []string{}) {
		t.Errorf("Wrong answer, test 2")
	}
	
	r = autocomplete.Input('a')
	if !reflect.DeepEqual(r, []string{})
	fmt.Println("searching i a", r)
}
