package datadatdatclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testErrorCodeNotFound = "NOT_FOUND"

// newTestServer creates a test server and configures an APIClient pointing to it.
func newTestServer(handler http.HandlerFunc) (*httptest.Server, *APIClient) {
	ts := httptest.NewServer(handler)
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: ts.URL}}
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

	result, resp, err := client.CommitsApi.GetCommit(context.Background(), "myrepo", "abc123").Execute()
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
		Code:    PtrString(testErrorCodeNotFound),
		Message: "Commit not found",
	}

	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(apiErr)
	})
	defer ts.Close()

	_, resp, err := client.CommitsApi.GetCommit(context.Background(), "myrepo", "nonexistent").Execute()
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}

	openAPIErr, ok := err.(*GenericOpenAPIError)
	if !ok {
		t.Fatalf("expected *GenericOpenAPIError type, got %T", err)
	}
	model, ok := openAPIErr.Model().(ApiError)
	if !ok {
		t.Fatal("expected ApiError model in error")
	}
	if model.GetCode() != testErrorCodeNotFound {
		t.Errorf("expected error code NOT_FOUND, got %q", model.GetCode())
	}
}

func TestGetCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    PtrString("INTERNAL_ERROR"),
			Message: "server blew up",
		})
	})
	defer ts.Close()

	_, resp, err := client.CommitsApi.GetCommit(context.Background(), "myrepo", "abc123").Execute()
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
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/v1/repositories/myrepo/commits") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		if ct := r.Header.Get("Content-Type"); !strings.Contains(ct, "application/json") {
			t.Errorf("expected application/json Content-Type, got %q", ct)
		}

		var incoming Commit
		if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if incoming.Id != "new-commit" {
			t.Errorf("expected commit id new-commit in body, got %q", incoming.Id)
		}

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

	result, resp, err := client.CommitsApi.CreateCommit(context.Background(), "myrepo").Commit(input).Execute()
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
			Code:    PtrString(testErrorCodeNotFound),
			Message: "repository not found",
		})
	})
	defer ts.Close()

	input := Commit{Id: "test", Properties: map[string]interface{}{}}
	_, resp, err := client.CommitsApi.CreateCommit(context.Background(), "nonexistent").Commit(input).Execute()
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}
}

func TestCreateCommit_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    PtrString("BAD_INPUT"),
			Message: "malformed commit",
		})
	})
	defer ts.Close()

	input := Commit{Id: "test", Properties: map[string]interface{}{}}
	_, resp, err := client.CommitsApi.CreateCommit(context.Background(), "myrepo").Commit(input).Execute()
	if err == nil {
		t.Fatal("expected error for 400 response")
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", resp.StatusCode)
	}
}

func TestCreateCommit_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://localhost:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.CommitsApi.CreateCommit(context.Background(), "myrepo").Execute()
	if err == nil {
		t.Fatal("expected error when commit body is not provided")
	}
	if !strings.Contains(err.Error(), "commit is required") {
		t.Errorf("expected 'commit is required' error, got %q", err.Error())
	}
}

// ---------------------------------------------------------------------------
// DeleteCommit (DELETE) - Success and Error
// ---------------------------------------------------------------------------

func TestDeleteCommit_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/v1/repositories/myrepo/commits/commit-to-delete") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.CommitsApi.DeleteCommit(context.Background(), "myrepo", "commit-to-delete").Execute()
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
			Code:    PtrString(testErrorCodeNotFound),
			Message: "commit not found",
		})
	})
	defer ts.Close()

	resp, err := client.CommitsApi.DeleteCommit(context.Background(), "myrepo", "nonexistent").Execute()
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}

	openAPIErr, ok := err.(*GenericOpenAPIError)
	if !ok {
		t.Fatalf("expected *GenericOpenAPIError type, got %T", err)
	}
	model, ok := openAPIErr.Model().(ApiError)
	if !ok {
		t.Fatal("expected ApiError model in error")
	}
	if model.GetCode() != testErrorCodeNotFound {
		t.Errorf("expected error code NOT_FOUND, got %q", model.GetCode())
	}
}

func TestDeleteCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiError{
			Code:    PtrString("INTERNAL_ERROR"),
			Message: "unexpected error",
		})
	})
	defer ts.Close()

	resp, err := client.CommitsApi.DeleteCommit(context.Background(), "myrepo", "abc123").Execute()
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// Request headers
// ---------------------------------------------------------------------------

func TestAPIRequest_UserAgentHeader(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		if ua != "OpenAPI-Generator/1.0.0/go" {
			t.Errorf("expected default User-Agent, got %q", ua)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test", Properties: map[string]interface{}{}})
	})
	defer ts.Close()

	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "repo", "id").Execute()
}

func TestAPIRequest_AcceptHeader(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		if accept != testContentTypeJSON {
			t.Errorf("expected Accept: application/json, got %q", accept)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test", Properties: map[string]interface{}{}})
	})
	defer ts.Close()

	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "repo", "id").Execute()
}

func TestAPIRequest_CustomDefaultHeader(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		custom := r.Header.Get("X-Custom-Header")
		if custom != "custom-value" {
			t.Errorf("expected X-Custom-Header=custom-value, got %q", custom)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test", Properties: map[string]interface{}{}})
	})
	defer ts.Close()

	client.GetConfig().AddDefaultHeader("X-Custom-Header", "custom-value")
	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "repo", "id").Execute()
}

// ---------------------------------------------------------------------------
// Per-request server override via context
// ---------------------------------------------------------------------------

func TestAPIRequest_ServerIndexOverride(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "from-secondary", Properties: map[string]interface{}{}})
	})
	defer ts.Close()

	// Configure a primary server pointing to a dead address and a secondary
	// server pointing at the test httptest server. With ContextServerIndex=1,
	// the request must route to the secondary.
	client.GetConfig().Servers = ServerConfigurations{
		{URL: "http://127.0.0.1:1"},
		{URL: ts.URL},
	}

	ctx := context.WithValue(context.Background(), ContextServerIndex, 1)
	result, _, err := client.CommitsApi.GetCommit(ctx, "repo", "id").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Id != "from-secondary" {
		t.Errorf("expected response from secondary server, got %q", result.Id)
	}
}

// ---------------------------------------------------------------------------
// URL path parameter encoding
// ---------------------------------------------------------------------------

func TestAPIRequest_PathParameterEncoding(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.RawPath
		if path == "" {
			path = r.URL.Path
		}
		if !strings.Contains(path, "my%20repo") {
			t.Errorf("expected 'my%%20repo' in path (PathEscape), got %q", path)
		}
		if !strings.Contains(path, "commit%2F1") {
			if !strings.Contains(r.URL.Path, "commit/1") {
				t.Errorf("expected 'commit%%2F1' or 'commit/1' in path, got raw=%q decoded=%q", r.URL.RawPath, r.URL.Path)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Commit{Id: "test", Properties: map[string]interface{}{}})
	})
	defer ts.Close()

	_, _, _ = client.CommitsApi.GetCommit(context.Background(), "my repo", "commit/1").Execute()
}

// ---------------------------------------------------------------------------
// ListCommits
// ---------------------------------------------------------------------------

func TestListCommits_Success(t *testing.T) {
	want := []Commit{
		{Id: "c1", Properties: map[string]interface{}{}},
		{Id: "c2", Properties: map[string]interface{}{}},
	}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/v1/repositories/alpha/commits") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.CommitsApi.ListCommits(context.Background(), "alpha").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 commits, got %d", len(got))
	}
}

func TestListCommits_WithTagFilter(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query()["tag"]; len(got) != 2 || got[0] != "production" || got[1] != "v1" {
			t.Errorf("expected tag=production&tag=v1, got %v", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode([]Commit{})
	})
	defer ts.Close()

	_, _, err := client.CommitsApi.ListCommits(context.Background(), "alpha").
		Tag([]string{"production", "v1"}).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestListCommits_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.CommitsApi.ListCommits(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// UpdateCommit
// ---------------------------------------------------------------------------

func TestUpdateCommit_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var in Commit
		_ = json.NewDecoder(r.Body).Decode(&in)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(in)
	})
	defer ts.Close()

	in := Commit{Id: "abc", Properties: map[string]interface{}{"tag": "v2"}}
	got, _, err := client.CommitsApi.UpdateCommit(context.Background(), "alpha", "abc").Commit(in).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Id != "abc" {
		t.Errorf("expected abc, got %q", got.Id)
	}
}

func TestUpdateCommit_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.CommitsApi.UpdateCommit(context.Background(), "alpha", "abc").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUpdateCommit_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	in := Commit{Id: "abc", Properties: map[string]interface{}{}}
	_, resp, err := client.CommitsApi.UpdateCommit(context.Background(), "alpha", "ghost").Commit(in).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetCommitStatus
// ---------------------------------------------------------------------------

func TestGetCommitStatus_Success(t *testing.T) {
	want := CommitStatus{LogicalSize: 100, ActualSize: 80, UniqueSize: 60, Ready: true}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/status") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.CommitsApi.GetCommitStatus(context.Background(), "alpha", "abc").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got.Ready {
		t.Error("expected Ready=true")
	}
}

func TestGetCommitStatus_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.CommitsApi.GetCommitStatus(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// CheckoutCommit
// ---------------------------------------------------------------------------

func TestCheckoutCommit_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/checkout") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.CommitsApi.CheckoutCommit(context.Background(), "alpha", "abc").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestCheckoutCommit_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	resp, err := client.CommitsApi.CheckoutCommit(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}
