package utils

import (
	"encoding/gob"
	"fmt"
	"os"
)

const DIRPATH = "./storage/bst.gob"

func Serialize(node *Node) error {

	object := bsf(node)

	file, err := os.Create(DIRPATH)
	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(&object)
		if err != nil {
			return err
		}
	}
	defer file.Close()
	return err
}

func Deserialize(object []ArrNode) ([]ArrNode, error) {
	file, err := os.Open(DIRPATH)
	if err != nil {
		return nil, err
	}

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&object)

	if err != nil {
		err = fmt.Errorf(err.Error() + "Hii")
		file.Close()
		return nil, err
	}

	file.Close()

	return object, nil
}
