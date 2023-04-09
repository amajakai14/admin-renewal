package db

import (
	"context"

	"github.com/amajakai14/admin-renewal/internal/desk"
)

type DeskRow struct {
	ID          int32  `db:"id"`
	TableName   string `db:"table_name"`
	IsOccupied  bool   `db:"is_occupied"`
	CorporationId string `db:"corporation_id"`
}

func (db *Database) PostDesk(ctx context.Context, d desk.Desk) (desk.Desk, error) {
	var deskRow DeskRow
	err := db.Client.GetContext(
		ctx,
		&deskRow,
		"INSERT INTO desk (table_name, is_occupied, corporation_id) VALUES ($1, $2, $3) RETURNING *",
		d.TableName,
		d.IsOccupied,
		d.CorporateId,
	)
	if err != nil {
		return desk.Desk{}, err
	}

	return deskRow.toDesk(), nil
}

func (db *Database) GetDesk(ctx context.Context, id int) (desk.Desk, error) {
	var deskRow DeskRow
	err := db.Client.GetContext(
		ctx,
		&deskRow,
		"SELECT * FROM desk WHERE id = $1",
		id,
	)
	if err != nil {
		return desk.Desk{}, err
	}

	return deskRow.toDesk(), nil
}

func (db *Database) GetDesks(ctx context.Context, corporationId string) ([]desk.Desk, error) {
	var deskRows []DeskRow
	err := db.Client.SelectContext(
		ctx,
		&deskRows,
		"SELECT * FROM desk WHERE corporation_id = $1",
		corporationId,
	)
	if err != nil {
		return nil, err
	}

	var desk []desk.Desk
	for _, deskRow := range deskRows {
		desk = append(desk, deskRow.toDesk())
	}

	return desk, nil
}

func (db *Database) UpdateDesk(ctx context.Context, d desk.Desk) error {
	deskRow := DeskRow{
		ID:          int32(d.ID),
		TableName:   d.TableName,
		IsOccupied:  d.IsOccupied,
		CorporationId: d.CorporateId,
	}
	_, err := db.Client.NamedExecContext(
		ctx,
		`UPDATE desk 
		SET table_name = :table_name, 
		is_occupied = :is_occupied
		WHERE id = :id`,
		deskRow,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) DeleteDesk(ctx context.Context, id int) error {
	_, err := db.Client.ExecContext(
		ctx,
		"DELETE FROM desk WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *DeskRow) toDesk() desk.Desk {
	return desk.Desk{
		ID:          int(db.ID),
		TableName:   db.TableName,
		IsOccupied:  db.IsOccupied,
		CorporateId: db.CorporationId,
	}
}
