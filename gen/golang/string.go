package golang

import "strings"

func upCase(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[0:1]) + s[1:]
	}
	return s
}
func lowCase(s string) string {
	if len(s) > 0 {
		return strings.ToLower(s[0:1]) + s[1:]
	}
	return s
}
