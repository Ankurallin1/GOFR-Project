package main

import (
	"bytes"
	"strings"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	// "go.mongodb.org/mongo-driver/bson/primitive"

)
func TestCreateStudent(t *testing.T) {
	// Create a request body with a sample student
	student := Student{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		CollegeName:  "Sample College",
		EnrollmentNo: "EN12345",
	}
	body, err := json.Marshal(student)
	if err != nil {
		t.Fatal(err)
	}

	// Create a POST request to create a new student
	req, err := http.NewRequest("POST", "/students", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Call the createStudent handler function
	createStudent(rec, req)

	// Check the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	responseBody := rec.Body.String()

if !strings.Contains(responseBody, "InsertedID") {
    t.Errorf("Handler returned unexpected body: InsertedID field not found")
}
}

func TestGetStudents(t *testing.T) {
	// Create a GET request to get all students
	req, err := http.NewRequest("GET", "/students", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Call the getStudents handler function
	getStudents(rec, req)

	// Check the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	
}

