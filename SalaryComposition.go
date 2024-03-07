//---after Abstraction And Polymorphism of Salary project(polymorphism1)
//We use Inheritance(Composition with Embeded struct)

package main

import "fmt"

const (
	BaseSalary       = 5600000
	ExtraHourRate    = 90000
	HourlySalaryRate = 110000
	ShiftSalaryRate  = 80000
	TaxRate          = 0.09
)

func main() {

	//---input have a Employee struct
	fullTimeEmployees := []FullTimeEmployee{
		{Employee: Employee{Id: 1, NationalCode: "3333", FullName: "Reza"}, ExtraHours: 5},

		{Employee: Employee{Id: 2, NationalCode: "4444", FullName: "Abbas"}, ExtraHours: 10},
	}
	partTimeEmployees := []PartTimeEmployee{
		{Employee: Employee{Id: 3, NationalCode: "6666", FullName: "Hasan"}, PartTimeHours: 30},
		{Employee: Employee{Id: 4, NationalCode: "0000", FullName: "Bagher"}, PartTimeHours: 50},
	}

	//2--implement polymorphism with interface
	var employees []EmployeeSalaryCalculator = []EmployeeSalaryCalculator{}

	for _, employee := range fullTimeEmployees {
		employees = append(employees, employee)
	}

	for _, employee := range partTimeEmployees {
		employees = append(employees, employee)
	}

	for _, employee := range employees {
		salary, tax := employee.SalaryCalculate()
		fmt.Printf("Employee(%T): %v \nSalary: %f\nTax:%f\n", employee, employee, salary, tax)
	}

}

//--Abstraction (Struct + interface)
//Polymorphism 1-input of func is removed

type EmployeeSalaryCalculator interface {
	SalaryCalculate() (salary float64, tax float64)
}

// ----composition with Embeded struct
type Employee struct {
	Id           int
	NationalCode string
	FullName     string
}

type FullTimeEmployee struct {
	Employee
	ExtraHours float64
}

type PartTimeEmployee struct {
	Employee
	PartTimeHours float64
}

// 1-edit code inside
func (employee FullTimeEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = BaseSalary + (employee.ExtraHours*ExtraHourRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

// 1-edit inside code
func (employee PartTimeEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = employee.PartTimeHours * HourlySalaryRate
	tax = salary * TaxRate
	salary += tax
	return
}
