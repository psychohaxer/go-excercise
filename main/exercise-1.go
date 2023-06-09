package main

import (
	"bufio"     // The bufio library is used to read input from the user
	"fmt"       // The fmt library is used to print messages to the screen
	"os"        // The os library is used to read input from the user
	"strconv"   // The strconv library is used for data type conversion
	"strings"   // The strings library is used to manipulate strings
)

// Book is a data type that stores book information
type Book struct {
	Category string
	Title    string
	Size     int
}

// indexCategory is the book category in the order according to the index
var indexCategory = []string{"6", "7", "0", "9", "4", "8", "1", "2", "5", "3"}

// parseInput converts the book notation into a book object
func parseInput(input string) Book {
	tokens := strings.Split(input, "")
	category := tokens[0]
	title := tokens[1]
	sizeArray := tokens[2:4]
	sizeStr := strings.Join(sizeArray, "")

	// Convert size to int
	size, _ := strconv.Atoi(sizeStr)
	return Book{category, title, size}
}

func main(){
	// Store input
	input, inputLine := "", ""
	var inputLineArr []string
	
	// Store book objects
	var booksUnsorted []Book
	var booksSorted []Book

	// Create scanner
	scanner := bufio.NewScanner(os.Stdin)
	
	// Receive input
	fmt.Println("Enter input:")
	for scanner.Scan() {
		input = scanner.Text()

		// Stop if input is empty
		if input == "" {
			break
		}

		inputLine += input
	}

	inputLineArr = strings.Split(inputLine, " ")

	// Loop for each inputLineArr
	for _, notation := range inputLineArr {
		book := parseInput(notation)
		booksUnsorted = append(booksUnsorted, book)
	}

	fmt.Println("Unsorted:", booksUnsorted)
	
	// Sort booksUnsorted based on Size
	for i := 0; i < len(booksUnsorted); i++ {
		for j := i + 1; j < len(booksUnsorted); j++ {
			if booksUnsorted[i].Size > booksUnsorted[j].Size {
				booksUnsorted[i], booksUnsorted[j] = booksUnsorted[j], booksUnsorted[i]
			}
		}
	}

	// Reverse the order of booksUnsorted
	for i := 0; i < len(booksUnsorted)/2; i++ {
		j := len(booksUnsorted) - i - 1
		booksUnsorted[i], booksUnsorted[j] = booksUnsorted[j], booksUnsorted[i]
	}

	// Loop for each indexCategory
	for _, category := range indexCategory {

		// Loop for each book in booksUnsorted
		for _, book := range booksUnsorted {
		
			// If the category is the same as the book category
			if category == book.Category {
				// Add the Book to booksSorted
				booksSorted = append(booksSorted, book)
			}
		}
	}

	fmt.Println("Sorted:", booksSorted)
	
}
