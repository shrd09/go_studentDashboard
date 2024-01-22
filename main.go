package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    "github.com/rs/cors"
    "github.com/shrd09/go_studentDashboard/adapters"
    "github.com/shrd09/go_studentDashboard/application"
    "github.com/shrd09/go_studentDashboard/handlers"

)

func main() {
    db, err := adapters.NewDatabase()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Success!")

    c := cors.Default()
    r := mux.NewRouter()

    // Enable CORS
    corsHandler := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Allow all origins for CORS
            w.Header().Set("Access-Control-Allow-Origin", "*")

            // Handle preflight requests
            if r.Method == http.MethodOptions {
                w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, POST, OPTIONS")
                w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
                w.Header().Set("Access-Control-Allow-Credentials", "true")
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusNoContent)  // Respond with 204 for preflight requests
                return
            }

            // For non-preflight requests, call the next handler
            next.ServeHTTP(w, r)
        })
    }

    r.Use(corsHandler)
    handler := c.Handler(r)

    // Initialize repositories
    // teacherRepo := adapters.NewTeacherRepository(db)
    enrollmentRepo := adapters.NewEnrollmentRepository(db)
    courseRepo := adapters.NewCourseRepository(db)
    teacherCourseRepo := adapters.NewTeacherCourseRepository(db)

    // Initialize application service
    enrollmentService := application.NewEnrollmentService(enrollmentRepo)
    courseService := application.NewCourseService(courseRepo)
    teacherCourseService := application.NewTeacherCourseService(teacherCourseRepo, enrollmentRepo)
    teacherService := application.NewTeacherService(teacherCourseRepo, enrollmentRepo)

    // Initialize HTTP handler
    enrollmentHandler := handlers.NewEnrollmentHandler(enrollmentService)
    courseHandler := handlers.NewCourseHandler(courseService)
    teacherCourseHandler := handlers.NewTeacherCourseHandler(teacherCourseService)
    teacherHandler := handlers.NewTeacherHandler(teacherService)

    // Defining routes
    r.HandleFunc("/courses", courseHandler.CreateCourse).Methods("POST")
    r.HandleFunc("/courses/{id}", courseHandler.DeleteCourse).Methods("DELETE")
    r.HandleFunc("/courses", courseHandler.GetCourses).Methods("GET")
    r.HandleFunc("/teacher-courses", teacherCourseHandler.AssignCourseToTeacher).Methods("POST")
    r.HandleFunc("/enrollments/{student_id}", enrollmentHandler.GetEnrollmentsByStudentID).Methods("GET")

	r.HandleFunc("/enrollments", enrollmentHandler.CreateEnrollment).Methods("POST")
    r.HandleFunc("/enrollments", enrollmentHandler.GetAllEnrollments).Methods("GET")
    r.HandleFunc("/teacher/courses/{teacher_id}", teacherCourseHandler.GetCoursesByTeacherID).Methods("GET")
    // r.HandleFunc("/teachers/{teacher_id}/courses/{course_id}/students", teacherHandler.GetStudentsByCourseID).Methods("GET")
    // r.HandleFunc("/teacher-courses/students/{teacher_id}/{course_id}", teacherCourseHandler.GetStudentsByTeacherAndCourseID).Methods("GET")
    // r.HandleFunc("/teacher-courses/students/enrolled/{teacher_id}", teacherCourseHandler.GetEnrolledStudentsByTeacherID).Methods("GET")
    r.HandleFunc("/teachers/{teacher_id}/enrolled-students", teacherHandler.GetEnrolledStudentsByTeacherID).Methods("GET")
    // r.HandleFunc("/teachers", teacherHandler.GetAllTeachers).Methods("GET")
    r.HandleFunc("/teachers/{teacher_id}/courses/{courseID}/enrolled-students", teacherHandler.GetEnrolledStudentsByCourseID).Methods("GET")
    http.Handle("/", r)
    log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

