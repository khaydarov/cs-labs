package main

import (
	"fmt"
	"testing"
)

var whitelist = []string{
	"test1",
	"test2",
	"test3",
	"test4",
}

var testCases = []struct {
	input  string
	expect bool
}{
	//{
	//	input:  "",
	//	expect: false,
	//},
	{
		input:  "test1@gmail.com",
		expect: true,
	},
	//{
	//	input:  "test5@gmail.com",
	//	expect: true,
	//},
	//{
	//	input:  "test6@gmail.com",
	//	expect: false,
	//},
}

func TestEmailWhitelist(t *testing.T) {
	tree := New(whitelist, &Sha256Hasher{})

	fmt.Println(tree.Verify("test1"))
	//for _, testCase := range testCases {
	//	t.Run(testCase.input, func(t *testing.T) {
	//		ok := tree.Verify(testCase.input)
	//
	//		if ok != testCase.expect {
	//			t.Error(fmt.Sprintf("expected %t, got %t", testCase.expect, ok))
	//		}
	//	})
	//}

}
