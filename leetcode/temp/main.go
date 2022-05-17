package main

import (
	"fmt"
	"sort"
)

// Solution 1
//
// 1. Sum of A[i], where A[i] is the length of account - when creating graph
// 2. Part with DFS: in worst case when all emails are unique it will be sum of A[i]
// 2.1 With sorting it is O(som of A[i] Log A[i])
//
// TC: O(Sum of A[i] log A[i])
// SC: O(Sum of A[i])
func accountsMerge1(accounts [][]string) [][]string {
	graph := make(map[string][]string)
	emailToNames := make(map[string]string)

	for _, account := range accounts {
		name := account[0]
		firstEmail := account[1]

		for _, email := range account[1:] {
			graph[firstEmail] = append(graph[firstEmail], email)
			graph[email] = append(graph[email], firstEmail)
			emailToNames[email] = name
		}
	}

	visited := make(map[string]bool)

	var dfs func(email string, current *[]string)
	dfs = func(email string, current *[]string) {
		visited[email] = true
		*current = append(*current, email)

		for _, to := range graph[email] {
			if !visited[to] {
				dfs(to, current)
			}
		}
	}

	var result [][]string
	for email, _ := range graph {
		var current []string
		if visited[email] {
			continue
		}

		dfs(email, &current)
		sort.Strings(current)
		current = append([]string{emailToNames[email]}, current...)

		result = append(result, current)
	}

	return result
}

// Solution 2
// Using Disjoint Set Union

type DisjointSet struct {
	parent 	[]int
}

func (s *DisjointSet) Find(x int) int {
	if s.parent[x] != x {
		return s.Find(s.parent[x])
	}

	return s.parent[x]
}

// Union joins two sets
func (s *DisjointSet) Union(x, y int) {
	xSet := s.Find(x)
	ySet := s.Find(y)

	if xSet == ySet {
		return
	}

	s.parent[xSet] = ySet
}

func Construct(n int) *DisjointSet {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DisjointSet{parent}
}

func accountsMerge2(accounts [][]string) [][]string {
	disjointSet := Construct(10)

	emailToName := make(map[string]string)
	emailToId := make(map[string]int)

	index := 0
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			emailToName[email] = name
			if _, ok := emailToId[email]; !ok {
				emailToId[email] = index
				index++
			}

			disjointSet.Union(emailToId[account[0]], emailToId[email])
		}
	}

	components := make(map[int][]string)
	for email, _ := range emailToName {
		idx := disjointSet.Find(emailToId[email])
		components[idx] = append(components[idx], email)
	}

	var result [][]string
	for _, emails := range components {
		sort.Strings(emails)

		emails = append([]string{emailToName[emails[0]]}, emails...)
		result = append(result, emails)
	}

	fmt.Println(result)
	return [][]string{}
}

func main() {
	accounts := [][]string{
		{"John","johnsmith@mail.com", "jonny@email"},
		{"John","johnsmith@mail.com","john00@mail.com"},
		{"Mary","mary@mail.com"},
		{"John","johnnybravo@mail.com", "john00@mail.com"},
	}

	fmt.Println(accountsMerge2(accounts))
}