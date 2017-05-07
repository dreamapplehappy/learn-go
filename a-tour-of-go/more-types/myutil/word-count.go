package myutil

import "strings"

func WordCount(s string) map[string]int {
	return map[string]int{s: int(len(strings.Fields(s)))}
}