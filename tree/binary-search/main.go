package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTreeSearch struct {
	root *Node
}

// Insert inserta un nuevo value en el árbol.
func (t *BinaryTreeSearch) Insert(value int) {
	newNode := &Node{value: value}

	if t.root == nil {
		t.root = newNode
		return
	}

	actual := t.root
	for {
		if value < actual.value {
			if actual.left == nil {
				actual.left = newNode
				break
			}
			actual = actual.left
		} else {
			if actual.right == nil {
				actual.right = newNode
				break
			}
			actual = actual.right
		}
	}
}

// Search busca un value en el árbol.
func (t *BinaryTreeSearch) Search(value int) bool {
	actual := t.root
	for actual != nil {
		if value == actual.value {
			return true
		} else if value < actual.value {
			actual = actual.left
		} else {
			actual = actual.right
		}
	}
	return false
}

// OrderedSearch realiza un recorrido en orden del árbol.
func (t *BinaryTreeSearch) OrderedSearch() {
	t.orderedSearch(t.root)
}

func (t *BinaryTreeSearch) orderedSearch(node *Node) {
	if node != nil {
		t.orderedSearch(node.left)
		fmt.Print(node.value, " ")
		t.orderedSearch(node.right)
	}
}

func main() {
	tree := &BinaryTreeSearch{}

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(80)

	fmt.Println("Recorrido en orden del árbol:")
	tree.OrderedSearch()

	fmt.Println("\nSearch 50:", tree.Search(50))
	fmt.Println("Search 100:", tree.Search(100))
}
