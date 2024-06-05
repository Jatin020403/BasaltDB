package utils

import (
	"encoding/gob"
	"io"
	"os"

	"github.com/pkg/errors"
)

func CheckFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

// Create If Not Exist
func CINEPartition(partition string) error {
	var DIRPATH = "./storage/" + partition + ".gob"

	if isFileExist := CheckFileExists(DIRPATH); !isFileExist {
		if _, err := os.Create(DIRPATH); err != nil {
			return err
		}
	} else {
		return errors.New("partition does not exist")
	}

	return nil
}

// Delete If Not Exist
func DINEPartition(partition string) error {
	var DIRPATH = "./storage/" + partition + ".gob"

	if isFileExist := CheckFileExists(DIRPATH); isFileExist {
		if err := os.Remove(DIRPATH); err != nil {
			return err
		}
	} else {
		return errors.New("partition does not exist")
	}

	return nil
}

func Serialize(partition string, node *Node) error {
	var DIRPATH = "./storage/" + partition + ".gob"

	object := bsf(node)

	file, err := os.Create(DIRPATH)

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
	var DIRPATH = "./storage/" + partition + ".gob"

	file, err := os.Open(DIRPATH)

	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New(partition + " Partition does not exist")
		}

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
