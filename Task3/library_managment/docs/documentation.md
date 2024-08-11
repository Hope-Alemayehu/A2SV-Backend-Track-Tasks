# Library Management System Documentation

## Overview

The Library Management System is a Go-based application designed to manage books and members in a library. The system allows users to add, remove, borrow, and return books, as well as list available and borrowed books.

### Controllers

The controllers package contains the logic for handling user input and interacting with the library services.

### Functions

- **ShowMenu**: Displays the main menu to the user.
- **GetInput**: Reads user input from the console.
- **HandleAddBook**: Adds a new book to the library.
- **HandleRemoveBook**: Removes a book from the library.
- **HandleBorrowBook**: Borrows a book from the library.
- **HandleReturnBook**: Returns a borrowed book to the library.
- **HandleListAvailableBooks**: Lists all available books in the library.
- **HandleListBorrowedBooks**: Lists all borrowed books by a member.
- **RunLibraryManagementSystem**: Runs the library management system.
  Models
  The models package contains the data structures for books and members.

### Structs

- **Book**: Represents a book with an ID, title, author, and status.
- **Member**: Represents a member with an ID, name, and borrowed books.

### Services

The services package contains the logic for managing books and members.

#### Interfaces

- **LibraryManager**: Defines the interface for library management.
- **Structs Library**: Implements the LibraryManager interface.
- **AddBook**: Adds a new book to the library.
- **RemoveBook**: Removes a book from the library.
- **BorrowBook**: Borrows a book from the library.
- **ReturnBook**: Returns a borrowed book to the library.
- **ListAvailableBooks**: Lists all available books in the library.
- **ListBorrowedBooks**: Lists all borrowed books by a member.

### Main

The main package contains the entry point for the application.

#### Functions

- **main**: Runs the library management system.

### Usage

To use the Library Management System, simply run the application and follow the prompts. The system will guide you through the available options.

### System Requirements

- Go programming language (version 1.15 or later)
- console or terminal to run the application

### Known Limitations

- The system does not persist data across sessions. All data is lost when the application is closed.
- The system does not have any error handling or validation for user input.
