package main

import (
	"errors"
	"fmt"
)

func LoadConfig(env string, keys ...string) error {
	defer fmt.Println("Arquivo de configuracao fechado!\n	")

	if env == "" {
		return errors.New("Ambiente nao especificado")
	}

	for _, k := range keys {
		if k == "DATABASE_URL" {

			fmt.Printf(" -> Chave carregada: %s\n", k)

		}
	}

	return nil
}

func main() {
	env := "production"
	keys := []string{"PORT", "DATABASE_URL", "LOG_LEVEL"}
	LoadConfig(env, keys...)

	err := LoadConfig("")
	if err != nil {
		fmt.Printf("ERRO CAPTURADO: %v\n", err)
	}

}
