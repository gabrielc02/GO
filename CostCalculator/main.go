package main

import "fmt"

type Instance struct {
	Name        string
	Type        string
	CostPerHour float64
}

func main() {

	AWSservers := []Instance{
		{Name: "Atena", Type: "t3.micro", CostPerHour: 0.02},
		{Name: "Perseu", Type: "lambda", CostPerHour: 0.50},
		{Name: "Hades", Type: "S3 bucket", CostPerHour: 1.20},
	}

	var total float64

	for _, srv := range AWSservers {
		cost := srv.CostPerHour * 24

		total += cost

		if cost > 10.0 {
			fmt.Printf("ALERTA: Instância [%-2s] acima do orcamento!\n", srv.Name)
		}

	}

	fmt.Printf("aily Total is: %-1f$\n", total)
}
