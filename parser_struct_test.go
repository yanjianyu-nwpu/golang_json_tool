package golang_json_tool

/*
func TestParseStruct(t *testing.T) {
	c := "type  TestA strcut"
	line := GetStringSplitBySpace(c)
	t.Log(line)

	s := "BypeA"
	//structNameWithOutFirst := s[1:]
	if bigEnReg.Match([]byte(s[:1])) {
		offset := int(s[0]) - int('A')
		sSmall := rune(int('a') + offset)
		s = string([]rune{sSmall}) + s[1:]
	}
	//offset := int(s[0]) - int('A')
	t.Log(s)
}


func TestParseStruct(t *testing.T) {
	testR := "type TestA struct {\nTable int64 `json:\"table,omitempty\"`\n B float32\n C *int64\n}"
	tt, err := ParserStruct(testR)
	t.Logf(fmt.Sprint(err))
	t.Logf(tt.CodeObjectName)
	t.Logf(tt.ObjectName)

	ii, _ := json.Marshal(&tt.Elems)
	t.Logf(string(ii))
	ii, _ = json.Marshal(&tt.RawElemList)
	t.Logf(string(ii))
}
*/
