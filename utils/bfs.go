package utils

import "github.com/Jatin020403/BasaltDB/models"

func bsf(node *models.Node) []models.ArrNode {
	if node == nil {
		return make([]models.ArrNode, 0)
	}

	arr := []models.ArrNode{}
	q := []models.Node{*node}
	var elem models.Node

	for len(q) != 0 {
		elem = q[0]
		q = q[1:]

		arrNode := models.ArrNode{Key: elem.Key, Value: elem.Value, Timestamp: elem.Timestamp}
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
