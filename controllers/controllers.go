package controllers

import (
	"net/http"
	"strconv"

	"trafilea/internal"
	"trafilea/models"

	"github.com/gin-gonic/gin"
)

// Healtcheck godoc
// @Summary Checks if the service is up
// @Description Checks if the service is up and returns the string "UP"
// @Produce application/json
// @Tags Healtcheck
// @Success 200 {string} string "UP"
// @Router /healthcheck [get]
func HealthCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "UP")
}

// Save number godoc
// @Summary Saves a number into a collection
// @Description Receives a number and saves it into a collection
// @Consume json
// @Param number body models.RequestBody true "A number between 1-100" example: {"number": 1}
// @Produce json
// @Tags Save
// @Success 200 {object} models.Response"
// @Router /multipliers [post]
func Save(context *gin.Context) {
	var reqBody models.RequestBody

	if err := context.BindJSON(&reqBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
		return
	}

	if reqBody.Number < 1 || reqBody.Number > 100 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "El número está fuera de rango (1-100)"})
		return
	}

	internal.Save(reqBody.Number)

	context.JSON(http.StatusOK, gin.H{"message": "Saved!"})
}

// Get a value godoc
// @Summary Retrieves a value from a specific number
// @Description Returns the value from a specific number between 1 and 100 if the number is multiplier from 3, 5 or both returns Type 1, Type 2 or Type 3 respectively
// @Param number query int true "Un numero entre 1 y 100"
// @Produce json
// @Tags Get a value
// @Success 200 {object} models.ResponseGet
// @Router /multipliers [get]
func Get(context *gin.Context) {
	numberStr := context.Query("number")
	if numberStr == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'number' es requerido"})
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "El valor del parámetro 'number' no es un entero válido"})
		return
	}

	if number < 1 || number > 100 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "El número está fuera de rango (1-100)"})
		return
	}

	response := internal.CheckInMap(number)

	context.JSON(http.StatusOK, gin.H{"message": response})
}

// Get collection value godoc
// @Summary Retrieves the values from a collection
// @Description Returns the values from a collection, if its empty returns null
// @Produce json
// @Tags Get collection
// @Success 200 {object} models.ResponseCollection
// @Router /multipliers/collection [get]
func GetAll(context *gin.Context) {
	data := internal.GetCollection()
	context.JSON(http.StatusOK, gin.H{"message": data})
}
