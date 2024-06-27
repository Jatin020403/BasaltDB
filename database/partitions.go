package database

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Jatin020403/BasaltDB/utils"
)

func CreateTemplate(partitionName string, n int) error {
	err := utils.InitialiseTemplate(partitionName, n)
	if err != nil {
		return errors.New("CreatePartition : " + err.Error())
	}
	return nil
}

func CreatePartition(partitionName string) error {
	err := utils.InitialisePartition(partitionName)
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

func RebalancePartition(partition string) error {

	oldPartition := "old_" + partition

	// Init the partition

	var err error
	err = CreateTemplate(oldPartition, 1)
	if err != nil {
		return err
	}

	// place the config

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	oldConfLocation := path.Join(wd, oldPartition, "config.yaml")
	confLocation := path.Join(wd, partition, "config.yaml")

	err = os.Rename(confLocation, oldConfLocation)
	if err != nil {
		return err
	}

	newConfLocation := path.Join(wd, partition, "new_config.yaml")

	err = os.Rename(newConfLocation, confLocation)
	if err != nil {
		return err
	}

	// Rename the old partition
	err = RenamePartition(partition, oldPartition)
	if err != nil {
		return err
	}

	// check file exists

	CheckPartLocExists(partition)

	// create old partition

	err = CreatePartition(partition)
	if err != nil {
		return err
	}

	// transfer data from old to old

	err = TransferPartitionData(oldPartition, partition)
	if err != nil {
		return err
	}

	// delete the old partition

	err = DeletePartition(oldPartition)
	if err != nil {
		return err
	}

	return nil
}

func TransferPartitionData(fromPartition string, toPartition string) error {

	fromConf, err := utils.ReadConfig(fromPartition)
	if err != nil {
		return err
	}

	for k := range fromConf.PartsMap {
		fromRoot, err := getRoot(fromPartition, k)
		if err != nil {
			return err
		}

		err = PartitionInsertTree(toPartition, fromRoot)
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckPartLocExists(partition string) error {

	conf, err := utils.ReadConfig(partition)
	if err != nil {
		return err
	}

	for _, v := range conf.PartsMap {
		err = utils.CheckPathExists(path.Dir(v.Loc))
		if err != nil {
			return err
		}
	}
	return nil
}

func RenamePartition(fromPartition string, toPartition string) error {

	// rename external folder
	err := os.Rename(fromPartition, toPartition)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}

	// rename internal files
	toConf, err := utils.ReadConfig(toPartition)
	if err != nil {
		return err
	}

	for k, v := range toConf.PartsMap {
		fromPath := v.Loc
		toPath := path.Join(path.Dir(fromPath), toPartition+"_"+fmt.Sprint(k)+".gob")
		err := os.Rename(fromPath, toPath)
		if err != nil {
			return err
		}
		toConf.PartsMap[k] = utils.Parts{Loc: toPath}
	}

	utils.WriteConfig(toPartition, toConf)

	return nil

}

func PartitionInsertTree(partition string, root *utils.Node) error {
	if root == nil {
		return nil
	}

	var err error

	err = InsertOne(partition, root.Key, root.Value)
	if err != nil {
		return err
	}

	err = PartitionInsertTree(partition, root.Left)
	if err != nil {
		return err
	}

	err = PartitionInsertTree(partition, root.Right)
	if err != nil {
		return err
	}

	return nil
}
