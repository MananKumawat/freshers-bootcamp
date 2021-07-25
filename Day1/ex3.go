package main

import (
	"fmt"
)

type Employee interface { // Interface for all Employee having common method salary
	salary() float64
}

type Fulltime struct{ // Struct for Fulltime
	basic float64
	NumberOfMonths int
}

type Contractor struct{ // Struct for Contractor
	basic float64
	NumberOfMonths int
}

type Freelencer struct{ // Struct for Freelencer
	basic float64
	NumberOfDays int
}

func (w Contractor) salary() float64 { // Salary for Contractor
	return 28.0 * w.basic * float64(w.NumberOfMonths)
}

func (w Fulltime) salary() float64 { // Salary for Fulltime
	return 28.0 * w.basic * float64(w.NumberOfMonths)
}

func (f Freelencer) salary() float64 { // Salary for Freelencer
	return 8.0 * f.basic * float64(f.NumberOfDays)
}

func TotalExpenses(employees []Employee) float64 {
	var total float64 = 0
	for _, employee := range employees {
		total += employee.salary()
	}
	return total
}

func main () {
	fulltime := Fulltime{500, 10}
	contractor := Contractor{100, 8}
	freelancer := Freelencer{10, 28}

	var employees = []Employee{fulltime, contractor, freelancer}

	fmt.Println(TotalExpenses(employees))
}