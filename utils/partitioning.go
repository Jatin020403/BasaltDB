package utils

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Parts struct {
	Loc string `yaml:"loc"`
}

type Config struct {
	PartCount int           `yaml:"partCount"`
	PartsMap  map[int]Parts `yaml:"parts"`
}

// Delete If Not Exist
func DINEPartition(partition string) error {
	err := CheckPathExists(partition)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("partition does not exist")
		}
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	confPath := path.Join(wd, partition, "config.yaml")
	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}

	var conf Config

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return err
	}

	if conf.PartCount != len(conf.PartsMap) {
		return errors.New("length of parts not same as location")
	}

	for _, v := range conf.PartsMap {

		if err := os.Remove(v.Loc); err != nil {
			return err
		}
	}

	err = os.RemoveAll(path.Join(wd, partition))
	if err != nil {
		return err
	}

	return nil
}

func InitialiseTemplate(partition string, n int) error {

	if err := os.Mkdir(partition, os.ModePerm); err != nil {
		return err
	}

	conf := Config{
		PartCount: n,
	}

	conf.PartsMap = make(map[int]Parts)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	for i := 0; i < conf.PartCount; i++ {
		part := partition + "_" + fmt.Sprint(i)
		path := path.Join(wd, partition, part+".gob")

		conf.PartsMap[i] = Parts{Loc: path}
	}

	err = WriteConfig(partition, conf)
	if err != nil {
		return err
	}
	return nil

}

func InitialisePartition(partition string) error {

	err := CheckPathExists(partition)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("partition not initialised")
		}
		return err
	}

	var conf Config

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	confPath := path.Join(wd, partition, "config.yaml")
	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return err
	}

	if conf.PartCount != len(conf.PartsMap) {
		return errors.New("length of parts not same as location")
	}

	for _, v := range conf.PartsMap {

		if _, err := os.Create(v.Loc); err != nil {
			return err
		}
	}

	return nil
}

func ReadConfig(partition string) (Config, error) {
	var conf Config

	wd, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}
	confPath := path.Join(wd, partition, "config.yaml")

	f, err := os.ReadFile(confPath)

	if err != nil {
		return Config{}, err
	}

	if err := yaml.Unmarshal(f, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}

func WriteConfig(partition string, conf Config) error {

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	out, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	confPath := path.Join(wd, partition, "config.yaml")
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

func GetPartNumber(partition string, key uint64) (int, error) {
	conf, err := ReadConfig(partition)
	if err != nil {
		return 0, err
	}

	return int(key % uint64(conf.PartCount)), nil
}

func GetPathFromPart(partition string, part int) (string, error) {
	conf, err := ReadConfig(partition)
	if err != nil {
		return "", err
	}

	path := conf.PartsMap[part].Loc

	return path, err
}
