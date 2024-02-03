package controllers

import (
	"net/http"

	"github.com/coolpythoncodes/nigerian-universities/utils"
	"github.com/gin-gonic/gin"
)

func GetAllUniversities(c *gin.Context) {
	univeristies, err := utils.ReadUniversitiesFromJSONFile(utils.JsonFileName)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "success",
		"data":    univeristies,
	})
}

func GetAllFederalUniversities(c *gin.Context) {
	univeristies, err := utils.ReadUniversitiesFromJSONFile(utils.JsonFileName)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	federalUniversities := utils.Filter(univeristies, "federal")

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "success",
		"data":    federalUniversities,
	})
}

func GetAllStateUniversities(c *gin.Context) {
	univeristies, err := utils.ReadUniversitiesFromJSONFile(utils.JsonFileName)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	stateUniversities := utils.Filter(univeristies, "state")

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "success",
		"data":    stateUniversities,
	})
}

func GetAllPrivateUniversities(c *gin.Context) {
	univeristies, err := utils.ReadUniversitiesFromJSONFile(utils.JsonFileName)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	privateUniversities := utils.Filter(univeristies, "private")

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "success",
		"data":    privateUniversities,
	})
}
