package main

import (
	"github.com/coolpythoncodes/nigerian-universities/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	// scraper.ScrapeUniversities()
	r := gin.Default()
	r.GET("/", controllers.GetAllUniversities)

	r.Run(":8080")

}
