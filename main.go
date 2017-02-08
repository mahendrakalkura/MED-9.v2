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

	if *action == "source_1_one" {
		source_1_one(settings)
	}

	if *action == "source_2_one" {
		source_2_one(settings)
	}

	if *action == "source_1_all" {
		source_1_all(settings)
	}

	if *action == "source_2_all" {
		source_2_all(settings)
	}

	if *action == "progress" {
		progress(settings)
	}

	if *action == "report" {
		report(settings)
	}
}
