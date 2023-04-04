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
	GetUserByID(ctx context.Context, id int) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id int) error
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

func (s *Service) GetUserByID(ctx context.Context, id int) (User, error) {
	return s.Store.GetUserByID(ctx, id)
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return s.Store.GetUserByEmail(ctx, email)
}

func (s *Service) UpdateUser(ctx context.Context,user User)  error {
	return s.Store.UpdateUser(ctx, user)
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	return s.Store.DeleteUser(ctx, id)
}
