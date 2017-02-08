package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
)

func source_1_one(settings *Settings) {
	fmt.Println("source_1_one()")

	database := get_database(settings)

	record := source_1_select_one(database)
	fmt.Printf("%-35s: %s\n", "Number", record.Number)
	fmt.Printf("%-35s: %s\n", "Street", record.Street)
	fmt.Printf("%-35s: %s\n", "City", record.City)
	fmt.Printf("%-35s: %s\n", "Zip", record.Zip)

	for _, typ := range Typs {
		source_1_2, err := get_source_1(settings, record.Street, record.Number, record.Zip, record.City, typ)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			panic(err)
		}
		fmt.Printf("%-27s Amt    : %s\n", typ[2], source_1_2.Amt)
		fmt.Printf("%-27s SedexId: %s\n", typ[2], source_1_2.SedexId)
	}
}