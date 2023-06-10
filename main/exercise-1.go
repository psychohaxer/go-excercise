package main

import (
	"strconv" // The strconv library is used for data type conversion
	"strings" // The strings library is used to manipulate strings
)

// Book is a data type that stores book information
type Book struct {
	Category string
	Title    string
	Size     int
}

// indexCategory is the book category in the order according to the index
var indexCategory = []string{"6", "7", "0", "9", "4", "8", "1", "2", "5", "3"}

// ParseInput converts the book notation into a book object
func ParseInput(input string) Book {
	tokens := strings.Split(input, "")
	category := tokens[0]
	title := tokens[1]
	sizeArray := tokens[2:4]
	sizeStr := strings.Join(sizeArray, "")

	// Convert size to int
	size, _ := strconv.Atoi(sizeStr)
	return Book{category, title, size}
}

// BooksParser converts the input line into an array of book objects
func BooksParser(inputLine string) []Book {
	var books []Book
	inputLineArr := strings.Split(inputLine, " ")

	for _, notation := range inputLineArr {
		book := ParseInput(notation)
		books = append(books, book)
	}

	return books
}

// BookSorter sort the books based on the category and size
func BookSorter(unsorted []Book) []Book {
	var sorted []Book

	// Sort unsorted Book based on Size
	for i := 0; i < len(unsorted); i++ {
		for j := i + 1; j < len(unsorted); j++ {
			if unsorted[i].Size > unsorted[j].Size {
				unsorted[i], unsorted[j] = unsorted[j], unsorted[i]
			}
		}
	}

	// Reverse the order of unsorted
	for i := 0; i < len(unsorted)/2; i++ {
		j := len(unsorted) - i - 1
		unsorted[i], unsorted[j] = unsorted[j], unsorted[i]
	}

	// Loop for each indexCategory
	for _, category := range indexCategory {

		// Loop for each book in booksUnsorted
		for _, book := range unsorted {

			// If the category is the same as the book category
			if category == book.Category {
				// Add the Book to booksSorted
				sorted = append(sorted, book)
			}
		}
	}
	return sorted
}

// StringToSortedBook converts the input string into a sorted array of book objects
func StringToSortedBook(input string) []Book {
	booksUnsorted := BooksParser(input)
	booksSorted := BookSorter(booksUnsorted)
	return booksSorted
}

// DisplayBooks converts the array of book objects into a string
func DisplayBooks(books []Book) string {
	var result string
	for _, book := range books {
		result += book.Category + book.Title + strconv.Itoa(book.Size) + " "
	}
	return result
}

// SortTheBooks is a function that sorts the books
func SortTheBooks(input string) string {
	booksSorted := StringToSortedBook(input)
	result := DisplayBooks(booksSorted)
	return result
}
