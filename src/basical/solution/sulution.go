package main

import (
	"fmt"
	"time"
)

type node struct {
	next   *node
	random *node
	value  int
}

func copyList(head *node) *node {
	if head == nil {
		return nil
	}
	newHead := node{
		value:  head.value,
		next:   nil,
		random: nil,
	}
	p := head.next
	pre := &newHead
	record := make(map[*node]*node)
	record[head] = pre
	for p != nil {
		newNode := &node{
			value:  p.value,
			next:   nil,
			random: nil,
		}
		pre.next = newNode
		record[p] = newNode
		p = p.next
		pre = pre.next
	}
	p = head
	q := &newHead
	for p != nil {
		if p.random != nil {
			q.random = record[p.random]
		}
		p = p.next
		q = q.next
	}
	return &newHead
}

func printList(head *node) {
	p := head
	for p != nil {
		fmt.Println(p.value, p.next, p.random)
		p = p.random
	}
}

func main() {
	head := &node{
		value:  0,
		next:   nil,
		random: nil,
	}

	node1 := &node{
		value:  1,
		next:   nil,
		random: nil,
	}

	node2 := &node{
		value:  2,
		next:   nil,
		random: nil,
	}
	head.next = node1
	head.random = node2
	node1.next = node2
	node1.random = head

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("ha")
	}()

	newHead := copyList(head)
	printList(head)
	printList(newHead)

}
