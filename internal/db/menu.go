package db

import (
	"context"
	"database/sql"

	"github.com/amajakai14/admin-renewal/internal/menu"
)

type MenuRow struct {
	ID            uint32         `db:"id"`
	MenuType      string         `db:"menu_type"`
	Price         uint32         `db:"price"`
	Available     bool           `db:"available"`
	HasImage      bool           `db:"has_image"`
	CreatedAt     string         `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
	CorporationID string         `db:"corporation_id"`
	MenuNameEN    sql.NullString `db:"menu_name_en"`
	MenuNameTH    sql.NullString `db:"menu_name_th"`
	Priority      sql.NullInt32  `db:"priority"`
}

func (d *Database) PostMenu(ctx context.Context, m menu.Menu) (menu.Menu, error) {
	menuRow := MenuRow{
		ID:            0,
		MenuType:      string(m.MenuType),
		Price:         m.Price,
		Available:     m.Available,
		HasImage:      m.HasImage,
		CorporationID: m.CorporationID,
		MenuNameEN:    toNullString(m.MenuNameEN),
		MenuNameTH:    toNullString(m.MenuNameTH),
		Priority:      toNullInt32(m.Priority),
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO menu
		(menu_type, price, available, has_image, corporation_id, menu_name_en, 
		menu_name_th, priority)
		VALUES
		(:menu_type, :price,:available, :has_image, :corporation_id,
		:menu_name_en,:menu_name_th,:priority)
		RETURNING id
		`,
		menuRow,
	)
	if err != nil {
		return menu.Menu{}, err
	}
	if rows.Next() {
		rows.Scan(&m.ID)
	}
	if err := rows.Close(); err != nil {
		return menu.Menu{}, err
	}
	return m, nil
}

func (d *Database) GetMenu(ctx context.Context, id uint32) (menu.Menu, error) {
	var menuRow MenuRow
	if err := d.Client.GetContext(
		ctx,
		&menuRow,
		"SELECT * FROM menu WHERE id = $1 LIMIT 1 ",
		id,
	); err != nil {
		return menu.Menu{}, err
	}
	return menuRow.toMenu(), nil
}

func (d *Database) GetMenus(ctx context.Context) ([]menu.Menu, error) {
	var menuRows []MenuRow
	if err := d.Client.SelectContext(
		ctx,
		&menuRows,
		"SELECT * FROM menus",
	); err != nil {
		return nil, err
	}
	menus := make([]menu.Menu, len(menuRows))
	for i, menuRow := range menuRows {
		menus[i] = menuRow.toMenu()
	}
	return menus, nil
}

func (d *Database) UpdateMenu(ctx context.Context, m menu.Menu) error {
	menuRow := MenuRow{
		ID:            m.ID,
		MenuType:      string(m.MenuType),
		Price:         m.Price,
		Available:     m.Available,
		HasImage:      m.HasImage,
		CorporationID: m.CorporationID,
		MenuNameEN:    toNullString(m.MenuNameEN),
		MenuNameTH:    toNullString(m.MenuNameTH),
		Priority:      toNullInt32(m.Priority),
	}
	_, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE menu
		SET menu_type = :menu_type, 
		price = :price, 
		available = :available, 
		has_image = :has_image, 
		menu_name_en = :menu_name_en, 
		menu_name_th = :menu_name_th, 
		priority = :priority
		WHERE id = :id
		`,
		menuRow,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteMenu(ctx context.Context, id uint32) error {
	_, err := d.Client.ExecContext(
		ctx,
		"DELETE FROM menu WHERE id = $1",
		id,
	)
	return err
}

func toMenus(menuRows []MenuRow) []menu.Menu {
	menus := make([]menu.Menu, len(menuRows))
	for i, menuRow := range menuRows {
		menus[i] = menuRow.toMenu()
	}
	return menus
}

func (m *MenuRow) toMenu() menu.Menu {
	return menu.Menu{
		ID:            m.ID,
		MenuNameTH:    toString(m.MenuNameTH),
		MenuNameEN:    toString(m.MenuNameEN),
		MenuType:      menu.MenuType(m.MenuType),
		Price:         m.Price,
		Available:     m.Available,
		HasImage:      m.HasImage,
		Priority:      toUInt32(m.Priority),
		CorporationID: m.CorporationID,
	}
}

func toString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func toUInt32(i sql.NullInt32) uint32 {
	if i.Valid {
		return uint32(i.Int32)
	}
	return 0
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func toNullInt32(i uint32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: int32(i),
		Valid: true,
	}
}
