package main

import (
	"database/sql"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
)

func select_source_1_random(database *sqlx.DB) Record {
	var record Record
	query := `
    SELECT *
    FROM records
    WHERE egeli_informatik_ch_co_amt IS NULL AND egeli_informatik_ch_co_sedex_id IS NULL
    ORDER BY RANDOM()
    LIMIT 1
    OFFSET 0
    `
	err := database.Get(&record, query)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	return record
}

func select_source_2_random(database *sqlx.DB) Record {
	var record Record
	query := `
    SELECT *
    FROM records
    WHERE tilbago_k_infinity_com_amt IS NULL AND tilbago_k_infinity_com_sedex_id IS NULL
    ORDER BY RANDOM()
    LIMIT 1
    OFFSET 0
    `
	err := database.Get(&record, query)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	return record
}

func update_source_1(database *sqlx.DB, amt string, record Record, source_1_2 Source12) {
	if amt == "bo" {
		record.EgeliInformatikChBoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChBoSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_bo_amt = :egeli_informatik_ch_bo_amt,
            egeli_informatik_ch_bo_sedex_id = :egeli_informatik_ch_bo_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "cf" {
		record.EgeliInformatikChBoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChBoSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_bo_amt = :egeli_informatik_ch_bo_amt,
            egeli_informatik_ch_bo_sedex_id = :egeli_informatik_ch_bo_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "co" {
		record.EgeliInformatikChCoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChCoSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_co_amt = :egeli_informatik_ch_co_amt,
            egeli_informatik_ch_co_sedex_id = :egeli_informatik_ch_co_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "cor" {
		record.EgeliInformatikChCorAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChCorSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_cor_amt = :egeli_informatik_ch_cor_amt,
            egeli_informatik_ch_cor_sedex_id = :egeli_informatik_ch_cor_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "dc" {
		record.EgeliInformatikChDcAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChDcSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_dc_amt = :egeli_informatik_ch_dc_amt,
            egeli_informatik_ch_dc_sedex_id = :egeli_informatik_ch_dc_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "eihi" {
		record.EgeliInformatikChEihiAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChEihiSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_eihi_amt = :egeli_informatik_ch_eihi_amt,
            egeli_informatik_ch_eihi_sedex_id = :egeli_informatik_ch_eihi_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "ic" {
		record.EgeliInformatikChIcAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChIcSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_ic_amt = :egeli_informatik_ch_ic_amt,
            egeli_informatik_ch_ic_sedex_id = :egeli_informatik_ch_ic_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "jp" {
		record.EgeliInformatikChJpAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChJpSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_jp_amt = :egeli_informatik_ch_jp_amt,
            egeli_informatik_ch_jp_sedex_id = :egeli_informatik_ch_jp_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "lro" {
		record.EgeliInformatikChLroAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChLroSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_lro_amt = :egeli_informatik_ch_lro_amt,
            egeli_informatik_ch_lro_sedex_id = :egeli_informatik_ch_lro_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "mso" {
		record.EgeliInformatikChMsoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChMsoSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_mso_amt = :egeli_informatik_ch_mso_amt,
            egeli_informatik_ch_mso_sedex_id = :egeli_informatik_ch_mso_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "ro" {
		record.EgeliInformatikChRoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChRoSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_Ro_amt = :egeli_informatik_ch_Ro_amt,
            egeli_informatik_ch_Ro_sedex_id = :egeli_informatik_ch_Ro_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
	if amt == "sao" {
		record.EgeliInformatikChSaoAmt = sql.NullString{String: source_1_2.Amt}
		record.EgeliInformatikChSaoSedexId = sql.NullString{String: source_1_2.SedexId}
		query := `
        UPDATE records
        SET
            egeli_informatik_ch_sao_amt = :egeli_informatik_ch_sao_amt,
            egeli_informatik_ch_sao_sedex_id = :egeli_informatik_ch_sao_sedex_id
        WHERE id = :id
        `
		database.NamedExec(query, record)
	}
}

func update_source_2(database *sqlx.DB, record Record, source_2 Source2) {
	record.TilbagoKInfinityComAmt = sql.NullString{String: source_2.Offices[0].Amt}
	record.TilbagoKInfinityComSedexId = sql.NullString{String: source_2.Offices[0].SedexId}
	query := `
    UPDATE records
    SET
        tilbago_k_infinity_com_amt = :tilbago_k_infinity_com_amt,
        tilbago_k_infinity_com_sedex_id = :tilbago_k_infinity_com_sedex_id
    WHERE id = :id
    `
	database.NamedExec(query, record)
}
