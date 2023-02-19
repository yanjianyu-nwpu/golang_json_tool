package golang_json_tool

type Object struct {
	ObjectName string // 结构体的名字

	CodeObjectName string // 生成代码的名字

	RawElemList []string // 输入的raw string

	Elems []*DataElem

	Name2Elem     map[string]*DataElem
	NormName2Elem map[string]*DataElem

	JsonN2Elem    map[string]*DataElem
	NormJson2Elem map[string]*DataElem

	SelectedStructName map[string]bool
}

func NewObject() *Object {
	return &Object{
		SelectedStructName: make(map[string]bool),
		JsonN2Elem:         make(map[string]*DataElem),
		Name2Elem:          make(map[string]*DataElem),
		NormName2Elem:      make(map[string]*DataElem),
		NormJson2Elem:      make(map[string]*DataElem),
	}
}
