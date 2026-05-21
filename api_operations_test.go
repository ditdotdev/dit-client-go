package datadatdatclient

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// ListOperations
// ---------------------------------------------------------------------------

func TestListOperations_Success(t *testing.T) {
	want := []Operation{
		{Id: "op-1", Type: "PULL", State: "RUNNING", Remote: "origin", CommitId: "abc"},
	}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/operations" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.OperationsApi.ListOperations(context.Background()).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0].Id != "op-1" {
		t.Errorf("unexpected operations: %+v", got)
	}
}

func TestListOperations_WithRepositoryFilter(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("repository"); got != "alpha" {
			t.Errorf("expected repository=alpha, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode([]Operation{})
	})
	defer ts.Close()

	_, _, err := client.OperationsApi.ListOperations(context.Background()).Repository("alpha").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestListOperations_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "boom"})
	})
	defer ts.Close()

	_, resp, err := client.OperationsApi.ListOperations(context.Background()).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected 500, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetOperation
// ---------------------------------------------------------------------------

func TestGetOperation_Success(t *testing.T) {
	want := Operation{Id: "op-1", Type: "PULL", State: "COMPLETE", Remote: "origin", CommitId: "abc"}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/v1/operations/op-1") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.OperationsApi.GetOperation(context.Background(), "op-1").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.State != "COMPLETE" {
		t.Errorf("expected COMPLETE, got %q", got.State)
	}
}

func TestGetOperation_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.OperationsApi.GetOperation(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// AbortOperation
// ---------------------------------------------------------------------------

func TestAbortOperation_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.OperationsApi.AbortOperation(context.Background(), "op-1").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestAbortOperation_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	resp, err := client.OperationsApi.AbortOperation(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetOperationProgress
// ---------------------------------------------------------------------------

func TestGetOperationProgress_Success(t *testing.T) {
	want := []ProgressEntry{
		{Id: 1, Type: "MESSAGE", Message: PtrString("started")},
		{Id: 2, Type: "PROGRESS", Percent: PtrInt32(50)},
	}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/progress") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.OperationsApi.GetOperationProgress(context.Background(), "op-1").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 entries, got %d", len(got))
	}
}

func TestGetOperationProgress_WithLastIdFilter(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("lastId"); got != "5" {
			t.Errorf("expected lastId=5, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode([]ProgressEntry{})
	})
	defer ts.Close()

	_, _, err := client.OperationsApi.GetOperationProgress(context.Background(), "op-1").LastId(5).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetOperationProgress_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.OperationsApi.GetOperationProgress(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// Pull
// ---------------------------------------------------------------------------

func TestPull_Success(t *testing.T) {
	want := Operation{Id: "op-pull", Type: "PULL", State: "RUNNING", Remote: "origin", CommitId: "abc"}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/pull") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var rp RemoteParameters
		_ = json.NewDecoder(r.Body).Decode(&rp)
		if rp.Provider != "s3" {
			t.Errorf("expected provider s3 in body, got %q", rp.Provider)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.OperationsApi.Pull(context.Background(), "alpha", "origin", "abc").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Id != "op-pull" {
		t.Errorf("expected op-pull, got %q", got.Id)
	}
}

func TestPull_WithMetadataOnly(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("metadataOnly"); got != "true" {
			t.Errorf("expected metadataOnly=true, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(Operation{Id: "x", Type: "PULL", State: "RUNNING", Remote: "r", CommitId: "c"})
	})
	defer ts.Close()

	_, _, err := client.OperationsApi.Pull(context.Background(), "alpha", "origin", "abc").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		MetadataOnly(true).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPull_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.OperationsApi.Pull(context.Background(), "alpha", "origin", "abc").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// Push
// ---------------------------------------------------------------------------

func TestPush_Success(t *testing.T) {
	want := Operation{Id: "op-push", Type: "PUSH", State: "RUNNING", Remote: "origin", CommitId: "abc"}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/push") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.OperationsApi.Push(context.Background(), "alpha", "origin", "abc").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Id != "op-push" {
		t.Errorf("expected op-push, got %q", got.Id)
	}
}

func TestPush_WithMetadataOnly(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("metadataOnly"); got != "true" {
			t.Errorf("expected metadataOnly=true, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(Operation{Id: "x", Type: "PUSH", State: "RUNNING", Remote: "r", CommitId: "c"})
	})
	defer ts.Close()

	_, _, err := client.OperationsApi.Push(context.Background(), "alpha", "origin", "abc").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		MetadataOnly(true).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPush_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.OperationsApi.Push(context.Background(), "alpha", "origin", "abc").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestPush_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "bad"})
	})
	defer ts.Close()

	_, resp, err := client.OperationsApi.Push(context.Background(), "alpha", "origin", "abc").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).
		Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}
