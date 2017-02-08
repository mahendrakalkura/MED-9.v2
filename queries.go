package main

import (
	"database/sql"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
)

func source_1_select_all(database *sqlx.DB, typ string) (int, []Record) {
	var statement string
	var row *sql.Row
	var count int
	var records []Record
	var err error

	statement = `
    SELECT COUNT(id) AS count
    FROM records
    WHERE egeli_informatik_ch_%s_amt IS NULL AND egeli_informatik_ch_%s_sedex_id IS NULL
    `
	statement = fmt.Sprintf(statement, typ[0], typ[0])
	row = database.QueryRow(statement)
	err = row.Scan(&count)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	statement = `
    SELECT *
    FROM records
    WHERE egeli_informatik_ch_%s_amt IS NULL AND egeli_informatik_ch_%s_sedex_id IS NULL
    ORDER BY id ASC
    `
	statement = fmt.Sprintf(statement, typ[0], typ[0])
	err = database.Select(&records, statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	return count, records
}

func source_1_select_one(database *sqlx.DB) Record {
	var record Record
	statement := `
    SELECT *
    FROM records
    WHERE egeli_informatik_ch_co_amt IS NULL AND egeli_informatik_ch_co_sedex_id IS NULL
    ORDER BY RANDOM()
    LIMIT 1
    OFFSET 0
    `
	err := database.Get(&record, statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	return record
}

func source_2_select_all(database *sqlx.DB) (int, []Record) {
	var statement string
	var row *sql.Row
	var count int
	var records []Record
	var err error

	statement = `
    SELECT COUNT(id) AS count
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    `
	row = database.QueryRow(statement)
	err = row.Scan(&count)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	statement = `
    SELECT *
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    ORDER BY id ASC
    `
	err = database.Select(&records, statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}

	return count, records
}

func source_2_select_one(database *sqlx.DB) Record {
	var record Record
	statement := `
    SELECT *
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    ORDER BY RANDOM()
    LIMIT 1
    OFFSET 0
    `
	err := database.Get(&record, statement)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	return record
}

func source_1_update(database *sqlx.DB, typ []string, record Record, source_1_2 Source12) {
	if typ[1] == "BO" {
		record.EgeliInformatikChBoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChBoSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "CF" {
		record.EgeliInformatikChBoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChBoSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "CO" {
		record.EgeliInformatikChCoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChCoSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "COR" {
		record.EgeliInformatikChCorAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChCorSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "DC" {
		record.EgeliInformatikChDcAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChDcSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "EIHI" {
		record.EgeliInformatikChEihiAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChEihiSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "IC" {
		record.EgeliInformatikChIcAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChIcSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "JP" {
		record.EgeliInformatikChJpAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChJpSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "LRO" {
		record.EgeliInformatikChLroAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChLroSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "MSO" {
		record.EgeliInformatikChMsoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChMsoSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "RO" {
		record.EgeliInformatikChRoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChRoSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	if typ[1] == "SAO" {
		record.EgeliInformatikChSaoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChSaoSedexId = sql.NullString{String: source_1_2.SedexId}
	}
	statement := `
    UPDATE records
    SET
        egeli_informatik_ch_%s_amt = :egeli_informatik_ch_%s_amt,
        egeli_informatik_ch_%s_sedex_id = :egeli_informatik_ch_%s_sedex_id
    WHERE id = :id
    `
	statement = fmt.Sprintf(statement, typ[0], typ[0])
	database.NamedExec(statement, record)
}

func source_2_update(database *sqlx.DB, record Record, source_2 Source2) {
	record.TilbagoKInfinityComAmt = sql.NullString{String: source_2.Offices[0].Amt}
	record.TilbagoKInfinityComSedexId = sql.NullString{String: source_2.Offices[0].SedexId}
	statement := `
    UPDATE records
    SET
        tilbago_k_infinity_com_amt = :tilbago_k_infinity_com_amt,
        tilbago_k_infinity_com_sedex_id = :tilbago_k_infinity_com_sedex_id
    WHERE id = :id
    `
	database.NamedExec(statement, record)
}
