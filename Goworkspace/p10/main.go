package main

import "fmt"

func encodeDirect(input []interface{}) []struct {
	Count   int
	Element interface{}
} {
	if len(input) == 0 {
		return nil
	}

	encoded := make([]struct {
		Count   int
		Element interface{}
	}, 0)

	count := 1
	currentElement := input[0]

	for i := 1; i < len(input); i++ {
		if input[i] == currentElement {
			count++
		} else {
			encoded = append(encoded, struct {
				Count   int
				Element interface{}
			}{
				Count:   count,
				Element: currentElement,
			})
			count = 1
			currentElement = input[i]
		}
	}

	encoded = append(encoded, struct {
		Count   int
		Element interface{}
	}{
		Count:   count,
		Element: currentElement,
	})

	return encoded
}

func main() {
	input := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}
	result := encodeDirect(input)
	fmt.Println(result)

	testEncodeDirect()
}

func testEncodeDirect() {
	input := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}
	expected := []struct {
		Count   int
		Element interface{}
	}{
		{4, 'a'},
		{1, 'b'},
		{2, 'c'},
		{2, 'a'},
		{1, 'd'},
		{4, 'e'},
	}

	result := encodeDirect(input)

	for i := range result {
		if result[i] != expected[i] {
			fmt.Printf("Test failed at index %d: Expected %v, but got %v\n", i, expected[i], result[i])
			return
		}
	}

	fmt.Println("Test passed!")
}
