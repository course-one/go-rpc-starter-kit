package common

import (
	"fmt"
)

// Student struct represents a student.
type Student struct {
	ID                  int
	FirstName, LastName string
}

// FullName returns the fullname of the student.
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

/*---------------*/

// College struct represents a college.
type College struct {
	database map[int]Student // private
}

// Add methods adds a student to the college (procedure).
func (c *College) Add(payload Student, reply *Student) error {

	// check if student already exists in the database
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("student with id '%d' already exists", payload.ID)
	}

	// add student `p` in the database
	c.database[payload.ID] = payload

	// set reply value
	*reply = payload

	// return `nil` error
	return nil
}

// Get methods returns a student with specific id (procedure).
func (c *College) Get(payload int, reply *Student) error {

	// get student with id `p` from the database
	result, ok := c.database[payload]

	// check if student exists in the database
	if !ok {
		return fmt.Errorf("student with id '%d' does not exist", payload)
	}

	// set reply value
	*reply = result

	// return `nil` error
	return nil
}

// NewCollege function returns a new instance of College (pointer).
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}
