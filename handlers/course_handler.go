package handlers

import (
    "net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
    "github.com/shrd09/go_studentDashboard/application"
)

// CourseHandler is a struct responsible for handling HTTP requests related to courses
type CourseHandler struct {
	courseService *application.CourseService
}

func NewCourseHandler(courseService *application.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
    var requestData map[string]string  //store data in key-value pairs where both keys and values are strings.

    // Decode JSON payload from the request body
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Extract CourseName from the decoded data
    CourseName, ok := requestData["course_name"]
    if !ok {
        http.Error(w, "Missing course_name in the request", http.StatusBadRequest)
        return
    }

    // Print the extracted CourseName for debugging
    fmt.Println("CourseName:", CourseName)

    err := h.courseService.CreateCourse(CourseName)
    if err != nil {
        http.Error(w, "Failed to create course", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Course created successfully"))
}

func (h *CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
    // Use mux.Vars to get route parameters
    vars := mux.Vars(r)
    courseID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid course ID", http.StatusBadRequest)
        return
    }

    err = h.courseService.DeleteCourse(courseID)
    if err != nil {
        http.Error(w, "Failed to delete course", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Course deleted successfully"))
}

func (h *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
    courses, err := h.courseService.GetCourses()
    if err != nil {
        http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
        return
    }

	 // Respond with the courses in JSON format
	 w.Header().Set("Content-Type", "application/json")
	 json.NewEncoder(w).Encode(courses)
}