package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestBodyUnmarshal(t *testing.T) {
	jsonData := []byte(`{"number":42}`)

	var requestBody RequestBody

	err := json.Unmarshal(jsonData, &requestBody)
	if err != nil {
		t.Fatalf("Error al deserializar JSON: %s", err)
	}

	expectedNumber := 42
	if requestBody.Number != expectedNumber {
		t.Errorf("Se esperaba que Number fuera %d pero es %d", expectedNumber, requestBody.Number)
	}
}

func TestResponseModel(t *testing.T) {
	response := Response{
		Message: "mensaje",
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Error al serializar el modelo a JSON: %v", err)
	}

	var decodedResponse Response
	err = json.Unmarshal(jsonBytes, &decodedResponse)
	if err != nil {
		t.Fatalf("Error al deserializar el JSON al modelo Response: %v", err)
	}

	assert.Equal(t, response.Message, decodedResponse.Message)
}
