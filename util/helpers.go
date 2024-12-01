package util

import (
	"runtime"
	"strings"
)

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

func ContainsAnyMatches(str string, match ...string) bool {
	s := strings.ToLower(str)
	for i := range match {
		if strings.Contains(s, strings.ToLower(match[i])) {
			return true
		}
	}
	return false
}

func ParsePatternsFromString(pattern string) []string {
	var formatted_patterns []string
	for _, v := range strings.Split(pattern, ",") {
		_v := strings.ToLower(v)
		if _v == "__os__" {
			formatted_patterns = append(formatted_patterns, runtime.GOOS)
		} else if _v == "__arch__" {
			formatted_patterns = append(formatted_patterns, runtime.GOARCH)
		} else {
			formatted_patterns = append(formatted_patterns, v)
		}
	}
	return formatted_patterns
}
