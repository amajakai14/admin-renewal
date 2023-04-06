package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/amajakai14/admin-renewal/internal/course"
)

type CourseRow struct {
	ID              int
	CourseName      string        `db:"course_name"`
	CoursePrice     int           `db:"course_price"`
	CourseTimeLimit int           `db:"course_time_limit"`
	CoursePriority  sql.NullInt32 `db:"course_priority"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       sql.NullTime  `db:"updated_at"`
	CorporationID   string        `db:"corporation_id"`
}

func (d *Database) PostCourse(ctx context.Context, c course.Course) (course.Course, error) {
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO courses
		(course_name, 
		course_price,
		course_time_limit,
		course_priority,
		corporation_id)
		VALUES
		(:course_name, 
		:course_price,
		:course_time_limit,
		:course_priority,
		:corporation_id)
		RETURNING id
		`,
		c,
	)
	if err != nil {
		return course.Course{}, err
	}
	if rows.Next() {
		rows.Scan(&c.ID)
	}
	if err := rows.Close(); err != nil {
		return course.Course{}, err
	}
	return c, nil
}

func (d *Database) GetCourse(ctx context.Context, id int) (course.Course, error) {
	var courseRow CourseRow
	if err := d.Client.GetContext(
		ctx,
		&courseRow,
		"SELECT * FROM courses WHERE id = ?",
		id,
	); err != nil {
		return course.Course{}, err
	}
	return toCourse(courseRow), nil
}

func (d *Database) GetCourses(ctx context.Context) ([]course.Course, error) {
	var courseRows []CourseRow
	if err := d.Client.SelectContext(
		ctx,
		&courseRows,
		"SELECT * FROM courses",
	); err != nil {
		return nil, err
	}
	return toCourses(courseRows), nil
}

func (d *Database) UpdateCourse(ctx context.Context, c course.Course) error {
	_, err := d.Client.NamedExecContext(
		ctx,
		`UPDATE courses
		SET
		course_name = :course_name,
		course_price = :course_price,
		course_time_limit = :course_time_limit,
		course_priority = :course_priority,
		updated_at = :updated_at
		WHERE id = :id
		`,
		c,
	)
	if err != nil {
		return err
	}
	return  nil
}

func (d *Database) DeleteCourse(ctx context.Context, id int) error {
	_, err := d.Client.ExecContext(
		ctx,
		"DELETE FROM courses WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}


func toCourses(courseRows []CourseRow) []course.Course {
	courses := make([]course.Course, len(courseRows))
	for i, courseRow := range courseRows {
		courses[i] = toCourse(courseRow)
	}
	return courses
}

func toCourse(courseRow CourseRow) course.Course {
	return course.Course{
		ID:              courseRow.ID,
		CourseName:      courseRow.CourseName,
		CoursePrice:     courseRow.CoursePrice,
		CourseTimeLimit: courseRow.CourseTimeLimit,
		CoursePriority:  int(courseRow.CoursePriority.Int32),
		CorporationID:   courseRow.CorporationID,
	}
}
