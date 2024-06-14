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

func CINEFile(partition string) error {
	var DIRPATH = "./" + partition + "/" + partition + ".gob"

	err := CheckPathExists(DIRPATH)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if _, err = os.Create(DIRPATH); err != nil {
		return err
	}

	return nil

}

// Create If Not Exist
func CINEPartition(partition string) error {
	var DIR = partition

	err := CheckPathExists(DIR)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if err = os.Mkdir(DIR, os.ModePerm); err != nil {
		return err
	}

	if err = CINEFile(partition); err != nil {
		return err
	}

	return nil
}

// Delete If Not Exist
func DINEPartition(partition string) error {
	var DIR = partition

	err := CheckPathExists(DIR)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if err = os.RemoveAll(DIR); err != nil {
		return err
	}

	return nil
}

func Serialize(partition string, node *Node) error {

	// Check Partition exists
	if err := CheckPathExists(partition); err != nil {
		return err
	}

	var DIRPATH = "./" + partition + "/" + partition + ".gob"

	// Check file exists
	if err := CheckPathExists(DIRPATH); err != nil {
		if os.IsNotExist(err) {
			err = CINEFile(DIRPATH)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

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
	// Check Partition exists
	if err := CheckPathExists(partition); err != nil {
		return nil, err
	}

	var DIRPATH = "./" + partition + "/" + partition + ".gob"

	// Check file exists
	if err := CheckPathExists(DIRPATH); err != nil {
		if os.IsNotExist(err) {
			err = CINEFile(DIRPATH)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

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
