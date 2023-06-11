package main

import (
	"fmt"
	"math"
	"strconv"
)

// StringToArray is used to convert a string to an array of strings.
func StringToArray(str string) []string {
	var arr []string
	for i := 0; i < len(str); i++ {
		arr = append(arr, string(str[i]))
	}

	return arr
}

// StringArrayToIntegerArray is used to convert an array of strings to an array of integers.
func StringArrayToIntegerArray(str []string) []int {
	var arr []int
	for i := 0; i < len(str); i++ {
		// Convert element at index i to an integer.
		temp, _ := strconv.Atoi(str[i])
		arr = append(arr, temp)
	}

	return arr
}

// ElementExtractor is used to extract elements from an array.
func ElementExtractor(arr []int, start int, end int) []int {
	return arr[start:end]
}

// ExtractElementsByDigitAndIndex is used to extract elements from an array by digit and index.
func ExtractElementsByDigitAndIndex(a []int, digit int, index int) int {
	var result []int
	var temp string
	counter := 0

	for i := 0; i < len(a); i++ {
		str := strconv.Itoa(a[i])
		temp += str

		if len(temp) == digit {
			if counter == index {
				num, _ := strconv.Atoi(temp)
				result = append(result, num)
			}
			counter++
			temp = ""
		}
	}

	return result[0]
}

func ElementCombiner(arr []int) int {

	if len(arr) > 1 {
		var str string

		// Iterate over temp to combine the elements.
		for i := 0; i < len(arr); i++ {
			// Append each element to str.
			str += strconv.Itoa(arr[i])
		}

		// Convert str to an integer.
		result, _ := strconv.Atoi(str)
		return result
	} else {
		return arr[0]
	}
}

// IsPreviousNumber is used to check if the current number is equal to the previous number minus 1.
func IsPreviousNumber(current int, previous int) bool {
	if current == previous+1 {
		return true
	} else {
		return false
	}
}

// NumberSeqDetector is used to detect the number sequence.
func NumberSeqDetector(arr []int) []int {
	var result, temp []int
	var current, previous int

	// Iterate over half of arr to use as digit.
	for digit := 1; digit < int(math.Ceil(float64(len(arr))/2)); digit++ {
		fmt.Println("\niterasi digit: ", digit)

		err := 0
		previous = ExtractElementsByDigitAndIndex(arr, digit, 0)
		temp = append(temp, previous)

		fmt.Println("temp: ", temp)

		// Iterate over arr to check the array elements.
		for index := 1; index < len(arr); index++ {
			fmt.Println("    iterasi index: ", index, "--------------")
			fmt.Println("    curr: ", current)
			fmt.Println("    prev: ", previous)
			fmt.Println("    temp: ", temp)

			if index+digit+digit >= len(arr) {
				break
			}

			// Extract the current element.
			// current = ElementCombiner(ElementExtractor(arr, index+digit, index+digit+digit))
			fmt.Println("\n    fetch: current")
			fmt.Println("    ExtractElementsByDigitAndIndex(arr,", digit, ",", index, ")")
			current = ExtractElementsByDigitAndIndex(arr, digit, index)
			fmt.Println("    curr: ", current)

			// Check if current + 1 = 9 or 99 or 999 etc.
			if current == int(math.Pow10(digit))-1 {
				// Append the current element to temp.
				temp = append(temp, current)

				// Increment digit.
				digit++

				// Set previous to current.
				previous = current

				// Continue to the next iteration.
				continue
			}

			fmt.Println("\n    isPrevious(", current, ",", previous, ")")
			// Check if the current element is equal to the previous element minus 1.

			if IsPreviousNumber(current, previous) {
				fmt.Println("    isPrevious: true")
				// Append the current element to temp.
				temp = append(temp, current)
			} else {
				fmt.Println("    isPrevious: false")
				err++
				fmt.Println("    err: ", err)
			}

			// Check if error > 2
			if err > 2 {
				// clear temp
				temp = nil
				break
			}

			// if all elements in arr have been iterated
			if index == len(arr) {
				fmt.Println("    status: break success")
				// break
			}
			// Set previous to current.
			previous = current
		}

		result = temp
	}

	return result
}

// func main() {
// 	in1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	// in2 := []int{2, 1, 2, 2, 2, 3, 2, 4, 2, 5}
// 	// in3 := []int{9, 8, 9, 9, 1, 0, 0, 1, 0, 1}

// 	fmt.Println("in1: ", in1)
// 	fmt.Println(NumberSeqDetector(in1))

// 	// fmt.Println("in2: ", in2)
// 	// fmt.Println(NumberSeqDetector(in2))

// 	// fmt.Println("in3: ", in3)
// 	// fmt.Println(NumberSeqDetector(in3))
// }

// MissingNumberDetector is used to detect the missing number.
// func MissingNumberDetector(arr []int) int {

// 	var current, previous int
// 	missing := 0

// 	// Iterate over half of arr to use as digit.
// 	for digit := 1; digit < int(math.Ceil(float64(len(arr))/2)); digit++ {

// 		err := 0
// 		previous = ExtractElementsByDigitAndIndex(arr, digit, 0)
// 		// temp = append(temp, previous)

// 		// Iterate over arr to check the array elements.
// 		for index := 0; index < len(arr); index += digit {

// 			if index+digit+digit > len(arr) {
// 				break
// 			}

// 			// Extract the current element.
// 			// current = ElementCombiner(ElementExtractor(arr, index+digit, index+digit+digit)
// 			fmt.Println("fetch: current")
// 			fmt.Println("ExtractElementsByDigitAndIndex(arr, ", digit, ",", index, ")")
// 			current = ExtractElementsByDigitAndIndex(arr, digit, index)
// 			fmt.Println("current: ", current)

// 			// Check if current + 1 = 9 or 99 or 999 etc.
// 			if current == int(math.Pow10(digit))-1 {
// 				// Append the current element to temp.
// 				// temp = append(temp, current)

// 				// Increment digit.
// 				digit++

// 				// Set previous to current.
// 				previous = current

// 				// Continue to the next iteration.
// 				continue
// 			}

// 			// Check if the current element is equal to the previous element minus 1.
// 			if IsPreviousNumber(current, previous) {
// 				// Append the current element to temp.
// 				// temp = append(temp, current)
// 				continue
// 			} else if IsPreviousNumber(current, previous+1) {
// 				missing = previous + 1
// 			} else {
// 				err++
// 			// }

// 			// Check if error > 2
// 			if err > 2 {
// 				break
// 			}

// 			// Set previous to current.
// 			previous = current
// 		}
// 	}
// 	return missing
// }
