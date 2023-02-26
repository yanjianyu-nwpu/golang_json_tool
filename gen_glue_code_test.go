package golang_json_tool

import "testing"

/*
func TestGenGlueCode(t *testing.T) {
	testSrc := "type TestA struct {\nTable int64 `json:\"table,omitempty\"`\n B float32\n C int64 \n }"
	testDst := "type TestB struct {\nTableA int `json:\"table\"`\n B few \n C int64 \n  }"

	k := GenCopyCode(testSrc, testDst)
	t.Log(k)
}
*/
func TestGenGlueCode(t *testing.T) {
	testSrc := "type TestA struct {\nA *int64 \n B float32\n C int64 \n D *u \n}"
	testDst := "type TestB struct {\nA int \n B float64 \n C *int \n  D u \n}"

	k := GenCopyCode(testSrc, testDst)
	t.Log(k)
}
