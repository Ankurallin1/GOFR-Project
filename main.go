package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
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
func createStudent(c *fiber.Ctx) error {
	var student Student
	if err := c.BodyParser(&student); err != nil {
		return err
	}

	collection := client.Database("testdb").Collection("students")
	result, err := collection.InsertOne(context.TODO(), student)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

// Function to get all students
func getStudents(c *fiber.Ctx) error {
	var students []Student
	collection := client.Database("testdb").Collection("students")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var student Student
		cursor.Decode(&student)
		students = append(students, student)
	}
	return c.JSON(students)
}

// Function to get a single student by ID
func getStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var student Student
	collection := client.Database("testdb").Collection("students")
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&student)
	if err != nil {
		return err
	}
	return c.JSON(student)
}

// Function to update a student by ID
func updateStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var updatedStudent Student
	if err := c.BodyParser(&updatedStudent); err != nil {
		return err
	}

	collection := client.Database("testdb").Collection("students")
	filter := bson.M{"_id": objectID}
	update := bson.D{{"$set", updatedStudent}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return c.JSON("Student updated successfully")
}

// Function to delete a student by ID
func deleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := client.Database("testdb").Collection("students")
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return c.JSON("Student deleted successfully")
}

func main() {
	app := fiber.New()

	// Define CRUD endpoints for students
	app.Get("/students", getStudents)
	app.Get("/students/:id", getStudent)
	app.Post("/students", createStudent)
	app.Put("/students/:id", updateStudent)
	app.Delete("/students/:id", deleteStudent)

	// Start the server
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
