package datadatdatclient

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
// ListRemoteCommits / GetRemoteCommit
//
// These two operations take Datadatdat-remote-parameters as a complex object
// header (in: header, schema: $ref RemoteParameters). The v7 generator
// encodes objects-in-headers using bracketed key names (e.g.
// "Datadatdat-remote-parameters[provider]"), which Go's net/http rejects as
// an invalid header field name. This is a genuine bug we need to resolve —
// either by changing the spec to encode remoteParameters as a JSON string or
// in the request body, or by patching the generator's header encoding for
// complex types.
//
// Until that's resolved (tracked in a follow-up issue), the only well-defined
// behavior we can test is that calling Execute() without the required
// DatadatdatRemoteParameters builder method returns the documented error.
// ---------------------------------------------------------------------------

func TestListRemoteCommits_MissingRequiredHeader(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RemotesApi.ListRemoteCommits(context.Background(), "alpha", "main").Execute()
	if err == nil {
		t.Fatal("expected error when DatadatdatRemoteParameters is missing")
	}
	if !strings.Contains(err.Error(), "datadatdatRemoteParameters is required") {
		t.Errorf("unexpected error text: %q", err.Error())
	}
}

func TestGetRemoteCommit_MissingRequiredHeader(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.RemotesApi.GetRemoteCommit(context.Background(), "alpha", "main", "abc").Execute()
	if err == nil {
		t.Fatal("expected error when DatadatdatRemoteParameters is missing")
	}
	if !strings.Contains(err.Error(), "datadatdatRemoteParameters is required") {
		t.Errorf("unexpected error text: %q", err.Error())
	}
}
