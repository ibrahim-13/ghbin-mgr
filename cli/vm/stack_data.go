package vm

type DataType string

const (
	DT_STRING DataType = "string"
	DT_INT    DataType = "int"
	DT_FLOAT  DataType = "float"
)

type Data struct {
	Type DataType
	data any
}

func NewData(t DataType, d any) Data {
	return Data{Type: t, data: d}
}

func (d Data) GetData() any {
	return d.data
}

func (d Data) GetInt() int {
	return d.data.(int)
}

func (d Data) GetFloat() float64 {
	return d.data.(float64)
}

func (d Data) GetString() string {
	return d.data.(string)
}
