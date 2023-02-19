package golang_json_tool

import "fmt"

var (
	dataType = map[string]bool{
		"int":     true,
		"uint":    true,
		"int8":    true,
		"uint8":   true,
		"int32":   true,
		"uint32":  true,
		"int64":   true,
		"uint64":  true,
		"float":   true,
		"float32": true,
		"float64": true,
	}
)

const (
	OKLine             = "\t%s.%s = %s.%s\n"
	OKLineWithTypeCase = "\t%s.%s = %s(%s.%s)\n" //强制转化
	NotOKLine          = "\t// %s.%s = %s.%s\n"
	NotOKLineOnlyDst   = "\t// %s.%s = \n"

	FuncLine = "func(%s *%s, %s *%s) { \n "
)

/*
	字段匹配规则
	1 name 一样 type一样，数值int float可以强制转化
	2 json name一样
	3 norName一样
	4 normJsonname一样


	注释内容
	1 只管被拷贝内容的需要
	2 不管数据type 需要

*/
func GenCopyCode(srcStr string, dstStr string) string {
	src, e0 := ParserStruct(srcStr)
	dst, e1 := ParserStruct(dstStr)

	if e0 != nil || e1 != nil {
		return ""
	}

	okContent := ""    // 匹配上的内容
	notOKContent := "" // 注释内容
	for _, elem := range dst.Elems {
		fieldName := elem.Name
		normFieldName := getNormName(fieldName)

		if srcElem, ok := src.Name2Elem[fieldName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, true)
			okContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}

		if srcElem, ok := src.NormJson2Elem[normFieldName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, true)
			okContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}

		jsonName := elem.JsonName
		normJsonName := getNormName(jsonName)

		if srcElem, ok := src.JsonN2Elem[jsonName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, true)
			okContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}

		if srcElem, ok := src.NormJson2Elem[normJsonName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, true)
			okContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}
	}

	for _, elem := range dst.Elems {
		fieldName := elem.Name
		normFieldName := getNormName(fieldName)
		if dst.SelectedStructName[fieldName] {
			continue
		}

		if srcElem, ok := src.Name2Elem[fieldName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, false)

			notOKContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}

		if srcElem, ok := src.NormJson2Elem[normFieldName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, false)
			notOKContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}

		jsonName := elem.JsonName
		normJsonName := getNormName(jsonName)
		if srcElem, ok := src.JsonN2Elem[jsonName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, false)
			notOKContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}

		if srcElem, ok := src.NormJson2Elem[normJsonName]; ok && srcElem != nil {
			tmp := getOKCode(src, dst, srcElem.Name, fieldName, false)
			notOKContent += tmp
			if tmp != "" {
				dst.SelectedStructName[fieldName] = true
				continue
			}
		}
	}

	for _, elem := range dst.Elems {
		fieldName := elem.Name
		if dst.SelectedStructName[fieldName] {
			continue
		}

		notOKContent += fmt.Sprintf(NotOKLineOnlyDst, dst.ObjectName, fieldName)
	}

	funcLine := fmt.Sprintf(FuncLine, src.CodeObjectName, src.ObjectName, dst.CodeObjectName, dst.ObjectName)
	return funcLine + okContent + notOKContent + "}"
}

// 匹配ok的代码，type ok 不
func getOKCode(src, dst *Object, srcFiledName, dstFieldName string, needTypeOK bool) string {
	srcObjName := src.ObjectName
	dstObjName := dst.ObjectName
	srcFildType := src.Name2Elem[srcFiledName].TypeStr
	dstFildType := dst.Name2Elem[dstFieldName].TypeStr

	//isOK := false
	// type 完全一样
	if srcFildType == dstFildType {
		res := fmt.Sprintf(OKLine, dstObjName, dstFieldName, srcObjName, srcFiledName)
		return res
	}

	if dataType[srcFildType] && dataType[dstFildType] {
		res := fmt.Sprintf(OKLineWithTypeCase, dstObjName, dstFieldName, dstFildType, srcObjName, srcFiledName)
		return res
	}

	if needTypeOK {
		return ""
	}

	return fmt.Sprintf(NotOKLine, dstObjName, dstFieldName, srcObjName, srcFiledName)
}
