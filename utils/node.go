package utils

type Node struct {
	Key       uint64
	Value     string
	Left      *Node
	Right     *Node
	Height    int
	Timestamp int64
}

type ArrNode struct {
	Key       uint64
	Value     string
	Timestamp int64
}
