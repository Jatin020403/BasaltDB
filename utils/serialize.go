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

func Serialize(partition string, part int, node *Node) error {

	// Check Partition exists
	if err := CheckPathExists(partition); err != nil {
		if os.IsNotExist(err) {
			return errors.New(partition + " partition does not exist")
		}

		return err
	}

	path, err := GetPathFromPart(partition, part)
	if err != nil {
		return err
	}

	// Check file exists
	if err := CheckPathExists(path); err != nil {
		if os.IsNotExist(err) {
			return errors.New(path + " does not exist")
		} else {
			return err
		}
	}

	object := bsf(node)

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

func Deserialize(partition string, part int, object []ArrNode) ([]ArrNode, error) {

	var err error

	// Check Partition exists
	if err = CheckPathExists(partition); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(partition + " partition does not exist")
		}
		return nil, err
	}

	path, err := GetPathFromPart(partition, part)
	if err != nil {
		return nil, err
	}

	// Check file exists
	if err := CheckPathExists(path); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(path + " does not exist")
		} else {
			return nil, err
		}
	}

	file, err := os.Open(path)

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
