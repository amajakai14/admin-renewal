package db

import (
	"context"

	maoc "github.com/amajakai14/admin-renewal/internal/menu_available_on_course"
)

type MenuAvailableOnCourseRow struct {
	CourseID      int    `db:"course_id"`
	MenuID        int    `db:"menu_id"`
	CorporationID string `db:"corporation_id"`
}

func (d *Database) PostCourseMapper(ctx context.Context, m maoc.MenuAvailableOnCourse) error {
	var menuAvailableOnCourseRows []MenuAvailableOnCourseRow
	err := d.deleteAll(ctx, m.CorporationID)
	if err != nil {
		return err
	}

	for courseID, menuIDs := range m.CourseOnMenu {
		for _, menuID := range menuIDs {
			menuAvailableOnCourseRows = append(menuAvailableOnCourseRows, MenuAvailableOnCourseRow{
				CourseID:      courseID,
				MenuID:        menuID,
				CorporationID: m.CorporationID,
			})
		}
	}

	_, err = d.Client.NamedExecContext(
		ctx,
		`
		INSERT INTO course_on_menu (
			course_id,
			menu_id,
			corporation_id
		) VALUES (
			:course_id,
			:menu_id,
			:corporation_id
		)
		`,
		menuAvailableOnCourseRows,
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetAll(ctx context.Context, corporationId string) (maoc.MenuAvailableOnCourse, error) {
	var menuAvailableOnCourseRows []MenuAvailableOnCourseRow
	err := d.Client.SelectContext(
		ctx,
		&menuAvailableOnCourseRows,
		`
		SELECT
			course_id,
			menu_id,
			corporation_id
		FROM course_on_menu
		WHERE corporation_id = $1
		`,
		corporationId,
	)
	if err != nil {
		return maoc.MenuAvailableOnCourse{}, err
	}

	return toMenuAvailableOnCourse(menuAvailableOnCourseRows, corporationId), nil
}

func (d *Database) deleteAll(ctx context.Context, corporationId string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM course_on_menu WHERE corporation_id = $1`,
		corporationId,
	)
	if err != nil {
		return err
	}
	return nil
}

func toMenuAvailableOnCourse(m []MenuAvailableOnCourseRow, corporationId string) maoc.MenuAvailableOnCourse {
	var menuAvailableOnCourse maoc.MenuAvailableOnCourse
	menuOnCourseMapper := make(map[int][]int)

	menuAvailableOnCourse.CorporationID = corporationId
	for _, menuAvailableOnCourseRow := range m {
		menuOnCourseMapper[menuAvailableOnCourseRow.CourseID] = append(menuOnCourseMapper[menuAvailableOnCourseRow.CourseID], menuAvailableOnCourseRow.MenuID)
	}
	menuAvailableOnCourse.CourseOnMenu = menuOnCourseMapper
	return menuAvailableOnCourse
}
