package db

import (
	"context"
	"testing"

	"github.com/amajakai14/admin-renewal/internal/menu"
	"github.com/stretchr/testify/assert"
)

func TestMenuDatabase(t *testing.T) {
	db, err := NewDatabase()
	assert.NoError(t, err)
	
	createMenu := menu.Menu{
		ID: 0,
		MenuNameTH: "ข้าวมันไก่",
		MenuNameEN: "Chicken Rice",
		MenuType: menu.MAIN_DISH,
		Price:      50,
		Available: true,
		HasImage:  false,
		Priority: 1,
		CorporationID: "test-corporation",
	}
	createdMenu, err := db.PostMenu(context.Background(), createMenu)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, createdMenu.ID)

	updateMenu := menu.Menu{
		ID: createdMenu.ID,
		MenuNameTH: "ข้าวมันไก่อร่อยเหาะ",
		MenuNameEN: "Chicken Rice Super delicious",
		MenuType: createdMenu.MenuType,
		Price: 	createdMenu.Price,
		HasImage:  true,
		Available: false,
		Priority: createdMenu.Priority,
		CorporationID: createdMenu.CorporationID,
	}
	err = db.UpdateMenu(context.Background(), updateMenu)
	assert.NoError(t, err)

	fetchedMenu, err := db.GetMenu(context.Background(), createdMenu.ID)
	assert.NoError(t, err)
	assert.Equal(t, updateMenu.MenuNameTH, fetchedMenu.MenuNameTH)
	assert.Equal(t, updateMenu.MenuNameEN, fetchedMenu.MenuNameEN)
	assert.Equal(t, updateMenu.HasImage, fetchedMenu.HasImage)
	assert.Equal(t, updateMenu.Available, fetchedMenu.Available)

	err = db.DeleteMenu(context.Background(), createdMenu.ID)
	assert.NoError(t, err)
	_, err = db.GetMenu(context.Background(), createdMenu.ID)
	assert.Error(t, err)
}
