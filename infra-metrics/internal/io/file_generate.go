package io

import (
	"encoding/json"
	"os"
)

// SAMPLE GENERATED WITH IA TO FUTURE REFERENCES
func WriterGenerate(header map[string]int, logs []LogFormat) error {
	// 1. Marshal the Header (Occurrences)
	headerData, err := json.MarshalIndent(header, "", "  ")
	if err != nil {
		return err
	}

	// 2. Marshal the Logs
	logsData, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return err
	}

	// 3. Combine them with a descriptive separator
	finalContent := append([]byte("--- SUMMARY ---\n"), headerData...)
	finalContent = append(finalContent, []byte("\n\n--- RAW LOGS ---\n")...)
	finalContent = append(finalContent, logsData...)

	// 4. Write the combined bytes to the file
	return os.WriteFile("./server_status_resume.log", finalContent, 0644)
}
