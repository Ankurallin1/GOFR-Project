package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define a struct to represent your data model
type Student struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	CollegeName  string `json:"collegeName,omitempty" bson:"collegeName,omitempty"`
	EnrollmentNo string `json:"enrollmentNo,omitempty" bson:"enrollmentNo,omitempty"`
}

var client *mongo.Client

// Connect to MongoDB
func init() {
	// Set up client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Create a MongoDB client
	var err error
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// Function to create a new student
func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	collection := client.Database("testdb").Collection("students")
	result, err := collection.InsertOne(context.TODO(), student)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

// Function to get all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []Student
	collection := client.Database("testdb").Collection("students")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var student Student
		cursor.Decode(&student)
		students = append(students, student)
	}
	json.NewEncoder(w).Encode(students)
}

// Function to get a single student by ID
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var student Student
	collection := client.Database("testdb").Collection("students")
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(student)
}

// Function to update a student by ID
func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	var updatedStudent Student
	_ = json.NewDecoder(r.Body).Decode(&updatedStudent)

	collection := client.Database("testdb").Collection("students")
	filter := bson.M{"_id": objectID}
	update := bson.D{{"$set", updatedStudent}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Student updated successfully")
}

// Function to delete a student by ID
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testdb").Collection("students")
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Student deleted successfully")
}

func main() {
	r := mux.NewRouter()

	// Define CRUD endpoints for students
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	// Start the server
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(":8080", r))
}
