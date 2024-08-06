package controllers

import (
	"bufio"
	"fmt"
	"library_managment/models"
	"library_managment/services"
	"os"
	"strconv"
	"strings"
)

func ShowMenu() {
	fmt.Println("1. Add a new book")
	fmt.Println("2. Remove an existing book")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List all available books")
	fmt.Println("6. List all borrowed books by a member")
	fmt.Println("7. Exit")
	fmt.Print("Choose an option: ")
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func HandleAddBook(library *services.Library) {
	fmt.Print("Enter book ID: ")
	id, _ := strconv.Atoi(GetInput())
	fmt.Print("Enter book title: ")
	title := GetInput()
	fmt.Print("Enter book author: ")
	author := GetInput()
	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func HandleRemoveBook(library *services.Library) {
	fmt.Print("Enter book ID: ")
	id, _ := strconv.Atoi(GetInput())
	library.RemoveBook(id)
	fmt.Println("Book removed successfully!")
}

func HandleBorrowBook(library *services.Library) {
	fmt.Print("Enter book ID: ")
	bookID, _ := strconv.Atoi(GetInput())
	fmt.Print("Enter member ID: ")
	memberID, _ := strconv.Atoi(GetInput())
	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func HandleReturnBook(library *services.Library) {
	fmt.Print("Enter book ID: ")
	bookID, _ := strconv.Atoi(GetInput())
	fmt.Print("Enter member ID: ")
	memberID, _ := strconv.Atoi(GetInput())
	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func HandleListAvailableBooks(library *services.Library) {
	books := library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
	} else {
		fmt.Println("Available books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func HandleListBorrowedBooks(library *services.Library) {
	fmt.Print("Enter member ID: ")
	memberID, _ := strconv.Atoi(GetInput())
	books := library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("Borrowed books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func RunLibraryManagementSystem() {
	library := services.NewLibrary()
	for {
		ShowMenu()
		choice := GetInput()
		switch choice {
		case "1":
			HandleAddBook(library)
		case "2":
			HandleRemoveBook(library)
		case "3":
			HandleBorrowBook(library)
		case "4":
			HandleReturnBook(library)
		case "5":
			HandleListAvailableBooks(library)
		case "6":
			HandleListBorrowedBooks(library)
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
		fmt.Println()
	}
}
