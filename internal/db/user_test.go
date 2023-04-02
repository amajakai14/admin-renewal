package db

import (
	"context"
	"testing"

	appUser "github.com/amajakai14/admin-renewal/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		user := &appUser.User{
			Name:"means",
			Email:"means@example.com",
			HashedPassword: "hashedpassword",
			Role: "admin",
			CorporationId: "test-corporation",
		}

		err = db.PostUser(context.Background(), user)
		assert.NotEqual(t, 0, user.ID)
		assert.NoError(t, err)

		newUser, err := db.GetUser(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, newUser.ID)
		assert.Equal(t, user.Name, newUser.Name)
		assert.Equal(t, user.Email, newUser.Email)
		assert.Equal(t, user.HashedPassword, newUser.HashedPassword)
		assert.Equal(t, user.Role, newUser.Role)
		assert.Equal(t, user.CorporationId, newUser.CorporationId)



		updateUser:= &appUser.User{
			ID: user.ID,
			Name:"new means",
			Email:"means@example.com",
			HashedPassword: "hashedpassword",
			Role: "admin",
			CorporationId: "test-corporation",
		}
		err = db.UpdateUser(context.Background(), updateUser)
		assert.NoError(t, err)

		newUser, err = db.GetUser(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Equal(t, updateUser.ID, newUser.ID)
		assert.Equal(t, updateUser.Name, newUser.Name)
		assert.Equal(t, updateUser.Email, newUser.Email)
		assert.Equal(t, updateUser.HashedPassword, newUser.HashedPassword)
		assert.Equal(t, updateUser.Role, newUser.Role)
		assert.Equal(t, updateUser.CorporationId, newUser.CorporationId)

		err = db.DeleteUser(context.Background(), user.ID)
		assert.NoError(t, err)

		_, err = db.GetUser(context.Background(), user.ID)
		assert.Error(t, err)
	})
}

