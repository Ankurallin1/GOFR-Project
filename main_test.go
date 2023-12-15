package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCRUDOperations(t *testing.T) {
	app := fiber.New()
	// Adding a cors for put and delete request	
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		return c.Next()
	})

	app.Get("/students", getStudents)
	app.Post("/students", createStudent)
	app.Get("/students/:id", getStudent)
	app.Put("/students/:id", updateStudent)
	app.Delete("/students/:id", deleteStudent)
	t.Run("CreateStudent", func(t *testing.T) {
		// Define a sample student for testing
		student := Student{
			Name:         "John Doe",
			Email:        "john.doe@example.com",
			CollegeName:  "Sample College",
			EnrollmentNo: "12345",
		}

		// Convert student to JSON
		studentJSON, err := json.Marshal(student)
		assert.NoError(t, err)

		// Create a POST request with the student JSON
		req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(studentJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		// Assert the status code is 200 OK
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("GetStudents", func(t *testing.T) {
		// Create a GET request to retrieve all students
		req := httptest.NewRequest(http.MethodGet, "/students", nil)

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		// Assert the status code is 200 OK
		assert.Equal(t, http.StatusOK, res.StatusCode)

	})

	
	t.Run("UpdateStudent", func(t *testing.T) {
		// Define a sample student for testing
		updatedStudent := Student{
			Name:         "Ronak",
			Email:        "Ronak.email@example.com",
			CollegeName:  "JUET College",
			EnrollmentNo: "201B365",
		}

		// Convert updated student to JSON
		updatedStudentJSON, err := json.Marshal(updatedStudent)
		assert.NoError(t, err)

		// Create a PUT request with the updated student JSON
		req := httptest.NewRequest(http.MethodPut, "/students/657ad262660b9c2342474c33", bytes.NewReader(updatedStudentJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("DeleteStudent", func(t *testing.T) {
		// Create a DELETE request to delete a student
		req := httptest.NewRequest(http.MethodDelete, "/students/657be72ea13c844c358f83b2", nil)

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

}
