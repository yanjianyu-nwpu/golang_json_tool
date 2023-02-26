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

/*
func TestGetUselessReg(t *testing.T) {
	t0 := []string{"{", "}", " \n", " ", "  "}
	for _, str := range t0 {
		if !uselessCodeReg.Match([]byte(str)) {
			t.Errorf("wrong %s", str)
		}
	}
}

func TestJsonTagReg(t *testing.T) {
	s := "`json:\"name\"`"

	r := jsonTagReg.Find([]byte(s))

	t.Log(string(r))
}
*/

func TestPtrReg(t *testing.T) {
	h := "*int"
	h0 := "int*64"

	if !ptrTagReg.Match([]byte(h)) {
		t.Errorf("wrong0")
	}
	if ptrTagReg.Match([]byte(h0)) {
		t.Errorf("w1")
	}

	jj := ptrTagReg.ReplaceAll([]byte(h), []byte(""))
	t.Logf(string(jj))
}
