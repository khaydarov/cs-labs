package main

var testCases = []struct {
	strings []string
	expect  [][]string
}{
	{
		[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
		[][]string{
			{"acef"},
			{"a", "z"},
			{"abc", "bcd", "xyz"},
			{"az", "ba"},
		},
	},
}
