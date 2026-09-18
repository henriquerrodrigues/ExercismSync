// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import("strings")

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
    if raw := strings.ReplaceAll(s, " ", "");len(raw) == 0{
        return ""
    }
	words := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(s, "_", " "), "-", " "))
	var acronym string
    for _, w := range words{
        
        acronym += strings.ToUpper(string(w[0]))
    }
	return acronym
}
