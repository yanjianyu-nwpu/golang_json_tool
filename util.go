package golang_json_tool

import (
	"regexp"
	"strings"
)

var (
	spaceReg       *regexp.Regexp
	enReg          *regexp.Regexp
	bigEnReg       *regexp.Regexp
	smallEnReg     *regexp.Regexp
	uselessCodeReg *regexp.Regexp
	jsonTagReg     *regexp.Regexp
)

func init() {
	spaceReg = regexp.MustCompile(`[\s]`)
	enReg = regexp.MustCompile(`[A-Za-z]+`)
	bigEnReg = regexp.MustCompile(`[A-Z]+`)
	smallEnReg = regexp.MustCompile(`[a-z]+`)
	jsonTagReg = regexp.MustCompile(`json:".+?"`)
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
