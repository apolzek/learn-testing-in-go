package main

import (
	"calc/calculator"
	"fmt"
)

type MyDBService struct{}

func (m MyDBService) GetSum() int {
	return 10
}

func main() {
	// Create an instance of MyDBService
	db := MyDBService{}

	// Call GetSumFromDB from calculator package
	result := calculator.GetSumFromDB(db)

	// Print the result
	fmt.Println("Result:", result)
}
