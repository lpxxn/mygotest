package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {

	a := NewANum(789)
	t.Log(a.String())
	t.Log(a.Value())

	b := NewANum(654)
	t.Log(b.String())
	t.Log(b.Value())

	rev := AddANum(a, b)
	t.Log(rev.String())
	t.Log(rev.Value())
	if rev.Value() != a.Value() + b.Value() {
		t.Error("err")
	}

	/*
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

	*/
}

func AddANum(a, b *ANum) *ANum {
	current := &ANum{}
	rev := current
	carry := 0
	for a != nil || b != nil {
		sum := carry
		if a != nil {
			sum += a.value
			a = a.next
		}
		if b != nil {
			sum += b.value
			b = b.next
		}
		current.next = &ANum{value: sum % 10}
		current = current.next
		carry = sum / 10
	}
	if carry > 0 {
		current.next = &ANum{value: carry}
	}
	return rev.next
}

type ANum struct {
	value int
	next  *ANum
}

func (a ANum) String() string {
	if a.next != nil {
		return fmt.Sprintf("%d -> %s", a.value, a.next.String())
	}
	return fmt.Sprint(a.value)
}

func (a ANum) Value() int {
	if a.next != nil {
		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", a.next.Value(), a.value))
		return v
	}
	return a.value
}

func NewANum(v int) *ANum {
	current := &ANum{}
	rev := current
	for v > 0 {
		current.next = &ANum{value: v % 10}
		current = current.next
		v /= 10
	}
	return rev.next
}

/*
func add(a, b *ListValue) *ListValue {
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
		nextList.Next = &ListValue{Val: v3 % 10}
		carry = v3 / 10
		nextList = nextList.Next
	}
	if carry > 0 {
		nextList.Next = &ListValue{Val: carry}
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

*/
