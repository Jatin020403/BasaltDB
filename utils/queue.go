package utils

import (
	"container/heap"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

var PQ PriorityQueue

func (PQ PriorityQueue) Len() int { return len(PQ) }

func (PQ PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if PQ[i] == nil {
		return false
	}
	if PQ[j] == nil {
		return true
	}
	return PQ[i].Timestamp > PQ[j].Timestamp
}

func (PQ PriorityQueue) Swap(i, j int) {
	PQ[i], PQ[j] = PQ[j], PQ[i]
}

func (PQ *PriorityQueue) Push(x any) {
	item := x.(*Node)
	*PQ = append(*PQ, item)
}

func (PQ *PriorityQueue) Pop() any {
	old := *PQ
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*PQ = old[0 : n-1]
	return item
}

func InitialiseQueue() {
	PQ = make(PriorityQueue, 50)
	heap.Init(&PQ)

}

// // This example creates a PriorityQueue with some items, adds and manipulates an item,
// // and then removes the items in priority order.
// func Example_priorityQueue() {
// 	// Some items and their priorities.
// 	items := map[string]int{
// 		"banana": 3, "apple": 2, "pear": 4,
// 	}

// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	PQ := make(PriorityQueue, len(items))
// 	i := 0
// 	for value, priority := range items {
// 		PQ[i] = &Item{
// 			value:    value,
// 			priority: priority,
// 			index:    i,
// 		}
// 		i++
// 	}
// 	heap.Init(&PQ)

// 	// Insert a new item and then modify its priority.
// 	item := &Item{
// 		value:    "orange",
// 		priority: 1,
// 	}
// 	heap.Push(&PQ, item)
// 	PQ.update(item, item.value, 5)

// 	// Take the items out; they arrive in decreasing priority order.
// 	for PQ.Len() > 0 {
// 		item := heap.Pop(&PQ).(*Item)
// 		fmt.Printf("%.2d:%s ", item.priority, item.value)
// 	}
// 	// Output:
// 	// 05:orange 04:pear 03:banana 02:apple
// }
