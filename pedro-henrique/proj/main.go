package main

import ( "pedro-henrique/proj/database"
		 "pedro-henrique/proj/service" )

func main() {

	// assure server closing at end of execution
	defer service.StopServer()

	// call db client constructor
	database.CreateClient()

	// call start server function
	service.StartServer()
}