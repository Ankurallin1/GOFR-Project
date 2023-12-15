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

	// Set up your CRUD endpoints for testing
	app.Get("/students", getStudents)
	app.Post("/students", createStudent)
	app.Get("/students/:id", getStudent)
	
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

	
}
