package main

import (
	"encoding/csv"
	"fmt"
	"github.com/charlesvdv/fuzmatch"
	"github.com/getsentry/raven-go"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"strconv"
)

func report(settings *Settings) {
	fmt.Println("report()")

	database := get_database(settings)

	total := records_select_report_total(database)
	if total == 0 {
		return
	}

	rows := records_select_report_records(database)

	file, create_err := os.Create("ADR461-REPORT.CSV")
	if create_err != nil {
		raven.CaptureErrorAndWait(create_err, nil)
		panic(create_err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	write_err := writer.Write(
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
	if write_err != nil {
		raven.CaptureErrorAndWait(write_err, nil)
		panic(write_err)
	}

	progress_bar := pb.StartNew(total)
	for rows.Next() {
		var record Record
		struct_scan_err := rows.StructScan(&record)
		if struct_scan_err != nil {
			raven.CaptureErrorAndWait(struct_scan_err, nil)
			panic(struct_scan_err)
		}

		ratios_amt := fuzmatch.PartialRatio(record.EgeliInformatikChCoAmt.String, record.TilbagoKInfinityComAmt.String)

		ratios_sedex_id := fuzmatch.PartialRatio(
			record.EgeliInformatikChCoSedexId.String, record.TilbagoKInfinityComSedexId.String,
		)

		write_err := writer.Write(
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
		if write_err != nil {
			raven.CaptureErrorAndWait(write_err, nil)
			panic(write_err)
		}

		progress_bar.Increment()
	}

	defer writer.Flush()
}
