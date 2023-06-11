package main

import (
	"reflect"
	"testing"
)

func TestStringToArray(t *testing.T) {
	// Test case 1
	input1 := "23242526272830"
	expected1 := []string{"2", "3", "2", "4", "2", "5", "2", "6", "2", "7", "2", "8", "3", "0"}
	result1 := StringToArray(input1)
	if !reflect.DeepEqual(expected1, result1) {
		t.Errorf("Expected %+v, but got %+v", expected1, result1)
	}

	// Test case 2
	input2 := "101102103104105106107108109111112113"
	expected2 := []string{"1", "0", "1", "1", "0", "2", "1", "0", "3", "1", "0", "4", "1", "0", "5", "1", "0", "6", "1", "0", "7", "1", "0", "8", "1", "0", "9", "1", "1", "1", "1", "1", "2", "1", "1", "3"}
	result2 := StringToArray(input2)
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("Expected %+v, but got %+v", expected2, result2)
	}

	// Test case 3
	input3 := "12346789"
	expected3 := []string{"1", "2", "3", "4", "6", "7", "8", "9"}
	result3 := StringToArray(input3)
	if !reflect.DeepEqual(expected3, result3) {
		t.Errorf("Expected %+v, but got %+v", expected3, result3)
	}

	// Test case 4
	input4 := "79101112"
	expected4 := []string{"7", "9", "1", "0", "1", "1", "1", "2"}
	result4 := StringToArray(input4)
	if !reflect.DeepEqual(expected4, result4) {
		t.Errorf("Expected %+v, but got %+v", expected4, result4)
	}

	// Test case 5
	input5 := "7891012"
	expected5 := []string{"7", "8", "9", "1", "0", "1", "2"}
	result5 := StringToArray(input5)
	if !reflect.DeepEqual(expected5, result5) {
		t.Errorf("Expected %+v, but got %+v", expected5, result5)
	}

	// Test case 6
	input6 := "9799100101102"
	expected6 := []string{"9", "7", "9", "9", "1", "0", "0", "1", "0", "1", "1", "0", "2"}
	result6 := StringToArray(input6)
	if !reflect.DeepEqual(expected6, result6) {
		t.Errorf("Expected %+v, but got %+v", expected6, result6)
	}

	// Test case 7
	input7 := "100001100002100003100004100006"
	expected7 := []string{"1", "0", "0", "0", "0", "1", "1", "0", "0", "0", "0", "2", "1", "0", "0", "0", "0", "3", "1", "0", "0", "0", "0", "4", "1", "0", "0", "0", "0", "6"}
	result7 := StringToArray(input7)
	if !reflect.DeepEqual(expected7, result7) {
		t.Errorf("Expected %+v, but got %+v", expected7, result7)
	}
}

func TestStringArrayToIntegerArray(t *testing.T) {
	// Test case 1
	input1 := []string{"2", "3", "2", "4", "2", "5", "2", "6", "2", "7", "2", "8", "3", "0"}
	expected1 := []int{2, 3, 2, 4, 2, 5, 2, 6, 2, 7, 2, 8, 3, 0}
	result1 := StringArrayToIntegerArray(input1)
	if !reflect.DeepEqual(expected1, result1) {
		t.Errorf("Expected %+v, but got %+v", expected1, result1)
	}

	// Test case 2
	input2 := []string{"1", "0", "1", "1", "0", "2", "1", "0", "3", "1", "0", "4", "1", "0", "5", "1", "0", "6", "1", "0", "7", "1", "0", "8", "1", "0", "9", "1", "1", "1", "1", "1", "2", "1", "1", "3"}
	expected2 := []int{1, 0, 1, 1, 0, 2, 1, 0, 3, 1, 0, 4, 1, 0, 5, 1, 0, 6, 1, 0, 7, 1, 0, 8, 1, 0, 9, 1, 1, 1, 1, 1, 2, 1, 1, 3}
	result2 := StringArrayToIntegerArray(input2)
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("Expected %+v, but got %+v", expected2, result2)
	}

	// Test case 3
	input3 := []string{"1", "2", "3", "4", "6", "7", "8", "9"}
	expected3 := []int{1, 2, 3, 4, 6, 7, 8, 9}
	result3 := StringArrayToIntegerArray(input3)
	if !reflect.DeepEqual(expected3, result3) {
		t.Errorf("Expected %+v, but got %+v", expected3, result3)
	}

	// Test case 4
	input4 := []string{"7", "9", "1", "0", "1", "1", "1", "2"}
	expected4 := []int{7, 9, 1, 0, 1, 1, 1, 2}
	result4 := StringArrayToIntegerArray(input4)
	if !reflect.DeepEqual(expected4, result4) {
		t.Errorf("Expected %+v, but got %+v", expected4, result4)
	}

	// Test case 5
	input5 := []string{"7", "8", "9", "1", "0", "1", "2"}
	expected5 := []int{7, 8, 9, 1, 0, 1, 2}
	result5 := StringArrayToIntegerArray(input5)
	if !reflect.DeepEqual(expected5, result5) {
		t.Errorf("Expected %+v, but got %+v", expected5, result5)
	}

	// Test case 6
	input6 := []string{"9", "7", "9", "9", "1", "0", "0", "1", "0", "1", "1", "0", "2"}
	expected6 := []int{9, 7, 9, 9, 1, 0, 0, 1, 0, 1, 1, 0, 2}
	result6 := StringArrayToIntegerArray(input6)
	if !reflect.DeepEqual(expected6, result6) {
		t.Errorf("Expected %+v, but got %+v", expected6, result6)
	}

	// Test case 7
	input7 := []string{"1", "0", "0", "0", "0", "1", "1", "0", "0", "0", "0", "2", "1", "0", "0", "0", "0", "3", "1", "0", "0", "0", "0", "4", "1", "0", "0", "0", "0", "6"}
	expected7 := []int{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6}
	result7 := StringArrayToIntegerArray(input7)
	if !reflect.DeepEqual(expected7, result7) {
		t.Errorf("Expected %+v, but got %+v", expected7, result7)
	}
}

func TestElementExtractor(t *testing.T) {
	// Test case 1
	arr1 := []int{1, 2, 3, 4, 5}
	start1 := 0
	end1 := 3
	expected1 := []int{1, 2, 3}
	result1 := ElementExtractor(arr1, start1, end1)
	if !reflect.DeepEqual(expected1, result1) {
		t.Errorf("Expected %+v, but got %+v", expected1, result1)
	}

	// Test case 2
	arr2 := []int{10, 20, 30, 40, 50, 60, 70}
	start2 := 2
	end2 := 5
	expected2 := []int{30, 40, 50}
	result2 := ElementExtractor(arr2, start2, end2)
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("Expected %+v, but got %+v", expected2, result2)
	}

	// Test case 3
	arr3 := []int{100, 200, 300}
	start3 := 1
	end3 := 2
	expected3 := []int{200}
	result3 := ElementExtractor(arr3, start3, end3)
	if !reflect.DeepEqual(expected3, result3) {
		t.Errorf("Expected %+v, but got %+v", expected3, result3)
	}

	// Test case 4
	arr4 := []int{5, 10, 15, 20, 25, 30, 35, 40}
	start4 := 3
	end4 := 7
	expected4 := []int{20, 25, 30, 35}
	result4 := ElementExtractor(arr4, start4, end4)
	if !reflect.DeepEqual(expected4, result4) {
		t.Errorf("Expected %+v, but got %+v", expected4, result4)
	}

	// Test case 5
	arr5 := []int{2, 4, 6, 8, 10}
	start5 := 0
	end5 := 0
	expected5 := []int{}
	result5 := ElementExtractor(arr5, start5, end5)
	if !reflect.DeepEqual(expected5, result5) {
		t.Errorf("Expected %+v, but got %+v", expected5, result5)
	}
}

func TestElementCombiner(t *testing.T) {
	// Test case 1
	arr1 := []int{1, 2, 3, 4, 5}
	expected1 := 12345
	result1 := ElementCombiner(arr1)
	if !reflect.DeepEqual(expected1, result1) {
		t.Errorf("Expected %+v, but got %+v", expected1, result1)
	}

	// Test case 2
	arr2 := []int{2, 4}
	expected2 := 24
	result2 := ElementCombiner(arr2)
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("Expected %+v, but got %+v", expected2, result2)
	}

	// Test case 3
	arr3 := []int{1, 0, 2, 1, 0, 3, 1, 0, 4, 1, 0, 5, 1, 0, 6}
	expected3 := 102103104105106
	result3 := ElementCombiner(arr3)
	if !reflect.DeepEqual(expected3, result3) {
		t.Errorf("Expected %+v, but got %+v", expected3, result3)
	}
}

// func TestNumberSeqDetector(t *testing.T) {
// 	// Test case 1
// 	input1 := []int{2, 3, 2, 4, 2, 6, 2, 7, 2, 8}
// 	expected1 := 25
// 	result1 := NumberSeqDetector(input1)
// 	if !reflect.DeepEqual(expected1, result1) {
// 		t.Errorf("Expected %+v, but got %+v", expected1, result1)
// 	}
// }

func TestExtractElementsByDigitAndIndex(t *testing.T) {
	tests := []struct {
		name   string
		a      []int
		digit  int
		index  int
		result int
	}{
		{
			name:   "One Digit - Index 0",
			a:      []int{1, 2, 3, 4, 5, 6},
			digit:  1,
			index:  0,
			result: 1,
		},
		{
			name:   "One Digit - Index 1",
			a:      []int{1, 2, 3, 4, 5, 6},
			digit:  1,
			index:  1,
			result: 2,
		},
		{
			name:   "Two Digits - Index 0",
			a:      []int{1, 2, 3, 4, 5, 6},
			digit:  2,
			index:  0,
			result: 12,
		},
		{
			name:   "Two Digits - Index 1",
			a:      []int{1, 2, 3, 4, 5, 6},
			digit:  2,
			index:  1,
			result: 34,
		},
		{
			name:   "Three Digits - Index 0",
			a:      []int{1, 2, 3, 4, 5, 6},
			digit:  3,
			index:  0,
			result: 123,
		},
		{
			name:   "Three Digits - Index 1",
			a:      []int{1, 2, 3, 4, 5, 6},
			digit:  3,
			index:  1,
			result: 456,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ExtractElementsByDigitAndIndex(test.a, test.digit, test.index)
			if !reflect.DeepEqual(result, test.result) {
				t.Errorf("Expected %v, but got %v", test.result, result)
			}
		})
	}
}

func TestIsPreviousNumber(t *testing.T) {
	tests := []struct {
		name     string
		current  int
		previous int
		expected bool
	}{
		{
			name:     "Current is previous number",
			current:  5,
			previous: 4,
			expected: true,
		},
		{
			name:     "Current is not previous number",
			current:  5,
			previous: 3,
			expected: false,
		},
		{
			name:     "Current is not previous number (equal values)",
			current:  5,
			previous: 5,
			expected: false,
		},
		{
			name:     "Current is previous number (negative numbers)",
			current:  -3,
			previous: -4,
			expected: true,
		},
		{
			name:     "Current is previous number (zero value)",
			current:  0,
			previous: -1,
			expected: true,
		},
		{
			name:     "Current is not previous number (zero value)",
			current:  0,
			previous: 1,
			expected: false,
		},
		{
			name:     "Current is previous number (large numbers)",
			current:  1000,
			previous: 999,
			expected: true,
		},
		{
			name:     "Current is not previous number (large numbers)",
			current:  1000,
			previous: 998,
			expected: false,
		},
		{
			name:     "Current is previous number (equal negative values)",
			current:  -10,
			previous: -11,
			expected: true,
		},
		{
			name:     "Current is not previous number (equal negative values)",
			current:  -10,
			previous: -9,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPreviousNumber(test.current, test.previous)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

// func TestMissingNumberDetector(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		arr      []int
// 		expected int
// 	}{
// 		{
// 			name:     "Missing number in the middle",
// 			arr:      []int{1, 2, 3, 5, 6},
// 			expected: 4,
// 		},
// 		{
// 			name:     "Missing number with larger digit",
// 			arr:      []int{1, 2, 3, 4, 5, 7, 8, 9, 1, 0, 1, 1, 1, 2, 1, 3, 1, 4, 1, 5},
// 			expected: 6,
// 		},
// 		{
// 			name:     "Exercise Test Case",
// 			arr:      []int{2, 3, 2, 4, 2, 5, 2, 6, 2, 7, 2, 8, 3, 0},
// 			expected: 29,
// 		},
// 		{
// 			name:     "Exercise Test Case",
// 			arr:      []int{1, 0, 1, 1, 0, 2, 1, 0, 3, 1, 0, 4, 1, 0, 5, 1, 0, 6, 1, 0, 7, 1, 0, 8, 1, 0, 9, 1, 1, 1, 1, 1, 2, 1, 1, 3},
// 			expected: 110,
// 		},
// 		{
// 			name:     "Exercise Test Case",
// 			arr:      []int{1, 2, 3, 4, 6, 7, 8, 9},
// 			expected: 5,
// 		},
// 		{
// 			name:     "Exercise Test Case",
// 			arr:      []int{7, 9, 1, 0, 1, 1, 1, 2},
// 			expected: 8,
// 		},
// 		{
// 			name:     "Exercise Test Case",
// 			arr:      []int{7, 8, 9, 1, 0, 1, 2},
// 			expected: 11,
// 		},
// 		{
// 			name:     "Exercise Test Case",
// 			arr:      []int{9, 7, 9, 9, 1, 0, 0, 1, 0, 1, 1, 0, 2},
// 			expected: 98,
// 		},
// 		{
// 			name:     "Exercise Test Case",
// 			arr:      []int{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6},
// 			expected: 10005,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			result := MissingNumberDetector(test.arr)
// 			if result != test.expected {
// 				t.Errorf("Expected %v, but got %v", test.expected, result)
// 			}
// 		})
// 	}
// }

func TestNumberSeqDetector(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "Exercise Test Case",
			arr:      []int{2, 3, 2, 4, 2, 5, 2, 6, 2, 7, 2, 8, 3, 0},
			expected: []int{23, 24, 25, 26, 27, 28, 30},
		},
		{
			name:     "Exercise Test Case",
			arr:      []int{1, 0, 1, 1, 0, 2, 1, 0, 3, 1, 0, 4, 1, 0, 5, 1, 0, 6, 1, 0, 7, 1, 0, 8, 1, 0, 9, 1, 1, 1, 1, 1, 2, 1, 1, 3},
			expected: []int{101, 102, 103, 104, 105, 106, 107, 108, 109, 111, 112, 113},
		},
		{
			name:     "Exercise Test Case",
			arr:      []int{1, 2, 3, 4, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 6, 7, 8, 9},
		},
		{
			name:     "Exercise Test Case",
			arr:      []int{7, 9, 1, 0, 1, 1, 1, 2},
			expected: []int{7, 9, 10, 11, 12},
		},
		{
			name:     "Exercise Test Case",
			arr:      []int{7, 8, 9, 1, 0, 1, 2},
			expected: []int{7, 8, 9, 10, 12},
		},
		{
			name:     "Exercise Test Case",
			arr:      []int{9, 7, 9, 9, 1, 0, 0, 1, 0, 1, 1, 0, 2},
			expected: []int{97, 99, 100, 101, 102},
		},
		{
			name:     "Exercise Test Case",
			arr:      []int{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6},
			expected: []int{100001, 100002, 100003, 100004, 100006},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := NumberSeqDetector(test.arr)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}

}
