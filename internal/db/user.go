package db

import (
	"context"
	"time"

	appUser "github.com/amajakai14/admin-renewal/internal/user"
)

type UserRow struct {
	ID            int
	Name          string
	Email         string
	Password      string
	Role          string
	EmailVerified bool      `db:"email_verified"`
	CreatedAt     time.Time `db:"created_at"`
	CorporationId string    `db:"corporation_id"`
}

func (d *Database) PostUser(ctx context.Context, user *appUser.User) error {
	var userRow UserRow
	userRow = UserRow{
		Name:          user.Name,
		Email:         user.Email,
		Password:      user.HashedPassword,
		Role:          user.Role,
		EmailVerified: false,
		CorporationId: user.CorporationId,
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO app_user
		(name, email, password, email_verified, role, corporation_id)
		VALUES
		(:name, :email, :password, :email_verified, :role, :corporation_id)
		RETURNING id
		`,
		userRow,
	)
	if err != nil {
		return err
	}
	if rows.Next() {
		rows.Scan(&user.ID)
	}
	if err := rows.Close(); err != nil {
		return err
	}
	return nil
}

func (d *Database) GetUser(ctx context.Context, id int) (*appUser.User, error) {
	var userRow UserRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT * FROM app_user WHERE id = $1`,
		id,
	)
	if err := row.Scan(
		&userRow.ID,
		&userRow.Name,
		&userRow.Email,
		&userRow.Password,
		&userRow.EmailVerified,
		&userRow.Role,
		&userRow.CreatedAt,
		&userRow.CorporationId,
	); err != nil {
		return nil, err
	}
	return toUser(userRow), nil
}

func toUser(userRow UserRow) *appUser.User {
	return &appUser.User{
		ID:             userRow.ID,
		Name:           userRow.Name,
		Email:          userRow.Email,
		HashedPassword: userRow.Password,
		Role:           userRow.Role,
		CreatedAt:      userRow.CreatedAt,
		CorporationId:  userRow.CorporationId,
		EmailVerified:  userRow.EmailVerified,
	}
}

func (d *Database) UpdateUser(ctx context.Context, user *appUser.User) error {
	userRow := UserRow{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Password:      user.HashedPassword,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
		CorporationId: user.CorporationId,
	}
	if _, err := d.Client.NamedQueryContext(
		ctx,
		`
		UPDATE app_user
		SET name = :name, 
		email = :email, 
		password = :password, 
		email_verified = :email_verified, 
		role = :role, 
		corporation_id = :corporation_id
		WHERE id = :id
		`,
		userRow,
	); err != nil {
		return err
	}
	return nil
}

func (d * Database) DeleteUser(ctx context.Context, id int) error {
	if _, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM app_user WHERE id = $1`,
		id,
	); err != nil {
		return err
	}
	return nil
}
