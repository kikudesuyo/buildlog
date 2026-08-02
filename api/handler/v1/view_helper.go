package v1

import (
	"strconv"
	"strings"
)

func incrementViewString(value string) string {
	digits := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, value)
	count, _ := strconv.ParseInt(digits, 10, 64)
	return strconv.FormatInt(count+1, 10)
}
