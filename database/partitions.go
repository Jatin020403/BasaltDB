package database

import (
	"errors"

	"github.com/Jatin020403/BasaltDB/utils"
)

func CreatePartition(partitionName string) error {
	err := utils.CINEPartition(partitionName)
	if err != nil {
		return errors.New("CreatePartition : " + err.Error())
	}
	return nil
}

func DeletePartition(partitionName string) error {
	err := utils.DINEPartition(partitionName)
	if err != nil {
		return errors.New("DeletePartition : " + err.Error())
	}
	return nil
}

/*
func GetAllPartitions() ([]string, error) {
	files, err := filepath.Glob("storage/*.gob")
	if err != nil {
		return []string{}, errors.New("GetAllPartitions : " + err.Error())
	}

	if len(files) == 0 {
		return []string{}, errors.New("GetAllPartitions : no partitions")
	}

	var pt []string

	for _, file := range files {
		pt = append(pt, strings.Split((strings.Split(file, "/")[1]), ".gob")[0])
	}
	return pt, nil
}
*/
