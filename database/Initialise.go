package database

import (
	"container/heap"

	"github.com/Jatin020403/BasaltDB/utils"
)

var partition_default string = "default"

func initialiseQueue() {
	utils.PQ = make(utils.PriorityQueue, 50)
	heap.Init(&utils.PQ)
}

func Initialise() {

	CreatePartition(partition_default)
	initialiseQueue()

}
