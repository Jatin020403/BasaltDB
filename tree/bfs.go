package tree

import "github.com/Jatin020403/BasaltDB/data"

func Bfs(node *Tree) []data.DataNode {
	if node == nil {
		return make([]data.DataNode, 0)
	}

	arr := []data.DataNode{}
	q := []Tree{*node}
	var elem Tree

	for len(q) != 0 {

		elem = q[0]
		q = q[1:]
		node := data.DataNode{Key: elem.Key, KeyString: elem.KeyString, Value: elem.Value, Timestamp: elem.Timestamp}
		arr = append(arr, node)

		if elem.Left != nil {
			q = append(q, *elem.Left)
		}
		if elem.Right != nil {
			q = append(q, *elem.Right)
		}

	}

	return arr
}
