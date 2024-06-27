package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/Jatin020403/BasaltDB/utils"
)

func InsertOne(partition string, key uint64, value string) error {

	if partition == "" {
		return errors.New("InsertOne : invalid partition")
	}

	if key == 0 {
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

	root = insert(root, utils.Node{Key: key, Value: value, Timestamp: time.Now().UnixNano()})

	if err := putRoot(partition, part, root); err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	return nil
}

func DeleteOne(partition string, key uint64) error {

	if partition == "" {
		return errors.New("DeleteOne : invalid partition")
	}

	if key == 0 {
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

	root = delete(root, key)

	if err := putRoot(partition, part, root); err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	return nil

}

func GetOne(partition string, key uint64) (string, error) {

	part, err := utils.GetPartNumber(partition, key)
	if err != nil {
		return "", err
	}

	root, err := getRoot(partition, part)
	if err != nil {
		return "", err
	}

	if node := get(root, key); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func GetTree(partition string) {

	conf, err := utils.ReadConfig(partition)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range conf.PartsMap {
		root, err := getRoot(partition, k)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		printRoot(root, "", true)
	}
}
