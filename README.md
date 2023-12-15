# Go CRUD Application with Fiber and MongoDB
The Go CRUD Application is a simple yet robust project built using the Fiber web framework and MongoDB database. 
It facilitates basic CRUD operations (Create, Read, Update, Delete) for managing student records. Fork, customize, 
and extend to meet your specific needs with the flexibility of Go and Fiber.
## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Automated Testing](#automated-testing)
- [Contributing](#contributing)


## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://go.dev/doc/install)
- [MongoDB](https://www.mongodb.com/docs/manual/administration/install-community/)
- [PostMan](https://www.postman.com/downloads/)

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/Ankurallin1/GOFR-Project
   
   cd GOFR-Project
2. Download Dependencies:

   ```bash
   go mod download

3. Verify Dependencies:

    ```bash
   cat go.sum
4. Run Project
    
    ```bash
    go run main.go

5. Open Server
    
    ```bash
    http://localhost:8080/students
    
## API Endpoints

The application provides RESTful API endpoints for CRUD operations on student records.

### Base URL

The base URL for all endpoints is:

### List of Endpoints

#### 1. **Get All Students**

- **Endpoint:**
  - `GET /students`

- **Description:**
  - Retrieves a list of all students.

- **Example Response:**
  ```json
  [    {      "id": "657ad262660b9c2342474c33",      "name": "ankur",      "email": "ankursri@gmail.com",      "collegeName": "JUET",      "enrollmentNo": "201B049"    },    {      "id": "657be72ea13c844c358f83b2",      "name": "John Doe",      "email": "john.doe@example.com",      "collegeName": "Sample College",      "enrollmentNo": "12345"    }  ]

![Screenshot (17)](https://github.com/Ankurallin1/GOFR-Project/assets/91478132/c5648829-eb6b-4ca5-9093-9be7d7e58edf)

#### 2. **Get Student by ID**

- **Endpoint:**
  - `GET /students/:id`
- **Parameters:**
  - `id`: The unique identifier for the student.

- **Description:**
  - Retrieves information about a specific student based on the provided ID.

- **Example Response:**
  ```json
  [    {      "id": "657ad262660b9c2342474c33",      "name": "ankur",      "email": "ankursri@gmail.com",      "collegeName": "JUET",      "enrollmentNo": "201B049"    }   ]  

![Screenshot (18)](https://github.com/Ankurallin1/GOFR-Project/assets/91478132/e359be2c-6ee7-40f5-93d0-b7ba02c1387a)

#### 3. **Create a New Student**

- **Endpoint:**
  - `POST /students`
- **Request Body:**
  ```json
  [    {      "name": "John Doe",      "email": "john.doe@example.com",      "collegeName": "Sample College",      "enrollmentNo": "12345"    }   ]  


- **Description:**
  - Creates a new student record.
    
- **Example Response:**
  ```json
  {      "id": "generated-student-id",      "name": "John Doe",      "email": "john.doe@example.com",      "collegeName": "Sample College",      "enrollmentNo": "12345"    }


![Screenshot (19)](https://github.com/Ankurallin1/GOFR-Project/assets/91478132/30226e36-e1a3-40f4-871a-781801ce3e4c)

#### 4. **Update Student by ID**

- **Endpoint:**
  - `PUT /students/:id`
- **Parameters:**
  - `id`: The unique identifier for the student.

- **Request Body:**
  ```json
  [    {      "name": "Updated Name",      "email": "updated.email@example.com",      "collegeName": "Updated College",      "enrollmentNo": "54321"    }   ]  

- **Description:**
  - Updates information about a specific student based on the provided ID.
    
- **Example Response:**
  ```json
  "Student updated successfully"
  
![Screenshot (21)](https://github.com/Ankurallin1/GOFR-Project/assets/91478132/a4a7ed03-7a37-4106-b33c-f21c1cd61316)

 
#### 5. **Delete Student by ID**

- **Endpoint:**
  - `DELETE /students/:id`

- **Parameters:**
  - `id`: The unique identifier for the student.

- **Description:**
  - Deletes a specific student record based on the provided ID.

- **Example Response:**
  ```json
  "Student deleted successfully"

![Screenshot (20)](https://github.com/Ankurallin1/GOFR-Project/assets/91478132/ced1449e-e7df-4f1c-a9ef-c3ff60c3ab77)

## Automated Testing

The project includes test cases that can be executed using Go's testing framework.

### Changes Required

- **Change the `id` in below lines of `main_test.go` :**

  - Line 79 for `UpdateStudent` 
     ```bash
     req := httptest.NewRequest(http.MethodPut, "/students/{id}", bytes.NewReader(updatedStudentJSON))
     
  - Line 91 for `DeleteStudent` 
     ```bash
     req := httptest.NewRequest(http.MethodDelete, "/students/{id}", nil)


### Running Tests

- Use the following command to run the tests:

  ```bash
  go test


<img style="padding-left: 100px;" width="1000" alt="test" src="https://github.com/Ankurallin1/GOFR-Project/assets/91478132/eb66b333-e06b-45cb-ab9f-00174340b2ea">


## Contributing

Feel free to contribute to this project. Fork the repository, make your changes, and submit a pull request.


