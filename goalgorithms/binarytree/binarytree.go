package main

import (
	"fmt"
	"container/list"
)


type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Value int
}

func (tree *TreeNode) String() string {
	if tree == nil {
		return ""
	}

	fmt.Println("value: ", tree.Value)
	left := tree.Left
	if left != nil {
		fmt.Println("Left :", left.Value)
		left.String()
	}
	right := tree.Right
	if right != nil {
		fmt.Println("Right : ", right.Value)
		right.String()
	}
	return ""
}

func InsertNodeTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		panic("can not insert into nil root")
	}

	if node.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = node
		} else {
			InsertNodeTree(tree.Right, node)
		}
	}

	if node.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = node
		} else {
			InsertNodeTree(tree.Left, node)
		}
	}

}

/*
*InitTree
*/
func InitTree(index int, values ...int) *TreeNode {
	rootNode := &TreeNode{Value: values[index]}
	for _, value := range values {
		node := &TreeNode{Value: value}
		InsertNodeTree(rootNode, node)
	}

	return rootNode
}


func main() {
	values := []int {6, 3, 9, 10, 5, 2, 4, 18}

	treeNode := InitTree(0, values...)
	fmt.Println(treeNode)

	//sort.Sort(sort.IntSlice(values))
	//fmt.Println(values)
	//
	//treeNode = InitTree(3, values...)
	//fmt.Println(treeNode)

	// 根左右
	fmt.Println("-----begin 根左右----")
	PreOrder(treeNode)
	fmt.Println("---")
	PerOrder2(treeNode)
	fmt.Println("-----end 根左右")

	// 左根右
	fmt.Println("----------")
	InOrder(treeNode)

	// 左右根
	fmt.Println("----------")
	AfterOrder(treeNode)
}

// 先序
// 根左右
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Value, "")
	PreOrder(tree.Left)
	PreOrder(tree.Right)

	//if tree.Left != nil {
	//	PreOrder(tree.Left)
	//}
	//
	//if tree.Right != nil {
	//	PreOrder(tree.Right)
	//}
}

type MyStack struct {
	List *list.List
}

func (stack *MyStack) pop() interface{} {
	if elem := stack.List.Back(); elem != nil {
		stack.List.Remove(elem)
		return elem.Value
	}
	return nil
}

func (stack *MyStack) push(elem interface{}) {
	stack.List.PushBack(elem)
}

func PerOrder2(tree *TreeNode) {
	stack := MyStack{List: list.New()}

	currentNode := tree

	for currentNode != nil {
		fmt.Println(currentNode.Value)
		if right := currentNode.Right; right != nil {
			stack.push(right)
		}
		if left := currentNode.Left; left != nil {
			stack.push(left)
		}
		ele := stack.pop()
		if ele == nil {
			currentNode = nil
		} else {

			currentNode = ele.(*TreeNode)
		}
	}
	//stack.push(tree)
	//
	//currentNode := stack.pop()
	//
	//for currentNode != nil {
	//	node, _ := currentNode.(*TreeNode)
	//	fmt.Println(node.Value)
	//	if right := node.Right; right != nil {
	//		stack.push(right)
	//	}
	//	if left := node.Left; left != nil {
	//		stack.push(left)
	//	}
	//	currentNode = stack.pop()
	//}
}

// 中序
// 左根右
func InOrder(tree *TreeNode) {
	if tree  == nil {
		return
	}
	InOrder(tree.Left)
	fmt.Println(tree.Value, "")

	InOrder(tree.Right)
}

// 后序遍历
func AfterOrder(tree *TreeNode) {
	if tree  == nil {
		return
	}
	AfterOrder(tree.Left)
	AfterOrder(tree.Right)
	fmt.Println(tree.Value, "")
}


/*
queue := make([]int, 0)
// Push
queue := append(queue, 1)
// Top (just get next element, don't remove it)
x = queue[0]
// Discard top element
queue = queue[1:]
// Is empty ?
if len(queue) == 0 {
    fmt.Println("Queue is empty !")


}

----

type stack []int

func (s stack) Push(v int) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, int) {
    // FIXME: What do we do if the stack is empty, though?

    l := len(s)
    return  s[:l-1], s[l-1]
}


func main(){
    s := make(stack,0)
    s = s.Push(1)
    s = s.Push(2)
    s = s.Push(3)

    s, p := s.Pop()
    fmt.Println(p)

}

https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang

https://studygolang.com/articles/11932
 */