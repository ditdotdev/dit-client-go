package ditclient

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetContext_Success(t *testing.T) {
	want := Context{
		Provider:   "docker",
		Properties: map[string]interface{}{"host": "unix:///var/run/docker.sock"},
	}

	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/context" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, resp, err := client.ContextsApi.GetContext(context.Background()).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	if got.Provider != "docker" {
		t.Errorf("expected provider docker, got %q", got.Provider)
	}
}

func TestGetContext_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "boom"})
	})
	defer ts.Close()

	_, resp, err := client.ContextsApi.GetContext(context.Background()).Execute()
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", resp.StatusCode)
	}
}

func TestGetContext_TransportError(t *testing.T) {
	// Point at a closed port to exercise the prepareRequest / callAPI
	// error path (no httptest.Server here).
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:1"}}
	client := NewAPIClient(cfg)

	_, _, err := client.ContextsApi.GetContext(context.Background()).Execute()
	if err == nil {
		t.Fatal("expected transport-level error for unreachable server")
	}
}
