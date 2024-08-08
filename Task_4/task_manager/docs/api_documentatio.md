# Task Manager API Documentation

## Overview

The **Task Manager API** provides a range of endpoints to manage tasks, including creating, updating, retrieving, and deleting tasks. This API helps you efficiently handle tasks in your application.

## Getting Started

To start using the Task Manager API, follow these steps:

1. **Set Up Your Environment:**

   - Ensure you have a working Go environment with the Gin framework installed.
   - Run the Task Manager API server locally or deploy it to your preferred environment.

2. **API Base URL:**

   - Local development: `http://localhost:8080`

3. **Authentication:**

   - No authentication required for these endpoints during this time.

4. **Content-Type:**
   - All requests and responses use JSON format.

## Endpoints

### 1. Create a Task

- **Endpoint:** `POST /tasks`
- **Description:** Creates a new task.
- **Request Body:**
  ```json
  {
    "id": "string", // Unique identifier for the task
    "title": "string", // Title of the task
    "description": "string", // Description of the task
    "due_date": "string", // Due date in ISO 8601 format
    "status": "string" // Status of the task (e.g., "Pending", "In Progress", "Completed")
  }
  ```
- **Response:**
  - **Status 200 OK:**
    ```json
    {
      "message": "Task Created"
    }
    ```
  - **Status 400 Bad Request:**
    ```json
    {
      "error": "Description of the error"
    }
    ```

### 2. Get All Tasks

- **Endpoint:** `GET /tasks`
- **Description:** Retrieves all tasks.
- **Response:**
  - **Status 200 OK:**
    ```json
    {
        "tasks": [
            {
                "id": "string",
                "title": "string",
                "description": "string",
                "due_date": "string",
                "status": "string"
            },
            ...
        ]
    }
    ```

### 3. Get a Task by ID

- **Endpoint:** `GET /tasks/:id`
- **Description:** Retrieves a specific task by its ID.
- **Response:**
  - **Status 200 OK:**
    ```json
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "due_date": "string",
      "status": "string"
    }
    ```
  - **Status 404 Not Found:**
    ```json
    {
      "error": "Task not found"
    }
    ```

### 4. Update a Task

- **Endpoint:** `PUT /tasks/:id`
- **Description:** Updates a specific task by its ID.
- **Request Body:**
  ```json
  {
    "title": "string", // Updated title of the task
    "description": "string", // Updated description of the task
    "due_date": "string", // Updated due date in ISO 8601 format
    "status": "string" // Updated status of the task
  }
  ```
- **Response:**
  - **Status 200 OK:**
    ```json
    {
      "message": "Task updated"
    }
    ```
  - **Status 404 Not Found:**
    ```json
    {
      "error": "Task not found"
    }
    ```

### 5. Delete a Task

- **Endpoint:** `DELETE /tasks/:id`
- **Description:** Deletes a specific task by its ID.
- **Response:**
  - **Status 200 OK:**
    ```json
    {
      "message": "Task deleted"
    }
    ```
  - **Status 404 Not Found:**
    ```json
    {
      "error": "Task not found"
    }
    ```

## Authentication

The Task Manager API currently does not require authentication. when authentication is added in the future, I will include a token in the `Authorization` header.

## Rate and Usage Limits

There are no specific rate limits for the Task Manager API at this time. However, it is advised to use the API responsibly and avoid excessive requests.

## Error Handling

Errors are returned in JSON format. The error response will contain a description of the problem:

- **Status 400 Bad Request:** Indicates that the request was malformed or contained invalid data.
- **Status 404 Not Found:** Indicates that the requested resource does not exist.
- **Status 500 Internal Server Error:** Indicates an unexpected server error.

## Need Help?

For additional help, refer to the following resources:

- [API Documentation](#) – Detailed API documentation and guidelines.
- [Support Forum](#) – Community support and discussions.
