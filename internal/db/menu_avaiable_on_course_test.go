package db

import (
	"context"
	"testing"

	mc "github.com/amajakai14/admin-renewal/internal/menu_available_on_course"
	"github.com/stretchr/testify/assert"
)

func TestMenuAvailableOnCourse(t *testing.T) {
	db, err := NewDatabase()
	assert.NoError(t, err)

	courseOnMenu := map[int][]int{
		1: []int{1, 2, 3, 4},
		2: []int{1, 2, 3, 4, 5, 6},
	}

	createMenuAvailableOnCourse := mc.MenuAvailableOnCourse{
		CourseOnMenu: courseOnMenu,
		CorporationID: "test-corporation",
	}

	err = db.PostCourseMapper(context.Background(), createMenuAvailableOnCourse)
	assert.NoError(t, err)

	createdMenuAvailableOnCourse, err := db.GetAll(context.Background(), createMenuAvailableOnCourse.CorporationID)
	assert.NoError(t, err)

	assert.Equal(t, createMenuAvailableOnCourse.CourseOnMenu, createdMenuAvailableOnCourse.CourseOnMenu)
	assert.Equal(t, createMenuAvailableOnCourse.CorporationID, createdMenuAvailableOnCourse.CorporationID)
}
