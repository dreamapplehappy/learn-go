package main

import "strings"

func WordCount(s string) map[string]int {
	return map[string]int{s: int(strings.Fields(s))}
}
