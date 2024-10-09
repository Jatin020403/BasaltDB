package database

import (
	"encoding/gob"
	"errors"
	"io"
	"os"

	"github.com/Jatin020403/BasaltDB/data"
)

func (partition Partition) putData(part int, object []data.DataNode) error {

	// Check Partition exists
	if err := CheckPathExists(partition.PartitionLoc); err != nil {
		if os.IsNotExist(err) {
			return errors.New(partition.Name + " partition does not exist")
		}

		return err
	}

	path := partition.Conf.PartsMap[part].PartLoc

	// Check file exists
	if err := CheckPathExists(path); err != nil {
		if os.IsNotExist(err) {
			return errors.New(path + " does not exist")
		} else {
			return err
		}
	}

	file, err := os.Create(path)

	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(&object)
		if err != nil {
			return errors.New("serialisation : " + err.Error())
		}
	}
	defer file.Close()
	return err
}

func (partition Partition) getData(part int) ([]data.DataNode, error) {

	var err error

	// Check Partition exists
	if err = CheckPathExists(partition.PartitionLoc); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(partition.Name + " partition does not exist")
		}
		return nil, err
	}

	pathPath := partition.Conf.PartsMap[part].PartLoc

	// Check file exists
	if err := CheckPathExists(pathPath); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(pathPath + " does not exist")
		} else {
			return nil, err
		}
	}

	file, err := os.Open(pathPath)

	if err != nil {
		return nil, errors.New("deserialisation : " + err.Error())
	}

	decoder := gob.NewDecoder(file)
	var object []data.DataNode
	err = decoder.Decode(&object)

	if errors.Is(err, io.EOF) {
		return []data.DataNode{}, nil
	}

	if err != nil {
		err = errors.New("deserialisation : " + err.Error())
		file.Close()
		return nil, err
	}

	file.Close()

	return object, nil
}

func CheckPathExists(path string) error {
	_, err := os.Stat(path)
	return err
}
