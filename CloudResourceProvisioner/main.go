package main

import (
	"fmt"
	"sync"
	"time"
)

func provisionServer(id int, ch chan<- string, wg *sync.WaitGroup) {
	// TODO: 1. Garanta que o wg.Done() seja chamado ao sair
	defer wg.Done()
	fmt.Printf("[Work] Iniciando criação do servidor %d...\n", id)
	time.Sleep(time.Duration(id) * 500 * time.Millisecond) // Simula tempos diferentes

	// TODO: 2. Envie o ID "server-id-X" para o canal
	ch <- fmt.Sprintf("server-id-%d", id)
}

func main() {
	var wg sync.WaitGroup
	serverChan := make(chan string)

	// Disparando o provisionamento de 5 servidores
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go provisionServer(i, serverChan, &wg)
	}

	// TODO: 3. Crie uma goroutine anônima que espera o WaitGroup (wg.Wait())

	// e depois fecha o canal (close(serverChan))
	// DICA: go func() { ... }()
	go func() {
		wg.Wait()
		close(serverChan)
	}() // FUncao anonima

	fmt.Println("[Main] Coletando IDs dos recursos criados...")

	// TODO: 4. Use um 'for res := range serverChan' para imprimir os resultados
	for res := range serverChan {
		fmt.Println(res)
	}
	fmt.Println("[Main] Provisionamento completo. Estado salvo.")
}
