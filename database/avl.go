package database

import (
	"fmt"
	"time"

	"github.com/Jatin020403/BasaltDB/utils"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(N *utils.Node) int {
	if N == nil {
		return 0
	}
	return N.Height
}

func NewNode(key string, value string, timestamp int64) *utils.Node {
	node := &utils.Node{Key: key, Value: value}
	node.Left = nil
	node.Right = nil
	node.Height = 1
	node.Timestamp = timestamp
	return node
}

func getRoot() *utils.Node {
	var arr []utils.ArrNode
	object, err := utils.Deserialize(arr)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var root *utils.Node

	for _, i := range object {
		root = insert(root, i.Key, i.Value, i.Timestamp)
	}

	return root
}

func putRoot(node *utils.Node) error {
	err := utils.Serialize(node)
	if err != nil {
		fmt.Println(err.Error())
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

func getBalanceFactor(N *utils.Node) int {
	if N == nil {
		return 0
	}
	return height(N.Left) - height(N.Right)
}

func InsertMany(object []utils.ArrNode) bool {

	root := getRoot()
	for _, i := range object {
		root = insert(root, i.Key, i.Value, time.Now().UnixNano())
	}

	if err := putRoot(root); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true

}

func InsertOne(key string, value string) bool {

	root := getRoot()
	root = insert(root, key, value, time.Now().UnixNano())

	if err := putRoot(root); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func insert(node *utils.Node, key string, value string, timestamp int64) *utils.Node {
	if node == nil {
		return NewNode(key, value, timestamp)
	}
	if key < node.Key {
		node.Left = insert(node.Left, key, value, timestamp)
	} else if key > node.Key {
		node.Right = insert(node.Right, key, value, timestamp)
	} else {
		if node.Timestamp < timestamp {
			return NewNode(key, value, timestamp)
		}
		return node
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	balanceFactor := getBalanceFactor(node)

	if balanceFactor > 1 {
		if key < node.Left.Key {
			return rightRotate(node)
		} else if key > node.Left.Key {
			node.Left = leftRotate(node.Left)
			return rightRotate(node)
		}
	}

	if balanceFactor < -1 {
		if key > node.Right.Key {
			return leftRotate(node)
		} else if key < node.Right.Key {
			node.Right = rightRotate(node.Right)
			return leftRotate(node)
		}
	}

	return node
}

func nodeWithMinimumValue(node *utils.Node) *utils.Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func DeleteNode(key string) bool {

	root := getRoot()
	root = delete(root, key)

	if err := putRoot(root); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true

}

func delete(root *utils.Node, key string) *utils.Node {

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
			temp := nodeWithMinimumValue(root.Right)
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
	balanceFactor := getBalanceFactor(root)

	if balanceFactor > 1 {
		if getBalanceFactor(root.Left) >= 0 {
			return rightRotate(root)
		} else {
			root.Left = leftRotate(root.Left)
			return rightRotate(root)
		}
	}
	if balanceFactor < -1 {
		if getBalanceFactor(root.Right) <= 0 {
			return leftRotate(root)
		} else {
			root.Right = rightRotate(root.Right)
			return leftRotate(root)
		}
	}

	return root
}

func GetOne(key string) (string, error) {
	root := getRoot()
	if node := get(root, key); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func get(node *utils.Node, key string) *utils.Node {
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

func GetAll() {
	root := getRoot()
	print(root, "", true)
}

func print(root *utils.Node, indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "   "
		} else {
			fmt.Print("L----")
			indent += "|  "
		}

		fmt.Println(" " + root.Key + "==>" + root.Value + " Timestamp: " + fmt.Sprint(root.Timestamp))
		print(root.Left, indent, false)
		print(root.Right, indent, true)
	}
}
