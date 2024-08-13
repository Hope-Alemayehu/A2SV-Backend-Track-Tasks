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

// create a libraryControler struct
type LibraryControler struct {
	LibraryInt services.LibraryManager
}

// create an instance of library manager and return struct
func NewLibraryController(LibraryInt services.LibraryManager) *LibraryControler {
	return &LibraryControler{LibraryInt: LibraryInt}
}

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

func HandleAddBook(library *LibraryControler) {
	fmt.Print("Enter book ID: ")
	id, _ := strconv.Atoi(GetInput())
	fmt.Print("Enter book title: ")
	title := GetInput()
	fmt.Print("Enter book author: ")
	author := GetInput()
	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	library.LibraryInt.AddBook(book)
	fmt.Println("Book added successfully!")
}

func HandleRemoveBook(library *LibraryControler) {
	fmt.Print("Enter book ID: ")
	id, _ := strconv.Atoi(GetInput())
	library.LibraryInt.RemoveBook(id)
	fmt.Println("Book removed successfully!")
}

func HandleBorrowBook(library *LibraryControler) {
	fmt.Print("Enter book ID: ")
	bookID, _ := strconv.Atoi(GetInput())
	fmt.Print("Enter member ID: ")
	memberID, _ := strconv.Atoi(GetInput())
	err := library.LibraryInt.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func HandleReturnBook(library *LibraryControler) {
	fmt.Print("Enter book ID: ")
	bookID, _ := strconv.Atoi(GetInput())
	fmt.Print("Enter member ID: ")
	memberID, _ := strconv.Atoi(GetInput())
	err := library.LibraryInt.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func HandleListAvailableBooks(library *LibraryControler) {
	books := library.LibraryInt.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
	} else {
		fmt.Println("Available books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func HandleListBorrowedBooks(library *LibraryControler) {
	fmt.Print("Enter member ID: ")
	memberID, _ := strconv.Atoi(GetInput())
	books := library.LibraryInt.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("Borrowed books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func RunLibraryManagementSystem(controller *LibraryControler) {

	for {
		ShowMenu()
		choice := GetInput()
		switch choice {
		case "1":
			HandleAddBook(controller)
		case "2":
			HandleRemoveBook(controller)
		case "3":
			HandleBorrowBook(controller)
		case "4":
			HandleReturnBook(controller)
		case "5":
			HandleListAvailableBooks(controller)
		case "6":
			HandleListBorrowedBooks(controller)
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
		fmt.Println()
	}
}
