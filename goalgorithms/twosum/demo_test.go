package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	a := NewNumValue(123)
	fmt.Println(a.String())
	fmt.Println(a.Value())

	b := NewNumValue(877)

	rev := addTwoNum(a, b)
	fmt.Println(rev.Value())

	rev = addTwoNum(b, NewNumValue(5))
	fmt.Println(rev.Value())
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

func addTwoNum(a, b *NumValue) *NumValue {
	current := &NumValue{}
	rev := current
	carry := 0
	for a != nil || b != nil {
		sum := 0
		if a != nil {
			sum += a.value
			a = a.Next
		}
		if b != nil {
			sum += b.value
			b = b.Next
		}
		current.Next = &NumValue{value: (sum + carry) % 10}
		current = current.Next
		carry = (sum + carry) / 10
	}
	if carry > 0 {
		current.Next = &NumValue{value: carry}
	}
	return rev.Next
}

type NumValue struct {
	value int
	Next  *NumValue
}

func (n *NumValue) String() string {
	if n.Next != nil {
		return fmt.Sprintf("%d -> %s", n.value, n.Next.String())
	}
	return fmt.Sprintln(n.value)

}

func (n *NumValue) Value() int {
	if n.Next != nil {
		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", n.Next.Value(), n.value))
		return v
	}
	return n.value
}

func NewNumValue(v int) *NumValue {
	current := &NumValue{value: 0}
	rev := current
	for v > 0 {
		current.Next = &NumValue{value: v % 10}
		current = current.Next
		v /= 10
	}
	return rev.Next
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
