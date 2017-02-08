package main

import (
	"database/sql"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type Item struct {
	source     string
	completed  string
	pending    string
	percentage string
}

func progress(settings *Settings) {
	fmt.Println("progress()")

	database := get_database(settings)

	var items []Item
	var statement string
	var row *sql.Row
	var err error
	var source string
	var total int64
	var pending int64
	var completed int64
	var percentage float64

	statement = `SELECT COUNT(id) AS count FROM records`
	row = database.QueryRow(statement)
	err = row.Scan(&total)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	for _, typ := range Typs {
		source = fmt.Sprintf("#1 %s (%s)", typ[2], typ[1])
		statement = `
        SELECT COUNT(id) AS count
        FROM records
        WHERE egeli_informatik_ch_%s_amt IS NULL AND egeli_informatik_ch_%s_sedex_id IS NULL
        `
		statement = fmt.Sprintf(statement, typ[0], typ[0])
		row = database.QueryRow(statement)
		err = row.Scan(&pending)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			panic(err)
		}
		completed = total - pending
		percentage = (float64(completed) * 100.00) / (float64(total) * 1.00)
		items = append(
			items,
			Item{
				source:     source,
				completed:  fmt.Sprintf("--%07s", strconv.FormatInt(completed, 10)),
				pending:    fmt.Sprintf("--%07s", strconv.FormatInt(pending, 10)),
				percentage: fmt.Sprintf("---%06.2f%%", percentage),
			},
		)
	}

	source = "#2"
	statement = `
    SELECT COUNT(id) AS count
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    `
	row = database.QueryRow(statement)
	err = row.Scan(&pending)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	completed = total - pending
	percentage = (float64(completed) * 100.00) / (float64(total) * 1.00)
	items = append(
		items,
		Item{
			source:     source,
			completed:  fmt.Sprintf("--%07s", strconv.FormatInt(completed, 10)),
			pending:    fmt.Sprintf("--%07s", strconv.FormatInt(pending, 10)),
			percentage: fmt.Sprintf("---%06.2f%%", percentage),
		},
	)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoFormatHeaders(false)
	table.SetAutoWrapText(false)
	table.SetHeader(
		[]string{
			"Source",
			"Completed",
			"Pending",
			"Percentage",
		},
	)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	for _, item := range items {
		table.Append(
			[]string{
				item.source,
				item.completed,
				item.pending,
				item.percentage,
			},
		)
	}
	table.Render()
}
