package stackmachine

import "errors"

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

func (d Data) CompareEq(v Data) (bool, error) {
	if d.Type != v.Type {
		return false, errors.New("data type does not match for jumpeq operation")
	}
	switch d.Type {
	case DT_INT:
		return d.GetInt() == v.GetInt(), nil
	case DT_FLOAT:
		return d.GetFloat() == v.GetFloat(), nil
	case DT_STRING:
		return d.GetString() == v.GetString(), nil
	default:
		return false, errors.New("data type comparison not implemented for jumpeq operation")
	}
}

func (d Data) CompareEqN(v Data) (bool, error) {
	if d.Type != v.Type {
		return false, errors.New("data type does not match for jumpeq operation")
	}
	switch d.Type {
	case DT_INT:
		return d.GetInt() != v.GetInt(), nil
	case DT_FLOAT:
		return d.GetFloat() != v.GetFloat(), nil
	case DT_STRING:
		return d.GetString() != v.GetString(), nil
	default:
		return false, errors.New("data type comparison not implemented for jumpeq operation")
	}
}
