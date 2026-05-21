package datadatdatclient

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// ListRepositories
// ---------------------------------------------------------------------------

func TestListRepositories_Success(t *testing.T) {
	want := []Repository{
		{Name: "alpha", Properties: map[string]interface{}{}},
		{Name: "beta", Properties: map[string]interface{}{}},
	}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/repositories" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.RepositoriesApi.ListRepositories(context.Background()).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 || got[0].Name != "alpha" {
		t.Errorf("unexpected repositories: %+v", got)
	}
}

func TestListRepositories_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "kaboom"})
	})
	defer ts.Close()

	_, resp, err := client.RepositoriesApi.ListRepositories(context.Background()).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// CreateRepository
// ---------------------------------------------------------------------------

func TestCreateRepository_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var incoming Repository
		if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
			t.Fatalf("bad body: %v", err)
		}
		if incoming.Name != "newrepo" {
			t.Errorf("expected newrepo, got %q", incoming.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(incoming)
	})
	defer ts.Close()

	in := Repository{Name: "newrepo", Properties: map[string]interface{}{}}
	got, resp, err := client.RepositoriesApi.CreateRepository(context.Background()).Repository(in).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}
	if got.Name != "newrepo" {
		t.Errorf("expected newrepo, got %q", got.Name)
	}
}

func TestCreateRepository_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RepositoriesApi.CreateRepository(context.Background()).Execute()
	if err == nil {
		t.Fatal("expected error when repository body not supplied")
	}
	if !strings.Contains(err.Error(), "repository is required") {
		t.Errorf("expected 'repository is required', got %q", err.Error())
	}
}

func TestCreateRepository_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ApiError{Code: PtrString("BAD_INPUT"), Message: "invalid"})
	})
	defer ts.Close()

	in := Repository{Name: "bad", Properties: map[string]interface{}{}}
	_, resp, err := client.RepositoriesApi.CreateRepository(context.Background()).Repository(in).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetRepository
// ---------------------------------------------------------------------------

func TestGetRepository_Success(t *testing.T) {
	want := Repository{Name: "alpha", Properties: map[string]interface{}{"owner": "team-a"}}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/v1/repositories/alpha") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.RepositoriesApi.GetRepository(context.Background(), "alpha").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "alpha" {
		t.Errorf("expected alpha, got %q", got.Name)
	}
}

func TestGetRepository_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Code: PtrString("NOT_FOUND"), Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.RepositoriesApi.GetRepository(context.Background(), "missing").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// UpdateRepository
// ---------------------------------------------------------------------------

func TestUpdateRepository_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var incoming Repository
		_ = json.NewDecoder(r.Body).Decode(&incoming)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(incoming)
	})
	defer ts.Close()

	in := Repository{Name: "alpha2", Properties: map[string]interface{}{}}
	got, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "alpha").Repository(in).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "alpha2" {
		t.Errorf("expected alpha2, got %q", got.Name)
	}
}

func TestUpdateRepository_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "alpha").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// DeleteRepository
// ---------------------------------------------------------------------------

func TestDeleteRepository_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.RepositoriesApi.DeleteRepository(context.Background(), "alpha").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestDeleteRepository_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	resp, err := client.RepositoriesApi.DeleteRepository(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetRepositoryStatus
// ---------------------------------------------------------------------------

func TestGetRepositoryStatus_Success(t *testing.T) {
	want := RepositoryStatus{LastCommit: PtrString("abc"), SourceCommit: PtrString("xyz")}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/status") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.RepositoriesApi.GetRepositoryStatus(context.Background(), "alpha").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetLastCommit() != "abc" {
		t.Errorf("expected lastCommit abc, got %q", got.GetLastCommit())
	}
}

func TestGetRepositoryStatus_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "no such repo"})
	})
	defer ts.Close()

	_, resp, err := client.RepositoriesApi.GetRepositoryStatus(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}
