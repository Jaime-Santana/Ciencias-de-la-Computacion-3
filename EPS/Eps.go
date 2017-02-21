package main

import (
	"fmt"
)

var lista[] Estructura

type Node struct{
	nombre string
	cedula string
	eps string
}

type Estructura struct{
	nom string
	conteo int
}

func (n *Node) String() string {
	return fmt.Sprint(n.nombre, "->" ,n.cedula)
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	buscar(node.eps)
	return node
}

func buscar(eps string){
	var token bool = false
	for i := 0; i < len(lista); i++ {
		if lista[i].nom == eps{
			token = true
			lista[i].conteo = lista[i].conteo+1
		}
	}
	if token==false{
		s := Estructura{nom:eps,conteo:1}
		lista = append(lista,s)
	}
} 

func main() {

  fmt.Println("Cola")
	q := NewQueue(1)
	q.Push(&Node{"Ana","1022411","Famisanar"})
	q.Push(&Node{"Maria","3516541","NuevaEps"})
	q.Push(&Node{"Jaime","3552441","Sisben"})
	q.Push(&Node{"Daniel","1022411","Famisanar"})
	q.Push(&Node{"Maria","3516541","NuevaEps"})
	q.Push(&Node{"Maria","3516541","NuevaEps"})
	q.Push(&Node{"Maria","3516541","NuevaEps"})
	fmt.Println(q.Pop(), q.Pop(), q.Pop(),q.Pop(),q.Pop(),q.Pop(),q.Pop())
	fmt.Println(lista)
}