package golang_json_tool

import "testing"

func TestTagReg(t *testing.T) {

	testR := "A int64 `json:\"A,omitemtpy\"`  `"

	r := tagReg.Find([]byte(testR))
	t.Log(string(r))
}

