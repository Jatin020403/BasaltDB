package database

import (
	"container/heap"
	"errors"
	"fmt"
	"time"

	"github.com/Jatin020403/BasaltDB/utils"
)

/*
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
*/

func InsertOne(partition string, key string, value string) error {

	if partition == "" {
		return errors.New("InsertOne : invalid partition")
	}

	if key == "" {
		return errors.New("InsertOne : invalid key")
	}

	if value == "" {
		return errors.New("InsertOne : invalid value")
	}

	root, err := getRoot(partition)
	if err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	hashedKey := utils.MurmurHash(key)
	root = insert(root, utils.Node{Key: hashedKey, Value: value, Timestamp: time.Now().UnixNano()})

	if err := putRoot(partition, root); err != nil {
		fmt.Println(err.Error())
		return errors.New("InsertOne : " + err.Error())
	}

	return nil
}

func InsertOneNew(key string, value string) bool {
	hashedKey := utils.MurmurHash(key)

	heap.Push(&utils.PQ, NewNode(hashedKey, value, time.Now().UnixNano()))

	return true
}

func MassInserts(partition string) bool {

	root, err := getRoot(partition)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	root = insert_pq(root)

	if err := putRoot(partition, root); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func DeleteOne(partition string, key string) error {

	if partition == "" {
		return errors.New("DeleteOne : invalid partition")
	}

	if key == "" {
		return errors.New("DeleteOne : invalid key")
	}

	root, err := getRoot(partition)
	if err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}
	hashedKey := utils.MurmurHash(key)
	root = delete(root, hashedKey)

	if err := putRoot(partition, root); err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

		return nil

}

func GetOne(partition string, key string) (string, error) {
	root, err := getRoot(partition)
	if err != nil {
		return "", err
	}
	hashedKey := utils.MurmurHash(key)
	if node := get(root, hashedKey); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func GetTree(partition string) {
	root, err := getRoot(partition)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	print(root, "", true)
}
