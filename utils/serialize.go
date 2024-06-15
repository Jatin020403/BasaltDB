package utils

import (
	"encoding/gob"
	"io"
	"os"

	"github.com/pkg/errors"
)

func CheckPathExists(path string) error {
	_, err := os.Stat(path)
	return err
}

// Create If Not Exist
func CINEPartition(partition string) error {

	err := CheckPathExists(partition)
	if err == nil {
		return errors.New("partition already exists")
	}
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if err = os.Mkdir(partition, os.ModePerm); err != nil {
		return err
	}

	PATH := "./" + partition + "/" + partition + ".gob"

	if _,err = os.Create(PATH); err != nil {
		return err
	}

	// Initialise_shard(partition)

	return nil
}

// Delete If Not Exist
func DINEPartition(partition string) error {
	err := CheckPathExists(partition)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("partition does not exist")
		}
		return err
	}

	if err = os.RemoveAll(partition); err != nil {
		return err
	}

	return nil
}

func Serialize(partition string, node *Node) error {

	// Check Partition exists
	if err := CheckPathExists(partition); err != nil {
		if os.IsNotExist(err) {
			return errors.New(partition + " partition does not exist")
		}

		return err
	}

	var PATH = "./" + partition + "/" + partition + ".gob"

	// Check file exists
	if err := CheckPathExists(PATH); err != nil {
		if os.IsNotExist(err) {
			return errors.New(partition + " shard does not exist")
		} else {
			return err
		}
	}

	object := bsf(node)

	file, err := os.Create(PATH)

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

func Deserialize(partition string, object []ArrNode) ([]ArrNode, error) {
	// Check Partition exists
	if err := CheckPathExists(partition); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(partition + " partition does not exist")
		}
		return nil, err
	}

	var PATH = "./" + partition + "/" + partition + ".gob"

	// Check file exists
	if err := CheckPathExists(PATH); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(partition + " shard does not exist")
		} else {
			return nil, err
		}
	}

	file, err := os.Open(PATH)

	if err != nil {
		return nil, errors.New("deserialisation : " + err.Error())
	}

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&object)

	if errors.Is(err, io.EOF) {
		return []ArrNode{}, nil
	}

	if err != nil {
		err = errors.New("deserialisation : " + err.Error())
		file.Close()
		return nil, err
	}

	file.Close()

	return object, nil
}
