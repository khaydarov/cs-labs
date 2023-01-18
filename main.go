package main

import (
	"fmt"
	"math"
)

func groupStrings(strings []string) [][]string {
	visited := make(map[string]bool)

	var groups [][]string

	for i, v := range strings {
		if _, ok := visited[v]; ok {
			continue
		}

		var group []string
		group = append(group, v)
		visited[v] = true

		for j := i + 1; j < len(strings); j++ {
			if isSameShiftingSequence(v, strings[j]) {
				group = append(group, strings[i])
				visited[strings[i]] = true
			}
		}
	}

	return groups
}

func isSameShiftingSequence(base, target string) bool {
	if len(base) != len(target) {
		return false
	}

	if len(base) == 1 && len(target) == 1 {
		return true
	}

	diff := int(target[0]) - int(base[0])

	for i := 1; i < len(base); i++ {
		baseCharPosition := int(base[i] - 'a')
		targetCharPosition := int(target[i] - 'a')

		fmt.Println(baseCharPosition, targetCharPosition, (baseCharPosition+diff)%26)

		if diff >= 0 && (baseCharPosition+diff)%26 != targetCharPosition {
			return false
		} else {
			if baseCharPosition > abs(diff) && (baseCharPosition+diff)%26 != targetCharPosition {
				return false
			} else if (26+baseCharPosition+diff)%26 != targetCharPosition {
				return false
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}

func main() {
	//r := isSameShiftingSequence("abc", "bcd")
	//fmt.Println(r)

	r := make([]uint32, math.MaxUint32)
	r[0] = 1

	fmt.Println(r[0])
}
