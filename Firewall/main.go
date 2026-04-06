package main

import (
	"fmt"
	"strconv" // Dica: Use strconv.Atoi para converter string para int
	"strings"
)

func main() {
	// Dados brutos que chegaram do seu sistema de log (Simulando o Ch 2 e 3)
	rawLogs := []string{"192.168.1.1,80,0.5", "10.0.0.5,443,0.1", "172.16.0.1,22,0.9", "1.1.1.1,8080,0.95"}

	// Ch 3: Lista de IPs que vamos bloquear
	var blacklist []string
	var isDangerous bool = false

	fmt.Println("Analisando logs de tráfego...")

	for _, log := range rawLogs {
		// Dica: Para simplificar, não vamos dar "split" na string agora.
		// Vamos fingir que você já pegou os valores.
		// Foque na LÓGICA de decisão e TIPOS.
		data := strings.Split(log, ",")
		ip := data[0]        // finja que aqui tem o IP
		portStr := data[1]   // finja que veio como string
		riskScore := data[2] // float64 (Ch 2)

		// --- DESAFIO CH 2: CONVERSÃO ---
		// TODO: Converta portStr para um inteiro chamado 'port' usando strconv.Atoi
		port, err := strconv.Atoi(portStr)
		// Se houver erro na conversão, apenas ignore este log (use 'continue')
		if err != nil {
			// Trate o erro aqui: logue, retorne ou pare a execução
			fmt.Println("Erro detectado:", err)
			continue
		}

		risk, err := strconv.ParseFloat(riskScore, 64)
		if err != nil {
			fmt.Println("Erro detectado", err)
		}

		// --- DESAFIO CH 4: DECISÃO DE ARQUITETO ---
		// Regra de bloqueio:
		// Se o riskScore for maior que 0.8 OU se a porta for a 22 (SSH)
		if risk > 0.8 || port == 22 {
			isDangerous = true
		}

		if isDangerous {
			// --- DESAFIO CH 6: APPEND ---
			// TODO: Adicione o IP na 'blacklist' usando a função append()
			blacklist = append(blacklist, ip)
			fmt.Printf("ALERTA: IP %s marcado para bloqueio!\n", ip)
		}
	}

	// Exibição Final
	fmt.Println("\n--- IPs para bloquear no Firewall (Iptables/NFTables) ---")
	fmt.Println(blacklist)
}
