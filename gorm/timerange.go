package gorm

import (
	"database/sql/driver"
	"fmt"
	"github.com/darrennong/utils/errors"
	"time"
)

type TimeRange struct {
	Begin time.Time
	End   time.Time
}

// String -> String Representation of Binary16
func (t *TimeRange) String() string {
	bytes, _ := t.MarshalJSON()
	return string(bytes)
}

// GormDataType -> sets type to binary(16)
func (t *TimeRange) GormDataType() string {
	return "text"
}

func (t *TimeRange) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("{\"Begin\":\"%s\",\"End\":\"%s\"}", t.Begin.Format("2006-01-02 15:04:05"), t.End.Format("2006-01-02 15:04:05")) // Format内即是你想转换的格式
	return []byte(stamp), nil
}

func (t *TimeRange) UnmarshalJSON(by []byte) error {
	var begin, end string
	var err error
	if _, err = fmt.Scanf("{\"Begin\":\"%s\",\"End\":\"%s\"}", &begin, &end); !errors.Assert(err) {
		return err
	}
	t.Begin, err = time.Parse("2006-01-02 15:04:05", begin)
	if errors.Assert(err) {
		t.End, err = time.Parse("2006-01-02 15:04:05", end)
	}
	return err
}

// Scan --> tells GORM how to receive from the database
func (t *TimeRange) Scan(value interface{}) error {
	var begin, end int64
	_, err := fmt.Scanf("%x|%x", &begin, &end)
	t.Begin = time.UnixMilli(begin)
	t.End = time.UnixMilli(end)
	return err
}

// Value -> tells GORM how to save into the database
func (t *TimeRange) Value() (driver.Value, error) {
	return fmt.Sprintf("%x|%x", t.Begin.UnixMilli(), t.End.UnixMilli()), nil
}
