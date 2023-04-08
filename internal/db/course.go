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
	CourseTimeLimit int           `db:"course_timelimit"`
	CoursePriority  sql.NullInt32 `db:"course_priority"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       sql.NullTime  `db:"updated_at"`
	CorporationID   string        `db:"corporation_id"`
}

func (d *Database) PostCourse(ctx context.Context, c course.Course) (course.Course, error) {
	courseRow := CourseRow{
		CourseName:      c.CourseName,
		CoursePrice:     c.CoursePrice,
		CourseTimeLimit: c.CourseTimeLimit,
		CoursePriority:  toNullInt32(c.CoursePriority),
		CorporationID:   c.CorporationID,
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO course
		(course_name, 
		course_price,
		course_timelimit,
		course_priority,
		corporation_id)
		VALUES
		(:course_name, 
		:course_price,
		:course_timelimit,
		:course_priority,
		:corporation_id)
		RETURNING id
		`,
		courseRow,
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
		"SELECT * FROM course WHERE id = $1",
		id,
	); err != nil {
		return course.Course{}, err
	}
	return toCourse(courseRow), nil
}

func (d *Database) GetCourses(ctx context.Context, corporationId string) ([]course.Course, error) {
	var courseRows []CourseRow
	if err := d.Client.SelectContext(
		ctx,
		&courseRows,
		"SELECT * FROM course WHERE corporation_id = $1",
		corporationId,
	); err != nil {
		return nil, err
	}
	return toCourses(courseRows), nil
}

func (d *Database) UpdateCourse(ctx context.Context, c course.Course) error {
	courseRow := CourseRow{
		ID:              c.ID,
		CourseName:      c.CourseName,
		CoursePrice:     c.CoursePrice,
		CourseTimeLimit: c.CourseTimeLimit,
		CoursePriority:  toNullInt32(c.CoursePriority),
		UpdatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
	}
	_, err := d.Client.NamedExecContext(
		ctx,
		`UPDATE course
		SET
		course_name = :course_name,
		course_price = :course_price,
		course_timelimit = :course_timelimit,
		course_priority = :course_priority,
		updated_at = :updated_at
		WHERE id = :id
		`,
		courseRow,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteCourse(ctx context.Context, id int) error {
	_, err := d.Client.ExecContext(
		ctx,
		"DELETE FROM course WHERE id = $1",
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
		CoursePriority:  uint32(courseRow.CoursePriority.Int32),
		CorporationID:   courseRow.CorporationID,
	}
}
