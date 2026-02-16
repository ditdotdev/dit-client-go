package Datadatdatclient

import (
	"net/http"
	"testing"
)

// ---------------------------------------------------------------------------
// NewAPIResponse
// ---------------------------------------------------------------------------

func TestNewAPIResponse_WithResponse(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
	}

	resp := NewAPIResponse(httpResp)
	if resp.Response != httpResp {
		t.Error("expected Response to be set")
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
}

func TestNewAPIResponse_NilResponse(t *testing.T) {
	resp := NewAPIResponse(nil)
	if resp.Response != nil {
		t.Error("expected nil Response")
	}
}

func TestNewAPIResponse_EmptyFields(t *testing.T) {
	httpResp := &http.Response{StatusCode: 200}
	resp := NewAPIResponse(httpResp)

	if resp.Message != "" {
		t.Errorf("expected empty Message, got %q", resp.Message)
	}
	if resp.Operation != "" {
		t.Errorf("expected empty Operation, got %q", resp.Operation)
	}
	if resp.RequestURL != "" {
		t.Errorf("expected empty RequestURL, got %q", resp.RequestURL)
	}
	if resp.Method != "" {
		t.Errorf("expected empty Method, got %q", resp.Method)
	}
	if resp.Payload != nil {
		t.Error("expected nil Payload")
	}
}

// ---------------------------------------------------------------------------
// NewAPIResponseWithError
// ---------------------------------------------------------------------------

func TestNewAPIResponseWithError_Message(t *testing.T) {
	resp := NewAPIResponseWithError("something failed")
	if resp.Message != "something failed" {
		t.Errorf("expected error message 'something failed', got %q", resp.Message)
	}
}

func TestNewAPIResponseWithError_NilResponse(t *testing.T) {
	resp := NewAPIResponseWithError("error")
	if resp.Response != nil {
		t.Error("expected nil embedded Response")
	}
}

func TestNewAPIResponseWithError_EmptyMessage(t *testing.T) {
	resp := NewAPIResponseWithError("")
	if resp.Message != "" {
		t.Errorf("expected empty message, got %q", resp.Message)
	}
}

// ---------------------------------------------------------------------------
// APIResponse field setting
// ---------------------------------------------------------------------------

func TestAPIResponse_FieldAssignment(t *testing.T) {
	resp := NewAPIResponse(&http.Response{StatusCode: 201})
	resp.Message = "created"
	resp.Operation = "CreateCommit"
	resp.RequestURL = "http://localhost:5001/v1/repositories/test/commits"
	resp.Method = "POST"
	resp.Payload = []byte(`{"id":"abc123"}`)

	if resp.Message != "created" {
		t.Errorf("expected Message=created, got %q", resp.Message)
	}
	if resp.Operation != "CreateCommit" {
		t.Errorf("expected Operation=CreateCommit, got %q", resp.Operation)
	}
	if resp.RequestURL != "http://localhost:5001/v1/repositories/test/commits" {
		t.Errorf("expected RequestURL, got %q", resp.RequestURL)
	}
	if resp.Method != "POST" {
		t.Errorf("expected Method=POST, got %q", resp.Method)
	}
	if string(resp.Payload) != `{"id":"abc123"}` {
		t.Errorf("expected Payload, got %q", resp.Payload)
	}
}
