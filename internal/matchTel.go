package internal

import (
	"regexp"
)

func MatchTel(tel string) bool {
	t := []byte(tel)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}
