package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateList(t *testing.T) {
	CreateList()

	expectedData := make(map[int]string)
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			expectedData[i] = "Type 3"
		case i%3 == 0 && i%5 != 0:
			expectedData[i] = "Type 1"
		case i%3 != 0 && i%5 == 0:
			expectedData[i] = "Type 2"
		default:
			expectedData[i] = fmt.Sprintf("%d", i)
		}
	}

	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("El mapa 'data' no se llenó correctamente. Se esperaba %v pero se obtuvo %v", expectedData, data)
	}
}

func TestGetData(t *testing.T) {
	expectedData := map[int]string{
		1: "Type 1",
		2: "2",
		3: "Type 1",
	}

	data = expectedData

	result := GetData()

	if !reflect.DeepEqual(result, expectedData) {
		t.Errorf("La función GetData() no devolvió el mapa esperado. Se esperaba %v pero se obtuvo %v", expectedData, result)
	}
}

func TestGetCollection(t *testing.T) {
	expectedData := []int{1, 2, 3}

	collection = expectedData

	result := GetCollection()

	if !reflect.DeepEqual(result, expectedData) {
		t.Errorf("La función GetData() no devolvió el mapa esperado. Se esperaba %v pero se obtuvo %v", expectedData, result)
	}
}

func TestSave(t *testing.T) {
	initialCollection := []int{1, 2, 3}
	collection = initialCollection
	result := Save(4)
	expectedResult := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Se esperaba %v pero se obtuvo %v", expectedResult, result)
	}

	initialCollection = []int{1, 2, 3}
	collection = initialCollection
	result = Save(2)
	expectedResult = []int{1, 2, 3}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Se esperaba %v pero se obtuvo %v", expectedResult, result)
	}
}

func TestCheckInMap(t *testing.T) {
	data = map[int]string{1: "Type 1", 2: "2", 3: "Type 3"}
	result := CheckInMap(3)
	expectedResult := "Type 3"
	if result != expectedResult {
		t.Errorf("Se esperaba %s pero se obtuvo %s \n", expectedResult, result)
	}

	data = map[int]string{1: "Type 1", 2: "2", 3: "Type 3"}
	result = CheckInMap(4)

	if result != "" {
		t.Errorf("Se esperaba %s pero se obtuvo %s \n", expectedResult, result)
	}
}
