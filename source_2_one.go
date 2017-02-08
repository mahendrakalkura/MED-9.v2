package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
)

func source_2_one(settings *Settings) {
	fmt.Println("source_2_one()")

	database := get_database(settings)

	record := source_2_select_one(database)
	fmt.Printf("%-7s: %s\n", "Number", record.Number)
	fmt.Printf("%-7s: %s\n", "Street", record.Street)
	fmt.Printf("%-7s: %s\n", "City", record.City)
	fmt.Printf("%-7s: %s\n", "Zip", record.Zip)

	source_2, err := get_source_2(settings, record.Street, record.Number, record.Zip, record.City)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	fmt.Printf("%-7s: %s\n", "Amt", source_2.Offices[0].Amt)
	fmt.Printf("%-7s: %s\n", "SedexId", source_2.Offices[0].SedexId)
}
