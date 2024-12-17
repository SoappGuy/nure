package repo

import (
	"fmt"
	"teacher_helper/internal/model"
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
	Query           string     `db:"query" param:"query" query:"query" form:"query"`
	OrderBy         string     `db:"order_by" param:"order_by" query:"order_by" form:"order_by"`
	IsDescending    bool       `db:"is_descending" param:"is_descending" query:"is_descending" form:"is_descending"`
	Gender          string     `db:"gender" param:"gender" query:"gender" form:"gender"`
	BirthdayFrom    model.Date `db:"birthday_from" param:"birthday_from" query:"birthday_from" form:"birthday_from"`
	BirthdayTo      model.Date `db:"birthday_to" param:"birthday_to" query:"birthday_to" form:"birthday_to"`
	FormOfEducation string     `db:"form_of_education" param:"form_of_education" query:"form_of_education" form:"form_of_education"`
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

	if q.BirthdayTo.Sub(q.BirthdayFrom) < 0 {
		q.BirthdayTo, q.BirthdayFrom = q.BirthdayFrom, q.BirthdayTo
	}

	return nil
}
