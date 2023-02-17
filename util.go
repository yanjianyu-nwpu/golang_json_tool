package golang_json_tool

import (
	"regexp"
	"strings"
)

var (
	spaceReg *regexp.Regexp
)

func init() {
	spaceReg = regexp.MustCompile(`[\s,]`)
}

// GetStringSplitBySpace split space string
func GetStringSplitBySpace(s string) []string {
	res := strings.Split(s, " ")
	tmp := make([]string, 0)

	for _, term := range res {
		if term == " " || term == "" {
			continue
		}

		termPure := spaceReg.ReplaceAll([]byte(term), []byte{})
		tmp = append(tmp, string(termPure))
	}

	return tmp
}
