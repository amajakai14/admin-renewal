package db

import (
	"context"
	"time"

	ch "github.com/amajakai14/admin-renewal/internal/channel"
	"github.com/google/uuid"
)

type ChannelRow struct {
	ID        string    `db:"id"`
	TableID   int       `db:"table_id"`
	Status    string    `db:"status"`
	TimeStart time.Time `db:"time_start"`
	TimeEnd   time.Time `db:"time_end"`
	CourseID  int       `db:"course_id"`
}

func (d *Database) PostChannel(ctx context.Context, channel ch.Channel) (ch.Channel, error) {
	err := d.generateChannelID(ctx, &channel)
	if err != nil {
		return ch.Channel{}, err
	}

	channelRow := ChannelRow{
		ID:        channel.ID,
		TableID:   channel.TableID,
		Status:    channel.Status,
		TimeStart: channel.TimeStart,
		TimeEnd:   channel.TimeEnd,
		CourseID:  channel.CourseID,
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`
		INSERT INTO channel (
			id,
			table_id,
			status,
			time_start,
			time_end,
			course_id
		) VALUES (
			:id,
			:table_id,
			:status,
			:time_start,
			:time_end,
			:course_id
		)
		`,
		channelRow,
	)
	if err != nil {
		return ch.Channel{}, err

	}
	if err := rows.Close(); err != nil {
		return ch.Channel{}, err
	}
	return channelRow.toChannel(), nil
}

func (d *Database) GetChannel(ctx context.Context, id string) (ch.Channel, error) {
	var channelRow ChannelRow
	row := d.Client.QueryRowContext(
		ctx,
		"SELECT * FROM channel WHERE id = $1",
		id,
	)
	if err := row.Scan(
		&channelRow.ID,
		&channelRow.TableID,
		&channelRow.Status,
		&channelRow.TimeStart,
		&channelRow.TimeEnd,
		&channelRow.CourseID,
	); err != nil {
		return ch.Channel{}, err
	}
	return channelRow.toChannel(), nil
}
func (d *Database) UpdateChannel(ctx context.Context, channel ch.Channel) error {
	channelRow := ChannelRow{
		ID:        channel.ID,
		TableID:   channel.TableID,
		Status:    channel.Status,
		TimeStart: channel.TimeStart,
		TimeEnd:   channel.TimeEnd,
		CourseID:  channel.CourseID,
	}

	_, err := d.Client.NamedQueryContext(
		ctx,
		`
		UPDATE channel SET
			table_id = :table_id,
			status = :status,
			time_start = :time_start,
			time_end = :time_end,
			course_id = :course_id
		WHERE id = :id
		`,
		channelRow,
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteChannel(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		"DELETE FROM channel WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChannelRow) toChannel() ch.Channel {
	return ch.Channel{
		ID:        c.ID,
		TableID:   c.TableID,
		Status:    c.Status,
		TimeStart: c.TimeStart,
		TimeEnd:   c.TimeEnd,
		CourseID:  c.CourseID,
	}
}

func (d *Database) generateChannelID(ctx context.Context, channel *ch.Channel) error {
	for {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		_, err = d.GetChannel(ctx, uuid.String())
		if err != nil {
			channel.ID = uuid.String()
			return nil
		}
	}
}
