package utils

import "strings"

func Split(s, sep string, parts *[]string) int {
	*parts = (*parts)[:0]
	if s == "" {
		return 0
	}
	sepLen := len(sep)
	if sepLen == 0 {
		for _, r := range s {
			*parts = append(*parts, string(r))
		}
		return len(*parts)
	}
	i := 0
	for {
		idx := strings.Index(s[i:], sep)
		if idx == -1 {
			*parts = append(*parts, s[i:])
			break
		}
		*parts = append(*parts, s[i:i+idx])
		i += idx + sepLen
	}
	return len(*parts)
}
