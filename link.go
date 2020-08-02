package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}



func main()  {

	node1 := &Node{
		Data: "hello",
	}
	node2 := &Node{
		Data: "world",
	}
	node1.Next = node2
	node3 := &Node{
		Data: "test",
	}
	node2.Next = node3
	currentNode := node1
	for  {
		fmt.Printf("data=%s\n", currentNode.Data)
		currentNode = currentNode.Next
		if currentNode.Next == nil {
			break
		}
	}
}


