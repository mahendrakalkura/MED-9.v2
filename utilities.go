package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
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
	timeout := time.Duration(30 * time.Second)

	proxy := get_proxy(settings.Proxies.Hostname, settings.Proxies.Ports)
	proxy_url, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	client.Timeout = timeout
	client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy_url)}

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
		panic(err)
	}
	return settings
}

func get_source_1(
	settings *Settings, street string, number string, zip string, city string, amt string,
) (Source12, error) {
	var source_1_1 Source11
	var source_1_2 Source12
	var request *http.Request
	var data url.Values
	var response *http.Response
	var path string
	var xpath *xmlpath.Path
	var value string
	var ok bool
	var err error

	place := fmt.Sprintf("%s %s %s %s", street, number, zip, city)

	client := get_http_client(settings)

	data = url.Values{}
	data.Add("action", "aemterfinden_suggestions")
	data.Add("place", place)

	request, err = http.NewRequest(
		"POST",
		"https://www.egeli-informatik.ch/prd/wp-admin/admin-ajax.php",
		bytes.NewBufferString(data.Encode()),
	)
	if err != nil {
		return source_1_2, errors.New("Error #1")
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("DNT", "1")
	request.Header.Add("Host", "www.egeli-informatik.ch")
	request.Header.Add("Origin", "https://www.egeli-informatik.ch")
	request.Header.Add("Referer", "https://www.egeli-informatik.ch/unsere_loesungen/forderungsmanagement/aemterfinden/")
	request.Header.Add("User-Agent", "Go")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	response, err = client.Do(request)
	if err != nil {
		return source_1_2, errors.New("Error #2")
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&source_1_1)
	if err != nil {
		panic(err)
		return source_1_2, errors.New("Error #3")
	}

	data = url.Values{}
	data.Add("place", place)
	data.Add("action", "aemterfinden_result")
	data.Add("addressObject[Aktiv]", strconv.FormatBool(source_1_1.Data[0].Aktiv))
	data.Add("addressObject[AlternativeSuchbegriffe][string]", source_1_1.Data[0].AlternativeSuchbegriffe.String)
	data.Add("addressObject[AlternativeSuchbegriffeAsSearchString]", source_1_1.Data[0].AlternativeSuchbegriffeAsSearchString)
	data.Add("addressObject[AlternativeSuchbegriffeAsString]", source_1_1.Data[0].AlternativeSuchbegriffeAsString)
	data.Add("addressObject[BfsNr]", source_1_1.Data[0].BfsNr)
	data.Add("addressObject[HausKey]", strconv.Itoa(source_1_1.Data[0].HausKey))
	data.Add("addressObject[HausNummer]", strconv.Itoa(source_1_1.Data[0].HausNummer))
	data.Add("addressObject[HausNummerAlpha]", source_1_1.Data[0].HausNummerAlpha)
	data.Add("addressObject[Kanton]", source_1_1.Data[0].Kanton)
	data.Add("addressObject[Land]", source_1_1.Data[0].Land)
	data.Add("addressObject[NameComplete]", source_1_1.Data[0].NameComplete)
	data.Add("addressObject[Onrp]", source_1_1.Data[0].Onrp)
	data.Add("addressObject[Ort]", source_1_1.Data[0].Ort)
	data.Add("addressObject[Postleitzahl]", source_1_1.Data[0].Postleitzahl)
	data.Add("addressObject[Quartier]", source_1_1.Data[0].Quartier)
	data.Add("addressObject[SprachCode]", source_1_1.Data[0].SprachCode)
	data.Add("addressObject[Stadtkreis]", source_1_1.Data[0].Stadtkreis)
	data.Add("addressObject[StrassenName]", source_1_1.Data[0].StrassenName)
	data.Add("amtTyp", amt)

	request, err = http.NewRequest(
		"POST",
		"https://www.egeli-informatik.ch/prd/wp-admin/admin-ajax.php",
		bytes.NewBufferString(data.Encode()),
	)
	if err != nil {
		return source_1_2, errors.New("Error #4")
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("DNT", "1")
	request.Header.Add("Host", "www.egeli-informatik.ch")
	request.Header.Add("Origin", "https://www.egeli-informatik.ch")
	request.Header.Add("Referer", "https://www.egeli-informatik.ch/unsere_loesungen/forderungsmanagement/aemterfinden/")
	request.Header.Add("User-Agent", "Go")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	response, err = client.Do(request)
	if err != nil {
		return source_1_2, errors.New("Error #5")
	}

	defer response.Body.Close()

	body_bytes, err := ioutil.ReadAll(response.Body)

	body_string := string(body_bytes)

	reader := strings.NewReader(body_string)
	root, err := xmlpath.ParseHTML(reader)
	if err != nil {
		return source_1_2, errors.New("Error #6")
	}

	path = `//li/div[@class="result"]/h2/text()`
	xpath = xmlpath.MustCompile(path)
	value, ok = xpath.String(root)
	if ok {
		source_1_2.Amt = get_text(value)
	}

	path = `//li/div[@class="result"]/div[@class="column"]/p[@class="eschkg_id"]/text()`
	xpath = xmlpath.MustCompile(path)
	value, ok = xpath.String(root)
	if ok {
		source_1_2.SedexId = get_text(value)
	}

	return source_1_2, nil
}

func get_source_2(settings *Settings, street string, number string, zip string, city string) (Source2, error) {
	var source_2 Source2
	var err error

	client := get_http_client(settings)

	request, err := http.NewRequest("GET", "http://tilbago.k-infinity.com:2607/dev/amtinfo", nil)
	if err != nil {
		return source_2, errors.New("Error #1")
	}

	request.Header.Add("Host", "tilbago.k-infinity.com:2607")
	request.Header.Add("User-Agent", "Go")

	query := request.URL.Query()
	query.Add("city", city)
	query.Add("number", number)
	query.Add("street", street)
	query.Add("zip", zip)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		return source_2, errors.New("Error #2")
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&source_2)
	if err != nil {
		return source_2, errors.New("Error #3")
	}

	return source_2, nil
}

func get_text(text string) string {
	texts := strings.SplitN(text, ":", 2)
	text = texts[len(texts)-1]
	text = strings.TrimSpace(text)
	return text
}
