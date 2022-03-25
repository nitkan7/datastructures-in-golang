package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (ll *linkedList) deleteNodeWithValue(val int) {
	fmt.Println("Deleting node with value ", val)
	iterator := ll.head
	for i := 0; i < ll.length; i++ {
		if iterator.data == val && i == 0 {
			ll.head = ll.head.next
			ll.length--
		} else if iterator.next.data == val {
			iterator.next = iterator.next.next
			ll.length--
		}
	}
}

func (ll *linkedList) append(n *node) {
	iterator := ll.head
	if ll.length == 0 {
		ll.head = n
		ll.length++
	} else {
		for i := 0; i < ll.length-1; i++ {
			iterator = iterator.next
		}
		iterator.next = n
		ll.length++
	}
}
func (l *linkedList) printLinkedList() {
	fmt.Println("Printing the List:")
	iterator := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", iterator.data)
		iterator = iterator.next
	}
	fmt.Println()
}

func main() {
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	ll := linkedList{}
	ll.append(node1)
	ll.append(node2)
	ll.append(node3)
	ll.printLinkedList()
	ll.deleteNodeWithValue(2)
	ll.printLinkedList()

}
