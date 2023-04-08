package db

import (
	"context"
	"testing"

	"github.com/amajakai14/admin-renewal/internal/course"
	"github.com/stretchr/testify/assert"
)

func TestCourseDatabase(t *testing.T) {
	db, err := NewDatabase()
	assert.NoError(t, err)

	createCourse := course.Course{
		CourseName: "standard",
		CoursePrice: 399,
		CourseTimeLimit: 60,
		CoursePriority: 1,
		CorporationID: "test-corporation",
	}
	createdCourse, err := db.PostCourse(context.Background(), createCourse)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, createdCourse.ID)

	updateCourse := course.Course{
		ID: createdCourse.ID,
		CourseName: "lunch buffet",
		CoursePrice: 299,
		CourseTimeLimit: 50,
		CoursePriority: 2,
	}
	err = db.UpdateCourse(context.Background(), updateCourse)
	assert.NoError(t, err)

	fetchedCourse, err := db.GetCourse(context.Background(), createdCourse.ID)
	assert.NoError(t, err)
	assert.Equal(t, updateCourse.CourseName, fetchedCourse.CourseName)
	assert.Equal(t, updateCourse.CoursePrice, fetchedCourse.CoursePrice)
	assert.Equal(t, updateCourse.CourseTimeLimit, fetchedCourse.CourseTimeLimit)
	assert.Equal(t, updateCourse.CoursePriority, fetchedCourse.CoursePriority)

	err = db.DeleteCourse(context.Background(), createdCourse.ID)
	assert.NoError(t, err)
	_, err = db.GetCourse(context.Background(), createdCourse.ID)
	assert.Error(t, err)
}

