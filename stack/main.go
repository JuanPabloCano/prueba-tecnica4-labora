package main

import "fmt"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(element interface{}) {
	s.items = append(s.items, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	element := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return element
}

func (s *Stack) Top() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[0]
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Tamaño de la pila:", stack.Size())
	fmt.Println("Tope de la pila:", stack.Top())
	fmt.Println("Vista previa de la pila:", stack.Peek())

	fmt.Println("Desapilando elementos:")
	for !stack.Empty() {
		fmt.Println(stack.Pop())
	}
}

// LIFO - Last In First Out
