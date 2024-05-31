package database

import (
	"container/heap"
	"fmt"
	"time"

	"github.com/Jatin020403/BasaltDB/utils"
)

func InsertLoop() {
	var lock bool
	lock = false
	t1 := time.Now()
	for {
		if !lock && utils.PQ.Len() > 0 {
			t1 = time.Now()
			lock = true
		}
		if utils.PQ.Len() > 20 || (time.Since(t1).Seconds() > 5 && lock) {
			MassInserts()
			lock = false
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Hiiiii " + fmt.Sprint(utils.PQ.Len()))
	}

}

func InsertOne(key string, value string) bool {

	root, err := getRoot()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	root = insert(root, utils.Node{Key: key, Value: value, Timestamp: time.Now().UnixNano()})

	if err := putRoot(root); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func InsertOneNew(key string, value string) bool {

	heap.Push(&utils.PQ, NewNode(key, value, time.Now().UnixNano()))

	return true
}

func MassInserts() bool {

	root, err := getRoot()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	root = insert_pq(root)

	if err := putRoot(root); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func DeleteNode(key string) bool {

	root, err := getRoot()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	root = delete(root, key)

	if err := putRoot(root); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true

}

func GetOne(key string) (string, error) {
	root, err := getRoot()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	if node := get(root, key); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func GetAll() {
	root, err := getRoot()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	print(root, "", true)
}
