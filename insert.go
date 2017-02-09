package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/lib/pq"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"os"
)

func insert(settings *Settings) {
	fmt.Println("insert()")

	database := get_database(settings)

	transaction, begin_err := database.Begin()
	if begin_err != nil {
		raven.CaptureErrorAndWait(begin_err, nil)
		panic(begin_err)
	}

	statement, prepare_err := transaction.Prepare(pq.CopyIn("records", "zip", "city", "street", "number"))
	if prepare_err != nil {
		raven.CaptureErrorAndWait(prepare_err, nil)
		panic(prepare_err)
	}

	file, open_err := os.Open("ADR461.CSV")
	if open_err != nil {
		raven.CaptureErrorAndWait(open_err, nil)
		panic(open_err)
	}

	stat, stat_err := file.Stat()
	if stat_err != nil {
		raven.CaptureErrorAndWait(stat_err, nil)
		panic(stat_err)
	}

	progress_bar := pb.New(int(stat.Size())).SetUnits(pb.U_BYTES)
	progress_bar.Start()

	proxy_reader := progress_bar.NewProxyReader(file)

	buffer_reader := bufio.NewReader(proxy_reader)

	csv_reader := csv.NewReader(buffer_reader)
	csv_reader.Comma = ';'
	csv_reader.Read()
	for {
		record, read_err := csv_reader.Read()
		if read_err == io.EOF {
			break
		}
		_, exec_err := statement.Exec(record[0], record[1], record[2], record[3])
		if exec_err != nil {
			raven.CaptureErrorAndWait(exec_err, nil)
			panic(exec_err)
		}
	}

	close_err := statement.Close()
	if close_err != nil {
		raven.CaptureErrorAndWait(close_err, nil)
		panic(close_err)
	}

	commit_err := transaction.Commit()
	if commit_err != nil {
		raven.CaptureErrorAndWait(commit_err, nil)
		panic(commit_err)
	}
}
