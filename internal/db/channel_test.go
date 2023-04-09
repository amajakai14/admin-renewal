package db

import (
	"context"
	"testing"
	"time"

	"github.com/amajakai14/admin-renewal/internal/channel"
	"github.com/amajakai14/admin-renewal/internal/course"
	"github.com/amajakai14/admin-renewal/internal/desk"
	"github.com/stretchr/testify/assert"
)

func TestChannelDatabase(t *testing.T) {
	db, err := NewDatabase()
	assert.NoError(t, err)

	createDesk := desk.Desk{
		TableName: "T_01",
		IsOccupied: false,
		CorporateId: "test-corporation",
	}

	createdDesk, err := db.PostDesk(context.Background(), createDesk)
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

	now := time.Now()
	duration := time.Duration(createdCourse.CourseTimeLimit) * time.Minute

	createChannel := channel.Channel{
		TableID: createdDesk.ID,
		Status: "ACTIVE",
		CourseID: createdCourse.ID,
		TimeStart: now,
		TimeEnd: now.Add(duration),
	}

	createdChannel, err := db.PostChannel(context.Background(), createChannel)
	assert.NoError(t, err)
	assert.NotEqual(t, "", createdChannel.ID)

	updateChannel := channel.Channel{
		ID: createdChannel.ID,
		TableID: createdDesk.ID,
		Status: "INACTIVE",
		CourseID: createdCourse.ID,
	}

	err = db.UpdateChannel(context.Background(), updateChannel)
	assert.NoError(t, err)

	fetchedChannel, err := db.GetChannel(context.Background(), createdChannel.ID)
	assert.NoError(t, err)
	assert.Equal(t, updateChannel.Status, fetchedChannel.Status)

	err = db.DeleteChannel(context.Background(), createdChannel.ID)
	assert.NoError(t, err)
	fetchedChannel, err = db.GetChannel(context.Background(), createdChannel.ID)
	assert.Error(t, err)
}
