package main

import (
	"strings"
)

func parseSshConfig(config []byte) ([]string, error) {
	content := string(config)

	lines := strings.Split(content, "\n")
	entries := []string{}

	for _, v := range lines {
		if strings.Contains(v, "Host ") {
			// remove all spaces from the start of the string
			v = strings.Trim(v, " \t")

			// ignore commented hosts
			if strings.HasPrefix(v, "#") {
				continue
			}

			// take only Host's name
			host, _ := strings.CutPrefix(v, "Host ")
			host = strings.Trim(host, " \"")
			entries = append(entries, host)
		}
	}

	return entries, nil
}
