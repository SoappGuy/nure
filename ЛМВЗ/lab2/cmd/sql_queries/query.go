package sql_queries

import (
	"context"
	"fmt"
	// "time"
)

// Student
// -------------------------------------
type ListStudentsArgs struct {
	Query        string
	IsDescending bool
	OrderBy      string
}

func (q *Queries) ListStudents(ctx context.Context, args ListStudentsArgs) ([]Student, error) {
	query := `
	SELECT
		student_id, firstname, middlename, lastname, gender, birthday, form_of_education, personal_file_number, note
	FROM
		Student
	WHERE
		firstname LIKE :query OR
		middlename LIKE :query OR
		lastname LIKE :query OR
		gender LIKE :query OR
		birthday LIKE :query OR
		form_of_education LIKE :query OR
		personal_file_number LIKE :query
	`
	if args.IsDescending {
		query += fmt.Sprintf(" ORDER BY %s DESC", args.OrderBy)
	} else {
		query += fmt.Sprintf(" ORDER BY %s ASC", args.OrderBy)
	}

	rows, err := q.db.NamedQueryContext(ctx, query, args)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.StudentID,
			&i.Firstname,
			&i.Middlename,
			&i.Lastname,
			&i.Gender,
			&i.Birthday,
			&i.PersonalFileNumber,
			&i.Note,
		); err != nil {
			return nil, err
		}
		students = append(students, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (q *Queries) GetStudent(ctx context.Context, id int64) (*Student, error) {
	query := `
	SELECT
		student_id, firstname, middlename, lastname, gender, birthday, form_of_education, personal_file_number, note
	FROM
		Student
	WHERE
		student_id = ?
	`
	rows, err := q.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.StudentID,
			&i.Firstname,
			&i.Middlename,
			&i.Lastname,
			&i.Gender,
			&i.Birthday,
			&i.PersonalFileNumber,
			&i.Note,
		); err != nil {
			return nil, err
		}
		students = append(students, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(students) != 1 {
		return nil, fmt.Errorf("Expected 1 student, but got %d", len(students))
	}

	return &students[0], nil
}

// -------------------------------------

// Foodstuff
// -------------------------------------
type ListFoodStuffArgs struct {
	Type string
	ID   string
}

func (q *Queries) ListFoodStuff(ctx context.Context, args ListFoodStuffArgs) ([]Foodstuff, error) {
	query := ` 
	SELECT 
		foodstuff_id, frequency, name, price, type 
	FROM 
		Foodstuff 
	WHERE 
		foodstuff_id LIKE :id AND type LIKE :type 
	ORDER BY frequency DESC `

	rows, err := q.db.NamedQueryContext(ctx, query, args)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foodstuffs []Foodstuff
	for rows.Next() {
		var i Foodstuff
		if err := rows.Scan(
			&i.FoodstuffID,
			&i.Frequency,
			&i.Name,
			&i.Price,
			&i.Type,
		); err != nil {
			return nil, err
		}
		foodstuffs = append(foodstuffs, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foodstuffs, nil
}

func (q *Queries) GetFoodStuff(ctx context.Context, id int64) (*Foodstuff, error) {
	query := ` SELECT foodstuff_id, frequency, name, price, type FROM Foodstuff WHERE foodstuff_id = ? `

	rows, err := q.db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foodstuffs []Foodstuff
	for rows.Next() {
		var i Foodstuff
		if err := rows.Scan(
			&i.FoodstuffID,
			&i.Frequency,
			&i.Name,
			&i.Price,
			&i.Type,
		); err != nil {
			return nil, err
		}
		foodstuffs = append(foodstuffs, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(foodstuffs) != 1 {
		return nil, fmt.Errorf("Expected 1 foodstuff, but got %d", len(foodstuffs))
	}

	return &foodstuffs[0], nil
}

func (q *Queries) IncrementFrequency(ctx context.Context, id int64) error {
	query := ` UPDATE Foodstuff SET frequency = frequency + 1 WHERE foodstuff_id = ? `

	_, err := q.db.ExecContext(ctx, query, id)
	return err
}

// -------------------------------------
