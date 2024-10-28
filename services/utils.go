package services

import "github.com/Gabriel-Weiss/go_tutorial/models"

func GetItemsByCategory(category string) []models.Item {
	var filteredItems []models.Item
	if category == "" {
		return models.Data
	}

	for _, item := range models.Data {
		if item.Category == category {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
