// interfaces/teacher_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/shrd09/go_studentDashboard/application"
)

type TeacherHandler struct {
	TeacherService *application.TeacherService
}

func NewTeacherHandler(teacherService *application.TeacherService) *TeacherHandler {
	return &TeacherHandler{TeacherService: teacherService}
}


// func (h *TeacherHandler) GetAllTeachers(w http.ResponseWriter, r *http.Request) {
//     teachers, err := h.TeacherService.GetAllTeachers()
//     if err != nil {
//         http.Error(w, "Failed to fetch teachers", http.StatusInternalServerError)
//         return
//     }

//     // Respond with the teachers in JSON format
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(teachers)
// }

func (h *TeacherHandler) GetEnrolledStudentsByTeacherID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacherID, _ := strconv.Atoi(vars["teacher_id"])

	students, err := h.TeacherService.GetEnrolledStudentsByTeacherID(teacherID)
	if err != nil {
		http.Error(w, "Failed to get enrolled students for the teacher", http.StatusInternalServerError)
		return
	}

	// Serialize and send students as JSON
	json.NewEncoder(w).Encode(students)
}



func (h *TeacherHandler) GetEnrolledStudentsByCourseID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    teacherID, err := strconv.Atoi(vars["teacher_id"])
    if err != nil {
        http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
        return
    }

    courseID, err := strconv.Atoi(vars["courseID"])
    if err != nil {
        http.Error(w, "Invalid course ID", http.StatusBadRequest)
        return
    }

    enrolledStudents, err := h.TeacherService.GetEnrolledStudentsByCourseID(teacherID, courseID)
    if err != nil {
        http.Error(w, "Failed to fetch enrolled students", http.StatusInternalServerError)
        return
    }

    // Respond with the enrolled students in JSON format
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(enrolledStudents)
}
