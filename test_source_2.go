package main

import (
	"fmt"
)

func test_source_2(settings *Settings) {
	fmt.Println("test_source_2()")
	fmt.Println("")

	database := get_database(settings)

	record := select_source_2_random(database)
	fmt.Println("Number :", record.Number)
	fmt.Println("Street :", record.Street)
	fmt.Println("City   :", record.City)
	fmt.Println("Zip    :", record.Zip)
	fmt.Println("")

	source_2, err := get_source_2(settings, record.Street, record.Number, record.Zip, record.City)
	if err != nil {
		panic(err)
	}
	fmt.Println("Amt    :", source_2.Offices[0].Amt)
	fmt.Println("SedexId:", source_2.Offices[0].SedexId)
}
