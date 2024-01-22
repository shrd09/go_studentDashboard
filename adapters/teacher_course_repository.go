package adapters

import (
	"gorm.io/gorm"
	"fmt"
	"github.com/shrd09/go_studentDashboard/domain"
)

type TeacherCourseRepository struct {
	DB *gorm.DB
}

func NewTeacherCourseRepository(db *gorm.DB) *TeacherCourseRepository {
	return &TeacherCourseRepository{DB: db}
}

func (r *TeacherCourseRepository) CreateTeacherCourse(tc *domain.TeacherCourse) error {
	return r.DB.Create(tc).Error
}


func (r *TeacherCourseRepository) GetCoursesByTeacherID(teacherID int) ([]domain.TeacherCourse, error) {
    var teacherCourses []domain.TeacherCourse
    if err := r.DB.Model(&domain.TeacherCourse{}).
        Select("course_id, year").  // Specifying the columns
        Where("teacher_id = ?", teacherID).
        Find(&teacherCourses).
        Error; err != nil {
        return nil, err
    }
    return teacherCourses, nil
}

func (r *TeacherCourseRepository) GetStudentsByTeacherAndCourseID(teacherID, courseID int) ([]domain.Student, error) {
    var students []domain.Student
    if err := r.DB.Debug().Model(&domain.Enrollment{}).
        Select("students.id, students.user_id, students.student_name, students.address, students.age").
        Joins("JOIN teacher_courses ON enrollments.course_id = teacher_courses.course_id AND teacher_courses.teacher_id = ?", teacherID).
        Joins("JOIN students ON enrollments.user_id = students.user_id").
        Where("teacher_courses.teacher_id = ? AND enrollments.course_id = ?", teacherID, courseID).
        Find(&students).
        Error; err != nil {
		fmt.Println("Error:", err)
        return nil, err
    }
	fmt.Println(students)
    return students, nil
}



