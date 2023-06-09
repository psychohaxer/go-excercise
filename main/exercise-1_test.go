package main

import (
	"reflect" // Package reflect implements run-time reflection, allowing program inspection and manipulation of objects with arbitrary types.
	"testing" // Package testing provides support for automated testing of Go packages with the testing framework.
)

func TestParseInput(t *testing.T) {
	// Test case 1
	input1 := "3A13"
	expected1 := Book{"3", "A", 13}
	result1 := parseInput(input1)
	if !reflect.DeepEqual(expected1, result1) {
		t.Errorf("Expected %+v, but got %+v", expected1, result1)
	}

	// Test case 2
	input2 := "5X19"
	expected2 := Book{"5", "X", 19}
	result2 := parseInput(input2)
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("Expected %+v, but got %+v", expected2, result2)
	}

	// Test case 3
	input3 := "9Y20"
	expected3 := Book{"9", "Y", 20}
	result3 := parseInput(input3)
	if !reflect.DeepEqual(expected3, result3) {
		t.Errorf("Expected %+v, but got %+v", expected3, result3)
	}

	// Additional test case with input "2C18"
	input4 := "2C18"
	expected4 := Book{"2", "C", 18}
	result4 := parseInput(input4)
	if !reflect.DeepEqual(expected4, result4) {
		t.Errorf("Expected %+v, but got %+v", expected4, result4)
	}

	// Additional test case with input "1N20"
	input5 := "1N20"
	expected5 := Book{"1", "N", 20}
	result5 := parseInput(input5)
	if !reflect.DeepEqual(expected5, result5) {
		t.Errorf("Expected %+v, but got %+v", expected5, result5)
	}

	// Additional test case with input "3N20"
	input6 := "3N20"
	expected6 := Book{"3", "N", 20}
	result6 := parseInput(input6)
	if !reflect.DeepEqual(expected6, result6) {
		t.Errorf("Expected %+v, but got %+v", expected6, result6)
	}
}

func TestMainFunction(t *testing.T) {
	// Test case with the given input "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	input := "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	expected := "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13"
	
	// Redirect the output to a buffer
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the main function
	main()

	// Close the write end of the pipe and restore the output
	w.Close()
	os.Stdout = oldOutput

	// Read the output from the read end of the pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Convert the output buffer to a string
	output := strings.TrimSpace(buf.String())

	// Compare the expected and actual output
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}
}
