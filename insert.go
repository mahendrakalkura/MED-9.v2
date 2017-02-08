package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/lib/pq"
	"io"
	"os"
)

func insert(settings *Settings) {
	fmt.Println("insert()")

	database := get_database(settings)

	transaction, err := database.Begin()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	statement, err := transaction.Prepare(pq.CopyIn("records", "zip", "city", "street", "number"))
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	file, _ := os.Open("ADR461.CSV")
	buffer := bufio.NewReader(file)
	resource := csv.NewReader(buffer)
	resource.Comma = ';'
	resource.Read()
	for {
		record, err := resource.Read()
		if err == io.EOF {
			break
		}
		_, err = statement.Exec(record[0], record[1], record[2], record[3])
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			panic(err)
		}
	}

	_, err = statement.Exec()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	err = statement.Close()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	err = transaction.Commit()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
}
