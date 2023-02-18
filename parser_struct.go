package golang_json_tool

import (
	"fmt"
	"strings"
)

func ParserStruct(rawString string) (*Object, error) {
	obj := &Object{}

	// 先split
	lines := strings.Split(rawString, "\n")
	//obj.RawElemList = lines
	if len(lines) == 0 {
		return nil, fmt.Errorf("rawstring is empty")
	}

	// 获取struct name    type Table struct 这种
	l0 := GetStringSplitBySpace(lines[0])
	if len(l0) <= 2 {
		return nil, fmt.Errorf("struct name is not type *** name")
	}

	// 生成对象的名字
	obj.ObjectName = l0[1]
	obj.CodeObjectName = strucName2CodeObjectName(obj.ObjectName)

	// 清洗数据 { } 这种
	rawDataElemStrList := make([]string, 0)
	for _, line := range lines {
		if !uselessCodeReg.Match([]byte(line)) {
			rawDataElemStrList = append(rawDataElemStrList, line)
		}
	}
	obj.RawElemList = rawDataElemStrList

	return obj, nil
}

// 生成对象名字，一般是结构体名字首字符大写变小写，如果不是那么在后面加上obj
func strucName2CodeObjectName(s string) string {
	if len(s) == 0 {
		return ""
	}

	//structNameWithOutFirst := s[1:]
	codeNameFirstAlphaCode := s
	if bigEnReg.Match([]byte(codeNameFirstAlphaCode[:1])) {
		offset := int(s[0]) - int('A')
		sSmall := rune(int('a') + offset)
		codeNameFirstAlphaCode = string([]rune{sSmall}) + codeNameFirstAlphaCode[1:]
	}

	if codeNameFirstAlphaCode == s {
		codeNameFirstAlphaCode += "Obj"
	}

	return codeNameFirstAlphaCode
}
