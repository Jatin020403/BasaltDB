package handlers

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
)

func CreatePartitionHandler(partitionName string) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = partition.CreatePartition()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(partition.Conf)
	fmt.Println("Partition " + partitionName + " created")
}

func DeletePartitionHandler(partitionName string) {
	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = partition.DeletePartition()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Partition " + partitionName + " deleted")
}

func InitHandler(partitionName string, n int) {

	partition := database.InitPartition(partitionName, "")

	err := partition.CreateTemplate(n)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Partition " + partitionName + " initialised")
}

func InitDefaultHandler(partitionName string, n int) {

	partition := database.InitPartition(partitionName, "")

	err := partition.CreateTemplate(n)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = partition.CreatePartition()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(partition)

	fmt.Println("Partition " + partitionName + " initialised")
}

func RebalancePartitionHandler(partitionName string) {
	partition, err := database.CollectPartition(partitionName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = database.RebalancePartition(partition)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func RenamePartitionHandler(fromPartitionName string, toPartitionName string) {

	fromPartition, err := database.CollectPartition(fromPartitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	toPartition, err := database.CollectPartition(toPartitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	database.RenamePartition(fromPartition, toPartition)

	fromPartition.DeletePartition()
}
