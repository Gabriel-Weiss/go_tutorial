package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Gabriel-Weiss/go_tutorial/models"
)

func GetAllCarts() ([]models.Cart, error) {
	var carts []models.Cart

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseUrl+"/carts", nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch carts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&carts); err != nil {
		return nil, fmt.Errorf("Failed to parse carts: %w", err)
	}

	return carts, nil
}

func GetCartById(id string) (models.Cart, error) {
	var cart models.Cart

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseUrl+"/carts/"+id, nil)
	if err != nil {
		return cart, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return cart, fmt.Errorf("Failed to fetch cart: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return cart, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&cart); err != nil {
		return cart, fmt.Errorf("Failed to parse cart: %w", err)
	}

	return cart, nil
}

func GetCartByUserId(userId string) ([]models.Cart, error) {
	var carts []models.Cart

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseUrl+"/carts/user/"+userId, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch user cart: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&carts); err != nil {
		return nil, fmt.Errorf("Failed to parse user cart: %w", err)
	}

	return carts, nil
}

func CreateCart(cart models.Cart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(cart)
	if err != nil {
		return fmt.Errorf("Failed to marshal cart data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BaseUrl+"/carts", strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to create cart: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func UpdateCart(cart models.Cart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(cart)
	if err != nil {
		return fmt.Errorf("Failed to marshal cart data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("%s/carts/%d", BaseUrl, cart.ID), strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to update cart: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func DeleteCart(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, BaseUrl+"/carts/"+id, nil)
	if err != nil {
		return fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to delete cart: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
