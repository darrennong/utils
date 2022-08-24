package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

// String -> String Representation of Binary16
func (t *Time) String() string {
	bytes, _ := json.Marshal(t.Time)
	return string(bytes)
}

// GormDataType -> sets type to binary(16)
func (t *Time) GormDataType() string {
	return "datetime(3)"
}

func (t *Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05")) // Format内即是你想转换的格式
	return []byte(stamp), nil
}

func (t *Time) UnmarshalJSON(by []byte) error {
	return json.Unmarshal(by, &(t.Time))
}

// Scan --> tells GORM how to receive from the database
func (t *Time) Scan(value interface{}) error {
	t.Time = value.(time.Time)
	return nil
}

// Value -> tells GORM how to save into the database
func (t *Time) Value() (driver.Value, error) {
	return t.Time, nil
}

func (t *Time) UnixTime(sec int64) {
	t.Time = time.Unix(sec, 0)
}
