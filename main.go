package main

import (
	"dep-test/db"
	"dep-test/usecase"
)

func main() {
	// to make use of the functionality from the use case package
	// we just need to create a new instance of the use case struct
	// the methods were added as receivers to this struct 
	//  for the ones that are exported (capital letter) they would be available here

	// but before just creating an instance of the struct, we need to add the dependencies
	// if not we would get an error when we try to use this

	// we need to put in a struct that implements the interface

	// dbSaver := db.DBSaver{}
	dbSaver := db.NewDB("databaseUrl")

	//  we have just done dependency injection for this usecase handler
	//  by putting in an implementation of the store

	// handler := usecase.Usecase{
	// 	Store: &dbSaver,
	// }

	handler := usecase.NewUseCase(dbSaver)

	// 
	handler.Createuser("moses")

	handler.Createuser("chuks")


}