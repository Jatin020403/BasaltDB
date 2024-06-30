package utils

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/Jatin020403/BasaltDB/models"
	"gopkg.in/yaml.v2"
)

func InitialiseConfig(partition models.Partition, n int) (models.Partition, error) {

	conf := models.Config{
		PartCount: n,
	}

	conf.PartsMap = make(map[int]models.Parts)

	for i := 0; i < conf.PartCount; i++ {
		part := partition.Name + "_" + fmt.Sprint(i)
		path := path.Join(partition.PartitionLoc, part+".gob")

		conf.PartsMap[i] = models.Parts{PartLoc: path}
	}

	partition.Conf = conf

	return partition, nil

}

func InitialiseParts(conf models.Config) error {

	for _, v := range conf.PartsMap {

		if _, err := os.Create(v.PartLoc); err != nil {
			return err
		}
	}

	return nil
}

func DeleteParts(conf models.Config) error {

	for _, v := range conf.PartsMap {

		if err := os.Remove(v.PartLoc); err != nil {
			return err
		}
	}

	return nil
}

func LoadConfig(partition models.Partition, confName string) (models.Partition, error) {

	wd, err := os.Getwd()
	if err != nil {
		return models.Partition{}, err
	}
	confPath := path.Join(wd, partition.Name, confName)

	f, err := os.ReadFile(confPath)

	if err != nil {
		return models.Partition{}, err
	}

	if err := yaml.Unmarshal(f, &partition.Conf); err != nil {
		return models.Partition{}, err
	}

	return partition, nil
}

func WriteConfig(partition models.Partition) error {

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
