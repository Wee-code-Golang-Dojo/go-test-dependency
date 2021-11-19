package db

import (
	"dep-test/models"
)

// interfaces are used to define behaviour
// it acts a contract between two different packages
// so you can have different structs that can implement this datastore interface
// The package that needs this interface to work doesn't need to know about the structs that implement it
// it only cares about some certain functionality
// e.g something like getUser, a use case function or a handler function just needs to get user at some point
type Datastore interface {
	GetUser(name string) (*models.User, error)
	SaveUser(*models.User) error
}