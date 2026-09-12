package main

import ( "strings" )
func parseLines(value string) []string {
	var lines []string

	for _, line := range strings.Split(value, "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	return lines
}