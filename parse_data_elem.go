package golang_json_tool

import (
	"strings"
)

func ParseDataElem(s string) *DataElem {
	elem := &DataElem{}

	terms := GetStringSplitBySpace(s)
	if len(terms) <= 1 {
		return nil
	}

	elem.Name = terms[0]
	rawTypeStr := terms[1]

	if ptrTagReg.Match([]byte(rawTypeStr)) {
		elem.IsPtrType = true
		rawTypeStr = strings.ReplaceAll(rawTypeStr, "*", "")
	}

	elem.TypeStr = rawTypeStr

	if len(terms) >= 3 {
		elem.JsonName = parseJsonTag(terms[2])
	}

	return elem
}

func parseJsonTag(s string) string {
	raw := jsonTagReg.Find([]byte(s))

	rawStr := string(raw)
	if rawStr == "" {
		return ""
	}

	rawStr = strings.ReplaceAll(rawStr, "json:", "")
	rawStr = strings.ReplaceAll(rawStr, "\"", "")

	terms := strings.Split(rawStr, ",")

	// -是没有的意思，不能混用
	if len(terms) > 0 && terms[0] != " " && terms[0] != "" && terms[0] != "-" {
		return terms[0]
	}

	return ""
}
