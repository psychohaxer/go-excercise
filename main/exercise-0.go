package main

import (
	"strconv" // The strconv library is used for data type conversion
)

// Check if an integer is a palindrome number
func IsPalindrome(num int) bool {
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
func CountPalindromes(start, end int) int {
	count := 0

	for num := start; num <= end; num++ {
		if IsPalindrome(num) {
			count++
		}
	}
	return count
}
