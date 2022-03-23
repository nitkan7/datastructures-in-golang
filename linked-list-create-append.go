package main

import "fmt"

//induvidual node of the linked list
type node struct{ 
	data int
	next *node
}

//the linkedList datastructure itself
type linkedList struct{
	head *node
	length int
}

//operations on the linked list - append - adds a node to the end of the list
func (l *linkedList) append(n *node){
	if(l.length==0){
		l.head=n
		l.length++
	}else{
	headcopy:=l.head
	for i := 0; i < l.length-1; i++ {
		
		headcopy=headcopy.next

	}
	headcopy.next=n
	l.length++
}
}

//operations on the linkedList - print the whole list from head to end
func (l *linkedList) printLinkedList(){
	iterator:=l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(iterator.data)
		iterator=iterator.next
	}
}

//main function
func main(){
	node1:=&node{data:1}
	node2:=&node{data:2}
	node3:=&node{data:3}
	ll:=linkedList{}
	ll.append(node1)
	ll.append(node2)
	ll.append(node3)
	ll.printLinkedList()

}
