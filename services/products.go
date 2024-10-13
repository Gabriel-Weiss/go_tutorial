package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gabriel-Weiss/go_tutorial/models"
)

const BaseUrl = "https://fakestoreapi.com"

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseUrl+"/products", nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch products: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("Failed to parse products: %w", err)
	}

	return products, nil
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseUrl+"/products/"+id, nil)
	if err != nil {
		return product, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return product, fmt.Errorf("Failed to fetch product: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return product, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return product, fmt.Errorf("Failed to parse product: %w", err)
	}

	return product, nil
}

func GetProductsInCategory(category string) ([]models.Product, error) {
	var products []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseUrl+"/products/category/"+category, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch products by category: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return products, fmt.Errorf("Failed to parse products by category: %w", err)
	}

	return products, nil
}

func CreateProduct(product models.Product) error { return nil }

func UpdateProduct(product models.Product) error { return nil }

func DeleteProduct(id string) error { return nil }
