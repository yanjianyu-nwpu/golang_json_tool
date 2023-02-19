package golang_json_tool

import (
	"fmt"
	"strings"
)

func ParserStruct(rawString string) (*Object, error) {
	obj := NewObject()

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
	for ind, line := range lines {
		if ind == 0 {
			continue
		}
		//fmt.Println(line)
		if line != "}" && line != "" {
			rawDataElemStrList = append(rawDataElemStrList, line)
		}
	}
	obj.RawElemList = rawDataElemStrList

	// 解析每个数据成员
	for _, dataRawSter := range rawDataElemStrList {
		//fmt.Println(dataRawSter)
		dataElem := ParseDataElem(dataRawSter)
		if dataElem == nil {
			continue
		}

		obj.Elems = append(obj.Elems, dataElem)
		if dataElem.Name == "" {
			continue
		}
		obj.Name2Elem[dataElem.Name] = dataElem
		normName := getNormName(dataElem.Name)
		if normName != "" {
			obj.NormName2Elem[normName] = dataElem
		}

		if dataElem.JsonName != "" {
			obj.JsonN2Elem[dataElem.JsonName] = dataElem
			normJsonName := getNormName(dataElem.JsonName)
			obj.NormJson2Elem[normJsonName] = dataElem
		}
	}

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

// 归一化的名字删除下划线，全部变成小写
func getNormName(s string) string {
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ToLower(s)
	return s
}
