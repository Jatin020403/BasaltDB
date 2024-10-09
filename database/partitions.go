package database

import (
	"errors"
	"os"
	"path"
)

type Partition struct {
	Name         string
	PartitionLoc string
	Conf         Config
}

type Config struct {
	PartCount int           `yaml:"partCount"`
	PartsMap  map[int]Parts `yaml:"parts"`
}

type Parts struct {
	PartLoc string `yaml:"loc"`
	Link    int    `yaml:"link"` // 0 - file, 1 - rpc
}

func InitPartition(name string, loc string) Partition {

	var partition Partition
	partition.Name = name
	partition.PartitionLoc = path.Join(loc, partition.Name)

	return partition
}

func CollectPartition(partitionName string) (Partition, error) {

	wd, err := os.Getwd()

	if err != nil {
		return Partition{}, err
	}

	partition := InitPartition(partitionName, wd)

	err = partition.LoadConfig("config.yaml")
	if err != nil {
		return Partition{}, nil
	}

	return partition, nil
}

func (partition *Partition) CreateTemplate(n int) error {

	var err error

	if err := os.Mkdir(partition.Name, os.ModePerm); err != nil {
		return err
	}

	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	partition.PartitionLoc = path.Join(wd, partition.Name)

	err = partition.InitialiseConfig(n)

	if err != nil {
		return errors.New("initConfig : " + err.Error())
	}

	err = partition.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (partition Partition) CreatePartition() error {

	err := CheckPathExists(partition.PartitionLoc)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("partition not initialised")
		}
		return err
	}

	conf := partition.Conf

	if conf.PartCount != len(conf.PartsMap) {
		return errors.New("length of parts not same as location")
	}

	err = InitialiseParts(conf)
	if err != nil {
		return errors.New("CreatePartition : " + err.Error())
	}
	// return errors.New("errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
	return nil
}

func (partition Partition) DeletePartition() error {
	err := CheckPathExists(partition.PartitionLoc)

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

	err = DeleteParts(conf)
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(partition.PartitionLoc))
	if err != nil {
		return err
	}

	return nil
}

func RebalancePartition(oldPartition Partition) error {

	var newPartition Partition
	var err error

	newPartition.Name = oldPartition.Name
	newPartition.PartitionLoc = oldPartition.PartitionLoc

	err = newPartition.LoadConfig("new_config.yaml")

	if err != nil {
		return err
	}

	err = oldPartition.CheckPartLocExists()

	if err != nil {
		return err
	}

	// Rename the old partition
	err = oldPartition.RenameInternalParts("old_" + oldPartition.Name)
	if err != nil {
		return err
	}

	err = newPartition.CreatePartition()
	if err != nil {
		return err
	}

	err = TransferPartitionData(oldPartition, newPartition)
	if err != nil {
		return err
	}

	err = DeleteParts(oldPartition.Conf)
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

func RenamePartition(fromPartition Partition, toPartition Partition) error {

	// rename external folder
	err := os.Rename(fromPartition.PartitionLoc, toPartition.PartitionLoc)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}

	// rename internal files
	err = fromPartition.RenameInternalParts(toPartition.Name)
	if err != nil {
		return err
	}

	return nil

}
