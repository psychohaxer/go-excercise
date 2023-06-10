package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	// Test case 1
	input1 := "3A13"
	expected1 := Book{"3", "A", 13}
	result1 := ParseInput(input1)
	if !reflect.DeepEqual(expected1, result1) {
		t.Errorf("Expected %+v, but got %+v", expected1, result1)
	}

	// Test case 2
	input2 := "5X19"
	expected2 := Book{"5", "X", 19}
	result2 := ParseInput(input2)
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("Expected %+v, but got %+v", expected2, result2)
	}

	// Test case 3
	input3 := "9Y20"
	expected3 := Book{"9", "Y", 20}
	result3 := ParseInput(input3)
	if !reflect.DeepEqual(expected3, result3) {
		t.Errorf("Expected %+v, but got %+v", expected3, result3)
	}

	// Additional test case with input "2C18"
	input4 := "2C18"
	expected4 := Book{"2", "C", 18}
	result4 := ParseInput(input4)
	if !reflect.DeepEqual(expected4, result4) {
		t.Errorf("Expected %+v, but got %+v", expected4, result4)
	}

	// Additional test case with input "1N20"
	input5 := "1N20"
	expected5 := Book{"1", "N", 20}
	result5 := ParseInput(input5)
	if !reflect.DeepEqual(expected5, result5) {
		t.Errorf("Expected %+v, but got %+v", expected5, result5)
	}

	// Additional test case with input "3N20"
	input6 := "3N20"
	expected6 := Book{"3", "N", 20}
	result6 := ParseInput(input6)
	if !reflect.DeepEqual(expected6, result6) {
		t.Errorf("Expected %+v, but got %+v", expected6, result6)
	}
}

func TestBooksParser(t *testing.T) {
	input := "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	expected := []Book{
		{"3", "A", 13},
		{"5", "X", 19},
		{"9", "Y", 20},
		{"2", "C", 18},
		{"1", "N", 20},
		{"3", "N", 20},
		{"1", "M", 21},
		{"1", "F", 14},
		{"9", "A", 21},
		{"3", "N", 21},
		{"0", "E", 13},
		{"5", "G", 14},
		{"8", "A", 23},
		{"9", "E", 22},
		{"3", "N", 14},
	}

	result := BooksParser(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %+v, but got %+v", expected, result)
	}
}

func TestBookSorter(t *testing.T) {
	unsorted := []Book{
		{"3", "A", 13},
		{"5", "X", 19},
		{"9", "Y", 20},
		{"2", "C", 18},
		{"1", "N", 20},
		{"3", "N", 20},
		{"1", "M", 21},
		{"1", "F", 14},
		{"9", "A", 21},
		{"3", "N", 21},
		{"0", "E", 13},
		{"5", "G", 14},
		{"8", "A", 23},
		{"9", "E", 22},
		{"3", "N", 14},
	}

	expected := []Book{
		{"0", "E", 13},
		{"9", "E", 22},
		{"9", "A", 21},
		{"9", "Y", 20},
		{"8", "A", 23},
		{"1", "M", 21},
		{"1", "N", 20},
		{"1", "F", 14},
		{"2", "C", 18},
		{"5", "X", 19},
		{"5", "G", 14},
		{"3", "N", 21},
		{"3", "N", 20},
		{"3", "N", 14},
		{"3", "A", 13},
	}

	result := BookSorter(unsorted)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %+v, but got %+v", expected, result)
	}
}

func TestDisplayBooks(t *testing.T) {
	books := []Book{
		{"0", "E", 13},
		{"9", "E", 22},
		{"9", "A", 21},
		{"9", "Y", 20},
		{"8", "A", 23},
		{"1", "M", 21},
		{"1", "N", 20},
		{"1", "F", 14},
		{"2", "C", 18},
		{"5", "X", 19},
		{"5", "G", 14},
		{"3", "N", 21},
		{"3", "N", 20},
		{"3", "A", 13},
		{"3", "N", 14},
	}

	expected := "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13 3N14 "

	result := DisplayBooks(books)
	if expected != result {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSortTheBooks(t *testing.T) {
	input := "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	expected := "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3N14 3A13 "

	result := SortTheBooks(input)
	if expected != result {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
