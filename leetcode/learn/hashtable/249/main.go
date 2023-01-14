package main

import (
	"fmt"
	"strconv"
)

func groupStrings(strings []string) [][]string {
	groups := make(map[string][]string)

	for _, v := range strings {
		key := shiftingSequence(v)
		fmt.Println(key, v)
		groups[key] = append(groups[key], v)
	}

	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func shiftingSequence(target string) string {
	var key string
	key += strconv.Itoa(len(target))

	for i := 1; i < len(target); i++ {
		previous := int(target[i-1] - 'a')
		current := int(target[i] - 'a')

		var distance int

		// ba; if current < prev :-> prev - current
		// az; if current > prev :-> 26 + current - previous
		if current < previous {
			distance = previous - current
		} else {
			distance = 26 - current + previous
		}

		key += strconv.Itoa(distance)
	}

	return key
}
