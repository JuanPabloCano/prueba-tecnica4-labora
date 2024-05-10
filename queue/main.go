package main

import "fmt"

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(element interface{}) {
	q.items = append(q.items, element)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	element := q.items[0]
	q.items = q.items[1:]
	return element
}

func (q *Queue) Front() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Tamaño de la cola:", queue.Size())
	fmt.Println("Frente de la cola:", queue.Front())

	fmt.Println("Desencolando elementos:")
	for !queue.Empty() {
		fmt.Println(queue.Dequeue())
	}
}

// FIFO - First In First Out
