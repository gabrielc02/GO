package main

import "fmt"

type Pod struct {
	Name      string
	Namespace string
	MemoryMB  int
}

func ValidatePod(currentMemory int, limit int, pod Pod) error {
	defer fmt.Println("--- Processamento do pod finalizado ---")
	// ERROR FORMAT
	if (currentMemory + pod.MemoryMB) > limit {
		return fmt.Errorf("estouro de memória: o pod [%s] precisa de %dMB, mas só restam %dMB",
			pod.Name, pod.MemoryMB, limit-currentMemory)
	}

	return nil

}

func main() {
	maxClusterMemory := 1000
	currentMemory := 0

	pods := []Pod{
		{Name: "Auth-1", Namespace: "AUTH", MemoryMB: 200},
		{Name: "DB-1", Namespace: "DB", MemoryMB: 900},
	}

	for _, p := range pods {
		err := ValidatePod(currentMemory, maxClusterMemory, p)

		if err != nil {
			fmt.Printf("Falha no Agendamento: %v\n", err)
			continue
		}

		currentMemory += p.MemoryMB
		fmt.Printf("Sucesso: Pod [%s] agendado!\n", p.Name)
	}
}
