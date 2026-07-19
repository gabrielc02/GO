package report

import "infra-audit/internal/inventory"

const (
	LevelCritical = "CRITICAL"
	LevelHigh     = "HIGH"
)

func Report(assets []inventory.Asset) []ReportedAsset {

	validInventory := []ReportedAsset{}

	for _, inv := range assets {

		if (inv.Status == "production") && (inv.CPUUsage > 80 || inv.MemoryUsage > 80) {
			var level string

			if inv.MemoryUsage > 90 && inv.CPUUsage > 90 {
				level = LevelCritical

			} else {
				level = LevelHigh
			}

			validInventory = append(validInventory, ReportedAsset{
				ID:            inv.ID,
				Hostname:      inv.Hostname,
				Status:        inv.Status,
				CPUUsage:      inv.CPUUsage,
				MemoryUsage:   inv.MemoryUsage,
				Region:        inv.Region,
				SeverityLevel: level,
			})
		}

	}

	return validInventory
}
