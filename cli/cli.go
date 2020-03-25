package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/babajka/babajka-analytics/babajka"
)

func main() {
	var secretPath, env string

	flag.StringVar(&secretPath, "secretPath", "", "Path to secret configuration file")
	// TODO: to consider retrieving env from a secret file.
	flag.StringVar(&env, "env", "", "Environment")

	flag.Parse()

	babajkaClient, err := babajka.NewClient(secretPath, env)
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
