package util

import (
	"database/sql/driver"
	"errors"
	"time"
)

type UnixTime int64

// implement sql.Scanner interface
func (t *UnixTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*t = UnixTime(v.Unix())
		return nil
	default:
		return errors.New("not a unix timestamp")
	}
}

// implement driver.Valuer interface
func (t UnixTime) Value() (driver.Value, error) {
	return time.Unix(int64(t), 0), nil
}
