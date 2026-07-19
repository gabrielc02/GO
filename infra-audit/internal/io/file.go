package io

import (
	"encoding/json"
	"fmt"
	"infra-audit/internal/inventory"
	"os"
)

func Reader() ([]inventory.Asset, error) {

	file, err := os.ReadFile("./inventory.json")
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo: %w", err)
	}

	var assets []inventory.Asset

	err = json.Unmarshal(file, &assets)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter JSON: %w", err)
	}

	return assets, nil
}
