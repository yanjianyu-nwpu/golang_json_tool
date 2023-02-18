package golang_json_tool

import "testing"

/*
import "testing"

func TestGetStringSplitBySpace(t *testing.T) {

	testR := "A int64 `json:\"A,omitemtpy\"`  `"

	tmp := GetStringSplitBySpace(testR)
	for _, tt := range tmp {
		t.Log(tt)
	}
}
*/

func TestGetUselessReg(t *testing.T) {
	t0 := []string{"{", "}", " \n", " ", "  "}
	for _, str := range t0 {
		if !uselessCodeReg.Match([]byte(str)) {
			t.Errorf("wrong %s", str)
		}
	}
}
