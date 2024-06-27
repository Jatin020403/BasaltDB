package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/Jatin020403/BasaltDB/utils"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(n *utils.Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func NewNode(key uint64, value string, timestamp int64) *utils.Node {
	node := &utils.Node{Key: key, Value: value}
	node.Left = nil
	node.Right = nil
	node.Height = 1
	node.Timestamp = timestamp
	return node
}

func getRoot(partition string, part int) (*utils.Node, error) {
	var arr []utils.ArrNode
	object, err := utils.Deserialize(partition, part, arr)

	if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("getRoot : " + err.Error())
	}

	if err != nil {
		return nil, errors.New("getRoot : " + err.Error())
	}

	var root *utils.Node

	for _, i := range object {
		root = insert(root, *NewNode(i.Key, i.Value, i.Timestamp))
	}

	return root, nil
}

func putRoot(partition string, part int, node *utils.Node) error {
	err := utils.Serialize(partition, part, node)
	if err != nil {
		return err
	}
	return nil
}

func rightRotate(y *utils.Node) *utils.Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1
	return x
}

func leftRotate(x *utils.Node) *utils.Node {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1
	return y
}

func get_balance_factor(N *utils.Node) int {
	if N == nil {
		return 0
	}
	return height(N.Left) - height(N.Right)
}

func insert(node *utils.Node, newNode utils.Node) *utils.Node {

	if node == nil {
		return NewNode(newNode.Key, newNode.Value, newNode.Timestamp)
	}
	if newNode.Key < node.Key {
		node.Left = insert(node.Left, newNode)
	} else if newNode.Key > node.Key {
		node.Right = insert(node.Right, newNode)
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

func node_with_minimum_value(node *utils.Node) *utils.Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func delete(root *utils.Node, key uint64) *utils.Node {

	if root == nil {
		return root
	}
	if key < root.Key {
		root.Left = delete(root.Left, key)
	} else if key > root.Key {
		root.Right = delete(root.Right, key)
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
			root.Right = delete(root.Right, temp.Key)
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

func get(node *utils.Node, key uint64) *utils.Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		return get(node.Left, key)
	} else if key > node.Key {
		return get(node.Right, key)
	}

	return node
}

func printRoot(root *utils.Node, indent string, last bool) {
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

	printRoot(root.Left, indent, false)
	printRoot(root.Right, indent, true)

}

func Print_inorder(root *utils.Node) {
	if root != nil {
		Print_inorder(root.Left)
		fmt.Println(fmt.Sprint(root.Key) + " : " + root.Value)
		Print_inorder(root.Right)
	}
}
