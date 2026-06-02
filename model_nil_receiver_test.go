package ditclient

import "testing"

// These tests exercise the "if o == nil { return zero }" branch of every
// generated Get / GetOk / Has accessor across every model. The openapi
// generator emits these checks defensively, so they are part of the
// statement count even though they are never hit in production code.
//
// Each test calls the accessor on a typed nil pointer and asserts the
// zero-value contract documented by the generated comments.

// ---------------------------------------------------------------------------
// ApiError
// ---------------------------------------------------------------------------

func TestNilReceiver_ApiError(t *testing.T) {
	var e *ApiError

	if got := e.GetCode(); got != "" {
		t.Errorf("GetCode: expected zero, got %q", got)
	}
	if p, ok := e.GetCodeOk(); p != nil || ok {
		t.Errorf("GetCodeOk: expected nil/false, got %v/%v", p, ok)
	}
	if e.HasCode() {
		t.Error("HasCode: expected false")
	}

	if got := e.GetMessage(); got != "" {
		t.Errorf("GetMessage: expected zero, got %q", got)
	}
	if p, ok := e.GetMessageOk(); p != nil || ok {
		t.Errorf("GetMessageOk: expected nil/false")
	}

	if got := e.GetDetails(); got != "" {
		t.Errorf("GetDetails: expected zero, got %q", got)
	}
	if p, ok := e.GetDetailsOk(); p != nil || ok {
		t.Errorf("GetDetailsOk: expected nil/false")
	}
	if e.HasDetails() {
		t.Error("HasDetails: expected false")
	}
}

// ---------------------------------------------------------------------------
// Commit
// ---------------------------------------------------------------------------

func TestNilReceiver_Commit(t *testing.T) {
	var c *Commit

	if got := c.GetId(); got != "" {
		t.Errorf("GetId: %q", got)
	}
	if p, ok := c.GetIdOk(); p != nil || ok {
		t.Errorf("GetIdOk")
	}
	if got := c.GetProperties(); got != nil {
		t.Errorf("GetProperties: %v", got)
	}
	if _, ok := c.GetPropertiesOk(); ok {
		t.Error("GetPropertiesOk: expected ok=false")
	}
}

// ---------------------------------------------------------------------------
// CommitStatus
// ---------------------------------------------------------------------------

func TestNilReceiver_CommitStatus(t *testing.T) {
	var s *CommitStatus

	if got := s.GetLogicalSize(); got != 0 {
		t.Errorf("GetLogicalSize: %d", got)
	}
	if _, ok := s.GetLogicalSizeOk(); ok {
		t.Error("GetLogicalSizeOk: expected ok=false")
	}

	if got := s.GetActualSize(); got != 0 {
		t.Errorf("GetActualSize: %d", got)
	}
	if _, ok := s.GetActualSizeOk(); ok {
		t.Error("GetActualSizeOk: expected ok=false")
	}

	if got := s.GetUniqueSize(); got != 0 {
		t.Errorf("GetUniqueSize: %d", got)
	}
	if _, ok := s.GetUniqueSizeOk(); ok {
		t.Error("GetUniqueSizeOk: expected ok=false")
	}

	if got := s.GetReady(); got {
		t.Errorf("GetReady: %v", got)
	}
	if _, ok := s.GetReadyOk(); ok {
		t.Error("GetReadyOk: expected ok=false")
	}

	if got := s.GetError(); got != "" {
		t.Errorf("GetError: %q", got)
	}
	if p, ok := s.GetErrorOk(); p != nil || ok {
		t.Error("GetErrorOk: expected nil/false")
	}
	if s.HasError() {
		t.Error("HasError: expected false")
	}
}

// ---------------------------------------------------------------------------
// Context
// ---------------------------------------------------------------------------

func TestNilReceiver_Context(t *testing.T) {
	var c *Context

	if got := c.GetProvider(); got != "" {
		t.Errorf("GetProvider: %q", got)
	}
	if p, ok := c.GetProviderOk(); p != nil || ok {
		t.Error("GetProviderOk: expected nil/false")
	}

	if got := c.GetProperties(); got != nil {
		t.Errorf("GetProperties: %v", got)
	}
	if _, ok := c.GetPropertiesOk(); ok {
		t.Error("GetPropertiesOk: expected ok=false")
	}
}

// ---------------------------------------------------------------------------
// Operation
// ---------------------------------------------------------------------------

func TestNilReceiver_Operation(t *testing.T) {
	var o *Operation

	if o.GetId() != "" {
		t.Error("GetId")
	}
	if p, ok := o.GetIdOk(); p != nil || ok {
		t.Error("GetIdOk")
	}
	if o.GetType() != "" {
		t.Error("GetType")
	}
	if p, ok := o.GetTypeOk(); p != nil || ok {
		t.Error("GetTypeOk")
	}
	if o.GetState() != "" {
		t.Error("GetState")
	}
	if p, ok := o.GetStateOk(); p != nil || ok {
		t.Error("GetStateOk")
	}
	if o.GetRemote() != "" {
		t.Error("GetRemote")
	}
	if p, ok := o.GetRemoteOk(); p != nil || ok {
		t.Error("GetRemoteOk")
	}
	if o.GetCommitId() != "" {
		t.Error("GetCommitId")
	}
	if p, ok := o.GetCommitIdOk(); p != nil || ok {
		t.Error("GetCommitIdOk")
	}
}

// ---------------------------------------------------------------------------
// ProgressEntry
// ---------------------------------------------------------------------------

func TestNilReceiver_ProgressEntry(t *testing.T) {
	var p *ProgressEntry

	if p.GetId() != 0 {
		t.Error("GetId")
	}
	if v, ok := p.GetIdOk(); v != nil || ok {
		t.Error("GetIdOk")
	}
	if p.GetType() != "" {
		t.Error("GetType")
	}
	if v, ok := p.GetTypeOk(); v != nil || ok {
		t.Error("GetTypeOk")
	}
	if p.GetMessage() != "" {
		t.Error("GetMessage")
	}
	if v, ok := p.GetMessageOk(); v != nil || ok {
		t.Error("GetMessageOk")
	}
	if p.HasMessage() {
		t.Error("HasMessage")
	}
	if p.GetPercent() != 0 {
		t.Error("GetPercent")
	}
	if v, ok := p.GetPercentOk(); v != nil || ok {
		t.Error("GetPercentOk")
	}
	if p.HasPercent() {
		t.Error("HasPercent")
	}
}

// ---------------------------------------------------------------------------
// Remote
// ---------------------------------------------------------------------------

func TestNilReceiver_Remote(t *testing.T) {
	var r *Remote

	if r.GetProvider() != "" {
		t.Error("GetProvider")
	}
	if p, ok := r.GetProviderOk(); p != nil || ok {
		t.Error("GetProviderOk")
	}
	if r.GetName() != "" {
		t.Error("GetName")
	}
	if p, ok := r.GetNameOk(); p != nil || ok {
		t.Error("GetNameOk")
	}
	if got := r.GetProperties(); got != nil {
		t.Errorf("GetProperties: %v", got)
	}
	if _, ok := r.GetPropertiesOk(); ok {
		t.Error("GetPropertiesOk")
	}
}

// ---------------------------------------------------------------------------
// RemoteParameters
// ---------------------------------------------------------------------------

func TestNilReceiver_RemoteParameters(t *testing.T) {
	var rp *RemoteParameters

	if rp.GetProvider() != "" {
		t.Error("GetProvider")
	}
	if p, ok := rp.GetProviderOk(); p != nil || ok {
		t.Error("GetProviderOk")
	}
	if got := rp.GetProperties(); got != nil {
		t.Errorf("GetProperties: %v", got)
	}
	if _, ok := rp.GetPropertiesOk(); ok {
		t.Error("GetPropertiesOk")
	}
}

// ---------------------------------------------------------------------------
// Repository
// ---------------------------------------------------------------------------

func TestNilReceiver_Repository(t *testing.T) {
	var r *Repository

	if r.GetName() != "" {
		t.Error("GetName")
	}
	if p, ok := r.GetNameOk(); p != nil || ok {
		t.Error("GetNameOk")
	}
	if got := r.GetProperties(); got != nil {
		t.Errorf("GetProperties: %v", got)
	}
	if _, ok := r.GetPropertiesOk(); ok {
		t.Error("GetPropertiesOk")
	}
}

// ---------------------------------------------------------------------------
// RepositoryStatus (all-optional)
// ---------------------------------------------------------------------------

func TestNilReceiver_RepositoryStatus(t *testing.T) {
	var s *RepositoryStatus

	if s.GetLastCommit() != "" {
		t.Error("GetLastCommit")
	}
	if p, ok := s.GetLastCommitOk(); p != nil || ok {
		t.Error("GetLastCommitOk")
	}
	if s.HasLastCommit() {
		t.Error("HasLastCommit")
	}
	if s.GetSourceCommit() != "" {
		t.Error("GetSourceCommit")
	}
	if p, ok := s.GetSourceCommitOk(); p != nil || ok {
		t.Error("GetSourceCommitOk")
	}
	if s.HasSourceCommit() {
		t.Error("HasSourceCommit")
	}
}

// ---------------------------------------------------------------------------
// Volume (Config optional)
// ---------------------------------------------------------------------------

func TestNilReceiver_Volume(t *testing.T) {
	var v *Volume

	if v.GetName() != "" {
		t.Error("GetName")
	}
	if p, ok := v.GetNameOk(); p != nil || ok {
		t.Error("GetNameOk")
	}
	if got := v.GetProperties(); got != nil {
		t.Errorf("GetProperties: %v", got)
	}
	if _, ok := v.GetPropertiesOk(); ok {
		t.Error("GetPropertiesOk")
	}
	if got := v.GetConfig(); got != nil {
		t.Errorf("GetConfig: %v", got)
	}
	if _, ok := v.GetConfigOk(); ok {
		t.Error("GetConfigOk")
	}
	if v.HasConfig() {
		t.Error("HasConfig")
	}
}

// ---------------------------------------------------------------------------
// VolumeStatus
// ---------------------------------------------------------------------------

func TestNilReceiver_VolumeStatus(t *testing.T) {
	var s *VolumeStatus

	if s.GetName() != "" {
		t.Error("GetName")
	}
	if p, ok := s.GetNameOk(); p != nil || ok {
		t.Error("GetNameOk")
	}
	if s.GetLogicalSize() != 0 {
		t.Error("GetLogicalSize")
	}
	if _, ok := s.GetLogicalSizeOk(); ok {
		t.Error("GetLogicalSizeOk")
	}
	if s.GetActualSize() != 0 {
		t.Error("GetActualSize")
	}
	if _, ok := s.GetActualSizeOk(); ok {
		t.Error("GetActualSizeOk")
	}
	if got := s.GetProperties(); got != nil {
		t.Errorf("GetProperties: %v", got)
	}
	if _, ok := s.GetPropertiesOk(); ok {
		t.Error("GetPropertiesOk")
	}
	if s.GetReady() {
		t.Error("GetReady")
	}
	if _, ok := s.GetReadyOk(); ok {
		t.Error("GetReadyOk")
	}
	if s.GetError() != "" {
		t.Error("GetError")
	}
	if p, ok := s.GetErrorOk(); p != nil || ok {
		t.Error("GetErrorOk")
	}
	if s.HasError() {
		t.Error("HasError")
	}
}
