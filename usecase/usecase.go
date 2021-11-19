package usecase

import (
	"dep-test/db"
	"dep-test/models"
	"fmt"
)

//  what you notice is here is that when we define the dependencies( the things this package needs to do its job)
//  we use interfaces not the actual struct
//  because we don't really need the struct, what we need is just the functionality
//  so any struct that can give us this functionality is okay for us (this is basically the idea behind mocking and testing)
type Usecase struct {
	Store db.Datastore
}


func NewUseCase(store db.Datastore) *Usecase {
	return &Usecase{
		Store: store,
	}
}

//  the createuser function as declared here is attached to the usecase struct as it's receiver
//  this means that wherever the usecase struct is created it would have all the methods we have attached like this one
//  this create user function needs some functionality to do it's work
//  it needs to find out if the name we are trying to create already exists and it not it can go ahead and create it
//  so it needs two things, a way to check for a single user
//  and a way to create a single user
func (U *Usecase) Createuser(name string) (*models.User, error) {
	// first we want to check if the user exists
	// here we are using the .GetUser functionality that has been defined in the interface
	_, err := U.Store.GetUser(name)
	if err != nil {
		fmt.Printf("user with %v already exist\n", name)
		return nil, err
	}
	user := models.User{
		Name: name,
	}

	err = U.Store.SaveUser(&user)
	if err != nil {
		fmt.Printf("err saving user %v\n", err)
		return nil, err
	}

	fmt.Println("successfully created the user")
	return &user, nil
}