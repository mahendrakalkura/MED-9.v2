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

func source_1_all(settings *Settings) {
	fmt.Println("source_1_all()")

	signal_channel := make(chan os.Signal)
	records_and_typs_channel := make(chan RecordAndTyp, settings.Others.Consumers*2)

	database := get_database(settings)

	for index := 1; index <= settings.Others.Consumers; index++ {
		go source_1_all_consumer(settings, database, records_and_typs_channel)
	}

	go source_1_all_producer(database, records_and_typs_channel)

	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	<-signal_channel

	close(records_and_typs_channel)
}

func source_1_all_consumer(settings *Settings, database *sqlx.DB, records_and_typs_channel chan RecordAndTyp) {
	for record_and_typ := range records_and_typs_channel {
		record := record_and_typ.Record
		typ := record_and_typ.Typ
		source_2, err := get_source_1(settings, record.Street, record.Number, record.Zip, record.City, typ)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
		} else {
			source_1_update(database, typ, record, source_2)
		}
	}
}

func source_1_all_producer(database *sqlx.DB, records_and_typs_channel chan RecordAndTyp) {
	for _, typ := range Typs {
		total, rows := source_1_select_all(database, typ)
		progress_bar := pb.StartNew(total)
		for rows.Next() {
			var record Record
			struct_scan_err := rows.StructScan(&record)
			if struct_scan_err != nil {
				raven.CaptureErrorAndWait(struct_scan_err, nil)
			} else {
				record_and_typ := RecordAndTyp{
					Record: record,
					Typ:    typ,
				}
				records_and_typs_channel <- record_and_typ
			}
			progress_bar.Increment()
		}
	}
}
