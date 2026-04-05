/* A project to contemple the dirst 3 caps of the book go idiomatics
 */

package main

import "fmt"

type Service struct {
	Name    string
	IP      string
	Latency int
	Online  bool
}

func main() {
	// Ch 3: Um Slice (lista dinâmica) de serviços
	cluster := []Service{
		{Name: "Auth-API", IP: "10.0.0.1", Latency: 15, Online: true},
		{Name: "Payment-GW", IP: "10.0.0.2", Latency: 500, Online: true},
		{Name: "DB-Primary", IP: "10.0.0.3", Latency: 0, Online: false}, // Caiu!
		{Name: "Cache-Redis", IP: "10.0.0.4", Latency: 5, Online: true},
	}

	// Ch 3: Um Map para sumarizar o estado do cluster (Status -> Quantidade)
	summary := make(map[string]int)

	fmt.Println("--- Relatório de Infraestrutura ---")

	// Ch 4: Loop para processar os dados
	for _, srv := range cluster {
		status := "Unknown"

		// --- DESAFIO CH 4: LÓGICA DE CONTROLE ---
		// TODO: Se Online for false, status é "CRITICAL"
		// TODO: Se Online for true mas Latency > 200, status é "WARNING"
		// TODO: Se Online for true e Latency <= 200, status é "OK"

		fmt.Printf("Serviço: %-12s | Status: [%-8s] | IP: %s\n", srv.Name, status, srv.IP)

		// TODO: Incremente o contador no Map 'summary' para o status encontrado
	}

	// Exibir o resumo final do Map
	fmt.Println("\n--- Resumo do Cluster ---")
	fmt.Printf("Status Final: %+v\n", summary)
}
