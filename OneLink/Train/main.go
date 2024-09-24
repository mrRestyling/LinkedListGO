package main

import "fmt"

type List struct {
	Val  int
	Node *List
}

func main() {

	arr := []int{1, 2, 3}

	myList := Add(arr)

	// for myList != nil {
	// 	fmt.Println(myList.Val)
	// 	myList = myList.Node
	// }

	resultList := Reverse(myList)

	for resultList != nil {
		fmt.Println(resultList.Val)
		resultList = resultList.Node
	}

}

func Add(arr []int) *List {

	if len(arr) == 0 {
		return nil
	}

	list := &List{Val: arr[0]}

	result := list

	for i := 1; i < len(arr); i++ {
		result.Node = &List{Val: arr[i]}
		result = result.Node
	}

	return list
}

func Reverse(list *List) *List {

	var result *List

	current := list

	for current != nil {
		next := current.Node
		current.Node = result
		result = current
		current = next
	}

	return result
}
