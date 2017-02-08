package main

import (
    "fmt"
)

func bootstrap(settings *Settings) {
    fmt.Println("bootstrap()")

    database := get_database(settings)

    statement := `
    DROP SCHEMA IF EXISTS public CASCADE;

    CREATE SCHEMA IF NOT EXISTS public;

    CREATE TABLE IF NOT EXISTS records
    (
        id INTEGER NOT NULL,
        zip CHARACTER VARYING(255) NOT NULL,
        city CHARACTER VARYING(255) NOT NULL,
        street CHARACTER VARYING(255) NOT NULL,
        number CHARACTER VARYING(255) NOT NULL,
        egeli_informatik_ch_co_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_co_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_cf_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_cf_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_dc_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_dc_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_ro_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_ro_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_jp_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_jp_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_lro_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_lro_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_cor_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_cor_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_bo_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_bo_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_eihi_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_eihi_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_sao_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_sao_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_ic_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_ic_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_mso_amt CHARACTER VARYING(255) DEFAULT NULL,
        egeli_informatik_ch_mso_sedex_id CHARACTER VARYING(255) DEFAULT NULL,
        tilbago_k_infinity_com_amt CHARACTER VARYING(255) DEFAULT NULL,
        tilbago_k_infinity_com_sedex_id CHARACTER VARYING(255) DEFAULT NULL
    );

    CREATE SEQUENCE records_id_sequence START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

    ALTER TABLE records ALTER COLUMN id SET DEFAULT nextval('records_id_sequence'::regclass);

    ALTER TABLE records ADD CONSTRAINT records_id_constraint PRIMARY KEY (id);

    CREATE INDEX records_street ON records USING btree (street);

    CREATE INDEX records_number ON records USING btree (number);

    CREATE INDEX records_zip ON records USING btree (zip);

    CREATE INDEX records_city ON records USING btree (city);
    `

    defer database.Close()

    database.MustExec(statement)
}
