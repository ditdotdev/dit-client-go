package ditclient

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// ListRemotes
// ---------------------------------------------------------------------------

func TestListRemotes_Success(t *testing.T) {
	want := []Remote{
		{Provider: "s3", Name: "main", Properties: map[string]interface{}{}},
		{Provider: "ssh", Name: "backup", Properties: map[string]interface{}{}},
	}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/v1/repositories/alpha/remotes") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.RemotesApi.ListRemotes(context.Background(), "alpha").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 remotes, got %d", len(got))
	}
}

func TestListRemotes_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "no repo"})
	})
	defer ts.Close()

	_, resp, err := client.RemotesApi.ListRemotes(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// CreateRemote
// ---------------------------------------------------------------------------

func TestCreateRemote_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		var in Remote
		_ = json.NewDecoder(r.Body).Decode(&in)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(in)
	})
	defer ts.Close()

	in := Remote{Provider: "s3", Name: "main", Properties: map[string]interface{}{}}
	got, _, err := client.RemotesApi.CreateRemote(context.Background(), "alpha").Remote(in).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "main" {
		t.Errorf("expected main, got %q", got.Name)
	}
}

func TestCreateRemote_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RemotesApi.CreateRemote(context.Background(), "alpha").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCreateRemote_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "bad"})
	})
	defer ts.Close()

	in := Remote{Provider: "s3", Name: "bad", Properties: map[string]interface{}{}}
	_, resp, err := client.RemotesApi.CreateRemote(context.Background(), "alpha").Remote(in).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetRemote
// ---------------------------------------------------------------------------

func TestGetRemote_Success(t *testing.T) {
	want := Remote{Provider: "s3", Name: "main", Properties: map[string]interface{}{}}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.RemotesApi.GetRemote(context.Background(), "alpha", "main").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "main" {
		t.Errorf("expected main, got %q", got.Name)
	}
}

func TestGetRemote_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.RemotesApi.GetRemote(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// UpdateRemote
// ---------------------------------------------------------------------------

func TestUpdateRemote_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		var in Remote
		_ = json.NewDecoder(r.Body).Decode(&in)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(in)
	})
	defer ts.Close()

	in := Remote{Provider: "s3", Name: "main2", Properties: map[string]interface{}{}}
	got, _, err := client.RemotesApi.UpdateRemote(context.Background(), "alpha", "main").Remote(in).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "main2" {
		t.Errorf("expected main2, got %q", got.Name)
	}
}

func TestUpdateRemote_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RemotesApi.UpdateRemote(context.Background(), "alpha", "main").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// DeleteRemote
// ---------------------------------------------------------------------------

func TestDeleteRemote_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.RemotesApi.DeleteRemote(context.Background(), "alpha", "main").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestDeleteRemote_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	resp, err := client.RemotesApi.DeleteRemote(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// ListRemoteCommits
//
// Post-#37: this is a POST that carries RemoteParameters in the request body
// instead of the original GET-with-object-header design that the v7
// generator couldn't encode without producing invalid HTTP header names.
// ---------------------------------------------------------------------------

func TestListRemoteCommits_Success(t *testing.T) {
	want := []Commit{{Id: "remote-1", Properties: map[string]interface{}{}}}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/v1/repositories/alpha/remotes/main/commits") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if ct := r.Header.Get("Content-Type"); !strings.Contains(ct, "application/json") {
			t.Errorf("expected application/json Content-Type, got %q", ct)
		}
		var got RemoteParameters
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("failed to decode body: %v", err)
		}
		if got.Provider != "s3" {
			t.Errorf("expected provider s3 in body, got %q", got.Provider)
		}
		if q := r.URL.Query()["tag"]; len(q) != 1 || q[0] != "v1" {
			t.Errorf("expected tag=v1 query param, got %v", q)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.RemotesApi.ListRemoteCommits(context.Background(), "alpha", "main").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		Tag([]string{"v1"}).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0].Id != "remote-1" {
		t.Errorf("unexpected commits: %+v", got)
	}
}

func TestListRemoteCommits_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "bad params"})
	})
	defer ts.Close()

	_, resp, err := client.RemotesApi.ListRemoteCommits(context.Background(), "alpha", "main").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestListRemoteCommits_MissingRequiredBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RemotesApi.ListRemoteCommits(context.Background(), "alpha", "main").Execute()
	if err == nil {
		t.Fatal("expected error when RemoteParameters is missing")
	}
	if !strings.Contains(err.Error(), "remoteParameters is required") {
		t.Errorf("unexpected error text: %q", err.Error())
	}
}

// ---------------------------------------------------------------------------
// GetRemoteCommit
// ---------------------------------------------------------------------------

func TestGetRemoteCommit_Success(t *testing.T) {
	want := Commit{Id: "abc123", Properties: map[string]interface{}{}}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var got RemoteParameters
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("failed to decode body: %v", err)
		}
		if got.Provider != "s3" {
			t.Errorf("expected provider s3 in body, got %q", got.Provider)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.RemotesApi.GetRemoteCommit(context.Background(), "alpha", "main", "abc123").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Id != "abc123" {
		t.Errorf("expected abc123, got %q", got.Id)
	}
}

func TestGetRemoteCommit_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.RemotesApi.GetRemoteCommit(context.Background(), "alpha", "main", "ghost").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

func TestGetRemoteCommit_MissingRequiredBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RemotesApi.GetRemoteCommit(context.Background(), "alpha", "main", "abc").Execute()
	if err == nil {
		t.Fatal("expected error when RemoteParameters is missing")
	}
	if !strings.Contains(err.Error(), "remoteParameters is required") {
		t.Errorf("unexpected error text: %q", err.Error())
	}
}
