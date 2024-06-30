package database

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Jatin020403/BasaltDB/models"
	"github.com/Jatin020403/BasaltDB/utils"
)

func CollectPartition(partitionName string) (models.Partition, error) {

	var partition models.Partition
	var err error

	partition.Name = partitionName

	wd, err := os.Getwd()

	if err != nil {
		return models.Partition{}, nil
	}

	partition.PartitionLoc = path.Join(wd, partition.Name)

	partition, err = utils.LoadConfig(partition, "config.yaml")
	if err != nil {
		return models.Partition{}, nil
	}

	return partition, nil
}

func CreateTemplate(partition models.Partition, n int) (models.Partition, error) {

	var err error

	if err := os.Mkdir(partition.Name, os.ModePerm); err != nil {
		return partition, err
	}

	wd, err := os.Getwd()

	if err != nil {
		return partition, err
	}

	partition.PartitionLoc = path.Join(wd, partition.Name)

	partition, err = utils.InitialiseConfig(partition, n)

	if err != nil {
		return partition, errors.New("initConfig : " + err.Error())
	}

	err = utils.WriteConfig(partition)
	if err != nil {
		return partition, err
	}
	return partition, nil
}

func CreatePartition(partition models.Partition) (models.Partition, error) {

	err := utils.CheckPathExists(partition.PartitionLoc)
	if err != nil {
		if os.IsNotExist(err) {
			return partition, errors.New("partition not initialised")
		}
		return partition, err
	}

	conf := partition.Conf

	if conf.PartCount != len(conf.PartsMap) {
		return partition, errors.New("length of parts not same as location")
	}

	err = utils.InitialiseParts(conf)
	if err != nil {
		return partition, errors.New("CreatePartition : " + err.Error())
	}
	return partition, nil
}

func DeletePartition(partition models.Partition) error {
	err := utils.CheckPathExists(partition.PartitionLoc)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("partition does not exist")
		}
		return err
	}

	conf := partition.Conf

	if conf.PartCount != len(conf.PartsMap) {
		return errors.New("length of parts not same as location")
	}

	err = utils.DeleteParts(conf)
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(partition.PartitionLoc))
	if err != nil {
		return err
	}

	return nil
}

func RebalancePartition(oldPartition models.Partition) error {

	var newPartition models.Partition
	var err error

	newPartition.Name = oldPartition.Name
	newPartition.PartitionLoc = oldPartition.PartitionLoc

	newPartition, err = utils.LoadConfig(newPartition, "new_config.yaml")

	if err != nil {
		return err
	}

	err = CheckPartLocExists(oldPartition)

	if err != nil {
		return err
	}

	// Rename the old partition
	oldPartition, err = RenameInternalParts(oldPartition, "old_"+oldPartition.Name)
	if err != nil {
		return err
	}

	newPartition, err = CreatePartition(newPartition)
	if err != nil {
		return err
	}

	err = TransferPartitionData(oldPartition, newPartition)
	if err != nil {
		return err
	}

	err = utils.DeleteParts(oldPartition.Conf)
	if err != nil {
		return err
	}

	oldConfName := path.Join(newPartition.PartitionLoc, "new_config.yaml")
	newConfName := path.Join(newPartition.PartitionLoc, "config.yaml")
	err = os.Rename(oldConfName, newConfName)
	if err != nil {
		return err
	}

	return nil

}

/*
func RebalancePartition(partition models.Partition) error {

	var oldPartition models.Partition
	var err error

	oldPartition.Name = "old_" + partition.Name
	oldPartition.PartitionLoc = path.Join(path.Dir(partition.PartitionLoc), "old_"+partition.Name)
	oldPartition.Conf = partition.Conf

	oldConfLocation := path.Join(oldPartition.PartitionLoc, "config.yaml")
	confLocation := path.Join(partition.PartitionLoc, "config.yaml")
	newConfLocation := path.Join(partition.PartitionLoc, "new_config.yaml")

	err = CheckPartLocExists(oldPartition)
	if err != nil {
		return err
	}

	partition, err = utils.LoadConfig(partition, "new_config.yaml")
	if err != nil {
		return err
	}

	err = CheckPartLocExists(partition)
	if err != nil {
		return err
	}

	// Init the partition

	oldPartition, err = CreateTemplate(oldPartition, 1)
	if err != nil {
		return err
	}

	// exchange configs

	err = os.Rename(confLocation, oldConfLocation)
	if err != nil {
		return err
	}

	err = os.Rename(newConfLocation, confLocation)
	if err != nil {
		return err
	}

	// create old partition

	oldPartition, err = CreatePartition(oldPartition)
	if err != nil {
		return err
	}

	// Rename the old partition
	err = RenamePartition(partition, oldPartition)
	if err != nil {
		return err
	}

	// transfer data from old to new

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
*/

func TransferPartitionData(fromPartition models.Partition, toPartition models.Partition) error {

	for k := range fromPartition.Conf.PartsMap {
		fromRoot, err := utils.GetRoot(fromPartition, k)
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

func CheckPartLocExists(partition models.Partition) error {
	for _, v := range partition.Conf.PartsMap {
		err := utils.CheckPathExists(path.Dir(v.PartLoc))
		if err != nil {
			return err
		}
	}
	return nil
}

func RenamePartition(fromPartition models.Partition, toPartition models.Partition) error {

	// rename external folder
	err := os.Rename(fromPartition.PartitionLoc, toPartition.PartitionLoc)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}

	// rename internal files
	_, err = RenameInternalParts(fromPartition, toPartition.Name)
	if err != nil {
		return err
	}

	return nil

}

func RenameInternalParts(partition models.Partition, toName string) (models.Partition, error) {

	for k, v := range partition.Conf.PartsMap {
		fromPath := v.PartLoc
		toPath := path.Join(path.Dir(fromPath), toName+"_"+fmt.Sprint(k)+".gob")
		err := os.Rename(fromPath, toPath)
		if err != nil {
			return models.Partition{}, err
		}
		partition.Conf.PartsMap[k] = models.Parts{PartLoc: toPath}
	}

	utils.WriteConfig(partition)
	return partition, nil
}

func PartitionInsertTree(partition models.Partition, root *models.Node) error {
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
