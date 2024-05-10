package main

import "fmt"

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type NodeRN struct {
	value int
	left  *NodeRN
	right *NodeRN
	color Color
}

type TreeRN struct {
	root *NodeRN
}

func (t *TreeRN) Insert(valor int) {
	t.root = insert(t.root, valor)
	t.root.color = Black // La raíz siempre debe ser negra
}

func insert(node *NodeRN, value int) *NodeRN {
	if node == nil {
		return &NodeRN{value: value, color: Red}
	}

	if value < node.value {
		node.left = insert(node.left, value)
	} else if value > node.value {
		node.right = insert(node.right, value)
	}

	// Restaurar las propiedades del árbol rojo-negro después de la inserción
	if isRed(node.right) && !isRed(node.left) {
		node = turnLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = turnRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		invertColors(node)
	}

	return node
}

func isRed(node *NodeRN) bool {
	if node == nil {
		return false
	}
	return node.color == Red
}

func turnLeft(node *NodeRN) *NodeRN {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = Red
	return x
}

func turnRight(node *NodeRN) *NodeRN {
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = Red
	return x
}

func invertColors(node *NodeRN) {
	node.color = !node.color
	node.left.color = !node.left.color
	node.right.color = !node.right.color
}

func (t *TreeRN) Search(value int) bool {
	return search(t.root, value)
}

func search(node *NodeRN, value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	} else if value < node.value {
		return search(node.left, value)
	} else {
		return search(node.right, value)
	}
}

func (t *TreeRN) InOrderTraversal() {
	inOrderTraversal(t.root)
	fmt.Println()
}

func inOrderTraversal(node *NodeRN) {
	if node != nil {
		inOrderTraversal(node.left)
		fmt.Print(node.value, " ")
		inOrderTraversal(node.right)
	}
}

func main() {
	tree := &TreeRN{}

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(80)

	fmt.Println("Recorrido en orden del árbol rojo-negro:")
	tree.InOrderTraversal()

	fmt.Println("Search 50:", tree.Search(50))
	fmt.Println("Search 100:", tree.Search(100))
}
