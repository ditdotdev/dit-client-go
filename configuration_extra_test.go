// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package ditclient

import (
	"context"
	"testing"
)

// ---------------------------------------------------------------------------
// getServerOperationIndex – the helper falls back to the global server index
// when nothing is in context for the endpoint. The invalid-type branch must
// return an error.
// ---------------------------------------------------------------------------

func TestServerURLWithContext_OperationServerIndexOverride(t *testing.T) {
	cfg := NewConfiguration()
	cfg.OperationServers = map[string]ServerConfigurations{
		"Endpoint": {{URL: "http://op-primary"}, {URL: "http://op-secondary"}},
	}
	ctx := context.WithValue(context.Background(), ContextOperationServerIndices,
		map[string]int{"Endpoint": 1})

	url, err := cfg.ServerURLWithContext(ctx, "Endpoint")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "http://op-secondary" {
		t.Errorf("expected http://op-secondary, got %q", url)
	}
}

func TestServerURLWithContext_OperationServerIndexFallback(t *testing.T) {
	// Endpoint absent from the operation indices map → fall back to
	// the global ContextServerIndex (which is also absent here).
	cfg := NewConfiguration()
	cfg.OperationServers = map[string]ServerConfigurations{
		"Endpoint": {{URL: "http://op-default"}},
	}
	ctx := context.WithValue(context.Background(), ContextOperationServerIndices,
		map[string]int{"OtherEndpoint": 1})

	url, err := cfg.ServerURLWithContext(ctx, "Endpoint")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "http://op-default" {
		t.Errorf("expected http://op-default, got %q", url)
	}
}

func TestServerURLWithContext_InvalidServerIndexType(t *testing.T) {
	cfg := NewConfiguration()
	ctx := context.WithValue(context.Background(), ContextServerIndex, "not-an-int")

	_, err := cfg.ServerURLWithContext(ctx, "endpoint")
	if err == nil {
		t.Fatal("expected error for invalid server index type")
	}
}

func TestServerURLWithContext_InvalidOperationServerIndicesType(t *testing.T) {
	cfg := NewConfiguration()
	ctx := context.WithValue(context.Background(), ContextOperationServerIndices, "not-a-map")

	_, err := cfg.ServerURLWithContext(ctx, "endpoint")
	if err == nil {
		t.Fatal("expected error for invalid operation server indices type")
	}
}

// ---------------------------------------------------------------------------
// getServerVariables / getServerOperationVariables
// ---------------------------------------------------------------------------

func TestServerURLWithContext_ServerVariables(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL: "https://{host}/api",
			Variables: map[string]ServerVariable{
				"host": {DefaultValue: "localhost"},
			},
		},
	}
	ctx := context.WithValue(context.Background(), ContextServerVariables,
		map[string]string{"host": "example.com"})

	url, err := cfg.ServerURLWithContext(ctx, "endpoint")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://example.com/api" {
		t.Errorf("expected https://example.com/api, got %q", url)
	}
}

func TestServerURLWithContext_OperationServerVariables(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL: "https://{host}/api",
			Variables: map[string]ServerVariable{
				"host": {DefaultValue: "localhost"},
			},
		},
	}
	ctx := context.WithValue(context.Background(), ContextOperationServerVariables,
		map[string]map[string]string{
			"endpoint": {"host": "ep.example.com"},
		})

	url, err := cfg.ServerURLWithContext(ctx, "endpoint")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://ep.example.com/api" {
		t.Errorf("expected https://ep.example.com/api, got %q", url)
	}
}

func TestServerURLWithContext_OperationServerVariablesFallback(t *testing.T) {
	// Endpoint not in operationVariables → fall back to ContextServerVariables.
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL: "https://{host}/api",
			Variables: map[string]ServerVariable{
				"host": {DefaultValue: "default"},
			},
		},
	}
	ctx := context.WithValue(context.Background(), ContextOperationServerVariables,
		map[string]map[string]string{
			"other": {"host": "other.example.com"},
		})
	ctx = context.WithValue(ctx, ContextServerVariables,
		map[string]string{"host": "fallback.example.com"})

	url, err := cfg.ServerURLWithContext(ctx, "endpoint")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://fallback.example.com/api" {
		t.Errorf("expected fallback URL, got %q", url)
	}
}

func TestServerURLWithContext_InvalidServerVariablesType(t *testing.T) {
	cfg := NewConfiguration()
	ctx := context.WithValue(context.Background(), ContextServerVariables, "not-a-map")

	_, err := cfg.ServerURLWithContext(ctx, "endpoint")
	if err == nil {
		t.Fatal("expected error for invalid ContextServerVariables type")
	}
}

func TestServerURLWithContext_InvalidOperationServerVariablesType(t *testing.T) {
	cfg := NewConfiguration()
	ctx := context.WithValue(context.Background(), ContextOperationServerVariables, "not-a-map")

	_, err := cfg.ServerURLWithContext(ctx, "endpoint")
	if err == nil {
		t.Fatal("expected error for invalid ContextOperationServerVariables type")
	}
}
