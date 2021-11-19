package db

import "dep-test/models"

//  here we want to implement the datastore interface

type DBSaver struct {
}

func NewDB(dbURL string) *DBSaver {
	// use the url to try and connect to the db
	return &DBSaver{}
}
// this line acts a code editor validation so that we know that our struct implements the interface
// what I'm doing here is creating a empty variable of the same type as the interface
// then we try to assign this to our struct
// if the struct implements the interface, then there would be no error
var _ Datastore = &DBSaver{}

func (d *DBSaver) GetUser(name string) (*models.User, error) {
	return nil, nil
}

func (d *DBSaver) SaveUser(*models.User) error {
	return nil
}