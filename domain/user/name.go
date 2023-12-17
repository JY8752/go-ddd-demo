package user

import (
	"errors"
	"reflect"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName, lastName string) (*FullName, error) {
	if err := firstNameValidation(firstName); err != nil {
		return nil, err
	}

	if err := lastNameValidation(lastName); err != nil {
		return nil, err
	}

	return &FullName{firstName: firstName, lastName: lastName}, nil
}

func (fn *FullName) FirstName() string {
	return fn.firstName
}

func (fn *FullName) LastName() string {
	return fn.lastName
}

func (fn *FullName) FullName() string {
	return fn.firstName + " " + fn.lastName
}

func (fn *FullName) Equals(other *FullName) bool {
	return reflect.DeepEqual(fn, other)
}

func (fn *FullName) String() string {
	return fn.FullName()
}

func firstNameValidation(firstName string) error {
	if firstName == "" {
		return errors.New("firstName is empty")
	}
	return nil
}

func lastNameValidation(lastName string) error {
	if lastName == "" {
		return errors.New("lastName is empty")
	}
	return nil
}
