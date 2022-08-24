package main_test

import (
	"sort"
	"testing"
)

func TestS(t *testing.T) {
	a := "abvadfasdf"
	t.Log(a)

	sList := []*Student{
		{ID: 5, Class: 2, Name: "a"},
		{ID: 1, Class: 1, Name: "b"},
		{ID: 3, Class: 3, Name: "a"},
		{ID: 7, Class: 1, Name: "a"},
		{ID: 2, Class: 2, Name: "a"},
		{ID: 5, Class: 1, Name: "a"},
		{ID: 4, Class: 3, Name: "a"},
	}
	sort.Slice(sList, func(i, j int) bool {
		if sList[i].Class == sList[j].Class {
			return sList[i].ID < sList[j].ID
		}
		return sList[i].Class < sList[j].Class
	})
	for _, item := range sList {
		t.Logf("%#v", item)
	}
}

type Student struct {
	ID    int64
	Class int32
	Name  string
}
