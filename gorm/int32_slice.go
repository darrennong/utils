package gorm

import (
	"database/sql/driver"
	"encoding/json"
)

type Int32s []int32

// Int32Array 自定义Gorm字段类型，必须使用指针类型
type Int32Array struct {
	*Int32s
}

// String -> String Representation of Binary16
func (arr *Int32Array) String() string {
	bytes, _ := json.Marshal(arr.Int32s)
	return string(bytes)
}

// GormDataType -> sets type to binary(16)
func (arr *Int32Array) GormDataType() string {
	return "text"
}

func (arr *Int32Array) MarshalJSON() ([]byte, error) {
	//items := any(arr.Int32s).([]int32)
	return json.Marshal(arr.Int32s)
}

func (arr *Int32Array) UnmarshalJSON(by []byte) error {
	return json.Unmarshal(by, &(arr.Int32s))
}

// Scan --> tells GORM how to receive from the database
func (arr *Int32Array) Scan(value interface{}) error {
	data, _ := value.([]byte)
	return json.Unmarshal(data, &(arr.Int32s))
}

// Value -> tells GORM how to save into the database
func (arr *Int32Array) Value() (driver.Value, error) {
	return arr.String(), nil
}

func (arr *Int32Array) Length() int {
	return len(*(arr.Int32s))
}
