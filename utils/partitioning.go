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
	PartCount int              `yaml:"partCount"`
	PartsMap  map[string]Parts `yaml:"parts"`
}

var defaultPartcount int = 3

// Create If Not Exist
func CINETemplate(partition string) error {

	err := CheckPathExists(partition)
	if err == nil {
		return errors.New("partition already exists")
	}

	if !os.IsNotExist(err) {
		return err
	}

	err = InitialiseTemplate(partition)

	return err
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

func InitialiseTemplate(partition string) error {

	if err := os.Mkdir(partition, os.ModePerm); err != nil {
		return err
	}

	conf := Config{
		PartCount: defaultPartcount,
	}

	conf.PartsMap = make(map[string]Parts)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	for i := 0; i < conf.PartCount; i++ {
		part := partition + "_" + fmt.Sprint(i)
		path := path.Join(wd, partition, part+".gob")

		conf.PartsMap[part] = Parts{Loc: path}
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

func GetPartNumber(partition string, key string) (int, error) {
	conf, err := GetConfig(partition)
	if err != nil {
		return 0, err
	}

	return int(MurmurHashInt(key) % uint64(conf.PartCount)), nil
}

func GetPathFromPart(partition string, part int) (string, error) {
	conf, err := GetConfig(partition)
	if err != nil {
		return "", err
	}

	path := conf.PartsMap[partition+"_"+fmt.Sprint(part)].Loc

	return path, err
}

func GetConfig(partition string) (Config, error) {
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
