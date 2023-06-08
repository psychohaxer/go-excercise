package main

import (
	"bufio"     // The bufio library is used to read input from the user
	"fmt"       // The fmt library is used to print messages to the screen
	"os"        // The os library is used to read input from the user
	"strconv"   // The strconv library is used for data type conversion
	"strings"   // The strings library is used to manipulate strings
)

// Check if an integer is a palindrome number
func isPalindrome(num int) bool {
	// Convert the integer to a string
	str := strconv.Itoa(num)
	length := len(str)

	// Loop from index 0 to half the length of the number
	// (rounded down if the length is odd).
	for i := 0; i < length/2; i++ {
		// Condition check to compare the digit at position i with
		// the digit at the opposite position (position length-1-i).
		// If the digits are not the same, it means the number is not a palindrome.
		if str[i] != str[length-1-i] {
			// If a different digit is found, the function immediately returns false,
			// indicating that the number is not a palindrome.
			return false
		}
	}
	return true
}

// Count the number of palindrome numbers between two integers
func countPalindromes(start, end int) int {
	count := 0

	for num := start; num <= end; num++ {
		if isPalindrome(num) {
			count++
		}
	}
	return count
}

// Main function
func main() {
	// Read input from the user
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter numbers in the format 'number1 number2':")

	// Loop until the user enters an empty line
	for scanner.Scan() {
		// Check if the user entered a blank line
		line := scanner.Text()
		if line == "" {
			break
		}

		// Split the input into two numbers
		input := strings.Split(line, " ")

		// Check if the input is valid
		// If the input does not consist of two numbers, then it is invalid
		if len(input) != 2 {
			fmt.Println("Invalid input format. Please enter two numbers separated by a space.")
			continue
		}

		// Convert to integers
		start, err1 := strconv.Atoi(input[0])
		end, err2 := strconv.Atoi(input[1])

		// Check if the input is valid
		// If there is an error during conversion, err1 or err2 will contain a non-nil value
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input format. Please enter two valid numbers.")
			continue
		}

		// Check if the input is valid
		// If the first number < 1 or the second number > 10^9 or the second number <= the first number, then it is invalid
		if start < 1 || end > 1e9 || start >= end {
			fmt.Println("Invalid number range. Make sure the first number >= 1, the second number <= 10^9, and the second number > the first number.")
			continue
		}

		// Count the number of palindrome numbers
		count := countPalindromes(start, end)
		fmt.Println(count)
	}

	// Check if there was an error while reading the input
	if err := scanner.Err(); err != nil {
		fmt.Println("An error occurred while reading the input:", err)
	}

}
