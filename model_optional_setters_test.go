package ditclient

import "testing"

// The existing accessor tests call Has*/GetOk* before Set* but skip the
// "value is set" branch of those accessors. These tests hit the
// "return true" / "return value, true" branch on every optional field.

func TestOptional_ApiError_GetCodeOk_AfterSet(t *testing.T) {
	e := NewApiError("m")
	e.SetCode("CODE")
	p, ok := e.GetCodeOk()
	if !ok {
		t.Error("expected ok=true")
	}
	if p == nil || *p != "CODE" {
		t.Errorf("expected *p=CODE, got %v", p)
	}
}

func TestOptional_ApiError_GetMessageOk_AfterSet(t *testing.T) {
	e := NewApiError("hello")
	p, ok := e.GetMessageOk()
	if !ok || p == nil || *p != "hello" {
		t.Errorf("expected ok/hello, got %v / %v", ok, p)
	}
}

func TestOptional_CommitStatus_GetErrorOk_AfterSet(t *testing.T) {
	s := NewCommitStatus(1, 1, 1, true)
	s.SetError("oops")
	if !s.HasError() {
		t.Error("expected HasError=true")
	}
	p, ok := s.GetErrorOk()
	if !ok || p == nil || *p != "oops" {
		t.Errorf("expected ok/oops, got %v / %v", ok, p)
	}
}

func TestOptional_ProgressEntry_HasGetMessage_AfterSet(t *testing.T) {
	p := NewProgressEntry(1, "T")
	p.SetMessage("hi")
	if !p.HasMessage() {
		t.Error("expected HasMessage=true after SetMessage")
	}
	if v, ok := p.GetMessageOk(); !ok || v == nil || *v != "hi" {
		t.Errorf("GetMessageOk: ok=%v v=%v", ok, v)
	}

	p.SetPercent(75)
	if !p.HasPercent() {
		t.Error("expected HasPercent=true after SetPercent")
	}
	if v, ok := p.GetPercentOk(); !ok || v == nil || *v != 75 {
		t.Errorf("GetPercentOk: ok=%v v=%v", ok, v)
	}
}

func TestOptional_RepositoryStatus_HasLastCommit_AfterSet(t *testing.T) {
	s := NewRepositoryStatus()
	s.SetLastCommit("abc")
	if !s.HasLastCommit() {
		t.Error("expected HasLastCommit=true")
	}
	if p, ok := s.GetLastCommitOk(); !ok || p == nil || *p != "abc" {
		t.Errorf("GetLastCommitOk: ok=%v p=%v", ok, p)
	}

	s.SetSourceCommit("xyz")
	if !s.HasSourceCommit() {
		t.Error("expected HasSourceCommit=true")
	}
	if p, ok := s.GetSourceCommitOk(); !ok || p == nil || *p != "xyz" {
		t.Errorf("GetSourceCommitOk: ok=%v p=%v", ok, p)
	}
}

func TestOptional_ApiError_HasCode_AfterSet(t *testing.T) {
	e := NewApiError("m")
	e.SetCode("c1")
	if !e.HasCode() {
		t.Error("expected HasCode=true")
	}
	e.SetDetails("d1")
	if !e.HasDetails() {
		t.Error("expected HasDetails=true")
	}
	if p, ok := e.GetDetailsOk(); !ok || p == nil || *p != "d1" {
		t.Errorf("GetDetailsOk: ok=%v p=%v", ok, p)
	}
}

func TestOptional_Volume_HasConfig_AfterSet(t *testing.T) {
	v := NewVolume("data", map[string]interface{}{})
	v.SetConfig(map[string]interface{}{"k": "v"})
	if !v.HasConfig() {
		t.Error("expected HasConfig=true")
	}
	if got, ok := v.GetConfigOk(); !ok || got == nil || got["k"] != "v" {
		t.Errorf("GetConfigOk: ok=%v got=%v", ok, got)
	}
}

func TestOptional_VolumeStatus_HasError_AfterSet(t *testing.T) {
	s := NewVolumeStatus("n", 1, 1, map[string]interface{}{}, true)
	s.SetError("err")
	if !s.HasError() {
		t.Error("expected HasError=true")
	}
	if p, ok := s.GetErrorOk(); !ok || p == nil || *p != "err" {
		t.Errorf("GetErrorOk: ok=%v p=%v", ok, p)
	}
}
