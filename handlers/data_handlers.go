package handlers

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
)

func InsertOneHandler(partitionName string, key string, value string) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = partition.InsertOne(key, value); err != nil {
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

	if err = partition.DeleteOne(key); err != nil {
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

	val, err := partition.GetOne(key)

	if err != nil {
		return "", fmt.Errorf("not found")
	}
	return val, nil
}

func GetAllHandler(partitionName string) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	partition.GetAll()
}
