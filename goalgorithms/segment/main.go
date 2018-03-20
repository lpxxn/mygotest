package main

import (
	"fmt"
)

type SegTree struct {
	MaxValue int
	Start int
	End int
	Left *SegTree
	Right *SegTree
}


func BuildTree(arr []int, l, r int) *SegTree{
	if l > r {
		return nil
	}
	if l == r {
		return &SegTree{Start:l, End:r, MaxValue:arr[l]}
	}

	root := &SegTree{Start:l, End:r}
	mid := (l + r)/2
	root.Left = BuildTree(arr, l, mid)
	root.Right = BuildTree(arr, mid + 1, r)
	root.MaxValue = max(root.Left, root.Right)
	return root
}

func (tree *SegTree) Query(start, end int) int {
	if tree.Left == nil && tree.Right == nil {
		return tree.MaxValue
	}

	if start <= tree.Start && end >= tree.End {
		return tree.MaxValue
	}

	mid := (tree.Start + tree.End)/2
	if (start < mid && end > mid) {
		return maxInt(tree.Left.Query(start, mid), tree.Right.Query(mid + 1, end))
	} else if start <= mid && end <= mid {
		return tree.Left.Query(start, end)
	} else {
		return tree.Right.Query(start, end)
	}

}

func max(n1, n2 *SegTree) int {
	if n1.MaxValue > n2.MaxValue {
		return n1.MaxValue
	}
	return n2.MaxValue
}

func maxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	arr := []int{1, 4, 2, 3}
	tree := BuildTree(arr, 0, len(arr) - 1)
	fmt.Println(tree.Query(2, 2))
}
/*

http://skyhigh233.com/blog/2016/10/03/seg-tree/
 */