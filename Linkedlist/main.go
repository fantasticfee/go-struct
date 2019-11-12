package main

import (
	"fmt"
	"go-struct/Linkedlist/singlinklist"
)

func main() {
	list := singlinklist.NewLinkList()
	for i := 1; i <= 10; i++ {
		newNode := singlinklist.NewLinkNode(i)
		if !(l.InsertToTail(newNode)) {
			fmt.Println("insert error")
		}
	}
	l.Println()

}
