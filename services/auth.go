package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func SignIn(username string, password string) (string, error) {
	var tokenResponse struct {
		Token string `json:"token"`
	}
	if username == "" || password == "" {
		return "", fmt.Errorf("Invalid username or password")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := map[string]string{
		"username": username,
		"password": password,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Failed to create payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BaseUrl+"/auth/login", strings.NewReader(string(jsonPayload)))
	if err != nil {
		return "", fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to sign in: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to sign in: received status code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("Failed to parse token: %w", err)
	}

	return tokenResponse.Token, nil
}

func SignUp(username string, password string, email string) error {
	if username == "" || password == "" || email == "" {
		return fmt.Errorf("Username, password and email are required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Failed to create payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BaseUrl+"/users", strings.NewReader(string(jsonPayload)))
	if err != nil {
		return fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to sign up: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to sign up: received status code %d", resp.StatusCode)
	}

	return nil
}
