package main

import (
	"reflect"
	"testing"
)

func TestSetUnion(t *testing.T) {
	s := Set{}
	s.Add(1)
	s.Add(2)

	s1 := Set{}
	s1.Add(3)
	s1.Add(4)

	r := s.Union(s1)

	expect := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(r.Values(), expect) {
		t.Errorf("Union result is not correct")
	}
}

func TestSet(t *testing.T) {
	s := Set{}
	s.Add(1)
	s.Add(2)
	s.Add(1)
	
	expect := []int{1, 2}
	if !reflect.DeepEqual(s.Values(), expect) {
		t.Errorf("Union result is not correct")
	}
}
