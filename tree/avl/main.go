package main

import (
	"fmt"
	"math"
)

type NodeAVL struct {
	value int
	left  *NodeAVL
	right *NodeAVL
	high  int
}

type TreeAVL struct {
	root *NodeAVL
}

// Función para obtener la high de un nodo.
func high(node *NodeAVL) int {
	if node == nil {
		return -1
	}
	return node.high
}

// Función para calcular el Balance de un nodo.
func Balance(node *NodeAVL) int {
	if node == nil {
		return 0
	}
	return high(node.left) - high(node.right)
}

// Función para rotar a la izquierda en el nodo dado.
func turnLeft(node *NodeAVL) *NodeAVL {
	right := node.right
	node.right = right.left
	right.left = node
	node.high = int(math.Max(float64(high(node.left)), float64(high(node.right))) + 1)
	right.high = int(math.Max(float64(high(right.left)), float64(high(right.right))) + 1)
	return right
}

// Función para rotar a la derecha en el nodo dado.
func turnRight(node *NodeAVL) *NodeAVL {
	right := node.left
	node.left = right.right
	right.right = node
	node.high = int(math.Max(float64(high(node.left)), float64(high(node.right))) + 1)
	right.high = int(math.Max(float64(high(right.left)), float64(high(right.right))) + 1)
	return right
}

// Función para insert un value en el árbol AVL.
func (t *TreeAVL) Insert(value int) {
	t.root = insert(t.root, value)
}

func insert(node *NodeAVL, value int) *NodeAVL {
	if node == nil {
		return &NodeAVL{value: value, high: 0}
	}

	if value < node.value {
		node.left = insert(node.left, value)
	} else if value > node.value {
		node.right = insert(node.right, value)
	} else {
		return node // El value ya existe
	}

	// Actualizar high del node actual
	node.high = int(math.Max(float64(high(node.left)), float64(high(node.right))) + 1)

	// Calcular el Balance del node
	balance := Balance(node)

	// Casos de desequilibrio y rotaciones
	if balance > 1 { // Desbalance hacia la izquierda
		if value < node.left.value {
			// Rotación simple a la derecha
			return turnRight(node)
		} else {
			// Rotación doble a la izquierda-derecha
			node.left = turnLeft(node.left)
			return turnRight(node)
		}
	}
	if balance < -1 { // Desbalance hacia la derecha
		if value > node.right.value {
			// Rotación simple a la izquierda
			return turnLeft(node)
		} else {
			// Rotación doble a la derecha-izquierda
			node.right = turnRight(node.right)
			return turnLeft(node)
		}
	}

	return node
}

// Función para search un value en el árbol AVL.
func (t *TreeAVL) Search(value int) bool {
	return search(t.root, value)
}

func search(node *NodeAVL, value int) bool {
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

// Función para realizar un recorrido en orden del árbol AVL.
func (t *TreeAVL) OrderedTraversal() {
	orderedTraversal(t.root)
	fmt.Println()
}

func orderedTraversal(node *NodeAVL) {
	if node != nil {
		orderedTraversal(node.left)
		fmt.Print(node.value, " ")
		orderedTraversal(node.right)
	}
}

func main() {
	tree := &TreeAVL{}

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(80)

	fmt.Println("Recorrido en orden del árbol AVL:")
	tree.OrderedTraversal()

	fmt.Println("Search 50:", tree.Search(50))
	fmt.Println("Search 100:", tree.Search(100))
}
