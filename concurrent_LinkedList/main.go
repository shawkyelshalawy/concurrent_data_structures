package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key  int
	next *Node
}

type list struct {
	head   *Node
	locker sync.Mutex
}

func (l *list) Init() {
	l.head = nil
}

func (l *list) Insert(key int) {
	newNode := &Node{key: key}
	l.locker.Lock()
	newNode.next = l.head
	l.head = newNode
	l.locker.Unlock()
}

func (l *list) Lookup(key int) bool {
	l.locker.Lock()
	defer l.locker.Unlock()
	curr := l.head
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *list) Delete(key int) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.head == nil {
		return
	}
	if l.head.key == key {
		l.head = l.head.next
		return
	}

	prev := l.head
	curr := l.head.next
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			return
		}
		prev = curr
		curr = curr.next
	}
}
func (l *list) InsertAfter(prevKey, newKey int) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	curr := l.head
	for curr != nil {
		if curr.key == prevKey {
			newNode := &Node{key: newKey, next: curr.next}
			curr.next = newNode
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *list) InsertBefore(nextKey, newKey int) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.head == nil {
		return false
	}

	if l.head.key == nextKey {
		newNode := &Node{key: newKey, next: l.head}
		l.head = newNode
		return true
	}

	prev := l.head
	curr := l.head.next

	for curr != nil {
		if curr.key == nextKey {
			newNode := &Node{key: newKey, next: curr}
			prev.next = newNode
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}
func (l *list) Print() {
	l.locker.Lock()
	defer l.locker.Unlock()

	curr := l.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.key)
		curr = curr.next
	}
	fmt.Println("nil")
}

func main() {
	list := &list{}
	list.Init()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			list.Insert(key)
		}(i)
	}

	wg.Wait()

	list.Print()

	for i := 0; i < 12; i++ {
		found := list.Lookup(i)
		fmt.Printf("Lookup %d: %v\n", i, found)
	}

	list.Delete(5)
	list.Print()
}
