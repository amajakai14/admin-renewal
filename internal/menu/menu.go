package menu

import "context"

type Menu struct {
	ID            uint32
	MenuNameTH    string
	MenuNameEN    string
	MenuType      MenuType
	Price         uint32
	Available     bool
	HasImage      bool
	Priority      uint32
	CorporationID string
}

type MenuType string

const (
	APPETIZER     MenuType = "APPETIZER"
	MAIN_DISH     MenuType = "MAIN_DISH"
	DESSERT       MenuType = "DESSERT"
	DRINK         MenuType = "DRINK"
	DEFAULT_PRICE uint32   = 0
)

type Store interface {
	PostMenu(context.Context, Menu) (Menu, error)
	GetMenu(context.Context, uint32) (Menu, error)
	GetMenus(context.Context) ([]Menu, error)
	UpdateMenu(context.Context, Menu) error
	DeleteMenu(context.Context, uint32) error
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) PostMenu(ctx context.Context, m Menu) (Menu, error) {
	return s.Store.PostMenu(ctx, m)
}

func (s *Service) GetMenu(ctx context.Context, id uint32) (Menu, error) {
	return s.Store.GetMenu(ctx, id)
}

func (s *Service) GetMenus(ctx context.Context) ([]Menu, error) {
	return s.Store.GetMenus(ctx)
}

func (s *Service) UpdateMenu(ctx context.Context, m Menu) error {
	return s.Store.UpdateMenu(ctx, m)
}

func (s *Service) DeleteMenu(ctx context.Context, id uint32) error {
	return s.Store.DeleteMenu(ctx, id)
}

