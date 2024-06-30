package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/Jatin020403/BasaltDB/models"
	"github.com/Jatin020403/BasaltDB/utils"
)

func InsertOne(partition models.Partition, key uint64, value string) error {

	if partition.Name == "" {
		return errors.New("InsertOne : invalid partition")
	}

	if key == 0 {
		return errors.New("InsertOne : invalid key")
	}

	if value == "" {
		return errors.New("InsertOne : invalid value")
	}

	part := int(key % uint64(partition.Conf.PartCount))

	root, err := utils.GetRoot(partition, part)
	if err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	root = utils.Insert(root, models.Node{Key: key, Value: value, Timestamp: time.Now().UnixNano()})

	if err := utils.PutRoot(partition, part, root); err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	return nil
}

func DeleteOne(partition models.Partition, key uint64) error {

	if partition.Name == "" {
		return errors.New("DeleteOne : invalid partition")
	}

	part := int(key % uint64(partition.Conf.PartCount))

	root, err := utils.GetRoot(partition, part)
	if err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	root = utils.Delete(root, key)

	if err := utils.PutRoot(partition, part, root); err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	return nil

}

func GetOne(partition models.Partition, key uint64) (string, error) {

	part := int(key % uint64(partition.Conf.PartCount))

	root, err := utils.GetRoot(partition, part)
	if err != nil {
		return "", err
	}

	if node := utils.Get(root, key); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func GetTree(partition models.Partition) {

	conf := partition.Conf

	for k := range conf.PartsMap {
		root, err := utils.GetRoot(partition, k)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		utils.PrintRoot(root, "", true)
	}
}
