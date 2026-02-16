package Datadatdatclient

import (
	"testing"
)

const testDefaultBasePath = "http://localhost:5001"

// ---------------------------------------------------------------------------
// NewConfiguration
// ---------------------------------------------------------------------------

func TestNewConfiguration_Defaults(t *testing.T) {
	cfg := NewConfiguration()

	if cfg.BasePath != testDefaultBasePath {
		t.Errorf("expected default BasePath http://localhost:5001, got %q", cfg.BasePath)
	}
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
	if cfg.Servers[0].Url != testDefaultBasePath {
		t.Errorf("expected server URL http://localhost:5001, got %q", cfg.Servers[0].Url)
	}
	if cfg.Servers[0].Description != "No description provided" {
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
// ServerUrl
// ---------------------------------------------------------------------------

func TestServerUrl_DefaultIndex(t *testing.T) {
	cfg := NewConfiguration()
	url, err := cfg.ServerUrl(0, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != testDefaultBasePath {
		t.Errorf("expected http://localhost:5001, got %q", url)
	}
}

func TestServerUrl_IndexOutOfRange_Negative(t *testing.T) {
	cfg := NewConfiguration()
	_, err := cfg.ServerUrl(-1, nil)
	if err == nil {
		t.Error("expected error for negative index")
	}
}

func TestServerUrl_IndexOutOfRange_TooLarge(t *testing.T) {
	cfg := NewConfiguration()
	_, err := cfg.ServerUrl(5, nil)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestServerUrl_VariableSubstitution(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = []ServerConfiguration{
		{
			Url:         "https://{host}:{port}/api",
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

	url, err := cfg.ServerUrl(0, map[string]string{"host": "example.com", "port": "443"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://example.com:443/api" {
		t.Errorf("expected https://example.com:443/api, got %q", url)
	}
}

func TestServerUrl_VariableSubstitution_DefaultValues(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = []ServerConfiguration{
		{
			Url:         "https://{host}:{port}/api",
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

	// Pass nil variables - should use defaults
	url, err := cfg.ServerUrl(0, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://localhost:8080/api" {
		t.Errorf("expected https://localhost:8080/api, got %q", url)
	}
}

func TestServerUrl_InvalidVariableValue(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = []ServerConfiguration{
		{
			Url:         "https://{host}/api",
			Description: "Test server",
			Variables: map[string]ServerVariable{
				"host": {
					DefaultValue: "localhost",
					EnumValues:   []string{"localhost", "example.com"},
				},
			},
		},
	}

	_, err := cfg.ServerUrl(0, map[string]string{"host": "invalid-host.com"})
	if err == nil {
		t.Error("expected error for invalid variable value")
	}
}

func TestServerUrl_NoEnumRestrictions(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = []ServerConfiguration{
		{
			Url:         "https://{host}/api",
			Description: "Test server",
			Variables: map[string]ServerVariable{
				"host": {
					DefaultValue: "localhost",
					EnumValues:   []string{}, // empty enum means any value is allowed
				},
			},
		},
	}

	url, err := cfg.ServerUrl(0, map[string]string{"host": "any-host.com"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://any-host.com/api" {
		t.Errorf("expected https://any-host.com/api, got %q", url)
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

func TestContextKey_Variables(t *testing.T) {
	if ContextOAuth2.String() != "auth token" {
		t.Errorf("expected 'auth token' for ContextOAuth2, got %q", ContextOAuth2.String())
	}
	if ContextBasicAuth.String() != "auth basic" {
		t.Errorf("expected 'auth basic' for ContextBasicAuth, got %q", ContextBasicAuth.String())
	}
	if ContextAccessToken.String() != "auth accesstoken" {
		t.Errorf("expected 'auth accesstoken' for ContextAccessToken, got %q", ContextAccessToken.String())
	}
	if ContextAPIKey.String() != "auth apikey" {
		t.Errorf("expected 'auth apikey' for ContextAPIKey, got %q", ContextAPIKey.String())
	}
}
