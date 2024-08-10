# Task Manager API Documentation

## Overview

This API allows you to manage tasks. You can create, retrieve, update, and delete tasks. Each task has an ID, title, description, due date, and status.

## Base URL

The base URL for accessing the API is: `http://localhost:8080`

## Endpoints

### Get All Tasks

- **URL**: `/tasks`
- **Method**: `GET`
- **Description**: Retrieves all tasks.
- **Response**:
  - **200 OK**: Returns a list of tasks.
  - **500 Internal Server Error**: Returns an error message.

### Get Task by ID

- **URL**: `/tasks/:id`
- **Method**: `GET`
- **Description**: Retrieves a task by its ID.
- **Parameters**:
  - `id` (string, required): The ID of the task to retrieve.
- **Response**:
  - **200 OK**: Returns the task with the specified ID.
  - **404 Not Found**: Returns an error message if the task is not found.
  - **500 Internal Server Error**: Returns an error message.

### Create a New Task

- **URL**: `/tasks`
- **Method**: `POST`
- **Description**: Creates a new task.
- **Request Body**:
  - `id` (string): The ID of the task.
  - `title` (string): The title of the task.
  - `description` (string): The description of the task.
  - `due_date` (string): The due date of the task.
  - `status` (string): The status of the task.
- **Response**:
  - **201 Created**: Returns the ID of the newly created task.
  - **400 Bad Request**: Returns an error message if the request body is invalid.
  - **500 Internal Server Error**: Returns an error message.

### Update Task by ID

- **URL**: `/tasks/:id`
- **Method**: `PUT`
- **Description**: Updates a task by its ID.
- **Parameters**:
  - `id` (string, required): The ID of the task to update.
- **Request Body**:
  - `title` (string): The title of the task.
  - `description` (string): The description of the task.
  - `due_date` (string): The due date of the task.
  - `status` (string): The status of the task.
- **Response**:
  - **200 OK**: Returns a success message.
  - **400 Bad Request**: Returns an error message if the request body is invalid.
  - **404 Not Found**: Returns an error message if the task is not found.
  - **500 Internal Server Error**: Returns an error message.

### Delete Task by ID

- **URL**: `/tasks/:id`
- **Method**: `DELETE`
- **Description**: Deletes a task by its ID.
- **Parameters**:
  - `id` (string, required): The ID of the task to delete.
- **Response**:
  - **200 OK**: Returns a success message.
  - **404 Not Found**: Returns an error message if the task is not found.
  - **500 Internal Server Error**: Returns an error message.

markdown
Copy code

## Example Requests

### Get All Tasks

```
curl -X GET http://localhost:8080/tasks
```

### Get Task by ID

```
curl -X GET http://localhost:8080/tasks/<id>
```

### Create a New Task

```
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
  "id": "1",
  "title": "Task 1",
  "description": "This is the first task",
  "due_date": "2024-08-15",
  "status": "Pending"
}'
```

### Update Task by ID

```
curl -X PUT http://localhost:8080/tasks/<id> \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2024-09-01",
  "status": "Completed"
}'
```

### Delete Task by ID

```
curl -X DELETE http://localhost:8080/tasks/<id>
```

## Data Models

## Task

The `Task` model represents a task in the system. It has the following fields:

- `id` (string): The unique identifier for the task.
- `title` (string): The title of the task.
- `description` (string): A detailed description of the task.
- `due_date` (string): The due date for the task.
- `status` (string): The current status of the task (e.g., "Pending", "Completed").

## Error Handling

The API returns standard HTTP status codes to indicate the success or failure of an API request. In case of an error, the response will contain a JSON object with an `error` field describing the issue.

- **400 Bad Request**: The request was invalid or cannot be served. An accompanying error message will explain why.
- **404 Not Found**: The requested resource could not be found.
- **500 Internal Server Error**: The server encountered an error and could not complete the request.

## Running the Application

To run the Task Manager API, follow these steps:

1.  Ensure you have [MongoDB](https://www.mongodb.com/) installed and running.
2.  Clone the repository.
3.  Navigate to the project directory.
4.  Run the following command to start the API server:

    ```
    go run main.go
    ```

## Server

The server will start and listen on [http://localhost:8080](http://localhost:8080).

## Database Configuration

The API uses MongoDB as the database. The connection details are configured in the `database` package. Ensure that MongoDB is running locally on the default port (27017) or adjust the connection settings as needed.

## Dependencies

The API relies on several Go packages, including:

- **Gin**: A web framework for Go.
- **MongoDB Go Driver**: The official MongoDB driver for Go.

To install the dependencies, run:

```
go mod tidy
```

This will ensure that all necessary packages are downloaded and included this project.
