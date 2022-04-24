package main

import "fmt"

//represents the individual node of a linked list.
type LinkedList struct{
  value int //represents the value held by the node
  next *LinkedList //represents the next node it points to
}

func main(){
  head:=LinkedList{1,nil}
  head.AddAtTail(2)
  head.display()
  start:=head.AddAtHead(0)
  start.display()
  start.AddAtIndex(2,10)
  start.display()
  start.DeleteAtIndex(1)
  start.display()
}

//performs adding a new node a the head operation.
func (head *LinkedList) AddAtHead(value int)*LinkedList{
  node:=LinkedList{value,nil}
  node.next=head
  head=&node
  return head
}

//performs adding new node at the end of the list.
func (head *LinkedList) AddAtTail(value int){
  node:=LinkedList{value,nil}
  for ;head.next!=nil; {
  head=head.next
  }
  head.next=&node
}

//prints the elements of the linked list.
func (head *LinkedList) display(){
  for ;head!=nil; {
    fmt.Println(*head)
    head=head.next
  }
}

//performs addition of node before a specified index
func (head *LinkedList) AddAtIndex(where,value int){
  node :=LinkedList{value,nil}
  for index := 0; index < where-1; index++ {
    head=head.next
  }
  node.next=head.next
  head.next=&node
}

//performs deletion of a node at the given index.
func (head *LinkedList) DeleteAtIndex(where int){
  for index := 0; index < where-1; index++ {
    head=head.next
  }
  node:=head.next.next
  head.next=node
}
