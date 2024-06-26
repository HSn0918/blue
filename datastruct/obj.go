package datastruct

const (
	Number = iota
	String
	List
	Json
)

type BlueType uint8

var BlueTypes = map[uint8]string{
	Number: "number",
	String: "string",
	List:   "list",
	Json:   "json",
}

var BlueTypes_ = map[string]uint8{
	"number": Number,
	"string": String,
	"list":   List,
	"json":   Json,
}

const (
	TypeMask    = 0x0F
	SubTypeMask = 0xF0
)

type BlueObj struct {
	Type uint8
}

func (obj *BlueObj) GetType() string {
	return BlueTypes[obj.Type&TypeMask]
}

func (obj *BlueObj) GetSubType() string {
	return BlueTypes[obj.Type&SubTypeMask]
}

type Value interface {
	Value() string
}
