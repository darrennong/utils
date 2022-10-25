package gorm

import (
	"database/sql/driver"
	"encoding/json"
)

type int32s []int32

// Int32Array 自定义Gorm字段类型，必须使用指针类型
type Int32Array struct {
	*int32s
}

// String -> String Representation of Binary16
func (arr *Int32Array) String() string {
	bytes, _ := json.Marshal(arr.int32s)
	return string(bytes)
}

// GormDataType -> sets type to binary(16)
func (arr *Int32Array) GormDataType() string {
	return "text"
}

func (arr *Int32Array) MarshalJSON() ([]byte, error) {
	items := ([]int32)(arr.int32s)
	return json.Marshal(items)
}

func (arr *Int32Array) UnmarshalJSON(by []byte) error {
	return json.Unmarshal(by, &(arr.int32s))
}

// Scan --> tells GORM how to receive from the database
func (arr *Int32Array) Scan(value interface{}) error {
	data, _ := value.([]byte)
	return json.Unmarshal(data, &(arr.int32s))
}

// Value -> tells GORM how to save into the database
func (arr *Int32Array) Value() (driver.Value, error) {
	return arr.String(), nil
}

func (arr *Int32Array) Length() int {
	return len(*(arr.int32s))
}

func NewInt32Array(len int) *Int32Array {
	return &Int32Array{
		make([]int32, len),
	}
}

func (arr *Int32Array) Set(i int, d int32) {
	arr.int32s[i] = d
}

func (arr *Int32Array) Get(i int) int32 {
	return arr.int32s[i]
}
