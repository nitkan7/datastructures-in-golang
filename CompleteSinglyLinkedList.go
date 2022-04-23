package main

import "fmt"

type LinkedList struct{
  value int
  next *LinkedList
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
func (head *LinkedList) AddAtHead(value int)*LinkedList{
  node:=LinkedList{value,nil}
  node.next=head
  head=&node
  return head
}

func (head *LinkedList) AddAtTail(value int){
  node:=LinkedList{value,nil}
  for ;head.next!=nil; {
  head=head.next
  }
  head.next=&node
}

func (head *LinkedList) display(){
  for ;head!=nil; {
    fmt.Println(*head)
    head=head.next
  }
}
func (head *LinkedList) AddAtIndex(where,value int){
  node :=LinkedList{value,nil}
  for index := 0; index < where-1; index++ {
    head=head.next
  }
  node.next=head.next
  head.next=&node
}
func (head *LinkedList) DeleteAtIndex(where int){
  for index := 0; index < where-1; index++ {
    head=head.next
  }
  node:=head.next.next
  head.next=node
}
