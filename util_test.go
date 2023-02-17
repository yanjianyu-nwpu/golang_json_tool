package golang_json_tool

import "testing"

func TestGetStringSplitBySpace(t *testing.T) {

	testR := "A int64 `json:\"A,omitemtpy\"`  `"

	tmp := GetStringSplitBySpace(testR)
	for _, tt := range tmp {
		t.Log(tt)
	}
}
