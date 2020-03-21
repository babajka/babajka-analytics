package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/babajka/babajka-analytics/babajka"
)

var secretPath string

func main() {
	flag.StringVar(&secretPath, "secretPath", "", "Path to secret configuration file")
	flag.Parse()

	babajkaClient, err := babajka.NewClient(secretPath)
	if err != nil {
		log.Fatal(err)
	}

	err = babajkaClient.UpdateAnalytics()
	if err != nil {
		fmt.Println("failed to update analytics:", err)
	}

	// metrics := babajka.GetContentMetrics()
	// printTable(metrics)
}
