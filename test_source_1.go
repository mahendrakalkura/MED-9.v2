package main

import (
	"fmt"
)

func test_source_1(settings *Settings) {
	fmt.Println("test_source_1()")
	fmt.Println("")

	database := get_database(settings)

	record := select_source_1_random(database)
	fmt.Println("Number : ", record.Number)
	fmt.Println("Street : ", record.Street)
	fmt.Println("City   : ", record.City)
	fmt.Println("Zip    : ", record.Zip)
	fmt.Println("")

	source_1_2, err := get_source_1(settings, record.Street, record.Number, record.Zip, record.City, "CO")
	if err != nil {
		panic(err)
	}
	fmt.Println("Amt    : ", source_1_2.Amt)
	fmt.Println("SedexId: ", source_1_2.SedexId)
}
