package audit

import (
	"infra-metrics/internal/io"
)

func AuditLog(logs []io.LogFormat) ([]io.LogFormat, error) {
	var log_resume []io.LogFormat

	for _, lev := range logs {

		if lev.Level == "ERROR" {
			log_resume = append(log_resume, lev)
		}

	}

	return log_resume, nil

}

func OcurrenceService(log_resume []io.LogFormat) (map[string]int, error) {
	log_ocurrence := make(map[string]int, 10)

	for _, ocu := range log_resume {
		log_ocurrence[ocu.Server]++
	}

	return log_ocurrence, nil
}
