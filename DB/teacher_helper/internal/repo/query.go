package repo

import (
	"github.com/jmoiron/sqlx"
)

type QueryRepo struct {
	db *sqlx.DB
}

func NewQueryRepo(db *sqlx.DB) *QueryRepo {
	return &QueryRepo{db: db}
}

func (r *QueryRepo) Query(query string) ([]map[string]any, error) {
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]map[string]any, 0)
	for rows.Next() {
		results := make(map[string]any)
		err = rows.MapScan(results)

		if err != nil {
			return nil, err
		}

		data = append(data, results)
	}

	return data, nil
}
