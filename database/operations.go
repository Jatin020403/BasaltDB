package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/Jatin020403/BasaltDB/tree"
	"github.com/Jatin020403/BasaltDB/utils"
)

func (partition Partition) InsertOne(unhashedKey string, value string) error {

	if partition.Name == "" {
		return errors.New("InsertOne : invalid partition")
	}

	key := utils.MurmurHashInt(unhashedKey)

	if value == "" {
		return errors.New("InsertOne : invalid value")
	}

	part := int(key % uint64(partition.Conf.PartCount))
	object, err := partition.getData(part)
	if err != nil {
		return err
	}

	root, err := tree.Deserialize(object)
	if err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	root = tree.Insert(root, *tree.NewNode(key, unhashedKey, value, time.Now().UnixMilli()))

	object, err = tree.Serialize(root)

	if err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	if err := partition.putData(part, object); err != nil {
		return errors.New("InsertOne : " + err.Error())
	}

	return nil
}

func (partition Partition) DeleteOne(unhashedKey string) error {

	if partition.Name == "" {
		return errors.New("DeleteOne : invalid partition")
	}

	key := utils.MurmurHashInt(unhashedKey)

	part := int(key % uint64(partition.Conf.PartCount))

	object, err := partition.getData(part)
	if err != nil {
		return err
	}

	root, err := tree.Deserialize(object)
	if err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	root = tree.Delete(root, key)

	object, err = tree.Serialize(root)
	if err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	if err := partition.putData(part, object); err != nil {
		return errors.New("DeleteOne : " + err.Error())
	}

	return nil
}

func (partition Partition) GetOne(unhashedKey string) (string, error) {

	key := utils.MurmurHashInt(unhashedKey)
	part := int(key % uint64(partition.Conf.PartCount))
	object, err := partition.getData(part)
	if err != nil {
		return "", err
	}

	root, err := tree.Deserialize(object)
	if err != nil {
		return "", err
	}

	if node := tree.Get(root, key); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func (partition Partition) GetAll() {

	conf := partition.Conf

	for i := range conf.PartsMap {
		object, err := partition.getData(i)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		root, err := tree.Deserialize(object)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tree.PrintRoot(root, "", true)
	}
}
