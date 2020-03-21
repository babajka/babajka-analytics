package main

import (
	"fmt"
	"sort"

	"github.com/babajka/babajka-analytics/babajka"
)

func printTable(metrics babajka.Metrics) {
	fmt.Printf("\nYandex Metrics for Production wir.by\n\n")
	fmt.Printf("%30sbe   ru   en\n", "")
	rows := []string{}
	for slug, metric := range metrics {
		tmp := fmt.Sprintf("%-30s", slug)
		for _, loc := range []string{"be", "ru", "en"} {
			tmp += fmt.Sprintf("%-3d  ", metric[loc])
		}
		tmp += "\n"
		rows = append(rows, tmp)
	}
	sort.Strings(rows)
	for _, row := range rows {
		fmt.Print(row)
	}
	fmt.Println()
}
