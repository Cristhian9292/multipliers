package internal

import (
	"fmt"
)

var data map[int]string
var collection []int

func CreateList() {
	data = make(map[int]string)
	for i := 1; i <= 100; i++ {
		output := fmt.Sprintf("%d", i)

		conditions := map[bool]string{
			i%3 == 0 && i%5 == 0: "Type 3",
			i%3 == 0 && i%5 != 0: "Type 1",
			i%3 != 0 && i%5 == 0: "Type 2",
		}

		for condition, result := range conditions {
			if condition {
				output = result
				break
			}
		}

		data[i] = output
	}
}

func GetData() map[int]string {
	return data
}

func GetCollection() []int {
	return collection
}

func Save(number int) []int {
	found := false
	for _, elem := range collection {
		if elem == number {
			found = true
			break
		}
	}

	if !found {
		fmt.Println(collection)
		collection = append(collection, number)
		return collection
	} else {
		fmt.Println(collection)
		return collection
	}
}

func CheckInMap(number int) string {
	if value, ok := data[number]; ok {
		fmt.Printf("El valor asociado a la clave %d es: %s\n", number, value)
		return value
	} else {
		fmt.Printf("No se encontró un valor para la clave %d\n", number)
		return ""
	}
}
