package Datadatdatclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestServer creates a test server and configures an APIClient pointing to it.
func newTestServer(handler http.HandlerFunc) (*httptest.Server, *APIClient) {
	ts := httptest.NewServer(handler)
	cfg := NewConfiguration()
	cfg.BasePath = ts.URL
	client := NewAPIClient(cfg)
	return ts, client
}

// ---------------------------------------------------------------------------
// GetCommit (GET) - Success and Error
// ---------------------------------------------------------------------------

func TestGetCommit_Success(t *testing.T) {
	commit := Commit{
		Id:         "abc123",
		Properties: map[string]interface{}{"tag": "v1.0"},
	}

	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and path
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/v1/repositories/myrepo/commits/abc123") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(commit)
	})
	defer ts.Close()

	result, resp, err := client.CommitsApi.GetCommit(context.Background(), "myrepo", "abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	if result.Id != "abc123" {
		t.Errorf("expected commit id abc123, got %q", result.Id)
	}
}

func TestGetCommit_NotFound(t *testing.T) {
	apiErr := ApiError{
		Code:    "NOT_FOUND",
		Message: "Commit not found",
	}

	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(apiErr)
	})
	defer ts.Close()

	_, resp, err := client.CommitsApi.GetCommit(context.Background(), "myrepo", "nonexistent")
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}

	// Verify error is a GenericOpenAPIError and contains the model
	openAPIErr, ok := err.(GenericOpenAPIError)
	if !ok {
		t.Fatal("expected GenericOpenAPIError type")
	}
	model, ok := openAPIErr.Model().(ApiError)
	if !ok {
		t.Fatal("expected ApiError model in error")
	}
	if model.Code != "NOT_FOUND" {
		t.Errorf("expected error code NOT_FOUND, got %q", model.Code)
	}
}

func TestGetCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    "INTERNAL_ERROR",
			Message: "server blew up",
		})
	})
	defer ts.Close()

	_, resp, err := client.CommitsApi.GetCommit(context.Background(), "myrepo", "abc123")
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// CreateCommit (POST) - Success and Error
// ---------------------------------------------------------------------------

func TestCreateCommit_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/v1/repositories/myrepo/commits") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		// Verify Content-Type
		if ct := r.Header.Get("Content-Type"); !strings.Contains(ct, "application/json") {
			t.Errorf("expected application/json Content-Type, got %q", ct)
		}

		// Decode request body
		var incoming Commit
		if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if incoming.Id != "new-commit" {
			t.Errorf("expected commit id new-commit in body, got %q", incoming.Id)
		}

		// Return created commit
		response := Commit{
			Id:         "new-commit",
			Properties: map[string]interface{}{"created": true},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(response)
	})
	defer ts.Close()

	input := Commit{
		Id:         "new-commit",
		Properties: map[string]interface{}{},
	}

	result, resp, err := client.CommitsApi.CreateCommit(context.Background(), "myrepo", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("expected status 201, got %d", resp.StatusCode)
	}
	if result.Id != "new-commit" {
		t.Errorf("expected commit id new-commit, got %q", result.Id)
	}
}

func TestCreateCommit_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    "NOT_FOUND",
			Message: "repository not found",
		})
	})
	defer ts.Close()

	input := Commit{Id: "test", Properties: map[string]interface{}{}}
	_, resp, err := client.CommitsApi.CreateCommit(context.Background(), "nonexistent", input)
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}
}

func TestCreateCommit_Unauthorized(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    "UNAUTHORIZED",
			Message: "authentication required",
		})
	})
	defer ts.Close()

	input := Commit{Id: "test", Properties: map[string]interface{}{}}
	_, resp, err := client.CommitsApi.CreateCommit(context.Background(), "myrepo", input)
	if err == nil {
		t.Fatal("expected error for 401 response")
	}
	if resp.StatusCode != 401 {
		t.Errorf("expected status 401, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// DeleteCommit (DELETE) - Success and Error
// ---------------------------------------------------------------------------

func TestDeleteCommit_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/v1/repositories/myrepo/commits/commit-to-delete") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.CommitsApi.DeleteCommit(context.Background(), "myrepo", "commit-to-delete")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected status 204, got %d", resp.StatusCode)
	}
}

func TestDeleteCommit_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    "NOT_FOUND",
			Message: "commit not found",
		})
	})
	defer ts.Close()

	resp, err := client.CommitsApi.DeleteCommit(context.Background(), "myrepo", "nonexistent")
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}

	openAPIErr, ok := err.(GenericOpenAPIError)
	if !ok {
		t.Fatal("expected GenericOpenAPIError type")
	}
	model, ok := openAPIErr.Model().(ApiError)
	if !ok {
		t.Fatal("expected ApiError model in error")
	}
	if model.Code != "NOT_FOUND" {
		t.Errorf("expected error code NOT_FOUND, got %q", model.Code)
	}
}

func TestDeleteCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    "INTERNAL_ERROR",
			Message: "unexpected error",
		})
	})
	defer ts.Close()

	resp, err := client.CommitsApi.DeleteCommit(context.Background(), "myrepo", "abc123")
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// Verify request headers and user agent
// ---------------------------------------------------------------------------

func TestAPIRequest_UserAgentHeader(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		if ua != "OpenAPI-Generator/1.0.0/go" {
			t.Errorf("expected default User-Agent, got %q", ua)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test"})
	})
	defer ts.Close()

	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "repo", "id")
}

func TestAPIRequest_AcceptHeader(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		if accept != "application/json" {
			t.Errorf("expected Accept: application/json, got %q", accept)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test"})
	})
	defer ts.Close()

	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "repo", "id")
}

func TestAPIRequest_CustomDefaultHeader(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		custom := r.Header.Get("X-Custom-Header")
		if custom != "custom-value" {
			t.Errorf("expected X-Custom-Header=custom-value, got %q", custom)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test"})
	})
	defer ts.Close()

	client.GetConfig().AddDefaultHeader("X-Custom-Header", "custom-value")
	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "repo", "id")
}

// ---------------------------------------------------------------------------
// Context-based authentication
// ---------------------------------------------------------------------------

func TestAPIRequest_BearerTokenAuth(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Bearer my-token" {
			t.Errorf("expected Bearer my-token, got %q", auth)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test"})
	})
	defer ts.Close()

	ctx := context.WithValue(context.Background(), ContextAccessToken, "my-token")
	_, _, _ = client.CommitsApi.GetCommit(ctx, "repo", "id")
}

func TestAPIRequest_BasicAuth(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			t.Error("expected basic auth to be set")
		}
		if user != "admin" || pass != "secret" {
			t.Errorf("expected admin:secret, got %s:%s", user, pass)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test"})
	})
	defer ts.Close()

	ctx := context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
		UserName: "admin",
		Password: "secret",
	})
	_, _, _ = client.CommitsApi.GetCommit(ctx, "repo", "id")
}

// ---------------------------------------------------------------------------
// URL path parameter encoding
// ---------------------------------------------------------------------------

func TestAPIRequest_PathParameterEncoding(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		// parameterToString + QueryEscape encodes spaces as "+" and "/" as "%2F"
		// The raw path should contain these encoded values
		path := r.URL.RawPath
		if path == "" {
			path = r.URL.Path
		}
		if !strings.Contains(path, "my+repo") {
			t.Errorf("expected 'my+repo' in path, got %q", path)
		}
		if !strings.Contains(path, "commit%2F1") {
			// Go's HTTP server may decode the path, check URL.Path as fallback
			if !strings.Contains(r.URL.Path, "commit/1") {
				t.Errorf("expected 'commit%%2F1' or 'commit/1' in path, got raw=%q decoded=%q", r.URL.RawPath, r.URL.Path)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test"})
	})
	defer ts.Close()

	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "my repo", "commit/1")
}
