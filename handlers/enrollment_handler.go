// interfaces/enrollment_handler.go
package handlers

import (
    "net/http"
    "strconv"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
    "github.com/shrd09/go_studentDashboard/application"
)

type EnrollmentHandler struct {
    enrollmentService *application.EnrollmentService
}

func NewEnrollmentHandler(enrollmentService *application.EnrollmentService) *EnrollmentHandler {
    return &EnrollmentHandler{enrollmentService: enrollmentService}
}

func (h *EnrollmentHandler) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	// Decode JSON payload from the request body
	var request struct {
		UserID   string  `json:"user_id"`
		CourseID string  `json:"course_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
        return
	}

	// Validate required fields
	if request.UserID == "" || request.CourseID == "" {
		http.Error(w, "Missing user_id or course_id in the request", http.StatusBadRequest)
		return
	}

	// Convert string values to integers if needed
    userID, err := strconv.Atoi(request.UserID)
    if err != nil {
        http.Error(w, "Invalid user_id in the request", http.StatusBadRequest)
        return
    }

    courseID, err := strconv.Atoi(request.CourseID)
    if err != nil {
        http.Error(w, "Invalid course_id in the request", http.StatusBadRequest)
        return
    }
	// Print the extracted data for debugging
	fmt.Printf("UserID: %d, CourseID: %d\n", userID, courseID)

	// Call the enrollment service
	err = h.enrollmentService.CreateEnrollment(userID, courseID)
	if err != nil {
		http.Error(w, "Failed to enroll in the course", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Enrollment successful"))
}

func (h *EnrollmentHandler) GetAllEnrollments(w http.ResponseWriter, r *http.Request) {
    enrollments, err := h.enrollmentService.GetAllEnrollments()
    if err != nil {
        http.Error(w, "Failed to retrieve enrollments", http.StatusInternalServerError)
        return
    }

    // Respond with the list of enrollments
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(enrollments)
}

func (h *EnrollmentHandler) GetEnrollmentsByStudentID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    studentID, _ := strconv.Atoi(vars["student_id"])
	fmt.Println("Received request for student ID:", studentID)
    enrollments, err := h.enrollmentService.GetEnrollmentsByStudentID(studentID)
    if err != nil {
		fmt.Println("Error getting enrollments:", err)
        http.Error(w, "Failed to get student enrollments", http.StatusInternalServerError)
        return
    }

    // Serialize and send enrollments as JSON
    json.NewEncoder(w).Encode(enrollments)
}

func (h *EnrollmentHandler) GetEnrollmentsByCourseID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    courseID, _ := strconv.Atoi(vars["course_id"])

    enrollments, err := h.enrollmentService.GetEnrollmentsByCourseID(courseID)
    if err != nil {
        http.Error(w, "Failed to get enrollments for the course", http.StatusInternalServerError)
        return
    }

    // Serialize and send enrollments as JSON
    json.NewEncoder(w).Encode(enrollments)
}


// Add other methods as needed

func (h *EnrollmentHandler) Show(w http.ResponseWriter, r *http.Request) {
	// Implement Show logic here
}

func (h *EnrollmentHandler) IndexByCourses(w http.ResponseWriter, r *http.Request) {
	// Implement IndexByCourses logic here
}

func (h *EnrollmentHandler) Index(w http.ResponseWriter, r *http.Request) {
	// Implement Index logic here
}

func (h *EnrollmentHandler) UpdateMarks(w http.ResponseWriter, r *http.Request) {
	// Implement UpdateMarks logic here
}
