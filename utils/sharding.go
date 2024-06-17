package utils

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type shards struct {
	Loc string `yaml:"loc"`
}

type Config struct {
	ShardCount int      `yaml:"shardCount"`
	Shards     []shards `yaml:"shards"`
}

var defaultShardcount int = 5

func InitialiseShard(partition string) error {

	for i := 0; i < defaultShardcount; i++ {
		if _, err := os.Create("./" + partition + "/" + partition + "_" + fmt.Sprint(i) + ".gob"); err != nil {
			return err
		}
	}

	c := Config{
		ShardCount: defaultShardcount,
	}

	out, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("./" + partition + "/" + "config.yaml")
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = io.WriteString(file, string(out))
	if err != nil {
		return err
	}

	return nil
}

func GetConfig(partition string) (Config, error) {
	var CONFIGPATH = "./" + partition + "/config.yaml"
	f, err := os.ReadFile(CONFIGPATH)
	var conf Config

	if err != nil {
		return Config{}, err
	}

	if err := yaml.Unmarshal(f, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}

func GetShard(partition string, key string) (int, error) {
	conf, err := GetConfig(partition)
	if err != nil {
		return 0, err
	}

	return int(MurmurHashInt(key) % uint64(conf.ShardCount)), nil
}
