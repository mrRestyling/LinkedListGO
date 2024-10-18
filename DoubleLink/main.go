package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type DoublyList struct {
	Head *Node
	Tail *Node
}

func main() {
	// CreateEx1()
	CreateEx2()

}

func CreateEx1() {

	head := &Node{Val: 1}
	second := &Node{Val: 2, Prev: head}
	head.Next = second

	third := &Node{Val: 3, Prev: second}
	second.Next = third

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func CreateEx2() {

	dl := NewDL()

	dl.AppendDl(1)
	dl.AppendDl(2)
	dl.AppendDl(3)
	dl.UpAppend(0)
	dl.Insert(666, 3)
	dl.Remove(666)

	dl.PrintALL()
}

func NewDL() *DoublyList {
	return &DoublyList{}
}

// Вставляем в конец списка
func (DL *DoublyList) AppendDl(val int) {
	newNode := &Node{Val: val}

	if DL.Head == nil {
		DL.Head = newNode
		DL.Tail = newNode
	} else {
		DL.Tail.Next = newNode
		newNode.Prev = DL.Tail
		DL.Tail = newNode
	}
}

// Вставляем в начало списка
func (DL *DoublyList) UpAppend(val int) {

	newNode := &Node{Val: val}

	if DL.Head == nil {
		DL.Head = newNode
		DL.Tail = newNode
	}

	DL.Head.Prev = newNode
	newNode.Next = DL.Head
	DL.Head = newNode
}

// Выводим весь список
func (DL *DoublyList) PrintALL() {
	current := DL.Head

	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

// Вставляем в любое место списка
func (DL *DoublyList) Insert(val int, index int) {
	if index < 0 {
		fmt.Println("Некорректный индекс")
		return
	}
	if index == 0 {
		DL.UpAppend(val)
		return
	}

	newNode := &Node{Val: val}

	current := DL.Head
	currentIndex := 0

	for current != nil && currentIndex < index-1 {
		current = current.Next
		currentIndex++
	}

	if current == nil {
		DL.AppendDl(val)
		return
	}

	newNode.Next = current.Next
	newNode.Prev = current
	current.Next = newNode

	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}

// Удаляем элемент из списка
func (DL *DoublyList) Remove(val int) {

	if DL.Head == nil {
		return
	}

	if DL.Head.Val == val {
		DL.Head = DL.Head.Next
		if DL.Head == nil {
			DL.Tail = nil
		}
		return
	}

	// Ищем элемент в списке:
	current := DL.Head

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			if current.Next == nil {
				DL.Tail = current
			}
			return
		}
		current = current.Next
	}
}
