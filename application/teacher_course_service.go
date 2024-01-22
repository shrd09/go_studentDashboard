// application/teacher_course_service.go
package application

import (
	"github.com/shrd09/go_studentDashboard/domain"
	"github.com/shrd09/go_studentDashboard/adapters"
)

type TeacherCourseService struct {
	TeacherCourseRepository *adapters.TeacherCourseRepository
	EnrollmentRepository    *adapters.EnrollmentRepository
}

func NewTeacherCourseService(teacherCourseRepo *adapters.TeacherCourseRepository, enrollmentRepo *adapters.EnrollmentRepository) *TeacherCourseService {
	return &TeacherCourseService{
		TeacherCourseRepository: teacherCourseRepo,
		EnrollmentRepository:    enrollmentRepo, // Add this line
	}
}

func (s *TeacherCourseService) AssignCourseToTeacher(teacherID, courseID int, year int) error {
	// Additional business logic, validation, etc. can be added here
	teacherCourse := &domain.TeacherCourse{
		TeacherID: teacherID,
		CourseID:  courseID,
		Year:      year,
	}

	return s.TeacherCourseRepository.CreateTeacherCourse(teacherCourse)
}

func (s *TeacherCourseService) GetCoursesByTeacherID(teacherID int) ([]domain.TeacherCourse, error) {
    return s.TeacherCourseRepository.GetCoursesByTeacherID(teacherID)
}

func (s *TeacherCourseService) GetEnrolledStudentsByTeacherID(teacherID int) ([]domain.Student, error) {
    // Get all courses taught by the teacher
    teacherCourses, err := s.TeacherCourseRepository.GetCoursesByTeacherID(teacherID)
    if err != nil {
        return nil, err
    }

    // Initialize a slice to store all enrolled students
    var enrolledStudents []domain.Student

    // Loop through each teacher course and fetch enrolled students
    for _, course := range teacherCourses {
        students, err := s.TeacherCourseRepository.GetStudentsByTeacherAndCourseID(teacherID, course.CourseID)
        if err != nil {
            return nil, err
        }
        enrolledStudents = append(enrolledStudents, students...)
    }

    return enrolledStudents, nil
}






func (s *TeacherCourseService) GetEnrollmentsByTeacherID(teacherID int) ([]domain.Enrollment, error) {
	// Get all courses taught by the teacher
	teacherCourses, err := s.TeacherCourseRepository.GetCoursesByTeacherID(teacherID)
	if err != nil {
		return nil, err
	}

	// Initialize a slice to store all enrollments
	var allEnrollments []domain.Enrollment

	// Loop through each teacher course and fetch enrollments
	for _, course := range teacherCourses {
		enrollments, err := s.EnrollmentRepository.GetEnrollmentsByCourseID(course.CourseID)
		if err != nil {
			return nil, err
		}
		allEnrollments = append(allEnrollments, enrollments...)
	}

	return allEnrollments, nil
}
