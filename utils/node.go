package utils

type Node struct {
	Key       string
	Value     string
	Left      *Node
	Right     *Node
	Height    int
	Timestamp int64
}

type ArrNode struct {
	Key       string
	Value     string
	Timestamp int64
}
