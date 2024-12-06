package repo

import (
	"fmt"
)

var allowedColumns = map[string]bool{
	"firstname":            true,
	"middlename":           true,
	"lastname":             true,
	"gender":               true,
	"birthday":             true,
	"form_of_education":    true,
	"personal_file_number": true,
	"phone":                true,
	"email":                true,
	"title":                true,
	"number_of_hours":      true,
	"":                     true,
}

type QueryParams struct {
	Query        string `db:"query" param:"query" query:"query" form:"query"`
	OrderBy      string `db:"order_by" param:"order_by" query:"order_by" form:"order_by"`
	IsDescending bool   `db:"is_descending" param:"is_descending" query:"is_descending" form:"is_descending"`
}

func (q *QueryParams) Validate(params ...QueryParams) error {
	defaults := new(QueryParams)
	if len(params) != 0 {
		defaults = &params[0]
	}

	if !allowedColumns[q.OrderBy] {
		return fmt.Errorf("Invalid order_by column: %s", q.OrderBy)
	}

	if q.OrderBy == "" {
		if defaults.OrderBy != "" {
			q.OrderBy = defaults.OrderBy
		} else {
			q.OrderBy = "lastname"
		}
	}

	if q.Query == "" {
		if defaults.Query != "" {
			q.Query = defaults.Query
		} else {
			q.Query = "%"
		}
	} else {
		q.Query = "%" + q.Query + "%"
	}

	return nil
}
