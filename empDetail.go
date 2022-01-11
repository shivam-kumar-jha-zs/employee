package main

import "fmt"

type empDetails struct {
	name string
	age  int
	dept string
}

func checkEmpDetails(xEmp empDetails) (empDetails, bool) {

	if xEmp.age < 22 {
		return empDetails{}, false
	}
	return xEmp, true
}
func main() {

	var emp []empDetails
	emp = append(emp, empDetails{age: 22, name: "abc", dept: "Golang"}, empDetails{age: 21, name: "sumit", dept: "java"})
	for i, _ := range emp {
		fmt.Println(checkEmpDetails(emp[i]))
	}

}
