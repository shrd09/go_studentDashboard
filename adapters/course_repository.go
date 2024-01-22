package adapters

import (
    "github.com/shrd09/go_studentDashboard/domain"
    "gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

// constructor function - creates new instance of CourseRepository
func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (r *CourseRepository) Create(course *domain.Course) error {
	return r.DB.Create(course).Error
}

// this code defines a repository (CourseRepository) responsible for handling database interactions related to the Course entity using the GORM library. The Create method is specifically used to persist a new Course record in the database.

func (r *CourseRepository) Delete(courseID int) error {
    return r.DB.Delete(&domain.Course{}, courseID).Error
}

func (r *CourseRepository) GetCourses() ([]domain.Course, error) {
    var courses []domain.Course
    if err := r.DB.Select("id, course_name").Find(&courses).Error; err != nil {
        return nil, err
    }
    return courses, nil
}
