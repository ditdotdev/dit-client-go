package datadatdatclient

import (
	"bytes"
	"encoding/json"
	"testing"
)

// The openapi-generator emits a Get/GetOk/Has/Set accessor triple for every
// optional field on every model, plus a NewX / NewXWithDefaults constructor
// and a NullableX wrapper. These tests exercise the patterns once per model
// type to cover the emitted code.

// ---------------------------------------------------------------------------
// ApiError — optional Code, required Message, optional Details
// ---------------------------------------------------------------------------

func TestApiError_AccessorsAndMarshalling(t *testing.T) {
	e := NewApiError("required message")

	if e.GetMessage() != "required message" {
		t.Errorf("expected Message, got %q", e.GetMessage())
	}
	if _, ok := e.GetMessageOk(); !ok {
		t.Error("GetMessageOk should return ok=true for required field")
	}

	// Optional Code: empty by default
	if e.HasCode() {
		t.Error("expected HasCode=false on fresh ApiError")
	}
	if _, ok := e.GetCodeOk(); ok {
		t.Error("expected GetCodeOk ok=false")
	}
	e.SetCode("X1")
	if !e.HasCode() || e.GetCode() != "X1" {
		t.Errorf("after SetCode: HasCode=%v Code=%q", e.HasCode(), e.GetCode())
	}

	// Optional Details
	e.SetDetails("trace")
	if !e.HasDetails() || e.GetDetails() != "trace" {
		t.Errorf("after SetDetails: HasDetails=%v Details=%q", e.HasDetails(), e.GetDetails())
	}

	e.SetMessage("updated")
	if e.GetMessage() != "updated" {
		t.Errorf("after SetMessage: %q", e.GetMessage())
	}

	// MarshalJSON / UnmarshalJSON round-trip
	data, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	var decoded ApiError
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if decoded.GetMessage() != "updated" || decoded.GetCode() != "X1" {
		t.Errorf("roundtrip lost fields: %+v", decoded)
	}
}

func TestApiError_WithDefaults(t *testing.T) {
	e := NewApiErrorWithDefaults()
	if e.HasCode() || e.HasDetails() {
		t.Error("expected fresh defaults")
	}
}

func TestApiError_UnmarshalRequiredMissing(t *testing.T) {
	// "message" is required — missing should fail decode.
	var e ApiError
	if err := json.Unmarshal([]byte(`{}`), &e); err == nil {
		t.Error("expected error when required 'message' is missing")
	}
}

func TestApiError_NilGetters(t *testing.T) {
	var e *ApiError
	if got := e.GetCode(); got != "" {
		t.Errorf("expected zero value from nil receiver, got %q", got)
	}
	if got := e.GetMessage(); got != "" {
		t.Errorf("expected zero value from nil receiver, got %q", got)
	}
	if _, ok := e.GetCodeOk(); ok {
		t.Error("nil receiver GetCodeOk should return ok=false")
	}
	if e.HasCode() {
		t.Error("nil receiver HasCode should be false")
	}
}

// ---------------------------------------------------------------------------
// NullableApiError
// ---------------------------------------------------------------------------

func TestNullableApiError_SetGetUnset(t *testing.T) {
	e := NewApiError("msg")
	n := NewNullableApiError(e)

	if !n.IsSet() {
		t.Error("expected IsSet=true after construction")
	}
	if n.Get() != e {
		t.Error("expected Get to return the original pointer")
	}

	n.Unset()
	if n.IsSet() {
		t.Error("expected IsSet=false after Unset")
	}
	if n.Get() != nil {
		t.Error("expected Get=nil after Unset")
	}

	n.Set(e)
	if !n.IsSet() {
		t.Error("expected IsSet=true after Set")
	}

	data, err := n.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	if !bytes.Contains(data, []byte(`"message":"msg"`)) {
		t.Errorf("expected message:msg in JSON, got %s", data)
	}

	var n2 NullableApiError
	if err := n2.UnmarshalJSON(data); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if !n2.IsSet() || n2.Get().GetMessage() != "msg" {
		t.Errorf("roundtrip failed: %+v", n2)
	}
}

// ---------------------------------------------------------------------------
// Commit — all-required fields
// ---------------------------------------------------------------------------

func TestCommit_Accessors(t *testing.T) {
	c := NewCommit("abc", map[string]interface{}{"k": "v"})

	if c.GetId() != "abc" {
		t.Errorf("expected GetId=abc, got %q", c.GetId())
	}
	if _, ok := c.GetIdOk(); !ok {
		t.Error("expected GetIdOk ok=true")
	}
	if got := c.GetProperties()["k"]; got != "v" {
		t.Errorf("expected properties[k]=v, got %v", got)
	}
	if _, ok := c.GetPropertiesOk(); !ok {
		t.Error("expected GetPropertiesOk ok=true")
	}

	c.SetId("xyz")
	if c.GetId() != "xyz" {
		t.Errorf("after SetId: %q", c.GetId())
	}
	c.SetProperties(map[string]interface{}{"new": true})
	if c.GetProperties()["new"] != true {
		t.Errorf("after SetProperties: %v", c.GetProperties())
	}
}

func TestCommit_WithDefaults(t *testing.T) {
	c := NewCommitWithDefaults()
	if c == nil {
		t.Fatal("expected non-nil")
	}
}

func TestCommit_NilGetters(t *testing.T) {
	var c *Commit
	if got := c.GetId(); got != "" {
		t.Errorf("expected empty from nil receiver, got %q", got)
	}
	if got := c.GetProperties(); got != nil {
		t.Errorf("expected nil properties from nil receiver, got %v", got)
	}
	if _, ok := c.GetIdOk(); ok {
		t.Error("nil receiver GetIdOk should return ok=false")
	}
}

func TestCommit_RoundTrip(t *testing.T) {
	c := NewCommit("abc", map[string]interface{}{"tag": "v1"})
	data, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var decoded Commit
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded.Id != "abc" {
		t.Errorf("roundtrip lost Id: %q", decoded.Id)
	}
}

func TestCommit_UnmarshalRequiredMissing(t *testing.T) {
	var c Commit
	if err := json.Unmarshal([]byte(`{}`), &c); err == nil {
		t.Error("expected error when required fields missing")
	}
}

func TestNullableCommit_RoundTrip(t *testing.T) {
	c := NewCommit("abc", map[string]interface{}{})
	n := NewNullableCommit(c)
	data, err := n.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	var n2 NullableCommit
	if err := n2.UnmarshalJSON(data); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	if !n2.IsSet() || n2.Get().Id != "abc" {
		t.Errorf("roundtrip failed: %+v", n2)
	}
	n2.Unset()
	if n2.IsSet() {
		t.Error("expected IsSet=false after Unset")
	}
}

// ---------------------------------------------------------------------------
// CommitStatus — all required except optional Error
// ---------------------------------------------------------------------------

func TestCommitStatus_Accessors(t *testing.T) {
	s := NewCommitStatus(1024, 800, 600, true)

	if s.GetLogicalSize() != 1024 || s.GetActualSize() != 800 || s.GetUniqueSize() != 600 {
		t.Errorf("size accessors wrong: logical=%d actual=%d unique=%d", s.GetLogicalSize(), s.GetActualSize(), s.GetUniqueSize())
	}
	if !s.GetReady() {
		t.Error("expected Ready=true")
	}
	if s.HasError() {
		t.Error("expected HasError=false on fresh CommitStatus")
	}
	s.SetError("oops")
	if !s.HasError() || s.GetError() != "oops" {
		t.Errorf("after SetError: %v %q", s.HasError(), s.GetError())
	}

	s.SetReady(false)
	if s.GetReady() {
		t.Error("expected Ready=false after SetReady(false)")
	}
}

// ---------------------------------------------------------------------------
// Context — all required
// ---------------------------------------------------------------------------

func TestContext_Accessors(t *testing.T) {
	c := NewContext("docker", map[string]interface{}{"sock": "/var/run/docker.sock"})
	if c.GetProvider() != "docker" {
		t.Errorf("expected provider=docker, got %q", c.GetProvider())
	}
	if c.GetProperties()["sock"] != "/var/run/docker.sock" {
		t.Errorf("properties not preserved: %v", c.GetProperties())
	}
	c.SetProvider("kube")
	if c.GetProvider() != "kube" {
		t.Errorf("after SetProvider: %q", c.GetProvider())
	}
}

// ---------------------------------------------------------------------------
// Operation — all required
// ---------------------------------------------------------------------------

func TestOperation_Accessors(t *testing.T) {
	op := NewOperation("op-1", "PUSH", "RUNNING", "origin", "abc")
	if op.GetId() != "op-1" || op.GetType() != "PUSH" || op.GetState() != "RUNNING" {
		t.Errorf("operation accessors wrong: %+v", op)
	}
	if op.GetRemote() != "origin" || op.GetCommitId() != "abc" {
		t.Errorf("remote/commitId wrong: %+v", op)
	}
	op.SetState("COMPLETE")
	if op.GetState() != "COMPLETE" {
		t.Errorf("after SetState: %q", op.GetState())
	}
}

// ---------------------------------------------------------------------------
// ProgressEntry — Id+Type required, Message+Percent optional
// ---------------------------------------------------------------------------

func TestProgressEntry_Accessors(t *testing.T) {
	p := NewProgressEntry(7, "PROGRESS")
	if p.GetId() != 7 || p.GetType() != "PROGRESS" {
		t.Errorf("required field accessors wrong")
	}

	if p.HasMessage() || p.HasPercent() {
		t.Error("expected optional fields to start unset")
	}
	p.SetMessage("halfway")
	p.SetPercent(50)
	if p.GetMessage() != "halfway" || p.GetPercent() != 50 {
		t.Errorf("after Set: msg=%q pct=%d", p.GetMessage(), p.GetPercent())
	}
}

// ---------------------------------------------------------------------------
// Remote / RemoteParameters
// ---------------------------------------------------------------------------

func TestRemote_Accessors(t *testing.T) {
	r := NewRemote("s3", "origin", map[string]interface{}{"bucket": "demo"})
	if r.GetProvider() != "s3" || r.GetName() != "origin" {
		t.Errorf("accessors wrong: %+v", r)
	}
	r.SetName("backup")
	if r.GetName() != "backup" {
		t.Errorf("after SetName: %q", r.GetName())
	}
}

func TestRemoteParameters_Accessors(t *testing.T) {
	rp := NewRemoteParameters("s3", map[string]interface{}{"region": "us-west-2"})
	if rp.GetProvider() != "s3" {
		t.Errorf("expected s3, got %q", rp.GetProvider())
	}
	rp.SetProvider("ssh")
	if rp.GetProvider() != "ssh" {
		t.Errorf("after SetProvider: %q", rp.GetProvider())
	}
}

// ---------------------------------------------------------------------------
// Repository / RepositoryStatus
// ---------------------------------------------------------------------------

func TestRepository_Accessors(t *testing.T) {
	r := NewRepository("alpha", map[string]interface{}{})
	if r.GetName() != "alpha" {
		t.Errorf("expected alpha, got %q", r.GetName())
	}
	r.SetName("beta")
	if r.GetName() != "beta" {
		t.Errorf("after SetName: %q", r.GetName())
	}
}

func TestRepositoryStatus_Accessors(t *testing.T) {
	s := NewRepositoryStatus()
	if s.HasLastCommit() || s.HasSourceCommit() {
		t.Error("expected fresh RepositoryStatus to have no commits")
	}
	s.SetLastCommit("abc")
	s.SetSourceCommit("xyz")
	if s.GetLastCommit() != "abc" || s.GetSourceCommit() != "xyz" {
		t.Errorf("accessors wrong: %+v", s)
	}
}

// ---------------------------------------------------------------------------
// Volume / VolumeStatus
// ---------------------------------------------------------------------------

func TestVolume_Accessors(t *testing.T) {
	v := NewVolume("data", map[string]interface{}{})
	if v.GetName() != "data" {
		t.Errorf("expected data, got %q", v.GetName())
	}
	if v.HasConfig() {
		t.Error("expected HasConfig=false")
	}
	v.SetConfig(map[string]interface{}{"key": "value"})
	if !v.HasConfig() {
		t.Error("expected HasConfig=true after SetConfig")
	}
	if got := v.GetConfig(); got["key"] != "value" {
		t.Errorf("GetConfig: %v", got)
	}
	v.SetName("data2")
	if v.GetName() != "data2" {
		t.Errorf("SetName not applied: %q", v.GetName())
	}
}

func TestVolumeStatus_Accessors(t *testing.T) {
	s := NewVolumeStatus("data", 1024, 800, map[string]interface{}{"k": "v"}, true)
	if s.GetName() != "data" || s.GetLogicalSize() != 1024 || !s.GetReady() {
		t.Errorf("accessors wrong: %+v", s)
	}
	if s.GetActualSize() != 800 {
		t.Errorf("expected ActualSize=800, got %d", s.GetActualSize())
	}
	if got := s.GetProperties()["k"]; got != "v" {
		t.Errorf("GetProperties: %v", got)
	}
	if s.HasError() {
		t.Error("expected HasError=false initially")
	}
	s.SetError("disk full")
	if !s.HasError() || s.GetError() != "disk full" {
		t.Errorf("after SetError: %v %q", s.HasError(), s.GetError())
	}
}

// ---------------------------------------------------------------------------
// Ptr helpers from utils.go
// ---------------------------------------------------------------------------

func TestPtrHelpers(t *testing.T) {
	if *PtrBool(true) != true {
		t.Error("PtrBool(true)")
	}
	if *PtrInt(42) != 42 {
		t.Error("PtrInt(42)")
	}
	if *PtrInt32(int32(5)) != int32(5) {
		t.Error("PtrInt32(5)")
	}
	if *PtrInt64(int64(99)) != int64(99) {
		t.Error("PtrInt64(99)")
	}
	if *PtrFloat32(float32(1.5)) != float32(1.5) {
		t.Error("PtrFloat32(1.5)")
	}
	if *PtrFloat64(float64(2.5)) != float64(2.5) {
		t.Error("PtrFloat64(2.5)")
	}
	if *PtrString("hi") != "hi" {
		t.Error("PtrString(hi)")
	}
}

func TestIsNil(t *testing.T) {
	if !IsNil(nil) {
		t.Error("IsNil(nil) should be true")
	}
	var p *int
	if !IsNil(p) {
		t.Error("IsNil(nil-pointer) should be true")
	}
	x := 5
	if IsNil(&x) {
		t.Error("IsNil(&5) should be false")
	}
	if IsNil(5) {
		t.Error("IsNil(5) should be false")
	}
}
