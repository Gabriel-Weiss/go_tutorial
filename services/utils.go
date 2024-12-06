package services

import (
	"fmt"

	"github.com/Gabriel-Weiss/go_tutorial/models"
)

func GetItemsByCategory(category string) []models.Item {
	var itemsByCategory []models.Item
	if category == "" {
		return models.Data
	}

	for _, item := range models.Data {
		if item.Category == category {
			itemsByCategory = append(itemsByCategory, item)
		}
	}

	if len(itemsByCategory) == 0 {
		return models.Data
	}

	return itemsByCategory
}

func GetItemByTitle(title string) (models.Item, error) {
	for _, item := range models.Data {
		if item.Title == title {
			return item, nil
		}
	}

	return models.Item{}, fmt.Errorf("item with title %s not found", title)
}

func GetItemById(id int) (models.Item, error) {
	for _, item := range models.Data {
		if item.Id == id {
			return item, nil
		}
	}

	return models.Item{}, fmt.Errorf("item with id %d not found", id)
}

func GetProductsInCart() ([]models.Item, error) {
	var productsInCart []models.Item

	for _, item := range models.Products {
		product, err := GetItemByTitle(item.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to find product with ID %s: %w", item.ID, err)
		}
		productsInCart = append(productsInCart, product)
	}

	return productsInCart, nil
}
