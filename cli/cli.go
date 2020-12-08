package main

import (
	"flag"
	"log"

	"github.com/babajka/babajka-analytics/babajka"
)

func main() {
	var secretPath, env string
	var enableSlack, printReport bool

	flag.StringVar(&secretPath, "secretPath", "", "Path to secret configuration file")
	// TODO: to consider retrieving env from a secret file.
	flag.StringVar(&env, "env", "", "Environment name")
	flag.BoolVar(&enableSlack, "enableSlack", false, "This options switches slack notifications on")
	flag.BoolVar(&printReport, "printReport", false, "This options produces detailed output")

	flag.Parse()

	babajkaClient, err := babajka.NewClient(secretPath, env, enableSlack)
	if err != nil {
		log.Fatal(err)
	}

	metrics, err := babajkaClient.UpdateAnalytics()
	if err != nil {
		log.Fatalf("failed to update analytics: %v", err)
	}

	if printReport {
		printTable(metrics)
	}
}
