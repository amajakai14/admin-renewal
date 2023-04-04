package db

import (
	"context"
	"testing"

	appUser "github.com/amajakai14/admin-renewal/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestUserDatabase(t *testing.T) {
	t.Run("test CRUD user", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		user := appUser.User{
			Name:"means",
			Email:"means@example.com",
			HashedPassword: "hashedpassword",
			Role: "admin",
			CorporationId: "test-corporation",
		}

		createdUser, err := db.PostUser(context.Background(), user)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, createdUser.ID) 

		newUser, err := db.GetUserByID(context.Background(), createdUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, createdUser.ID, newUser.ID)
		assert.Equal(t, createdUser.Name, newUser.Name)
		assert.Equal(t, createdUser.Email, newUser.Email)
		assert.Equal(t, createdUser.HashedPassword, newUser.HashedPassword)
		assert.Equal(t, createdUser.Role, newUser.Role)
		assert.Equal(t, createdUser.CorporationId, newUser.CorporationId)

		updateUser := appUser.User{
			ID: createdUser.ID,
			Name:"new means",
			Email:"means@example.com",
			HashedPassword: "hashedpassword",
			Role: "admin",
			CorporationId: "test-corporation",
		}
		err = db.UpdateUser(context.Background(), updateUser)
		assert.NoError(t, err)

		newUser, err = db.GetUserByID(context.Background(), createdUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, updateUser.ID, newUser.ID)
		assert.Equal(t, updateUser.Name, newUser.Name)
		assert.Equal(t, updateUser.Email, newUser.Email)
		assert.Equal(t, updateUser.HashedPassword, newUser.HashedPassword)
		assert.Equal(t, updateUser.Role, newUser.Role)
		assert.Equal(t, updateUser.CorporationId, newUser.CorporationId)

		err = db.DeleteUser(context.Background(), createdUser.ID)
		assert.NoError(t, err)

		_, err = db.GetUserByID(context.Background(), createdUser.ID)
		assert.Error(t, err)
	})
}

