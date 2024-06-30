package models

type Partition struct {
	Name string
	PartitionLoc string
	Conf Config
}

type Config struct {
	PartCount int           `yaml:"partCount"`
	PartsMap  map[int]Parts `yaml:"parts"`
}

type Parts struct {
	PartLoc  string `yaml:"loc"`
	Link int    `yaml:"link"` // 0 - file, 1 - rpc
}
