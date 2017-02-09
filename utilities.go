package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/xmlpath.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func get_database(settings *Settings) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings.SQLX.Hostname,
		settings.SQLX.Port,
		settings.SQLX.Username,
		settings.SQLX.Password,
		settings.SQLX.Database,
	)
	database := sqlx.MustConnect("postgres", dsn)
	return database
}

func get_http_client(settings *Settings) *http.Client {
	client := &http.Client{}

	timeout := time.Duration(60 * time.Second)
	client.Timeout = timeout

	if len(settings.Proxies.Hostname) > 0 && len(settings.Proxies.Ports) > 0 {
		proxy := get_proxy(settings.Proxies.Hostname, settings.Proxies.Ports)
		proxy_url, err := url.Parse(proxy)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy_url)}
	}

	return client
}

func get_proxy(hostname string, ports []int) string {
	port := get_random_number(ports[0], ports[1]+1)
	return fmt.Sprintf("https://%s:%d", hostname, port)
}

func get_random_number(minimum int, maximum int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(maximum-minimum) + minimum
}

func get_settings() *Settings {
	var settings = &Settings{}
	_, err := toml.DecodeFile("settings.toml", settings)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	return settings
}

func get_source_1(
	settings *Settings, street string, number string, zip string, city string, typ []string,
) (Source12, error) {
	var source_1_1 Source11
	var source_1_2 Source12

	place := fmt.Sprintf("%s %s %s %s", street, number, zip, city)

	client := get_http_client(settings)

	data_1 := url.Values{}
	data_1.Add("action", "aemterfinden_suggestions")
	data_1.Add("place", place)

	request_1, new_request_1_err := http.NewRequest(
		"POST",
		"https://www.egeli-informatik.ch/prd/wp-admin/admin-ajax.php",
		bytes.NewBufferString(data_1.Encode()),
	)
	if new_request_1_err != nil {
		raven.CaptureErrorAndWait(new_request_1_err, nil)
		return source_1_2, errors.New("new_request_1_err")
	}

	request_1.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request_1.Header.Add("DNT", "1")
	request_1.Header.Add("Host", "www.egeli-informatik.ch")
	request_1.Header.Add("Origin", "https://www.egeli-informatik.ch")
	request_1.Header.Add("Referer", "https://www.egeli-informatik.ch/unsere_loesungen/forderungsmanagement/aemterfinden/")
	request_1.Header.Add("User-Agent", "Go")
	request_1.Header.Add("X-Requested-With", "XMLHttpRequest")

	response_1, do_1_err := client.Do(request_1)
	if do_1_err != nil {
		raven.CaptureErrorAndWait(do_1_err, nil)
		return source_1_2, errors.New("do_1_err")
	}

	defer response_1.Body.Close()

	new_decoder_err := json.NewDecoder(response_1.Body).Decode(&source_1_1)
	if new_decoder_err != nil {
		raven.CaptureErrorAndWait(new_decoder_err, nil)
		return source_1_2, errors.New("new_decoder_err")
	}

	if source_1_1.TotalHits == 0 {
		source_1_2.Amt = "404"
		source_1_2.SedexId = "404"
		return source_1_2, nil
	}

	if len(source_1_1.Data) == 0 {
		source_1_2.Amt = "404"
		source_1_2.SedexId = "404"
		return source_1_2, nil
	}

	Data := source_1_1.Data[0]

	data_2 := url.Values{}
	data_2.Add("place", place)
	data_2.Add("action", "aemterfinden_result")
	data_2.Add("addressObject[Aktiv]", strconv.FormatBool(Data.Aktiv))
	data_2.Add("addressObject[AlternativeSuchbegriffeAsSearchString]", Data.AlternativeSuchbegriffeAsSearchString)
	data_2.Add("addressObject[AlternativeSuchbegriffeAsString]", Data.AlternativeSuchbegriffeAsString)
	data_2.Add("addressObject[BfsNr]", Data.BfsNr)
	data_2.Add("addressObject[HausKey]", strconv.Itoa(Data.HausKey))
	data_2.Add("addressObject[HausNummer]", strconv.Itoa(Data.HausNummer))
	data_2.Add("addressObject[HausNummerAlpha]", Data.HausNummerAlpha)
	data_2.Add("addressObject[Kanton]", Data.Kanton)
	data_2.Add("addressObject[Land]", Data.Land)
	data_2.Add("addressObject[NameComplete]", Data.NameComplete)
	data_2.Add("addressObject[Onrp]", Data.Onrp)
	data_2.Add("addressObject[Ort]", Data.Ort)
	data_2.Add("addressObject[Postleitzahl]", Data.Postleitzahl)
	data_2.Add("addressObject[Quartier]", Data.Quartier)
	data_2.Add("addressObject[SprachCode]", Data.SprachCode)
	data_2.Add("addressObject[Stadtkreis]", Data.Stadtkreis)
	data_2.Add("addressObject[StrassenName]", Data.StrassenName)
	data_2.Add("amtTyp", typ[1])

	request_2, new_request_2_err := http.NewRequest(
		"POST",
		"https://www.egeli-informatik.ch/prd/wp-admin/admin-ajax.php",
		bytes.NewBufferString(data_2.Encode()),
	)
	if new_request_2_err != nil {
		raven.CaptureErrorAndWait(new_request_2_err, nil)
		return source_1_2, errors.New("new_request_2_err")
	}

	request_2.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request_2.Header.Add("DNT", "1")
	request_2.Header.Add("Host", "www.egeli-informatik.ch")
	request_2.Header.Add("Origin", "https://www.egeli-informatik.ch")
	request_2.Header.Add("Referer", "https://www.egeli-informatik.ch/unsere_loesungen/forderungsmanagement/aemterfinden/")
	request_2.Header.Add("User-Agent", "Go")
	request_2.Header.Add("X-Requested-With", "XMLHttpRequest")

	response_2, do_2_err := client.Do(request_2)
	if do_2_err != nil {
		raven.CaptureErrorAndWait(do_2_err, nil)
		return source_1_2, errors.New("do_2_err")
	}

	defer response_2.Body.Close()

	body_bytes, read_all_err := ioutil.ReadAll(response_2.Body)
	if read_all_err != nil {
		raven.CaptureErrorAndWait(read_all_err, nil)
		return source_1_2, errors.New("read_all_err")
	}

	body_string := string(body_bytes)

	reader := strings.NewReader(body_string)
	root, parse_html_err := xmlpath.ParseHTML(reader)
	if parse_html_err != nil {
		raven.CaptureErrorAndWait(parse_html_err, nil)
		return source_1_2, errors.New("parse_html_err")
	}

	path_1 := `//li/div[@class="result"]/h2/text()`
	xpath_1 := xmlpath.MustCompile(path_1)
	value_1, ok_1 := xpath_1.String(root)
	if ok_1 {
		source_1_2.Amt = get_text(value_1)
	}

	path_2 := `//li/div[@class="result"]/div[@class="column"]/p[@class="eschkg_id"]/text()`
	xpath_2 := xmlpath.MustCompile(path_2)
	value_2, ok_2 := xpath_2.String(root)
	if ok_2 {
		source_1_2.SedexId = get_text(value_2)
	}

	return source_1_2, nil
}

func get_source_2(settings *Settings, street string, number string, zip string, city string) (Source2, error) {
	var source_2 Source2

	client := get_http_client(settings)

	request, new_request_err := http.NewRequest("GET", "http://tilbago.k-infinity.com:2607/dev/amtinfo", nil)
	if new_request_err != nil {
		raven.CaptureErrorAndWait(new_request_err, nil)
		return source_2, errors.New("new_request_err")
	}

	request.Header.Add("Host", "tilbago.k-infinity.com:2607")
	request.Header.Add("User-Agent", "Go")

	query := request.URL.Query()
	query.Add("city", city)
	query.Add("number", number)
	query.Add("street", street)
	query.Add("zip", zip)
	request.URL.RawQuery = query.Encode()

	response, do_err := client.Do(request)
	if do_err != nil {
		raven.CaptureErrorAndWait(do_err, nil)
		return source_2, errors.New("do_err")
	}

	defer response.Body.Close()

	new_decoder_err := json.NewDecoder(response.Body).Decode(&source_2)
	if new_decoder_err != nil {
		raven.CaptureErrorAndWait(new_decoder_err, nil)
		return source_2, errors.New("new_decoder_err")
	}

	if len(source_2.Offices) == 0 {
		source_2.Offices = []Source2Office{
			{
				Amt:     strconv.Itoa(*source_2.Code),
				SedexId: strconv.Itoa(*source_2.Code),
			},
		}
	}

	return source_2, nil
}

func get_text(text string) string {
	texts := strings.SplitN(text, ":", 2)
	text = texts[len(texts)-1]
	text = strings.TrimSpace(text)
	return text
}
