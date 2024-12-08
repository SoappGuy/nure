package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalParam(param string) error {
	t, err := time.Parse(`2006-01-02`, param)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalParam() (string, error) {
	return time.Time(d).Format(`2006-01-02`), nil
}

func (d Date) Format(format string) string {
	return time.Time(d).Format(format)
}

// Scan: Implement the Scanner interface for reading from the database
func (d *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*d = Date(v)
		return nil

	case string:
		t, err := time.Parse(`2006-01-02`, v)
		if err != nil {
			return fmt.Errorf("failed to parse Date from string: %w", err)
		}
		*d = Date(t)
		return nil

	case []byte:
		t, err := time.Parse(`2006-01-02`, string(v))
		if err != nil {
			return fmt.Errorf("failed to parse Date from bytes: %w", err)
		}
		*d = Date(t)
		return nil
	default:
		return fmt.Errorf("unsupported type for Date: %T", value)
	}
}

// Value: Implement the Valuer interface for writing to the database
func (d Date) Value() (driver.Value, error) {
	return time.Time(d).Format(`2006-01-02`), nil
}

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (nt *NullTime) UnmarshalParam(param string) error {
	t, err := time.Parse(`15:04`, param)
	if err != nil {
		return err
	}
	(*nt).Time = t.AddDate(1, 0, 0)
	(*nt).Valid = true
	return nil
}

func (nt NullTime) MarshalParam() (string, error) {
	return nt.Time.Format(`15:04`), nil
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
