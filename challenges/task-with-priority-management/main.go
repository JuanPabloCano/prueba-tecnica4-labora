package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FibNode struct {
	task     string
	priority int
	degree   int
	marked   bool
	parent   *FibNode
	child    *FibNode
	left     *FibNode
	right    *FibNode
}

type FibonacciHeap struct {
	min   *FibNode
	count int
}

func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{}
}

func (h *FibonacciHeap) Insert(task string, priority int) *FibNode {
	node := &FibNode{
		task:     task,
		priority: priority,
		left:     nil,
		right:    nil,
	}
	node.left = node
	node.right = node
	h.mergeIntoRootList(node)
	if h.min == nil || node.priority < h.min.priority {
		h.min = node
	}
	h.count++
	return node
}

func (h *FibonacciHeap) Minimum() *FibNode {
	return h.min
}

func (h *FibonacciHeap) RemoveMin() *FibNode {
	node := h.min
	if node != nil {
		if node.child != nil {
			start := node.child
			child := start
			for {
				next := child.right
				h.mergeIntoRootList(child)
				child.parent = nil
				child = next
				if child == start {
					break
				}
			}
		}
		// Remove node from the root list
		node.left.right = node.right
		node.right.left = node.left
		if node == node.right {
			h.min = nil
		} else {
			h.min = node.right
			h.consolidate()
		}
		h.count--
	}
	return node
}

func (h *FibonacciHeap) mergeIntoRootList(node *FibNode) {
	if h.min == nil {
		h.min = node
	} else {
		node.left = h.min
		node.right = h.min.right
		h.min.right.left = node
		h.min.right = node
	}
}

func (h *FibonacciHeap) consolidate() {
	degreeMap := make(map[int]*FibNode)
	var nodes []*FibNode
	for node := h.min; node != nil; node = node.right {
		nodes = append(nodes, node)
		if node == h.min.left {
			break
		}
	}
	for _, node := range nodes {
		deg := node.degree
		for {
			exist, ok := degreeMap[deg]
			if !ok {
				break
			}
			if node.priority > exist.priority {
				node, exist = exist, node
			}
			h.link(exist, node)
			delete(degreeMap, deg)
			deg++
		}
		degreeMap[deg] = node
	}
	minNode := h.min
	for _, node := range degreeMap {
		if node.priority < minNode.priority {
			minNode = node
		}
	}
	h.min = minNode
}

func (h *FibonacciHeap) link(y, x *FibNode) {
	y.left.right = y.right
	y.right.left = y.left
	y.parent = x
	if x.child == nil {
		x.child = y
		y.right = y
		y.left = y
	} else {
		y.left = x.child
		y.right = x.child.right
		x.child.right.left = y
		x.child.right = y
	}
	y.marked = false
	x.degree++
}

func main() {
	fh := NewFibonacciHeap()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Task Manager using Fibonacci Heap")
	fmt.Println("Commands: add <task> <priority>, min, remove, exit")

	for scanner.Scan() {
		input := scanner.Text()
		switch {
		case input == "exit":
			return

		case input == "min":
			minimum := fh.Minimum()
			if minimum != nil {
				fmt.Printf("Task with highest priority: %s (Priority %d)\n", minimum.task, minimum.priority)
			} else {
				fmt.Println("No tasks in the heap.")
			}

		case input == "remove":
			removeMin := fh.RemoveMin()
			if removeMin != nil {
				fmt.Printf("Completed task: %s (Priority %d)\n", removeMin.task, removeMin.priority)
			} else {
				fmt.Println("No tasks to complete.")
			}

		case strings.HasPrefix(input, "add"):
			parts := strings.Fields(input)
			if len(parts) < 3 {
				fmt.Println("Invalid command. Try 'add <task> <priority>'.")
				continue
			}
			priority, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				fmt.Println("Invalid priority. It must be an integer.")
				continue
			}
			task := strings.Join(parts[1:len(parts)-1], " ")
			fh.Insert(task, priority)
			fmt.Println("Added task:", task, "with priority", priority)

		default:
			fmt.Println("Unknown command:", input)
		}
		fmt.Println("Enter command:")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
