package main

import "fmt"

type aRRR []int

func (ar *aRRR) RemoveEle(v int) {
	pA := *ar
	lenAr := len(pA)
	for i := lenAr - 1; i >= 0; i-- {
		if pA[i] > v {
			pA = append(pA[:i], pA[i+1:]...)
		}
	}
	*ar = pA
}

// no
func (ar aRRR) RemoveEle2(v int) {
	for i := len(ar) - 1; i >= 0; i-- {
		if ar[i] > v {
			ar = append(ar[:i], ar[i+1:]...)
		}
	}
}

func main() {
	ar := aRRR{1, 10, 2, 3, 8, 4, 2, 15}
	ar.RemoveEle(5)
	fmt.Println(ar)

	ar2 := aRRR{1, 10, 2, 3, 8, 4, 2, 15}
	ar2.RemoveEle2(5)
	fmt.Println(ar2)


	l := StudentList{{Name: "test", Age: 18}, {Name: "zhang", Age: 8}, {Name: "san", Age: 19}, {Name: "li", Age: 5}, {Name: "si", Age: 10}}
	l2 := StudentList{{Name: "test", Age: 18}, {Name: "zhang", Age: 8}, {Name: "san", Age: 19}, {Name: "li", Age: 5}, {Name: "si", Age: 10}}
	removeStudent(&l, 9)
	fmt.Println(l)
	fmt.Println(l2)
	l2.removeStudent(9)
	fmt.Println(l2)
}

type Student struct {
	Name string
	Age int64
}

type StudentList []*Student

func removeStudent(s *StudentList, age int64) {
	if len(*s) == 0 {
		return
	}
	index := 0
	for _, item := range *s {
		if item.Age < age {
			(*s)[index] = item
			index++
		}
	}
	*s = (*s)[:index]
}

func (s *StudentList) removeStudent(age int64) {
	if len(*s) == 0 {
		return
	}
	tmp := *s
	index := 0
	for _, item := range tmp {
		if item.Age < age {
			tmp[index] = item
			index++
		}
	}
	*s = tmp[:index]
}