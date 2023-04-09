package menuavailableoncourse

import "context"

type MenuAvailableOnCourse struct {
	CourseOnMenu map[int][]int
	CorporationID string
}

type Store interface {
	PostCourseMapper(context.Context, []MenuAvailableOnCourse) ([]MenuAvailableOnCourse, error)
	GetAll(context.Context, string) ([]MenuAvailableOnCourse, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) PostCourseMapper(ctx context.Context, m []MenuAvailableOnCourse) ([]MenuAvailableOnCourse, error) {
	return s.Store.PostCourseMapper(ctx, m)
}

func (s *Service) GetAll(ctx context.Context, corporationId string) ([]MenuAvailableOnCourse, error) {
	return s.Store.GetAll(ctx, corporationId)
}

