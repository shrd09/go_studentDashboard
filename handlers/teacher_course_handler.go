// interfaces/teacher_course_handler.go
package handlers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shrd09/go_studentDashboard/application"
)

type TeacherCourseHandler struct {
	teacherCourseService *application.TeacherCourseService
}

func NewTeacherCourseHandler(teacherCourseService *application.TeacherCourseService) *TeacherCourseHandler {
	return &TeacherCourseHandler{teacherCourseService: teacherCourseService}
}

func (h *TeacherCourseHandler) AssignCourseToTeacher(w http.ResponseWriter, r *http.Request) {
    // Decode JSON payload from the request body
    var requestData map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Extract necessary data from the decoded payload
    teacherID, ok := requestData["teacher_id"].(float64)
    if !ok {
        http.Error(w, "Missing or invalid teacher_id in the request", http.StatusBadRequest)
        return
    }

    courseID, ok := requestData["course_id"].(float64)
    if !ok {
        http.Error(w, "Missing or invalid course_id in the request", http.StatusBadRequest)
        return
    }

    year, ok := requestData["year"].(float64)
    if !ok {
        http.Error(w, "Missing or invalid year in the request", http.StatusBadRequest)
        return
    }

    // Convert float64 values to integers
    teacherIDInt := int(teacherID)
    courseIDInt := int(courseID)
    yearInt := int(year)

    // Call the service to assign the course to the teacher
    err := h.teacherCourseService.AssignCourseToTeacher(teacherIDInt, courseIDInt, yearInt)
    if err != nil {
        http.Error(w, "Failed to assign course to teacher", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Course assigned to teacher successfully"))
}


func (h *TeacherCourseHandler) GetCoursesByTeacherID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    teacherID, _ := strconv.Atoi(vars["teacher_id"])

    courses, err := h.teacherCourseService.GetCoursesByTeacherID(teacherID)
    if err != nil {
        http.Error(w, "Failed to get teacher courses", http.StatusInternalServerError)
        return
    }

    // Serialize and send courses as JSON
    json.NewEncoder(w).Encode(courses)
}


func (h *TeacherCourseHandler) GetEnrolledStudentsByTeacherID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    teacherID, _ := strconv.Atoi(vars["teacher_id"])

    enrolledStudents, err := h.teacherCourseService.GetEnrolledStudentsByTeacherID(teacherID)
    if err != nil {
        http.Error(w, "Failed to get enrolled students for the teacher", http.StatusInternalServerError)
        return
    }

    // Serialize and send enrolled students as JSON
    json.NewEncoder(w).Encode(enrolledStudents)
}


