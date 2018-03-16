package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode{
	curr := head
	var last *ListNode = nil

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = last
		last = curr
		curr = nextTemp
	}
	return last
}


func main() {
	l1 := &ListNode{Val:1, Next: nil}
	l2 := &ListNode{Val:2, Next: nil}
	l3 := &ListNode{Val:3, Next: nil}
	l4 := &ListNode{Val:4, Next: nil}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	fmt.Println(l1)
	reverseList(l1)
	fmt.Println(l1)
}


/*
https://segmentfault.com/a/1190000004524683
Node reversalLinkedList(Node node){
    Node last = null;
    Node next = null;
    while(node != null){
        next = node.getNext();//1 拿到下个节点的引用，为了提供给第4步使node向链表后方移动
        node.setNext(last);//2 将当前节点指向上一个节点
        last = node;//3 将last指向当前节点，提供给下次循环的第2步
        node = next;//4 将当前节点的引用（即游标）指向下一个节点
    }
    /*
    * 循环完成后，node和next都指向了原节点D的next指向的位置，即null
    * 而last指向了上述null前面的位置，即节点D
    * 将last返回，则得到链表ABCD反转后链表DCBA的头节点D
    */
//return last;
//}

// */
