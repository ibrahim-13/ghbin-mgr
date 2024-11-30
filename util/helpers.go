package util

import "strings"

func ContainsAllMatches(str string, match ...string) bool {
	s := strings.ToLower(str)
	matchCount := 0
	for i := range match {
		if strings.Contains(s, strings.ToLower(match[i])) {
			matchCount += 1
		}
	}
	return len(match) == matchCount
}
