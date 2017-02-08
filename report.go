package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/charlesvdv/fuzmatch"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"strconv"
)

func report(settings *Settings) {
	fmt.Println("report()")

	database := get_database(settings)

	var statement string
	var row *sql.Row
	var rows *sqlx.Rows
	var count int
	var record Record
	var file *os.File
	var ratios_amt int
	var ratios_sedex_id int
	var err error

	statement = `
    SELECT COUNT(id)
    FROM records
    WHERE
        egeli_informatik_ch_co_amt != tilbago_k_infinity_com_amt
        OR
        egeli_informatik_ch_co_sedex_id != tilbago_k_infinity_com_sedex_id
    `
	row = database.QueryRow(statement)
	err = row.Scan(&count)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	if count == 0 {
		return
	}

	file, err = os.Create("ADR461-REPORT.CSV")
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	defer file.Close()

	statement = `
    SELECT *
    FROM records
    WHERE
        egeli_informatik_ch_co_amt != tilbago_k_infinity_com_amt
        OR
        egeli_informatik_ch_co_sedex_id != tilbago_k_infinity_com_sedex_id
    ORDER BY id ASC
    `
	rows, err = database.Queryx(statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	writer := csv.NewWriter(file)

	err = writer.Write(
		[]string{
			"Street",
			"Number",
			"Zip",
			"City",
			"Amt wählen - Source 1",
			"Amt wählen - Source 2",
			"Amt wählen - Match %",
			"Sedex ID - Source 1",
			"Sedex ID - Source 2",
			"Sedex ID - Match %",
		},
	)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	bar := pb.StartNew(count)
	for rows.Next() {
		err = rows.StructScan(&record)
		ratios_amt = fuzmatch.PartialRatio(record.EgeliInformatikChCoAmt.String, record.TilbagoKInfinityComAmt.String)
		ratios_sedex_id = fuzmatch.PartialRatio(
			record.EgeliInformatikChCoSedexId.String, record.TilbagoKInfinityComSedexId.String,
		)
		err := writer.Write(
			[]string{
				record.Street,
				record.Number,
				record.Zip,
				record.City,
				record.EgeliInformatikChCoAmt.String,
				record.EgeliInformatikChCoSedexId.String,
				strconv.Itoa(ratios_amt),
				record.TilbagoKInfinityComAmt.String,
				record.TilbagoKInfinityComSedexId.String,
				strconv.Itoa(ratios_sedex_id),
			},
		)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			panic(err)
		}
		bar.Increment()
	}

	defer writer.Flush()
}
