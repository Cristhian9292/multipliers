package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"trafilea/internal"
	"trafilea/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var data = internal.GetData()

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/healthcheck", func(context *gin.Context) {
		context.String(200, "UP")
	})
	return router
}

func setupRouterSave() *gin.Engine {
	router := gin.Default()
	router.POST("/multipliers", Save)
	return router
}

func setupRouterGet() *gin.Engine {
	router := gin.Default()

	// Definir la ruta GET para "/multipliers"
	router.GET("/multipliers", func(context *gin.Context) {
		// Obtener el parámetro de consulta "number"
		numberStr := context.Query("number")
		if numberStr == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'number' es requerido"})
			return
		}

		// Convertir el parámetro de consulta a un entero
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El valor del parámetro 'number' no es un entero válido"})
			return
		}

		// Validar el número
		if number < 1 || number > 100 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El número está fuera de rango (1-100)"})
			return
		}

		// Llamar a la función que maneja la lógica de negocio
		response := internal.CheckInMap(number)

		// Responder con la respuesta obtenida
		context.JSON(http.StatusOK, gin.H{"message": response})
	})

	return router
}

func setupRouterGetAll() *gin.Engine {
	router := gin.Default()

	// Ruta GET para "/multipliers/collection"
	router.GET("/multipliers/collection", func(context *gin.Context) {
		// Obtener la colección de datos
		data := internal.GetCollection()

		// Responder con la colección de datos
		context.JSON(http.StatusOK, gin.H{"message": data})
	})

	return router
}

func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "UP", w.Body.String())
}

func TestSave(t *testing.T) {
	// Configurar el enrutador
	router := setupRouterSave()

	// Crear un nuevo ResponseRecorder
	w := httptest.NewRecorder()

	// Crear una solicitud POST a "/multipliers" con un número válido en el cuerpo JSON
	reqBody := models.RequestBody{Number: 5}
	reqBodyJSON, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/multipliers", bytes.NewBuffer(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// Servir la solicitud utilizando el enrutador
	router.ServeHTTP(w, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodificar el cuerpo de la respuesta en una estructura JSON
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error al decodificar el cuerpo de la respuesta: %v", err)
	}

	// Verificar si la respuesta es la esperada
	expectedResponse := map[string]string{"message": "Saved!"}
	assert.Equal(t, expectedResponse, response)
}

func TestGet(t *testing.T) {
	// Configurar el enrutador
	router := setupRouterGet()

	// Crear un nuevo ResponseRecorder
	w := httptest.NewRecorder()

	// Crear una solicitud GET a "/multipliers" con un número válido como parámetro de consulta
	req, _ := http.NewRequest("GET", "/multipliers?number=3", nil)

	// Servir la solicitud utilizando el enrutador
	router.ServeHTTP(w, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAll(t *testing.T) {
	// Configurar el enrutador
	router := setupRouterGetAll()

	// Crear un nuevo ResponseRecorder
	w := httptest.NewRecorder()

	// Crear una solicitud GET a "/multipliers/collection"
	req, _ := http.NewRequest("GET", "/multipliers/collection", nil)

	// Servir la solicitud utilizando el enrutador
	router.ServeHTTP(w, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodificar el cuerpo de la respuesta en una estructura JSON
	var response map[string][]int
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error al decodificar el cuerpo de la respuesta: %v", err)
	}

	// Verificar si la respuesta es la esperada
	expectedResponse := map[string][]int{"message": {5}} // Ajusta el valor esperado según tu lógica de negocio
	assert.Equal(t, expectedResponse, response)
}
