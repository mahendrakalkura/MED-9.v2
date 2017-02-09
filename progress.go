package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func progress(settings *Settings) {
	fmt.Println("progress()")

	database := get_database(settings)

	total := records_select_total(database)

	var items []Item

	for _, typ := range Typs {
		source, completed, pending, percentage := source_1_select_progress(database, typ, total)
		items = append(
			items,
			Item{
				source:     source,
				completed:  completed,
				pending:    pending,
				percentage: percentage,
			},
		)
	}

	source, completed, pending, percentage := source_2_select_progress(database, total)
	items = append(
		items,
		Item{
			source:     source,
			completed:  completed,
			pending:    pending,
			percentage: percentage,
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
