package main

// import (
// 	"fmt"
// )

// type List struct {
// 	Val  int
// 	Node *List
// }

// func main() {

// 	arr := []int{1, 2, 3}

// 	myList := Add(arr)

// 	// for myList != nil {
// 	// 	fmt.Printf("%v\n", myList)
// 	// 	myList = myList.Node
// 	// }

// 	// revList := Reverse(myList)

// 	revList := Reverse2(myList)

// 	// revList := Reverse3(myList)

// 	for revList != nil {
// 		fmt.Printf("%v\n", revList)
// 		revList = revList.Node
// 	}

// }

// func Add(arr []int) *List {

// 	if len(arr) == 0 {
// 		return nil
// 	}

// 	list := &List{Val: arr[0]}
// 	current := list

// 	for i := 1; i < len(arr); i++ {

// 		current.Node = &List{Val: arr[i]}
// 		current = current.Node
// 	}

// 	return list
// }

// func Reverse(list *List) *List {

// 	if list == nil {
// 		return nil
// 	}

// 	if list.Node == nil {
// 		return list
// 	}

// 	newList := Reverse(list.Node)
// 	list.Node.Node = list
// 	list.Node = nil

// 	return newList

// }

// func Reverse2(list *List) *List {

// 	var prev *List
// 	current := list

// 	for current != nil {
// 		next := current.Node
// 		current.Node = prev
// 		prev = current
// 		current = next
// 	}

// 	return prev

// }
