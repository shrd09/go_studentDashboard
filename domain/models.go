package domain
import (
	"time"
)


// User represents the user structure
type User struct {
	ID        int       `gorm:"primarykey"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

// Admin represents the admin structure
type Admin struct {
	ID   int    `gorm:"id"`
	UserID int
	Name string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

// Student represents the student structure
type Student struct {
	ID         int       `gorm:"primarykey"`
	UserID     int       `gorm:"unique;not null"`
	StudentName string    
	Address    string    
	Age        int       
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

// Teacher represents the teacher structure
type Teacher struct {
	ID         int    `gorm:"primarykey"`
	UserID     int    `gorm:"unique;not null"`
	TeacherName string 
	PhoneNo    string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

// Course represents the course structure
type Course struct {
	ID        int    `gorm:"primarykey"`
	CourseName string `gorm:"not null"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

// Enrollment represents the enrollment structure
type Enrollment struct {
	ID        int       `gorm:"primarykey"`
	UserID    int       `gorm:"not null"`
	CourseID  int       `gorm:"not null"`
	Marks     int       
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type TeacherCourse struct {
	ID        int     `gorm:"primarykey" json:"id"`
	TeacherID int 	   `gorm:"not null" json:"teacher_id"`
	CourseID  int		`gorm:"not null" json:"course_id"`
	Year      int
	CreatedAt time.Time 
	UpdatedAt time.Time 
}