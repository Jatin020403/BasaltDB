package utils

// import "github.com/Jatin020403/BasaltDB/models"

// type PriorityQueue []*models.Node

// var PQ PriorityQueue

// func (PQ PriorityQueue) Len() int { return len(PQ) }

// func (PQ PriorityQueue) Less(i, j int) bool {
// 	if PQ[i] == nil {
// 		return false
// 	}
// 	if PQ[j] == nil {
// 		return true
// 	}
// 	return PQ[i].Timestamp > PQ[j].Timestamp
// }

// func (PQ PriorityQueue) Swap(i, j int) {
// 	PQ[i], PQ[j] = PQ[j], PQ[i]
// }

// func (PQ *PriorityQueue) Push(x any) {
// 	item := x.(*models.Node)
// 	*PQ = append(*PQ, item)
// }

// func (PQ *PriorityQueue) Pop() any {
// 	old := *PQ
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil // avoid memory leak
// 	*PQ = old[0 : n-1]
// 	return item
// }
