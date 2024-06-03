package main

import (
	"github.com/Jatin020403/BasaltDB/cmd"
	"github.com/Jatin020403/BasaltDB/database"
)

func main() {
	database.Initialise()
	// go database.InsertLoop()
	cmd.Execute()

}
