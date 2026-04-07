package main

import (
	"errors"
	"fmt"
)

type Node struct {
	ID          int
	Maintenance bool
}

func SetMaintenanceValue(n Node) {
	n.Maintenance = true
	fmt.Printf(" [Value] Mudei a COPIA do Node %d para true\n", n.ID)
}

func SetMaintenancePointer(n *Node) error {
	if n == nil {
		return errors.New("ERRO: Node inexistente")
	}

	n.Maintenance = true
	fmt.Printf(" [Pointer] Mudei o ORIGINAL do Node %d para true\n", n.ID)
	return nil
}

func main() {
	node1 := Node{ID: 101, Maintenance: false}

	fmt.Println("Status Inicial:", node1.Maintenance)

	SetMaintenanceValue(node1)
	fmt.Println("Após Valor:", node1.Maintenance)

	err := SetMaintenancePointer(&node1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Após Ponteiro:", node1.Maintenance)
	var nodeNulo *Node
	err = SetMaintenancePointer(nodeNulo)
	if err != nil {
		fmt.Printf("Segurança: %v\n", err)
	}
}
