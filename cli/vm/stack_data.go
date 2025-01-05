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
