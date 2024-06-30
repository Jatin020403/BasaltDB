package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/Jatin020403/BasaltDB/models"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(n *models.Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func NewNode(key uint64, value string, timestamp int64) *models.Node {
	node := &models.Node{Key: key, Value: value}
	node.Left = nil
	node.Right = nil
	node.Height = 1
	node.Timestamp = timestamp
	return node
}

func GetRoot(partition models.Partition, part int) (*models.Node, error) {
	var arr []models.ArrNode
	object, err := Deserialize(partition, part, arr)

	if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("getRoot : " + err.Error())
	}

	if err != nil {
		return nil, errors.New("getRoot : " + err.Error())
	}

	var root *models.Node

	for _, i := range object {
		root = Insert(root, *NewNode(i.Key, i.Value, i.Timestamp))
	}

	return root, nil
}

func PutRoot(partition models.Partition, part int, node *models.Node) error {
	err := Serialize(partition, part, node)
	if err != nil {
		return err
	}
	return nil
}

func rightRotate(y *models.Node) *models.Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1
	return x
}

func leftRotate(x *models.Node) *models.Node {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1
	return y
}

func get_balance_factor(N *models.Node) int {
	if N == nil {
		return 0
	}
	return height(N.Left) - height(N.Right)
}

func Insert(node *models.Node, newNode models.Node) *models.Node {

	if node == nil {
		return NewNode(newNode.Key, newNode.Value, newNode.Timestamp)
	}
	if newNode.Key < node.Key {
		node.Left = Insert(node.Left, newNode)
	} else if newNode.Key > node.Key {
		node.Right = Insert(node.Right, newNode)
	} else {
		if newNode.Timestamp > node.Timestamp {
			node.Value = newNode.Value
			node.Timestamp = newNode.Timestamp
			return node
		}
		return node
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	balanceFactor := get_balance_factor(node)

	if balanceFactor > 1 {
		if newNode.Key < node.Left.Key {
			return rightRotate(node)
		} else if newNode.Key > node.Left.Key {
			node.Left = leftRotate(node.Left)
			return rightRotate(node)
		}
	}

	if balanceFactor < -1 {
		if newNode.Key > node.Right.Key {
			return leftRotate(node)
		} else if newNode.Key < node.Right.Key {
			node.Right = rightRotate(node.Right)
			return leftRotate(node)
		}
	}

	return node
}

func node_with_minimum_value(node *models.Node) *models.Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func Delete(root *models.Node, key uint64) *models.Node {

	if root == nil {
		return root
	}
	if key < root.Key {
		root.Left = Delete(root.Left, key)
	} else if key > root.Key {
		root.Right = Delete(root.Right, key)
	} else {
		if root.Left == nil || root.Right == nil {
			temp := root.Left
			if temp == nil {
				temp = root.Right
			}
			if temp == nil {
				root = nil
			} else {
				*root = *temp
			}
		} else {
			temp := node_with_minimum_value(root.Right)
			root.Key = temp.Key
			root.Value = temp.Value
			root.Timestamp = temp.Timestamp
			root.Right = Delete(root.Right, temp.Key)
		}
	}
	if root == nil {
		return root
	}
	root.Height = 1 + max(height(root.Left), height(root.Right))
	balanceFactor := get_balance_factor(root)

	if balanceFactor > 1 {
		if get_balance_factor(root.Left) >= 0 {
			return rightRotate(root)
		} else {
			root.Left = leftRotate(root.Left)
			return rightRotate(root)
		}
	}
	if balanceFactor < -1 {
		if get_balance_factor(root.Right) <= 0 {
			return leftRotate(root)
		} else {
			root.Right = rightRotate(root.Right)
			return leftRotate(root)
		}
	}

	return root
}

func Get(node *models.Node, key uint64) *models.Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		return Get(node.Left, key)
	} else if key > node.Key {
		return Get(node.Right, key)
	}

	return node
}

func PrintRoot(root *models.Node, indent string, last bool) {
	if root == nil {
		return
	}

	fmt.Print(indent)
	if last {
		fmt.Print("R----")
		indent += "   "
	} else {
		fmt.Print("L----")
		indent += "|  "
	}
	fmt.Println(" " + fmt.Sprint(root.Key) + "==>" + root.Value + " Timestamp: " + fmt.Sprint(root.Timestamp))

	PrintRoot(root.Left, indent, false)
	PrintRoot(root.Right, indent, true)

}

func Print_inorder(root *models.Node) {
	if root != nil {
		Print_inorder(root.Left)
		fmt.Println(fmt.Sprint(root.Key) + " : " + root.Value)
		Print_inorder(root.Right)
	}
}
