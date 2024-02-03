package main

import (
	"github.com/coolpythoncodes/nigerian-universities/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	// scraper.ScrapeUniversities()
	r := gin.Default()
	r.GET("/", controllers.GetAllUniversities)
	r.GET("/federal", controllers.GetAllFederalUniversities)
	r.GET("/state", controllers.GetAllStateUniversities)
	r.GET("/private", controllers.GetAllPrivateUniversities)

	r.Run(":8080")

}
