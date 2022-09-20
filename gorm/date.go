package gorm

import "fmt"

type Date struct {
	Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", d.Format("2006-01-02")) // Format内即是你想转换的格式
	return []byte(stamp), nil
}
