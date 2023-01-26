package main

// Given rooms that each room[i] has a distinct keys in it. Each key opens a room with
// number on it.

// [[1], [2], [3], []]
// All rooms are locked except room 0. Start from here.

// Approach
// starting from rooms[0] adding all keys given there. We add to stack
// iteration over stack and unlock the room with key we have and add to stack found in the room. Every time we visit a room we mark it as visited. At the end we check if all rooms are opened
// TC: O(N) - number of rooms
// SC: O(N) - N for stack and N for seen/visited

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() int {
	if s.Empty() {
		return -1
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return top
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func canVisitAllRooms(rooms [][]int) bool {
	seen := make(map[int]bool)
	stack := Stack{}

	seen[0] = true
	for i := 0; i < len(rooms[0]); i++ {
		seen[rooms[0][i]] = true
		stack.Push(rooms[0][i])
	}

	for !stack.Empty() {
		key := stack.Pop()

		for i := 0; i < len(rooms[key]); i++ {
			if _, ok := seen[rooms[key][i]]; !ok {
				seen[rooms[key][i]] = true
				stack.Push(rooms[key][i])
			}
		}
	}

	for i := 0; i < len(rooms); i++ {
		if _, ok := seen[i]; !ok {
			return false
		}
	}

	return true
}