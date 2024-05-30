package utils

func bsf(node *Node) []ArrNode {
	if node == nil {
		return make([]ArrNode, 0)
	}

	arr := []ArrNode{}
	q := []Node{*node}
	var elem Node

	for len(q) != 0 {
		elem = q[0]
		q = q[1:]

		arrNode := ArrNode{Key: elem.Key, Value: elem.Value, Timestamp: elem.Timestamp}
		arr = append(arr, arrNode)

		if elem.Left != nil {
			q = append(q, *elem.Left)
		}
		if elem.Right != nil {
			q = append(q, *elem.Right)
		}
	}

	return arr

}
