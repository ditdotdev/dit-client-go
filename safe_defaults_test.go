package datadatdatclient

import (
	"net/http"
	"sync"
	"testing"
	"time"
)

// ---------------------------------------------------------------------------
// NewAPIClientWithDefaults
// ---------------------------------------------------------------------------

func TestNewAPIClientWithDefaults_AssignsFreshClient(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClientWithDefaults(cfg)

	if client.GetConfig().HTTPClient == nil {
		t.Fatal("expected a non-nil *http.Client")
	}
	if client.GetConfig().HTTPClient == http.DefaultClient {
		t.Error("expected a non-shared *http.Client, got http.DefaultClient")
	}
	if client.GetConfig().HTTPClient.Timeout != DefaultClientTimeout {
		t.Errorf("expected timeout %v, got %v", DefaultClientTimeout, client.GetConfig().HTTPClient.Timeout)
	}
}

func TestNewAPIClientWithDefaults_PreservesCustomClient(t *testing.T) {
	cfg := NewConfiguration()
	custom := &http.Client{Timeout: 7 * time.Second}
	cfg.HTTPClient = custom

	client := NewAPIClientWithDefaults(cfg)

	if client.GetConfig().HTTPClient != custom {
		t.Error("expected caller-supplied HTTPClient to be preserved")
	}
	if client.GetConfig().HTTPClient.Timeout != 7*time.Second {
		t.Errorf("expected preserved timeout 7s, got %v", client.GetConfig().HTTPClient.Timeout)
	}
}

func TestNewAPIClientWithDefaults_ReturnsWorkingClient(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClientWithDefaults(cfg)

	// The defaulted constructor should still wire up the service stubs.
	if client.CommitsApi == nil || client.RemotesApi == nil || client.VolumesApi == nil {
		t.Error("expected service stubs to be initialized")
	}
}

// ---------------------------------------------------------------------------
// SetDefaultHeader / DefaultHeaders
// ---------------------------------------------------------------------------

func TestSetDefaultHeader_WritesEntry(t *testing.T) {
	cfg := NewConfiguration()
	SetDefaultHeader(cfg, "X-Test", "value")

	if cfg.DefaultHeader["X-Test"] != "value" {
		t.Errorf("expected DefaultHeader[X-Test]=value, got %q", cfg.DefaultHeader["X-Test"])
	}
}

func TestSetDefaultHeader_InitializesNilMap(t *testing.T) {
	cfg := &Configuration{} // intentionally not via NewConfiguration to exercise the nil-map branch
	SetDefaultHeader(cfg, "X-Test", "value")

	if cfg.DefaultHeader == nil {
		t.Fatal("expected DefaultHeader map to be initialized")
	}
	if cfg.DefaultHeader["X-Test"] != "value" {
		t.Errorf("expected X-Test=value, got %q", cfg.DefaultHeader["X-Test"])
	}
}

func TestSetDefaultHeader_Overwrites(t *testing.T) {
	cfg := NewConfiguration()
	SetDefaultHeader(cfg, "X-Test", "first")
	SetDefaultHeader(cfg, "X-Test", "second")

	if cfg.DefaultHeader["X-Test"] != "second" {
		t.Errorf("expected X-Test=second, got %q", cfg.DefaultHeader["X-Test"])
	}
}

func TestDefaultHeaders_ReturnsSnapshot(t *testing.T) {
	cfg := NewConfiguration()
	SetDefaultHeader(cfg, "X-First", "1")
	SetDefaultHeader(cfg, "X-Second", "2")

	snap := DefaultHeaders(cfg)
	if len(snap) != 2 {
		t.Errorf("expected 2 entries in snapshot, got %d", len(snap))
	}
	if snap["X-First"] != "1" || snap["X-Second"] != "2" {
		t.Errorf("unexpected snapshot contents: %v", snap)
	}

	// Mutating the snapshot must not affect the underlying map.
	snap["X-Third"] = "3"
	if _, ok := cfg.DefaultHeader["X-Third"]; ok {
		t.Error("mutating the snapshot leaked into cfg.DefaultHeader")
	}
}

func TestDefaultHeaders_EmptyOnFreshConfig(t *testing.T) {
	cfg := NewConfiguration()
	snap := DefaultHeaders(cfg)
	if len(snap) != 0 {
		t.Errorf("expected empty snapshot, got %d entries", len(snap))
	}
}

// TestConcurrentSetAndSnapshot verifies that SetDefaultHeader and
// DefaultHeaders can run concurrently without tripping Go's race detector.
// Run with `go test -race ./...` to actually exercise the data-race
// detector — the test still passes (no panic, no result corruption) under
// plain `go test`, but the value is in the -race run.
func TestConcurrentSetAndSnapshot(t *testing.T) {
	cfg := NewConfiguration()

	const goroutines = 8
	const iterations = 100

	var wg sync.WaitGroup
	wg.Add(goroutines * 2)

	// Writers
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				SetDefaultHeader(cfg, "X-Worker", "value")
			}
		}()
	}

	// Readers
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_ = DefaultHeaders(cfg)
			}
		}()
	}

	wg.Wait()
}
