# Library Management System Documentation

## Overview

This project implements a simple console-based library management system in Go. It demonstrates the use of structs, interfaces, methods, slices, and maps.

## Features

- Add a new book
- Remove an existing book
- Borrow a book
- Return a book
- List all available books
- List all borrowed books by a member

## Directory Structure

library_management/
├── main.go
├── controllers/
│ └── library_controller.go
├── models/
│ └── book.go
│ └── member.go
├── services/
│ └── library_service.go
├── docs/
│ └── documentation.md
└── go.mod

markdown
Copy code

## Running the Application

1. Navigate to the `library_management` directory.
2. Run `go run main.go`.

## Implementation Details

- `Book` struct represents
