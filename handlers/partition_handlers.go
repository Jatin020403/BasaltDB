package handlers

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/Jatin020403/BasaltDB/models"
)

func CreatePartitionHandler(partitionName string) {

	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = database.CreatePartition(partition)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Partition " + partitionName + " created")

}

func DeletePartitionHandler(partitionName string) {
	partition, err := database.CollectPartition(partitionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = database.DeletePartition(partition)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Partition " + partitionName + " deleted")
}

func InitHandler(partitionName string, n int) {

	var partition models.Partition
	partition.Name = partitionName

	partition, err := database.CreateTemplate(partition, n)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Partition " + partitionName + " initialised")
}

func InitDefaultHandler(partitionName string, n int) {

	var partition models.Partition
	partition.Name = partitionName

	partition, err := database.CreateTemplate(partition, n)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	partition, err = database.CreatePartition(partition)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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

	database.DeletePartition(fromPartition)
}
