package main

import "fmt"

type Pod struct {
	Name      string
	Namespace string
	MemoryMB  int
}

func main() {
	var currentMemory int
	maxClusterMemory := 1000
	podsPerNamespace := make(map[string]int)
	Pod := []Pod{
		{Name: "Pod_SP01", Namespace: "AUTH", MemoryMB: 240},
		{Name: "Pod_SP02", Namespace: "AUTH", MemoryMB: 540},
		{Name: "Pod_RJ01", Namespace: "DB", MemoryMB: 450},
		{Name: "Pod_RJ02", Namespace: "DB", MemoryMB: 30},
		{Name: "Pod_BH", Namespace: "WEB", MemoryMB: 120},
	}

	for _, pod := range Pod {

		if (currentMemory + pod.MemoryMB) > maxClusterMemory {
			fmt.Printf("REJEITADO: Pod [%s] excede a capacidade do cluster!\n", pod.Name)
			continue
		}

		currentMemory += pod.MemoryMB
		podsPerNamespace[pod.Namespace]++
	}

	fmt.Println(podsPerNamespace)
	fmt.Printf("Memoria total utilizada: %d MB\n", currentMemory)
}
