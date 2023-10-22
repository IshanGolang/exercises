package main

import (
	"fmt"
)

func encodeModified(input []interface{}) []interface{} {
	if len(input) == 0 {
		return []interface{}{}
	}

	encoded := make([]interface{}, 0)
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count == 1 {
				encoded = append(encoded, input[i-1])
			} else {
				encoded = append(encoded, struct {
					Count   int
					Element interface{}
				}{
					Count:   count,
					Element: input[i-1],
				})
			}
			count = 1
		}
	}

	if count == 1 {
		encoded = append(encoded, input[len(input)-1])
	} else {
		encoded = append(encoded, struct {
			Count   int
			Element interface{}
		}{
			Count:   count,
			Element: input[len(input)-1],
		})
	}

	return encoded
}

func main() {
	input := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}
	result := encodeModified(input)
	fmt.Println(result)

	// Test the function
	testEncodeModified()
}

func testEncodeModified() {
	input := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}
	expected := []interface{}{
		struct {
			Count   int
			Element interface{}
		}{
			Count:   4,
			Element: 'a',
		},
		'b',
		struct {
			Count   int
			Element interface{}
		}{
			Count:   2,
			Element: 'c',
		},
		struct {
			Count   int
			Element interface{}
		}{
			Count:   2,
			Element: 'a',
		},
		'd',
		struct {
			Count   int
			Element interface{}
		}{
			Count:   4,
			Element: 'e',
		},
	}

	result := encodeModified(input)

	for i := range result {
		if result[i] != expected[i] {
			fmt.Printf("Test failed at index %d: Expected %v, but got %v\n", i, expected[i], result[i])
			return
		}
	}

	fmt.Println("Test passed!")
}
