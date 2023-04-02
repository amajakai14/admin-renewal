package user

import (
	"context"
	"time"
)

type User struct {
	ID int 
	Name string
	Email string
	HashedPassword string
	Role string
	EmailVerified bool
	CreatedAt time.Time
	CorporationId string
}

type Store interface {
	PostUser(ctx context.Context, user User) (User, error) 
	GetUser(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, user User, id string) (User, error)
	DeleteUser(ctx context.Context, id string) error
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{store}
}

func (s *Service) PostUser(ctx context.Context,user User) (User, error) {
	return s.Store.PostUser(ctx, user)
}

func (s *Service) GetUser(ctx context.Context, id string) (User, error) {
	return s.Store.GetUser(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context,user User, id string) (User, error) {
	return s.Store.UpdateUser(ctx, user, id)
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.Store.DeleteUser(ctx, id)
}
