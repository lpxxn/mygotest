package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	p1 := NewListValue(231)
	t.Log(p1.String())
	t.Log(p1.Values())
	p2 := NewListValue(111)

	rev := add(p1, p2)
	t.Log(rev)
	t.Log(rev.Values())

	rev = add(p1, NewListValue(789))
	t.Log(rev)
	t.Log(rev.Values())
}

func add(a, b *ListValue) *ListValue {
	if a == nil && b == nil {
		return nil
	}
	newList := &ListValue{}
	nextList := newList
	carry := 0

	for a != nil || b != nil {
		v1, v2 := 0, 0
		if a != nil {
			v1 = a.Val
			a = a.Next
		}
		if b != nil {
			v2 = b.Val
			b = b.Next
		}
		v3 := v1 + v2 + carry
		nextList.Next = &ListValue{
			Val: v3 % 10,
		}
		carry = v3 / 10
		nextList = nextList.Next
	}

	if carry > 0 {
		nextList.Next = &ListValue{
			Val: carry,
		}
	}

	return newList.Next
}

func NewListValue(v int) *ListValue {
	rev := &ListValue{Val: 0}
	next := rev
	for v > 0 {
		next.Next = &ListValue{
			Val: v % 10,
		}
		next = next.Next
		v /= 10
	}
	return rev.Next
}

type ListValue struct {
	Val  int
	Next *ListValue
}

func (l *ListValue) Values() int {
	if l.Next != nil {
		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", l.Next.Values(), l.Val))
		return v
	}
	return l.Val
}

func (l *ListValue) String() string {
	rev := strconv.Itoa(l.Val)
	if l.Next != nil {
		rev += " -> " + l.Next.String()
	}
	return rev
}
