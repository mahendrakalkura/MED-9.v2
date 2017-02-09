package main

import (
	"database/sql"
)

type Settings struct {
	Proxies SettingsProxies `toml:"proxies"`
	Sentry  SettingsSentry  `toml:"sentry"`
	SQLX    SettingsSQLX    `toml:"sqlx"`
}

type SettingsProxies struct {
	Hostname string `toml:"hostname"`
	Ports    []int  `toml:"ports"`
}

type SettingsSentry struct {
	Dsn string `toml:"dsn"`
}

type SettingsSQLX struct {
	Database string `toml:"database"`
	Hostname string `toml:"hostname"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
}

type Record struct {
	Id                           int            `db:"id"`
	Zip                          string         `db:"zip"`
	City                         string         `db:"city"`
	Street                       string         `db:"street"`
	Number                       string         `db:"number"`
	EgeliInformatikChCoAmt       sql.NullString `db:"egeli_informatik_ch_co_amt"`
	EgeliInformatikChCoSedexId   sql.NullString `db:"egeli_informatik_ch_co_sedex_id"`
	EgeliInformatikChCfAmt       sql.NullString `db:"egeli_informatik_ch_cf_amt"`
	EgeliInformatikChCfSedexId   sql.NullString `db:"egeli_informatik_ch_cf_sedex_id"`
	EgeliInformatikChDcAmt       sql.NullString `db:"egeli_informatik_ch_dc_amt"`
	EgeliInformatikChDcSedexId   sql.NullString `db:"egeli_informatik_ch_dc_sedex_id"`
	EgeliInformatikChRoAmt       sql.NullString `db:"egeli_informatik_ch_ro_amt"`
	EgeliInformatikChRoSedexId   sql.NullString `db:"egeli_informatik_ch_ro_sedex_id"`
	EgeliInformatikChJpAmt       sql.NullString `db:"egeli_informatik_ch_jp_amt"`
	EgeliInformatikChJpSedexId   sql.NullString `db:"egeli_informatik_ch_jp_sedex_id"`
	EgeliInformatikChLroAmt      sql.NullString `db:"egeli_informatik_ch_lro_amt"`
	EgeliInformatikChLroSedexId  sql.NullString `db:"egeli_informatik_ch_lro_sedex_id"`
	EgeliInformatikChCorAmt      sql.NullString `db:"egeli_informatik_ch_cor_amt"`
	EgeliInformatikChCorSedexId  sql.NullString `db:"egeli_informatik_ch_cor_sedex_id"`
	EgeliInformatikChBoAmt       sql.NullString `db:"egeli_informatik_ch_bo_amt"`
	EgeliInformatikChBoSedexId   sql.NullString `db:"egeli_informatik_ch_bo_sedex_id"`
	EgeliInformatikChEihiAmt     sql.NullString `db:"egeli_informatik_ch_eihi_amt"`
	EgeliInformatikChEihiSedexId sql.NullString `db:"egeli_informatik_ch_eihi_sedex_id"`
	EgeliInformatikChSaoAmt      sql.NullString `db:"egeli_informatik_ch_sao_amt"`
	EgeliInformatikChSaoSedexId  sql.NullString `db:"egeli_informatik_ch_sao_sedex_id"`
	EgeliInformatikChIcAmt       sql.NullString `db:"egeli_informatik_ch_ic_amt"`
	EgeliInformatikChIcSedexId   sql.NullString `db:"egeli_informatik_ch_ic_sedex_id"`
	EgeliInformatikChMsoAmt      sql.NullString `db:"egeli_informatik_ch_mso_amt"`
	EgeliInformatikChMsoSedexId  sql.NullString `db:"egeli_informatik_ch_mso_sedex_id"`
	TilbagoKInfinityComAmt       sql.NullString `db:"tilbago_k_infinity_com_amt"`
	TilbagoKInfinityComSedexId   sql.NullString `db:"tilbago_k_infinity_com_sedex_id"`
}

type Source11 struct {
	TotalHits int            `json:"totalHits"`
	Data      []Source11Data `json:"data"`
}

type Source11Data struct {
	Aktiv                                 bool   `json:"Aktiv"`
	AlternativeSuchbegriffeAsSearchString string `json:"AlternativeSuchbegriffeAsSearchString"`
	AlternativeSuchbegriffeAsString       string `json:"AlternativeSuchbegriffeAsString"`
	BfsNr                                 string `json:"BfsNr"`
	HausKey                               int    `json:"HausKey"`
	HausNummer                            int    `json:"HausNummer"`
	HausNummerAlpha                       string `json:"HausNummerAlpha"`
	Kanton                                string `json:"Kanton"`
	Land                                  string `json:"Land"`
	NameComplete                          string `json:"NameComplete"`
	Onrp                                  string `json:"Onrp"`
	Ort                                   string `json:"Ort"`
	Postleitzahl                          string `json:"Postleitzahl"`
	Quartier                              string `json:"Quartier"`
	SprachCode                            string `json:"SprachCode"`
	Stadtkreis                            string `json:"Stadtkreis"`
	StrassenName                          string `json:"StrassenName"`
}

type Source12 struct {
	Amt     string
	SedexId string
}

type Source2 struct {
	Offices []Source2Office `json:"offices"`
	Code    *int            `json:"code,omitempty"`
}

type Source2Office struct {
	Amt     string `json:"amt"`
	SedexId string `json:"sedexId"`
}

type Item struct {
	source     string
	completed  string
	pending    string
	percentage string
}
