package utils

import (
	"strings"

	"github.com/coolpythoncodes/nigerian-universities/models"
)

func Filter(universities []models.Universities, univerisityType string) []models.Universities {
	u := make([]models.Universities, 0)

	for _, uni := range universities {
		if strings.ToLower(uni.Type) == univerisityType {
			u = append(u, uni)
		}
	}

	return u
}
