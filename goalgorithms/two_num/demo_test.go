package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {

	v1 := 123
	n1 := NewNum(v1)
	t.Log(n1.String())
	t.Log(n1.Value())
	if n1.Value() != v1 {
		t.Fatal()
	}

	n2 := NewNum(789)

	n3 := AddTowNum(n1, n2)
	if n3.Value() != v1 + n2.Value() {
		t.Fatal()
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

func AddTowNum(a, b *Num) *Num {
	current := &Num{}
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
		current.next = &Num{value: sum % 10}
		carry = sum / 10
		current = current.next
	}
	if carry > 0 {
		current.next = &Num{value: carry}
	}
	return rev.next
}

type Num struct {
	value int
	next  *Num
}

func NewNum(v int) *Num {
	current := &Num{}
	rev := current
	for v > 0 {
		current.next = &Num{value: v % 10}
		v = v / 10
		current = current.next
	}
	return rev.next
}

func (n Num) String() string {
	if n.next != nil {
		return fmt.Sprintf("%d -> %s", n.value, n.next.String())
	}
	return fmt.Sprint(n.value)
}

func (n Num) Value() int {
	if n.next != nil {
		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", n.next.Value(), n.value))
		return v
	}
	return n.value
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
