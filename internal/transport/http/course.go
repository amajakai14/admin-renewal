package http

import (
	"context"

	"github.com/amajakai14/admin-renewal/internal/course"
	"github.com/gin-gonic/gin"
)

type CourseService interface {
	PostCourse(context.Context, course.Course) (course.Course, error)
	GetCourse(context.Context, int) (course.Course, error)
	GetCourses(context.Context) ([]course.Course, error)
	UpdateCourse(context.Context, string, course.Course) error
	DeleteCourse(context.Context, int) error
}

func (h *Handler) PostCourse(c *gin.Context) {

}

func (h *Handler) GetCourse(c *gin.Context) {

}

func (h *Handler) GetCourses(c *gin.Context) {

}

func (h *Handler) UpdateCourse(c *gin.Context) {

}

func (h *Handler) DeleteCourse(c *gin.Context) {

}
