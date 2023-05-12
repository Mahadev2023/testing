// Package is two method as below
package word

import "strings"

//UseCount take string and map word and that count
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count method take string and return the count of word
//int as out
func Count(s string) int {

	return len(strings.Fields(s))
}
