// application/teacher_service.go
package application

import (
	"github.com/shrd09/go_studentDashboard/domain"
	"github.com/shrd09/go_studentDashboard/adapters"
	"fmt"
)

// type TeacherService struct {
// 	TeacherRepository *adapters.TeacherRepository
// }

// func NewTeacherService(teacherRepo *adapters.TeacherRepository) *TeacherService {
// 	return &TeacherService{TeacherRepository: teacherRepo}
// }

type TeacherService struct {
	TeacherRepository *adapters.TeacherRepository
	TeacherCourseRepository *adapters.TeacherCourseRepository
	EnrollmentRepository    *adapters.EnrollmentRepository
}

func NewTeacherService(teacherCourseRepo *adapters.TeacherCourseRepository, enrollmentRepo *adapters.EnrollmentRepository) *TeacherService {
	return &TeacherService{
		TeacherCourseRepository: teacherCourseRepo,
		EnrollmentRepository:    enrollmentRepo,
		// TeacherRepository:      teacherRepo,
	}
}


func (s *TeacherService) GetEnrolledStudentsByTeacherID(teacherID int) ([]domain.Student, error) {
	// Get all courses taught by the teacher
	teacherCourses, err := s.TeacherCourseRepository.GetCoursesByTeacherID(teacherID)
	if err != nil {
		return nil, err
	}

	// Initialize a slice to store all students
	var enrolledStudents []domain.Student

	// Loop through each teacher course and fetch enrolled students
	for _, course := range teacherCourses {
		enrollments, err := s.EnrollmentRepository.GetEnrollmentsByCourseID(course.CourseID)
		if err != nil {
			return nil, err
		}

		// Fetch students using the user_id from enrollments
		for _, enrollment := range enrollments {
			// Assuming you have a UserRepository to fetch user details
			student, err := s.EnrollmentRepository.GetStudentByID(enrollment.UserID)
			if err != nil {
				return nil, err
			}

			enrolledStudents = append(enrolledStudents, student)
		}
	}

	return enrolledStudents, nil
}

// func (s *TeacherService) GetAllTeachers() ([]domain.Teacher, error) {
//     return s.TeacherRepository.GetAllTeachers()
// }


func (s *TeacherService) GetEnrolledStudentsByCourseID(teacherID, courseID int) ([]domain.Student, error) {
	// Get all TeacherCourse records for the specified teacher and course
	teacherCourses, err := s.TeacherCourseRepository.GetStudentsByTeacherAndCourseID(teacherID, courseID)
	fmt.Println(teacherCourses)
	if err != nil {
		return nil, err
	}
	return teacherCourses, nil
}
