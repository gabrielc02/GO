package main

import (
	"fmt"
	"infra-audit/internal/io"
	"infra-audit/internal/report"
)

func main() {
	assets, err := io.Reader()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reported := report.Report(assets)

	for _, rep := range reported {
		fmt.Printf("REPORTED - %v \n", rep)
	}

}
