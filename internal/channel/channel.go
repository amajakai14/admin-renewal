package channel

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ID        string
	TableID   int
	Status    string
	TimeStart time.Time
	TimeEnd   time.Time
	CourseID  int
}

type Store interface {
	PostChannel(context.Context, Channel) (Channel, error)
	GetChannel(context.Context, string) (Channel, error)
	UpdateChannel(context.Context, Channel) error
	DeleteChannel(context.Context, string) error
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{store}
}

func (s *Service) PostChannel(ctx context.Context, channel Channel) (Channel, error) {
	for {
		channel.ID = uuid.New().String()
		_, err := s.GetChannel(ctx, channel.ID)
		if err == nil {
			break
		}
	}
	return s.Store.PostChannel(ctx, channel)
}

func (s *Service) GetChannel(ctx context.Context, id string) (Channel, error) {
	return s.Store.GetChannel(ctx, id)
}

func (s *Service) UpdateChannel(ctx context.Context, channel Channel) error {
	return s.Store.UpdateChannel(ctx, channel)
}

func (s *Service) DeleteChannel(ctx context.Context, id string) error {
	return s.Store.DeleteChannel(ctx, id)
}
