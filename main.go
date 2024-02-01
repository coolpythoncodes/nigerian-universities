package main

import (
	"github.com/coolpythoncodes/nigerian-universities/scraper"
	"github.com/gin-gonic/gin"
)

func main() {
	scraper.ScrapeUniversities()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")

}
