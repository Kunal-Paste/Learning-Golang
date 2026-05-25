package greet

import "strings"

//exporting function always start with capital letter
func Hello(name string) string {
	clean := normalizeName(name)

	return "Hello " + clean
}

func normalizeName(name string) string {
	n := strings.TrimSpace(name)

	if n == "" {
		return "Guest"
	}

	return strings.ToUpper(name)
}
