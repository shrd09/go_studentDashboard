// application/enrollment_service.go
package application

import "github.com/shrd09/go_studentDashboard/domain"
import "github.com/shrd09/go_studentDashboard/adapters"

type EnrollmentService struct {
	EnrollmentRepository *adapters.EnrollmentRepository
}

func NewEnrollmentService(enrollmentRepo *adapters.EnrollmentRepository) *EnrollmentService {
	return &EnrollmentService{EnrollmentRepository: enrollmentRepo}
}

func (s *EnrollmentService) CreateEnrollment(userID, courseID int) error {

	enrollment := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
	}

	return s.EnrollmentRepository.Create(enrollment)
}

func (s *EnrollmentService) GetAllEnrollments() ([]domain.Enrollment, error) {
    return s.EnrollmentRepository.GetAllEnrollments()
}

func (s *EnrollmentService) GetEnrollmentsByStudentID(studentID int) ([]domain.Enrollment, error) {
    return s.EnrollmentRepository.GetEnrollmentsByStudentID(studentID)
}

func (s *EnrollmentService) GetEnrollmentsByCourseID(courseID int) ([]domain.Enrollment, error) {
    return s.EnrollmentRepository.GetEnrollmentsByCourseID(courseID)
}


func (s *EnrollmentService) UpdateMarks(enrollmentID uint, marks int) error {
	return s.EnrollmentRepository.UpdateMarks(enrollmentID, marks)
}

// alreadyEnrolled, err := s.EnrollmentRepository.CheckEnrollment(userID, courseID)
	// if err != nil {
	// 	return err
	// }

	// if alreadyEnrolled {
	// 	return errors.New("user is already enrolled in the course")
	// }