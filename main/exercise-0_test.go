package main

import (
	"testing" // Package testing provides support for automated testing of Go packages with the testing framework.
)

func TestIsPalindrome(t *testing.T) {
	// Test case with a palindrome number
	if !IsPalindrome(121) {
		t.Errorf("IsPalindrome(121) = false, want true")
	}

	// Test case with a non-palindrome number
	if IsPalindrome(123) {
		t.Errorf("IsPalindrome(123) = true, want false")
	}
}

func TestCountPalindromes(t *testing.T) {
	// Test case with range 1 to 10, there should be 9 palindrome numbers
	count := CountPalindromes(1, 10)
	expected := 9
	if count != expected {
		t.Errorf("CountPalindromes(1, 10) = %d, want %d", count, expected)
	}

	// Test case with range 99 to 100, there should be 1 palindrome number
	count = CountPalindromes(99, 100)
	expected = 1
	if count != expected {
		t.Errorf("CountPalindromes(99, 100) = %d, want %d", count, expected)
	}

	// Test case with range 21 to 31, there should be 1 palindrome number
	count = CountPalindromes(21, 31)
	expected = 1
	if count != expected {
		t.Errorf("CountPalindromes(21, 31) = %d, want %d", count, expected)
	}

	// Test case with range 1 to 100, there should be 18 palindrome numbers
	count = CountPalindromes(1, 100)
	expected = 18
	if count != expected {
		t.Errorf("CountPalindromes(1, 100) = %d, want %d", count, expected)
	}

	// Test case with range 1000 to 2000, there should be 108 palindrome numbers
	count = CountPalindromes(1000, 2000)
	expected = 108
	if count != expected {
		t.Errorf("CountPalindromes(1000, 2000) = %d, want %d", count, expected)
	}

	// Additional test case with range 1 to 10, the expected output is 9
	count = CountPalindromes(1, 10)
	expected = 9
	if count != expected {
		t.Errorf("CountPalindromes(1, 10) = %d, want %d", count, expected)
	}

	// Additional test case with range 99 to 100, the expected output is 9
	count = CountPalindromes(99, 100)
	expected = 9
	if count != expected {
		t.Errorf("CountPalindromes(99, 100) = %d, want %d", count, expected)
	}

	// Additional test case with range 21 to 31, the expected output is 1
	count = CountPalindromes(21, 31)
	expected = 1
	if count != expected {
		t.Errorf("CountPalindromes(21, 31) = %d, want %d", count, expected)
	}
}

// Unit test for the main functions in the main package
func TestMainFunctions(t *testing.T) {
	// Test case to ensure that the main functions do not produce any errors
	main()
}
