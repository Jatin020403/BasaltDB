package utils

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type shards struct{
	loc string `yaml:"loc"`

}

type defaultConfig struct {
	ShardCount int8	`yaml:"shardCount"`
	Shards []shards `yaml:"shards"`
}

var defaultShardcount int = 3

func Initialise_shard(partition string) error {

	for i := 0; i < defaultShardcount; i++ {
		if _, err := os.Create("./" + partition + "/" + partition + "_" + fmt.Sprint(i)); err != nil {
			return err
		}
	}

	c := defaultConfig{
		ShardCount: int8(defaultShardcount),
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

	fmt.Println(out)

	return nil
}
