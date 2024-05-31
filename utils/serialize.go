package utils

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
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

	if errors.Is(err, io.EOF) {
		fmt.Println("File Empty")
		return []ArrNode{}, nil
	}

	if err != nil {
		err = errors.New(err.Error() + " : Deserialisation failed")
		file.Close()
		return nil, err
	}

	file.Close()

	return object, nil
}
