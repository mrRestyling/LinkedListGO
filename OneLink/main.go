package main

import "fmt"

type List struct {
	Val  int
	Next *List
}

func main() {
	// printList()
	addInList()

}

func printList() {

	// Создаем звенья списка:
	first := &List{Val: 1}

	second := &List{Val: 2}
	// Присваиваем к первому звену второе:
	first.Next = second
	// Присваиваем к третьему звену первое(со вторым):
	third := &List{Val: 3, Next: first}

	// Печатаем результат *используя цикл*:
	for third != nil {
		fmt.Print(third.Val, " ")
		third = third.Next
	}
}

func addInList() {

	foot := &List{Val: 666}
	head := &List{Val: 2, Next: foot}

	ListADD(head, 3)
	ListADD(head, 4)
	ListADD(head, 5)
	ListDEL(head, 666)

	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

func ListADD(head *List, val int) {

	newNode := &List{Val: val}

	for head.Next != nil {
		head = head.Next
	}

	head.Next = newNode
}

func ListDEL(myHEAD *List, myVAL int) *List {
	if myHEAD == nil {
		return nil
	}
	if myHEAD.Val == myVAL {
		return myHEAD.Next
	}

	for myHEAD.Next != nil {
		if myHEAD.Next.Val == myVAL {
			myHEAD.Next = myHEAD.Next.Next
			return myHEAD
		}
		myHEAD = myHEAD.Next
	}
	return myHEAD
}
