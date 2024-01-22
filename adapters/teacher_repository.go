package adapters

import (
    "github.com/shrd09/go_studentDashboard/domain"
    "gorm.io/gorm"
)

type TeacherRepository struct {
	DB *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{DB: db}
}

func (r *TeacherRepository) GetAllTeachers() ([]domain.Teacher, error) {
    var teachers []domain.Teacher
    if err := r.DB.Select("id, teacher_name").Find(&teachers).Error; err != nil {
        return nil, err
    }
    return teachers, nil
}