package io

import (
	"bufio"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReaderExtractor() []LogFormat {
	dat, err := os.Open("./server_status.log")
	check(err)
	defer dat.Close()

	re := regexp.MustCompile(`^(.{19}) \[(\w+)\] (.+) on (.+)$`)

	var logs []LogFormat
	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if len(matches) == 5 {
			entry := LogFormat{
				Timestamp: matches[1],
				Level:     matches[2],
				Message:   matches[3],
				Server:    matches[4],
			}
			logs = append(logs, entry)
		}
	}

	return logs
}
