package main

import (
	"flag"
	"github.com/getsentry/raven-go"
)

func main() {
	action := flag.String("action", "", "")

	flag.Parse()

	var settings *Settings

	settings = get_settings()

	raven.SetDSN(settings.Sentry.Dsn)

	if *action == "bootstrap" {
		bootstrap(settings)
	}

	if *action == "insert" {
		insert(settings)
	}

	if *action == "test_source_1" {
		test_source_1(settings)
	}

	if *action == "test_source_2" {
		test_source_2(settings)
	}

	if *action == "queue_1" {
		queue_1(settings)
	}

	if *action == "queue_2" {
		queue_2(settings)
	}

	if *action == "workers_1" {
		workers_1(settings)
	}

	if *action == "workers_2" {
		workers_2(settings)
	}

	if *action == "report" {
		report(settings)
	}
}
