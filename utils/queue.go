package utils

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
