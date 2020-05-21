package main

// Node definition struct
type Node struct {
	value    int
	previous *Node
	next     *Node
}

// DoubleLinkedList definition struct
type DoubleLinkedList struct {
	head *Node
	tail *Node
}

// AddNode is to add a new node to the end of the list
func (dll *DoubleLinkedList) AddNode(node Node) {
	if dll.head == nil {
		dll.head = &node
		dll.tail = &node
	} else {
		dll.tail.next = &node
		node.previous = dll.tail
		dll.tail = &node
	}
}

// AddToBeginning helps to add a node to the beginning of the list
func (dll *DoubleLinkedList) AddToBeginning(node Node) {
	if dll.head == nil {
		dll.head = &node
		dll.tail = &node
	} else {
		node.next = dll.head
		dll.head.previous = &node
		dll.head = &node
	}
}

// Length of the list is returned
func (dll *DoubleLinkedList) Length() int {
	var length int = 0
	currentNode := dll.head
	if currentNode != nil {
		for currentNode != nil {
			currentNode = currentNode.next
			length++
		}
	}
	return length
}

// WalkList helps to walk the list
func (dll *DoubleLinkedList) WalkList() {
	currentNode := dll.head
	print("\nWalking the list: ")
	for currentNode != nil {
		print(currentNode.value, " ")
		currentNode = currentNode.next
	}
}

// WalkBackwards walks the list from the end
func (dll *DoubleLinkedList) WalkBackwards() {
	currentNode := dll.tail
	print("\nWalking the list backwards: ")
	for currentNode != nil {
		print(currentNode.value, " ")
		currentNode = currentNode.previous
	}
}

func main() {
	dll := DoubleLinkedList{}

	dll.AddNode(Node{value: 5})
	dll.AddNode(Node{value: 57})
	dll.AddToBeginning(Node{value: 4})
	dll.AddToBeginning(Node{value: 13})

	dll.WalkList()
	dll.WalkBackwards()
	println("\nLength of the list:", dll.Length())
}
