package db

import (
	"context"
	"time"

	ch "github.com/amajakai14/admin-renewal/internal/channel"
)

type ChannelRow struct {
	ID        string
	TableID   int
	Status    string
	TimeStart time.Time `db:"time_start"`
	TimeEnd   time.Time `db:"time_end"`
	CourseID  int       `db:"course_id"`
}

func (d *Database) PostChannel(ctx context.Context, channel ch.Channel) (ch.Channel, error) {
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
	return toChannel(channelRow), nil
}
func (d *Database) GetChannel(ctx context.Context, id string) (ch.Channel, error) {
	var channel ChannelRow
	row := d.Client.QueryRowContext(
		ctx,
		"SELECT * FROM channel WHERE id = $1",
		id,
	)
	if err := row.Scan(
		&channel.ID,
		&channel.TableID,
		&channel.Status,
		&channel.TimeStart,
		&channel.TimeEnd,
		&channel.CourseID,
	); err != nil {
		return ch.Channel{}, err
	}
	return toChannel(channel), nil
}
func (d *Database) UpdateChannel() {}
func (d *Database) DeleteChannel() {}
func toChannel(channelRow ChannelRow) ch.Channel {
	return ch.Channel{
		ID:        channelRow.ID,
		TableID:   channelRow.TableID,
		Status:    channelRow.Status,
		TimeStart: channelRow.TimeStart,
		TimeEnd:   channelRow.TimeEnd,
		CourseID:  channelRow.CourseID,
	}
}
