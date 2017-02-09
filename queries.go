package main

import (
	"database/sql"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func records_select_total(database *sqlx.DB) int64 {
	statement := `SELECT COUNT(id) FROM records`
	row := database.QueryRow(statement)
	var total int64
	err := row.Scan(&total)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	return total
}

func records_select_report_total(database *sqlx.DB) int {
	statement := `
    SELECT COUNT(id)
    FROM records
    WHERE
        egeli_informatik_ch_co_amt != tilbago_k_infinity_com_amt
        OR
        egeli_informatik_ch_co_sedex_id != tilbago_k_infinity_com_sedex_id
    `
	row := database.QueryRow(statement)
	var total int
	err := row.Scan(&total)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	return total
}

func records_select_report_records(database *sqlx.DB) *sqlx.Rows {
	statement := `
    SELECT *
    FROM records
    WHERE
        egeli_informatik_ch_co_amt != tilbago_k_infinity_com_amt
        OR
        egeli_informatik_ch_co_sedex_id != tilbago_k_infinity_com_sedex_id
    ORDER BY id ASC
    `
	rows, err := database.Queryx(statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	return rows
}

func source_1_select_progress(database *sqlx.DB, typ []string, total int64) (string, string, string, string) {
	source := fmt.Sprintf("#1 - %s (%s)", typ[2], typ[1])

	query := `
    SELECT COUNT(id)
    FROM records
    WHERE egeli_informatik_ch_%s_amt IS NULL AND egeli_informatik_ch_%s_sedex_id IS NULL
    `
	statement := fmt.Sprintf(query, typ[0], typ[0])
	row := database.QueryRow(statement)
	var pending int64
	err := row.Scan(&pending)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	pending_string := fmt.Sprintf("%07s", strconv.FormatInt(pending, 10))

	completed := total - pending
	completed_string := fmt.Sprintf("--%07s", strconv.FormatInt(completed, 10))

	percentage := (float64(completed) * 100.00) / (float64(total) * 1.00)
	percentage_string := fmt.Sprintf("---%06.2f%%", percentage)

	return source, completed_string, pending_string, percentage_string
}

func source_1_select_all(database *sqlx.DB, typ []string) (int, *sqlx.Rows) {
	query_total := `
    SELECT COUNT(id)
    FROM records
    WHERE egeli_informatik_ch_%s_amt IS NULL AND egeli_informatik_ch_%s_sedex_id IS NULL
    `
	statement_total := fmt.Sprintf(query_total, typ[0], typ[0])
	row_total := database.QueryRow(statement_total)
	var total int
	err_total := row_total.Scan(&total)
	if err_total != nil {
		raven.CaptureErrorAndWait(err_total, nil)
	}

	query_star := `
    SELECT *
    FROM records
    WHERE egeli_informatik_ch_%s_amt IS NULL AND egeli_informatik_ch_%s_sedex_id IS NULL
    ORDER BY id ASC
    `
	statement_star := fmt.Sprintf(query_star, typ[0], typ[0])
	rows, queryx_err := database.Queryx(statement_star)
	if queryx_err != nil {
		raven.CaptureErrorAndWait(queryx_err, nil)
	}
	return total, rows
}

func source_1_select_one(database *sqlx.DB) Record {
	statement := `
    SELECT *
    FROM records
    WHERE egeli_informatik_ch_co_amt IS NULL AND egeli_informatik_ch_co_sedex_id IS NULL
    ORDER BY RANDOM()
    LIMIT 1
    OFFSET 0
    `
	var record Record
	err := database.Get(&record, statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	return record
}

func source_1_update(database *sqlx.DB, typ []string, record Record, source_1_2 Source12) {
	if typ[1] == "BO" {
		record.EgeliInformatikChBoAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChBoSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "CF" {
		record.EgeliInformatikChCfAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChCfSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "CO" {
		record.EgeliInformatikChCoAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChCoSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "COR" {
		record.EgeliInformatikChCorAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChCorSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "DC" {
		record.EgeliInformatikChDcAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChDcSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "EIHI" {
		record.EgeliInformatikChEihiAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChEihiSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "IC" {
		record.EgeliInformatikChIcAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChIcSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "JP" {
		record.EgeliInformatikChJpAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChJpSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "LRO" {
		record.EgeliInformatikChLroAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChLroSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "MSO" {
		record.EgeliInformatikChMsoAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChMsoSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "RO" {
		record.EgeliInformatikChRoAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChRoSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	if typ[1] == "SAO" {
		record.EgeliInformatikChSaoAmt = sql.NullString{String: source_1_2.Amt, Valid: true}
		record.EgeliInformatikChSaoSedexId = sql.NullString{String: source_1_2.SedexId, Valid: true}
	}
	query := `
    UPDATE records
    SET
        egeli_informatik_ch_%s_amt = :egeli_informatik_ch_%s_amt,
        egeli_informatik_ch_%s_sedex_id = :egeli_informatik_ch_%s_sedex_id
    WHERE id = :id
    `
	statement := fmt.Sprintf(query, typ[0], typ[0], typ[0], typ[0])
	database.NamedExec(statement, record)
}

func source_2_select_progress(database *sqlx.DB, total int64) (string, string, string, string) {
	source := "#2"

	statement := `
    SELECT COUNT(id)
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    `
	row := database.QueryRow(statement)
	var pending int64
	err := row.Scan(&pending)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	pending_string := fmt.Sprintf("%07s", strconv.FormatInt(pending, 10))

	completed := total - pending
	completed_string := fmt.Sprintf("--%07s", strconv.FormatInt(completed, 10))

	percentage := (float64(completed) * 100.00) / (float64(total) * 1.00)
	percentage_string := fmt.Sprintf("---%06.2f%%", percentage)

	return source, completed_string, pending_string, percentage_string
}

func source_2_select_all(database *sqlx.DB) (int, *sqlx.Rows) {
	statement_total := `
    SELECT COUNT(id)
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    `
	row_total := database.QueryRow(statement_total)
	var total int
	err_total := row_total.Scan(&total)
	if err_total != nil {
		raven.CaptureErrorAndWait(err_total, nil)
	}

	statement_star := `
    SELECT *
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    ORDER BY id ASC
    `
	rows, queryx_err := database.Queryx(statement_star)
	if queryx_err != nil {
		raven.CaptureErrorAndWait(queryx_err, nil)
	}
	return total, rows
}

func source_2_select_one(database *sqlx.DB) Record {
	statement := `
    SELECT *
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    ORDER BY RANDOM()
    LIMIT 1
    OFFSET 0
    `
	var record Record
	err := database.Get(&record, statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	return record
}

func source_2_update(database *sqlx.DB, record Record, source_2 Source2) {
	record.TilbagoKInfinityComAmt = sql.NullString{String: source_2.Offices[0].Amt, Valid: true}
	record.TilbagoKInfinityComSedexId = sql.NullString{String: source_2.Offices[0].SedexId, Valid: true}
	statement := `
    UPDATE records
    SET
        tilbago_k_infinity_com_amt = :tilbago_k_infinity_com_amt,
        tilbago_k_infinity_com_sedex_id = :tilbago_k_infinity_com_sedex_id
    WHERE id = :id
    `
	database.NamedExec(statement, record)
}
