package code

// DateTime file string
var DateTime = []byte(`package db

import (
"database/sql/driver"
"fmt"
"time"
)

type DateTime time.Time

func Now() DateTime {
 return DateTime(time.Now())
}

func (d DateTime) Value() (driver.Value, error) {
	t := time.Time(d)
	return t.Format(time.RFC3339), nil
}

func (d *DateTime) Scan(src interface{}) error {
	var source string
	switch src := src.(type) {
		case string:
			source = src
		case []byte:
			source = string(src)
		default:
			return fmt.Errorf("incompatible type")
	}

	dateTime, err := time.Parse(time.RFC3339, source)
	if err != nil {
		return err
	}

	*d = DateTime(dateTime)

	return nil
}`)
