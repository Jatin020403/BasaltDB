package handlers

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/Jatin020403/BasaltDB/utils"
)

func InsertOneHandler(partitionName string, key string, value string) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	hashedKey := utils.MurmurHashInt(key)

	if err = database.InsertOne(partition, hashedKey, value); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert successful")

}

func DeleteOneHandler(partitionName string, key string) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	hashedKey := utils.MurmurHashInt(key)

	if err = database.DeleteOne(partition, hashedKey); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("delete success")
	}

}

func GetOneHandler(partitionName string, key string) (string, error) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	hashedKey := utils.MurmurHashInt(key)

	part := int(hashedKey % uint64(partition.Conf.PartCount))

	root, err := utils.GetRoot(partition, part)
	if err != nil {
		return "", err
	}

	if node := utils.Get(root, hashedKey); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("not found")
}

func GetTreeHandler(partitionName string) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conf := partition.Conf

	for k := range conf.PartsMap {
		root, err := utils.GetRoot(partition, k)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		utils.PrintRoot(root, "", true)
	}
}
