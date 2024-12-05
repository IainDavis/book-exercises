package exercises

import (
	"fmt"
)

// header struct
type LinkedList[T comparable] struct {
	head *ListNode[T]
	tail *ListNode[T]
}

// node struct
type ListNode[T comparable] struct {
	value T
	next  *ListNode[T]
}

// node factory
func newListNode[T comparable](value T) *ListNode[T] {
	newNode := ListNode[T]{
		value: value,
		next:  nil,
	}

	return &newNode
}

func (l LinkedList[T]) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList[T]) Add(value T) {
	newNode := newListNode(value)

	if l.isEmpty() {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList[T]) Insert(value T, target int) {

	fmt.Printf("inserting %v at index %d\n", value, target)

	lTrav := l.head
	var lPrev *ListNode[T]

	for ; target >= 0; lPrev, lTrav, target = lTrav, lTrav.next, target-1 {
		if target == 1 && lTrav == nil {
			// insert at end of list
			l.Add(value)
			break
		} else if target > 0 && lTrav == nil {
			fmt.Println("Failed to add -- target out-of-bounds for list")
			break
		} else if target == 0 {
			// insert at head, or mid-list
			newNode := newListNode(value)
			newNode.next = lTrav
			if lTrav == l.head {
				l.head = newNode
			}
			if lTrav == l.tail {
				l.tail = newNode
			}
			if lPrev != nil {
				lPrev.next = newNode
			}
			break
		} else {
			continue
		}
	}

	if target == -1 {
		fmt.Println("Insert failed: index beyond end of list")
	}

}

func (l *LinkedList[T]) Print() {
	if l.head == nil {
		fmt.Println("<empty list>")
		return
	}

	lTrav := l.head
	for {
		fmt.Println((*lTrav).value)
		lTrav = lTrav.next
		if lTrav == nil {
			break
		}
	}
}

func Exercise_8_3() {
	var list LinkedList[int]
	list.Add(2)
	list.Add(3)
	list.Add(7)
	list.Add(22)

	fmt.Println("Before insert")
	list.Print()
	fmt.Println()

	fmt.Println("Insert into list middle")
	list.Insert(12, 2)
	list.Print()
	fmt.Println()

	fmt.Println("Insert at list head")
	list.Insert(15, 0)
	list.Print()
	fmt.Println()

	fmt.Println("Insert at list tail")
	list.Insert(200, 6)
	list.Print()
	fmt.Println()

	fmt.Println("Insert at out-of-range index")
	list.Insert(240, 75)
	list.Print()
	fmt.Println()
}
