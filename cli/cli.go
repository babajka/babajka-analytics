package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/babajka/babajka-analytics/babajka"
)

func main() {
	var secretPath, env string
	var enableSlack bool

	flag.StringVar(&secretPath, "secretPath", "", "Path to secret configuration file")
	// TODO: to consider retrieving env from a secret file.
	flag.StringVar(&env, "env", "", "Environment name")
	flag.BoolVar(&enableSlack, "enableSlack", false, "This options switches slack notifications on")

	flag.Parse()

	babajkaClient, err := babajka.NewClient(secretPath, env, enableSlack)
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
