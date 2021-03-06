package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"os/signal"
	"syscall"
)

func source_2_all(settings *Settings) {
	fmt.Println("source_2_all()")

	signal_channel := make(chan os.Signal)
	records_channel := make(chan Record, settings.Others.Consumers*2)

	database := get_database(settings)

	for index := 1; index <= settings.Others.Consumers; index++ {
		go source_2_all_consumer(settings, database, records_channel)
	}

	go source_2_all_producer(database, records_channel)

	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	<-signal_channel

	close(records_channel)
}

func source_2_all_consumer(settings *Settings, database *sqlx.DB, records_channel chan Record) {
	for record := range records_channel {
		source_2, err := get_source_2(settings, record.Street, record.Number, record.Zip, record.City)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
		} else {
			source_2_update(database, record, source_2)
		}
	}
}

func source_2_all_producer(database *sqlx.DB, records_channel chan Record) {
	total, rows := source_2_select_all(database)
	progress_bar := pb.StartNew(total)
	for rows.Next() {
		var record Record
		struct_scan_err := rows.StructScan(&record)
		if struct_scan_err != nil {
			raven.CaptureErrorAndWait(struct_scan_err, nil)
		} else {
			records_channel <- record
		}
		progress_bar.Increment()
	}
}
