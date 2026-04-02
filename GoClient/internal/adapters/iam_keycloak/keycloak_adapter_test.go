package iam_keycloak

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestValidateToken_Success(t *testing.T) {
	// 1. Configure the "Mock Keycloak"
	// This server will intercept the adapter's request and return whatever we tell it to.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the adapter sent the token correctly in the Header
		if r.Header.Get("Authorization") != "Bearer token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"sub": "123e4567-e89b-12d3-a456-426614174000", "email": "test@weiclothe.com"}`)
	}))
	defer mockServer.Close()

	// 2. Instantiate the adapter, and pass it the mock server URL
	adapter := NewKeycloakAdapter(mockServer.URL, "weiclothe", "dummy-client", "dummy-secret")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 3. Execute the function 
	uid, err := adapter.ValidateToken(ctx, "token")

	// 4. Assertions 
	if err != nil {
		t.Fatalf("Expected success, but got error: %v", err)
	}

	expectedUID := "123e4567-e89b-12d3-a456-426614174000"
	if uid != expectedUID {
		t.Errorf("Extracted UID was incorrect. Expected %s, got %s", expectedUID, uid)
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	// 1. Configure the Mock Server 
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate Keycloak rejecting the token 
		// Return an HTTP 401 Unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer mockServer.Close()

	// 2. Instantiate the adapter using the mock server URL
	adapter := NewKeycloakAdapter(mockServer.URL, "weiclothe", "dummy-client", "dummy-secret")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 3. Execute the function with a invalid token
	uid, err := adapter.ValidateToken(ctx, "invalid_token")

	// 4. Assertions 
	if err == nil {
		t.Fatalf("Expected an error for an invalid token, but the function returned success")
	}

	if uid != "" {
		t.Errorf("Expected an empty UID for invalid token, but got: %s", uid)
	}
}