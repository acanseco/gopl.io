package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Adress    string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.ID = 1
	dilbert.Name = "Dilbert"
	dilbert.Adress = "123 Main St."
	dilbert.DoB = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)
	dilbert.Position = "Engineer"
	dilbert.Salary = 100000
	dilbert.ManagerID = 2

	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	var employeeOfTheMonth *Employee = &dilbert

	employeeOfTheMonth.Position += " (proactive team player)"

	fmt.Println(dilbert)
	fmt.Println(*EmployeeByID(1))
}

func EmployeeByID(id int) *Employee {
	return &Employee{
		ID:        1,
		Name:      "Dilbert",
		Adress:    "123 Main St.",
		DoB:       time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC),
		Position:  "Engineer",
		Salary:    100000,
		ManagerID: 2,
	}
}
