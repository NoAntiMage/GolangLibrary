package models

import "time"

type Employee struct {
	EmpNo     int
	BirthDate time.Time
	FistName  string
	LastName  string
	Gender    string
	HireDate  time.Time
	LeaveDate time.Time
}
