package database

import (
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

	part, err := utils.GetPartNumber(partition, key)
	if err != nil {
		return err
	}

	root, err := getRoot(partition, part)
	if err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	hashedKey := utils.MurmurHashMod(key)
	root = insert(root, utils.Node{Key: hashedKey, Value: value, Timestamp: time.Now().UnixNano()})

	if err := putRoot(partition, part, root); err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	return nil
}

func DeleteOne(partition string, key string) error {

	if partition == "" {
		return errors.New("DeleteOne : invalid partition")
	}

	if key == "" {
		return errors.New("DeleteOne : invalid key")
	}

	part, err := utils.GetPartNumber(partition, key)
	if err != nil {
		return err
	}

	root, err := getRoot(partition, part)
	if err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	hashedKey := utils.MurmurHashMod(key)
	root = delete(root, hashedKey)

	if err := putRoot(partition, part, root); err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	return nil

}

func GetOne(partition string, key string) (string, error) {

	part, err := utils.GetPartNumber(partition, key)
	if err != nil {
		return "", err
	}

	root, err := getRoot(partition, part)
	if err != nil {
		return "", err
	}

	hashedKey := utils.MurmurHashMod(key)
	if node := get(root, hashedKey); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func GetTree(partition string) {

	root, err := getAllRoot(partition)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	print(root, "", true)
}
