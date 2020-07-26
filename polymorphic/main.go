package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int //第一個是 method 名稱，第二個是 return type
}

// 長期雇工
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

// 合約雇工
type Contract struct {
	empId    int
	basicpay int
}

// 長期雇工 實作 CalculateSalary()
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// 合約雇工實作 CalculateSalary()
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/

// 必須傳 slice 並且 type 是 SalaryCalculator interface
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		// v 必須都有實作 CalculateSalary()
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	// var employees []SalaryCalculator{}
	// employees = append(pemp1, pemp2, cemp1)
	totalExpense(employees)
}

// Total Expense Per Month $14050
