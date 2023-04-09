package db

import (
	"context"
	"testing"

	"github.com/amajakai14/admin-renewal/internal/desk"
	"github.com/stretchr/testify/assert"
)

func TestDeskDatabase(t *testing.T) {
	db, err := NewDatabase()
	assert.NoError(t, err)

	createDesk := desk.Desk{
		TableName: "T_01",
		IsOccupied: false,
		CorporateId: "test-corporation",
	}

	createdDesk, err := db.PostDesk(context.Background(), createDesk)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, createdDesk.ID)

	updateDesk := desk.Desk{
		ID: createdDesk.ID,
		TableName: "T_02",
		IsOccupied: true,
	}

	err = db.UpdateDesk(context.Background(), updateDesk)
	assert.NoError(t, err)

	fetchedDesk, err := db.GetDesk(context.Background(), createdDesk.ID)
	assert.NoError(t, err)
	assert.Equal(t, updateDesk.TableName, fetchedDesk.TableName)
	assert.Equal(t, updateDesk.IsOccupied, fetchedDesk.IsOccupied)

	err = db.DeleteDesk(context.Background(), createdDesk.ID)
	assert.NoError(t, err)
	_, err = db.GetDesk(context.Background(), createdDesk.ID)
	assert.Error(t, err)
}
