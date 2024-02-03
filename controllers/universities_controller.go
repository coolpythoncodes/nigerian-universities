package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/coolpythoncodes/nigerian-universities/models"
	"github.com/coolpythoncodes/nigerian-universities/utils"
	"github.com/gin-gonic/gin"
)

func GetAllUniversities(c *gin.Context) {
	univeristies := []models.Universities{}
	universitiesJson, err := os.ReadFile(utils.JsonFileName)

	if err != nil {
		restErr := utils.NewInternalServerError("unable to get data")
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}
	if err := json.Unmarshal(universitiesJson, &univeristies); err != nil {
		restErr := utils.NewInternalServerError("unable to get data unmarshale")
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "success",
		"data":    univeristies,
	})
}
