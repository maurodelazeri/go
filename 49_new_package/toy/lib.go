package toy

import (
	"fmt"
	"strconv"
)

// Toy is the datatype
type Toy struct {
	Name string
}

// NewToy returns a new Toy type. It supports the following optional params:
func NewToy(v ...interface{}) *Toy {
	if len(v) < 1 || v[0] == nil {
		return nil
	}

	r := &Toy{Name: ""}

	// if multiple parameters, only the last one assigned
	for i := range v {
		switch v[i].(type) {
		case string:
			r.Name = v[i].(string) // convert interface to string
		case []byte:
			r.Name = string(v[i].([]byte)) // convert interface to []byte
		case int:
			r.Name = strconv.Itoa(v[i].(int)) // convert interface to int
		}
	}

	return r
}

// Length returns the length as a string
func (r *Toy) Length() string {
	return strconv.Itoa(len(r.Name))
}

// GetName returns a string
func (r *Toy) GetName() string {
	return r.Name
}

// NewToys returns a slice of Toy
func NewToys(args ...interface{}) ([]Toy, error) {
	manyToys := []Toy{}
	if len(args) < 1 || args[0] == nil {
		return nil, fmt.Errorf("Nothing or blank")
	}

	// if multiple parameters, return a slice of Toy
	for i := range args {
		toy1 := Toy{Name: ""}
		switch args[i].(type) {
		case string:
			//toy1.Name = args[i].(string) // convert interface to string
			toy1 = Toy{Name: args[i].(string)} // this way works too
		case []byte:
			toy1.Name = string(args[i].([]byte)) // convert interface to []byte
		case int:
			toy1.Name = strconv.Itoa(args[i].(int)) // convert interface to int
		}

		manyToys = append(manyToys, toy1)
	}

	return manyToys, nil
}
