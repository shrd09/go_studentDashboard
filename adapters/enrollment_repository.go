package adapters

import (
    "github.com/shrd09/go_studentDashboard/domain"
    "gorm.io/gorm"
)

type EnrollmentRepository struct {
	DB *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) *EnrollmentRepository {
	return &EnrollmentRepository{DB: db}
}

func (r *EnrollmentRepository) Create(enrollment *domain.Enrollment) error {
	return r.DB.Create(enrollment).Error
}


func (r *EnrollmentRepository) GetAllEnrollments() ([]domain.Enrollment, error) {
    var enrollments []domain.Enrollment
    if err := r.DB.Select("id, user_id, course_id, marks").Find(&enrollments).Error; err != nil {
        return nil, err
    }
    return enrollments, nil
}


// Add a new method to get enrollments by student ID
func (r *EnrollmentRepository) GetEnrollmentsByStudentID(studentID int) ([]domain.Enrollment, error) {
    var enrollments []domain.Enrollment
    if err := r.DB.Select("course_id, marks").Where("user_id = ?", studentID).Find(&enrollments).Error; err != nil {
        return nil, err
    }
    return enrollments, nil
}


// Add a new method to get enrollments by course ID
func (r *EnrollmentRepository) GetEnrollmentsByCourseID(courseID int) ([]domain.Enrollment, error) {
    var enrollments []domain.Enrollment
    if err := r.DB.Select("id, user_id, course_id, marks").Where("course_id = ?", courseID).Find(&enrollments).Error; err != nil {
        return nil, err
    }
    return enrollments, nil
}



func (r *EnrollmentRepository) GetStudentByID(studentID int) (domain.Student, error) {
	var student domain.Student
	if err := r.DB.Select("id, user_id, student_name, address, age").
		Where("user_id = ?", studentID).
		First(&student).Error; err != nil {
		return domain.Student{}, err
	}

	return student, nil
}




func (r *EnrollmentRepository) UpdateMarks(enrollmentID uint, marks int) error {
	return r.DB.Model(&domain.Enrollment{}).Where("id = ?", enrollmentID).Update("marks", marks).Error
}