package rssfilter

import (
	"regexp"
	"strings"
)

type predicate func(item string) bool

func match(pattern, str string) (bool, error) {
	if pattern[0] == '/' && pattern[len(pattern)-1] == '/' {
		// match against regex
		return regexp.Match(pattern[1:len(pattern)-1], []byte(str))
	} else {
		// match against substring
		return strings.Contains(str, pattern), nil
	}
	return false, nil
}
