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
}
