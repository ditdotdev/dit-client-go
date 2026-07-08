// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package ditclient

import (
	"context"
	"testing"
)

const testDefaultBasePath = "http://localhost:5001"

// ---------------------------------------------------------------------------
// NewConfiguration
// ---------------------------------------------------------------------------

func TestNewConfiguration_Defaults(t *testing.T) {
	cfg := NewConfiguration()

	if cfg.UserAgent != "OpenAPI-Generator/1.0.0/go" {
		t.Errorf("expected default UserAgent, got %q", cfg.UserAgent)
	}
	if cfg.Debug {
		t.Error("expected Debug to be false by default")
	}
	if cfg.DefaultHeader == nil {
		t.Error("expected DefaultHeader map to be initialized")
	}
	if len(cfg.DefaultHeader) != 0 {
		t.Errorf("expected empty DefaultHeader, got %d entries", len(cfg.DefaultHeader))
	}
	if cfg.HTTPClient != nil {
		t.Error("expected HTTPClient to be nil by default (set on NewAPIClient)")
	}
}

func TestNewConfiguration_ServersInitialized(t *testing.T) {
	cfg := NewConfiguration()

	if len(cfg.Servers) != 1 {
		t.Fatalf("expected 1 server, got %d", len(cfg.Servers))
	}
	if cfg.Servers[0].URL != testDefaultBasePath {
		t.Errorf("expected server URL http://localhost:5001, got %q", cfg.Servers[0].URL)
	}
	if cfg.Servers[0].Description != "Local Dit server (default)" {
		t.Errorf("expected server description, got %q", cfg.Servers[0].Description)
	}
}

// ---------------------------------------------------------------------------
// AddDefaultHeader
// ---------------------------------------------------------------------------

func TestAddDefaultHeader_Single(t *testing.T) {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("X-Custom", "value1")

	if cfg.DefaultHeader["X-Custom"] != "value1" {
		t.Errorf("expected X-Custom=value1, got %q", cfg.DefaultHeader["X-Custom"])
	}
}

func TestAddDefaultHeader_Multiple(t *testing.T) {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("X-First", "one")
	cfg.AddDefaultHeader("X-Second", "two")

	if len(cfg.DefaultHeader) != 2 {
		t.Errorf("expected 2 headers, got %d", len(cfg.DefaultHeader))
	}
	if cfg.DefaultHeader["X-First"] != "one" {
		t.Errorf("expected X-First=one, got %q", cfg.DefaultHeader["X-First"])
	}
	if cfg.DefaultHeader["X-Second"] != "two" {
		t.Errorf("expected X-Second=two, got %q", cfg.DefaultHeader["X-Second"])
	}
}

func TestAddDefaultHeader_Overwrite(t *testing.T) {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("X-Custom", "original")
	cfg.AddDefaultHeader("X-Custom", "updated")

	if cfg.DefaultHeader["X-Custom"] != "updated" {
		t.Errorf("expected X-Custom=updated, got %q", cfg.DefaultHeader["X-Custom"])
	}
}

// ---------------------------------------------------------------------------
// ServerURL
// ---------------------------------------------------------------------------

func TestServerURL_DefaultIndex(t *testing.T) {
	cfg := NewConfiguration()
	url, err := cfg.ServerURL(0, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != testDefaultBasePath {
		t.Errorf("expected http://localhost:5001, got %q", url)
	}
}

func TestServerURL_IndexOutOfRange_Negative(t *testing.T) {
	cfg := NewConfiguration()
	_, err := cfg.ServerURL(-1, nil)
	if err == nil {
		t.Error("expected error for negative index")
	}
}

func TestServerURL_IndexOutOfRange_TooLarge(t *testing.T) {
	cfg := NewConfiguration()
	_, err := cfg.ServerURL(5, nil)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestServerURL_VariableSubstitution(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://{host}:{port}/api",
			Description: "Test server",
			Variables: map[string]ServerVariable{
				"host": {
					DefaultValue: "localhost",
					EnumValues:   []string{"localhost", "example.com"},
				},
				"port": {
					DefaultValue: "8080",
					EnumValues:   []string{"8080", "443"},
				},
			},
		},
	}

	url, err := cfg.ServerURL(0, map[string]string{"host": "example.com", "port": "443"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://example.com:443/api" {
		t.Errorf("expected https://example.com:443/api, got %q", url)
	}
}

func TestServerURL_VariableSubstitution_DefaultValues(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://{host}:{port}/api",
			Description: "Test server",
			Variables: map[string]ServerVariable{
				"host": {
					DefaultValue: "localhost",
					EnumValues:   []string{},
				},
				"port": {
					DefaultValue: "8080",
					EnumValues:   []string{},
				},
			},
		},
	}

	url, err := cfg.ServerURL(0, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://localhost:8080/api" {
		t.Errorf("expected https://localhost:8080/api, got %q", url)
	}
}

func TestServerURL_InvalidVariableValue(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://{host}/api",
			Description: "Test server",
			Variables: map[string]ServerVariable{
				"host": {
					DefaultValue: "localhost",
					EnumValues:   []string{"localhost", "example.com"},
				},
			},
		},
	}

	_, err := cfg.ServerURL(0, map[string]string{"host": "invalid-host.com"})
	if err == nil {
		t.Error("expected error for invalid variable value")
	}
}

func TestServerURL_NoEnumRestrictions(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://{host}/api",
			Description: "Test server",
			Variables: map[string]ServerVariable{
				"host": {
					DefaultValue: "localhost",
					EnumValues:   []string{}, // empty enum means any value is allowed
				},
			},
		},
	}

	url, err := cfg.ServerURL(0, map[string]string{"host": "any-host.com"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://any-host.com/api" {
		t.Errorf("expected https://any-host.com/api, got %q", url)
	}
}

// ---------------------------------------------------------------------------
// ServerURLWithContext
// ---------------------------------------------------------------------------

func TestServerURLWithContext_NilContext(t *testing.T) {
	// The generated ServerURLWithContext explicitly handles a nil context as
	// a fast-path that skips index/variable lookup. We test that fast-path
	// here, which is why staticcheck's SA1012 is suppressed.
	cfg := NewConfiguration()
	url, err := cfg.ServerURLWithContext(nil, "any") //nolint:staticcheck // SA1012: exercising the explicit nil-context fast path
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != testDefaultBasePath {
		t.Errorf("expected default URL, got %q", url)
	}
}

func TestServerURLWithContext_ServerIndexOverride(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{URL: "http://primary"},
		{URL: "http://secondary"},
	}
	ctx := context.WithValue(context.Background(), ContextServerIndex, 1)
	url, err := cfg.ServerURLWithContext(ctx, "any")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "http://secondary" {
		t.Errorf("expected secondary URL, got %q", url)
	}
}

// ---------------------------------------------------------------------------
// contextKey
// ---------------------------------------------------------------------------

func TestContextKey_String(t *testing.T) {
	key := contextKey("token")
	if key.String() != "auth token" {
		t.Errorf("expected 'auth token', got %q", key.String())
	}
}

func TestContextKey_ServerIndex(t *testing.T) {
	if ContextServerIndex.String() != "auth serverIndex" {
		t.Errorf("expected 'auth serverIndex', got %q", ContextServerIndex.String())
	}
}
