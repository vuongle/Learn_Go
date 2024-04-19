package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// Create enum for status. Data type of enum is int therefore, use iota. iota starts by 0, ...
// Therefore,  TodoStatus has value 0 | 1 | 2
type TodoStatus int

const (
	StatusDoing TodoStatus = iota
	StatusDone
	StatusDeleted
)

var allStatuses = [3]string{"Doing", "Done", "Deleted"}

// Define a method "String" with value receiver(item TodoStatus)
// Logic: Convery int to string
func (item TodoStatus) String() string {
	return allStatuses[item]
}

// Define a function to parse a string to int
// This function returns 2 value: int(TodoStatus) and error
func parseStringToTodoStatus(s string) (TodoStatus, error) {
	for i := range allStatuses {
		if allStatuses[i] == s {
			return TodoStatus(i), nil
		}
	}

	return TodoStatus(0), errors.New("invalid status string")
}

// Override/Implement a method "Scan" with pointer receiver(item *TodoStatus)
// Logic: read value from db (string type) and map to TodoStatus (int type)
// This method is automatically called by GORM
func (item *TodoStatus) Scan(value interface{}) error {
	fmt.Println("------------ Scan() CALLED")
	// convert value from db to byte array
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	// then convert byte array to string
	v, err := parseStringToTodoStatus(string(b))
	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	// change value of the pointer. use syntax "* + pointer variable"
	// first: the pointer points to int
	// now: points to string
	*item = v

	// no error -> return nil. If there is any errors -> already returned above
	return nil
}

// Override/Implement a method with pointer receiver to map from TodoStatus (int type) to string in db
func (item *TodoStatus) Value() (driver.Value, error) {
	fmt.Println("------------ Value() CALLED")
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

// Override/Implement MarshalJSON() of a struct
// This method is automatically called when replying data to client. It's called json encoding
// Logic: convert TodoStatus (int type) to json string(this json string includes " + value + ") in byte form
func (item *TodoStatus) MarshalJSON() ([]byte, error) {
	fmt.Println("------------ MarshalJSON() CALLED")
	if item == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// Override/Implement UnmarshalJSON() of a struct
// Convert json string to TodoStatus (int type). It's called json decoding
func (item *TodoStatus) UnmarshalJSON(data []byte) error {
	fmt.Println("------------ UnmarshalJSON() CALLED")
	str := strings.ReplaceAll(string(data), "\"", "")

	intVal, err := parseStringToTodoStatus(str)
	if err != nil {
		return err
	}

	*item = intVal

	return nil
}
