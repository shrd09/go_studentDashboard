package application

import "github.com/shrd09/go_studentDashboard/domain"
import "github.com/shrd09/go_studentDashboard/adapters"

type CourseService struct {
	CourseRepository *adapters.CourseRepository
}

func NewCourseService(courseRepo *adapters.CourseRepository) *CourseService {
	return &CourseService{CourseRepository: courseRepo}
}

func (s *CourseService) CreateCourse(CourseName string) error {
	course := &domain.Course{
		CourseName: CourseName,
	}

	return s.CourseRepository.Create(course)
}

func (s *CourseService) DeleteCourse(courseID int) error {
    return s.CourseRepository.Delete(courseID)
}

func (s *CourseService) GetCourses() ([]domain.Course, error) {
    return s.CourseRepository.GetCourses()
}
