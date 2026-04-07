package main

import (
	"errors"
	"fmt"
)

type Deployable interface {
	Name() string
	Check() error
}

type WebApp struct {
	URL    string
	Status string
}

func (w WebApp) Name() string { return w.URL }

func (w WebApp) Check() error {
	if w.Status == "offline" {
		return errors.New("sistema está offline")
	}
	return nil
}

type Database struct {
	Connection string
	Engine     string
}

func (d Database) Name() string { return d.Connection }

func (d Database) Check() error {
	if d.Engine == "" {
		return errors.New("engine de banco de dados não informada")
	}
	return nil
}

func RunPreDeploy(resources ...Deployable) {
	for _, res := range resources {
		fmt.Printf("\n--- Analisando: %s ---\n", res.Name())

		err := res.Check()
		if err != nil {
			fmt.Printf("DEPLOY CANCELADO para [%s]: %v\n", res.Name(), err)
			continue
		}

		fmt.Printf("%s passou nos testes!\n", res.Name())
	}
}

func main() {
	app := WebApp{URL: "https://minha-api.com", Status: "online"}
	dbError := Database{Connection: "postgres://localhost", Engine: ""}
	dbOk := Database{Connection: "mysql://prod", Engine: "mysql"}

	listaDeDeploy := []Deployable{app, dbError, dbOk}

	RunPreDeploy(listaDeDeploy...)
}
