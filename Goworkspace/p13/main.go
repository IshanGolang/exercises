package main

import (
	"fmt"
)

type Encoded struct {
	Count   int
	Element interface{}
}

func decode(encodedList []Encoded) []interface{} {
	decoded := make([]interface{}, 0)

	for _, item := range encodedList {
		if item.Count == 1 {
			decoded = append(decoded, item.Element)
		} else {
			for i := 0; i < item.Count; i++ {
				decoded = append(decoded, item.Element)
			}
		}
	}

	return decoded
}

func main() {
	encodedList := []Encoded{
		{4, 'a'},
		{1, 'b'},
		{2, 'c'},
		{2, 'a'},
		{1, 'd'},
		{4, 'e'},
	}

	decodedList := decode(encodedList)
	fmt.Println(decodedList)

	testDecode()
}

func testDecode() {
	encodedList := []Encoded{
		{4, 'a'},
		{1, 'b'},
		{2, 'c'},
		{2, 'a'},
		{1, 'd'},
		{4, 'e'},
	}

	expected := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}

	decodedList := decode(encodedList)

	for i := range decodedList {
		if decodedList[i] != expected[i] {
			fmt.Printf("Test failed at index %d: Expected %v, but got %v\n", i, expected[i], decodedList[i])
			return
		}
	}

	fmt.Println("Test passed!")
}
