package G10_collection_datatypes

import "encoding/json"

type CalcFunc func(int) int




type Employee struct {
	ID int
	Name string
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}


type Employer struct {
	firstName string `json:"firstName"`
	lastName string `json:"lastName"`
	salary int `json:"salary"`
	fullTime bool `json:"fulltime,omitempty"`
}


func TestCreate() {
	var e1 Employee
	e1.ID = 1
	e2 := Employee{1,"James"}
	e3 := Employee{ID: 1, Name: "James"}
	e2.ID = 2
	e3.ID = 3
}


func TestEmbedding() {
	w1 := Wheel{Circle{Point{1,2},3},4}
	w1.Radius = 2
}

func TestJSON() {
	employees := []Employer{
		{"f1","l1",5000,true},
		{"f2","l2",5000,false},
		{"f2","l2",5000, false},
	}

	// from struct to json
	data, err := json.Marshal(employees)
	println(data,err)
	fData, fErr := json.MarshalIndent(employees,"","  ")
	println(fData,fErr)


	// from json to struct
	var e []struct{firstName string} // will only take firstName as attribute
	json.Unmarshal(fData, &e)
}