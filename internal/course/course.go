package course

import "context"

type Course struct {
	ID              int
	CourseName      string
	CoursePrice     int
	CourseTimeLimit int
	CoursePriority  uint32
	CorporationID   string
}

const (
	TIMIT_LIMIT_DEFAULT = 60
)

type Store interface {
	PostCourse(context.Context, Course) (Course, error)
	GetCourse(context.Context, int) (Course, error)
	GetCourses(context.Context, string) ([]Course, error)
	UpdateCourse(context.Context, Course) error
	DeleteCourse(context.Context, int) error
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) PostCourse(ctx context.Context, c Course) (Course, error) {
	return s.Store.PostCourse(ctx, c)
}

func (s *Service) GetCourse(ctx context.Context, id int) (Course, error) {
	return s.Store.GetCourse(ctx, id)
}

func (s *Service) GetCourses(ctx context.Context, corporationId string) ([]Course, error) {
	return s.Store.GetCourses(ctx, corporationId)
}

func (s *Service) UpdateCourse(ctx context.Context, c Course) error {
	return s.Store.UpdateCourse(ctx, c)
}

func (s *Service) DeleteCourse(ctx context.Context, id int) error {
	return s.Store.DeleteCourse(ctx, id)
}
