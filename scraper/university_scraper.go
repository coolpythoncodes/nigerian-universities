package scraper

import (
	"fmt"

	"github.com/coolpythoncodes/nigerian-universities/models"
	"github.com/coolpythoncodes/nigerian-universities/utils"
	"github.com/gocolly/colly"
)

var universities []models.Universities

func ScrapeUniversities() {
	pagesToScrape := map[string]string{
		"https://www.nuc.edu.ng/nigerian-univerisities/federal-univeristies/1": "Federal",
		"https://www.nuc.edu.ng/nigerian-univerisities/state-univerisity/1":    "State",
		"https://www.nuc.edu.ng/nigerian-univerisities/private-univeristies/1": "Private",
	}
	c := colly.NewCollector(
		colly.AllowedDomains("www.nuc.edu.ng", "nuc.edu.ng"),
	)

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		universityType := pagesToScrape[e.Request.URL.String()]
		// fmt.Println("universityType", universityType)

		nameOfUniversity := e.ChildText(".column-2")
		nameOfChancellor := e.ChildText(".column-3")
		universityWebsite := e.ChildText(".column-4 a")
		yearOfEstablishment := e.ChildText(".column-5")

		university := models.Universities{
			Name:                nameOfUniversity,
			ViceChancellor:      nameOfChancellor,
			YearOfEstablishment: yearOfEstablishment,
			Url:                 universityWebsite,
			Type:                universityType,
		}

		universities = append(universities, university)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for url := range pagesToScrape {
		c.Visit(url)
	}
	utils.WriteJSON("nigerian-university.json", universities)
}
