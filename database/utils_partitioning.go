package database

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/Jatin020403/BasaltDB/data"
	"gopkg.in/yaml.v2"
)

func (partition *Partition) assignConf(conf Config) {
	partition.Conf = conf
}

func InitialiseParts(conf Config) error {

	for _, v := range conf.PartsMap {
		if _, err := os.Create(v.PartLoc); err != nil {
			return err
		}
	}

	return nil
}

func DeleteParts(conf Config) error {

	for _, v := range conf.PartsMap {

		if err := os.Remove(v.PartLoc); err != nil {
			return err
		}
	}

	return nil
}

func (partition *Partition) InitialiseConfig(n int) error {

	conf := Config{
		PartCount: n,
	}

	conf.PartsMap = make(map[int]Parts)

	for i := 0; i < conf.PartCount; i++ {
		part := partition.Name + "_" + fmt.Sprint(i)
		path := path.Join(partition.PartitionLoc, part+".gob")
		conf.PartsMap[i] = Parts{PartLoc: path}
	}

	partition.assignConf(conf)

	return nil
}

func (partition *Partition) LoadConfig(confName string) error {

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	confPath := path.Join(wd, partition.Name, confName)

	f, err := os.ReadFile(confPath)

	if err != nil {
		return err
	}

	var conf Config

	if err = yaml.Unmarshal(f, &conf); err != nil {
		return err
	}

	partition.assignConf(conf)
	// partition.Conf = conf

	return nil
}

func (partition Partition) WriteConfig() error {

	out, err := yaml.Marshal(partition.Conf)
	if err != nil {
		return err
	}

	confPath := path.Join(partition.PartitionLoc, "config.yaml")
	confFile, err := os.Create(confPath)
	if err != nil {
		return err
	}
	defer confFile.Close()

	_, err = io.WriteString(confFile, string(out))
	if err != nil {
		return err
	}

	return nil
}

func TransferPartitionData(fromPartition Partition, toPartition Partition) error {

	for k := range fromPartition.Conf.PartsMap {

		fromData, err := fromPartition.getData(k)

		if err != nil {
			return err
		}

		toDataLength := toPartition.Conf.PartCount
		var toDataPartition = make([][]data.DataNode, toDataLength)

		for i := range toDataPartition {
			toDataPartition[i] = make([]data.DataNode, 0)
		}

		for _, data := range fromData {
			part := int(data.Key % uint64(toDataLength))
			toDataPartition[part] = append(toDataPartition[part], data)
		}

		for part, dataArr := range toDataPartition {

			err = toPartition.putData(part, dataArr)

			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (partition Partition) CheckPartLocExists() error {
	for _, v := range partition.Conf.PartsMap {
		err := CheckPathExists(path.Dir(v.PartLoc))
		if err != nil {
			return err
		}
	}
	return nil
}

func (partition *Partition) RenameInternalParts(toName string) error {

	for k, v := range partition.Conf.PartsMap {
		fromPath := v.PartLoc
		toPath := path.Join(path.Dir(fromPath), toName+"_"+fmt.Sprint(k)+".gob")
		err := os.Rename(fromPath, toPath)
		if err != nil {
			return err
		}
		partition.Conf.PartsMap[k] = Parts{PartLoc: toPath}
	}

	err := partition.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

/*
func PartitionInsertTree(rootOld *Node, rootNew *Node) error {
	if root == nil {
		return nil
	}

	var err error

	part := int(root.key % uint64(partition.Conf.PartCount))

	root, err = tree.Deserialize(partition, part)
	if err != nil {
		return err
	}

	err = partition.PartitionInsertTree(root.Left)
	if err != nil {
		return err
	}

	err = partition.PartitionInsertTree(root.Right)
	if err != nil {
		return err
	}

	return nil
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

*/
