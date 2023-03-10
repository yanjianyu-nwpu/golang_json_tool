package golang_json_tool

import "regexp"

var (
	tagReg *regexp.Regexp
)

func init() {

	tagReg = regexp.MustCompile("`.*?`")
}

type DataElem struct {
	Name      string
	TypeStr   string
	JsonName  string
	Omitempty bool
	IsPtrType bool
	Tags      map[string]string
	RawString string
	RawTerms  []string
}
