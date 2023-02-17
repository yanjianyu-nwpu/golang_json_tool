package golang_json_tool

type Object struct {
	ObjectName string

	RawElemList []string

	Elems []*DataElem

	Name2Elem     map[string]*DataElem
	NormName2Elem map[string]*DataElem

	JsonN2Elem    map[string]*DataElem
	NormJson2Elem map[string]*DataElem
}
