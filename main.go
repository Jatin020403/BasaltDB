package main

import (
	"github.com/Jatin020403/BasaltDB/cmd"
	"github.com/Jatin020403/BasaltDB/utils"
)

func main() {
	utils.InitialiseQueue()
	// go database.InsertLoop()
	cmd.Execute()
	// database.GetAll()

}
