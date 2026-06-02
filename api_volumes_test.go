package ditclient

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// ListVolumes
// ---------------------------------------------------------------------------

func TestListVolumes_Success(t *testing.T) {
	want := []Volume{
		{Name: "data", Properties: map[string]interface{}{}},
		{Name: "logs", Properties: map[string]interface{}{}},
	}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/v1/repositories/alpha/volumes") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.VolumesApi.ListVolumes(context.Background(), "alpha").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 volumes, got %d", len(got))
	}
}

func TestListVolumes_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "no repo"})
	})
	defer ts.Close()

	_, resp, err := client.VolumesApi.ListVolumes(context.Background(), "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// CreateVolume
// ---------------------------------------------------------------------------

func TestCreateVolume_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var in Volume
		_ = json.NewDecoder(r.Body).Decode(&in)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(in)
	})
	defer ts.Close()

	in := Volume{Name: "data", Properties: map[string]interface{}{}}
	got, _, err := client.VolumesApi.CreateVolume(context.Background(), "alpha").Volume(in).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "data" {
		t.Errorf("expected data, got %q", got.Name)
	}
}

func TestCreateVolume_MissingBody(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:0"}}
	client := NewAPIClient(cfg)

	_, _, err := client.VolumesApi.CreateVolume(context.Background(), "alpha").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCreateVolume_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "bad"})
	})
	defer ts.Close()

	in := Volume{Name: "data", Properties: map[string]interface{}{}}
	_, resp, err := client.VolumesApi.CreateVolume(context.Background(), "alpha").Volume(in).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetVolume
// ---------------------------------------------------------------------------

func TestGetVolume_Success(t *testing.T) {
	want := Volume{Name: "data", Properties: map[string]interface{}{}}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.VolumesApi.GetVolume(context.Background(), "alpha", "data").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "data" {
		t.Errorf("expected data, got %q", got.Name)
	}
}

func TestGetVolume_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.VolumesApi.GetVolume(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// DeleteVolume
// ---------------------------------------------------------------------------

func TestDeleteVolume_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.VolumesApi.DeleteVolume(context.Background(), "alpha", "data").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestDeleteVolume_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	resp, err := client.VolumesApi.DeleteVolume(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// ActivateVolume
// ---------------------------------------------------------------------------

func TestActivateVolume_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/activate") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.VolumesApi.ActivateVolume(context.Background(), "alpha", "data").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestActivateVolume_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	resp, err := client.VolumesApi.ActivateVolume(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// DeactivateVolume
// ---------------------------------------------------------------------------

func TestDeactivateVolume_Success(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/deactivate") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer ts.Close()

	resp, err := client.VolumesApi.DeactivateVolume(context.Background(), "alpha", "data").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestDeactivateVolume_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	resp, err := client.VolumesApi.DeactivateVolume(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// GetVolumeStatus
// ---------------------------------------------------------------------------

func TestGetVolumeStatus_Success(t *testing.T) {
	want := VolumeStatus{
		Name:        "data",
		LogicalSize: 1024,
		ActualSize:  512,
		Properties:  map[string]interface{}{},
		Ready:       true,
	}
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/status") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})
	defer ts.Close()

	got, _, err := client.VolumesApi.GetVolumeStatus(context.Background(), "alpha", "data").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got.Ready {
		t.Error("expected Ready=true")
	}
	if got.LogicalSize != 1024 {
		t.Errorf("expected logicalSize=1024, got %d", got.LogicalSize)
	}
}

func TestGetVolumeStatus_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ApiError{Message: "missing"})
	})
	defer ts.Close()

	_, resp, err := client.VolumesApi.GetVolumeStatus(context.Background(), "alpha", "ghost").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}
