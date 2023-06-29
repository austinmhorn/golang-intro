package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"rsc.io/quote"
)

func main() {
	fmt.Println("-- Initiating...")

	fmt.Println("Go Quote: " + quote.Go())

	// Create a new list and put some strings in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	l.PushBack(5)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("Enter # to search for: ")

	var num int

	// Taking input from user
	fmt.Scanln(&num)

	i, found := sort.Find(l.Len(), func(i int) int {
		return strings.Compare(strconv.Itoa(num), strconv.Itoa(i))
	})
	if found {
		fmt.Printf("Found %s at entry %d\n", strconv.Itoa(num), i)
	} else {
		fmt.Printf("%s not found, would insert at %d", strconv.Itoa(num), i)
	}

	fmt.Println("-- Terminating...")
}
