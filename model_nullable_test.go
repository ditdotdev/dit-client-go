package datadatdatclient

import (
	"encoding/json"
	"testing"
)

// Every model type gets a NullableX wrapper with the same Set/Get/IsSet/Unset
// + Marshal/Unmarshal surface. These tests exercise the wrapper for every
// type once so the boilerplate is actually executed.

func TestNullableCommitStatus_Lifecycle(t *testing.T) {
	val := NewCommitStatus(100, 80, 60, true)
	n := NewNullableCommitStatus(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return json.Marshal(n) },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableContext_Lifecycle(t *testing.T) {
	val := NewContext("docker", map[string]interface{}{})
	n := NewNullableContext(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableOperation_Lifecycle(t *testing.T) {
	val := NewOperation("op", "PUSH", "RUNNING", "remote", "c1")
	n := NewNullableOperation(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableProgressEntry_Lifecycle(t *testing.T) {
	val := NewProgressEntry(1, "MESSAGE")
	n := NewNullableProgressEntry(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableRemote_Lifecycle(t *testing.T) {
	val := NewRemote("s3", "origin", map[string]interface{}{})
	n := NewNullableRemote(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableRemoteParameters_Lifecycle(t *testing.T) {
	val := NewRemoteParameters("s3", map[string]interface{}{})
	n := NewNullableRemoteParameters(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableRepository_Lifecycle(t *testing.T) {
	val := NewRepository("repo", map[string]interface{}{})
	n := NewNullableRepository(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableRepositoryStatus_Lifecycle(t *testing.T) {
	val := NewRepositoryStatus()
	n := NewNullableRepositoryStatus(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableVolume_Lifecycle(t *testing.T) {
	val := NewVolume("data", map[string]interface{}{})
	n := NewNullableVolume(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

func TestNullableVolumeStatus_Lifecycle(t *testing.T) {
	val := NewVolumeStatus("data", 1024, 800, map[string]interface{}{}, true)
	n := NewNullableVolumeStatus(val)
	checkNullableLifecycle(t, n, func() (interface{}, error) { return n.MarshalJSON() },
		func(data []byte) error { return n.UnmarshalJSON(data) })
}

// ---------------------------------------------------------------------------

// nullable is the common interface every NullableX wrapper implements.
type nullable interface {
	IsSet() bool
	Unset()
}

func checkNullableLifecycle(t *testing.T, n nullable, marshal func() (interface{}, error), unmarshal func([]byte) error) {
	t.Helper()
	if !n.IsSet() {
		t.Error("expected IsSet=true after construction")
	}

	raw, err := marshal()
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	data, ok := raw.([]byte)
	if !ok {
		t.Fatalf("marshal returned %T, want []byte", raw)
	}
	if err := unmarshal(data); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !n.IsSet() {
		t.Error("expected IsSet=true after unmarshal")
	}

	n.Unset()
	if n.IsSet() {
		t.Error("expected IsSet=false after Unset")
	}
}

// ---------------------------------------------------------------------------
// XWithDefaults constructors — exercised separately to bump coverage on the
// generated NewXWithDefaults variants for every model.
// ---------------------------------------------------------------------------

func TestWithDefaultsConstructors_AllNonNil(t *testing.T) {
	if NewCommitStatusWithDefaults() == nil {
		t.Error("CommitStatus")
	}
	if NewContextWithDefaults() == nil {
		t.Error("Context")
	}
	if NewOperationWithDefaults() == nil {
		t.Error("Operation")
	}
	if NewProgressEntryWithDefaults() == nil {
		t.Error("ProgressEntry")
	}
	if NewRemoteWithDefaults() == nil {
		t.Error("Remote")
	}
	if NewRemoteParametersWithDefaults() == nil {
		t.Error("RemoteParameters")
	}
	if NewRepositoryWithDefaults() == nil {
		t.Error("Repository")
	}
	if NewRepositoryStatusWithDefaults() == nil {
		t.Error("RepositoryStatus")
	}
	if NewVolumeWithDefaults() == nil {
		t.Error("Volume")
	}
	if NewVolumeStatusWithDefaults() == nil {
		t.Error("VolumeStatus")
	}
}

// ---------------------------------------------------------------------------
// GetXOk variants on optional fields for the models we hadn't yet exercised
// ---------------------------------------------------------------------------

func TestModel_OkVariants(t *testing.T) {
	// ApiError.GetDetailsOk
	e := NewApiError("msg")
	if _, ok := e.GetDetailsOk(); ok {
		t.Error("expected GetDetailsOk ok=false on fresh ApiError")
	}
	e.SetDetails("trace")
	if _, ok := e.GetDetailsOk(); !ok {
		t.Error("expected GetDetailsOk ok=true after SetDetails")
	}

	// CommitStatus required-field GetXOk
	cs := NewCommitStatus(1, 2, 3, true)
	if _, ok := cs.GetLogicalSizeOk(); !ok {
		t.Error("CommitStatus.GetLogicalSizeOk")
	}
	if _, ok := cs.GetActualSizeOk(); !ok {
		t.Error("CommitStatus.GetActualSizeOk")
	}
	if _, ok := cs.GetUniqueSizeOk(); !ok {
		t.Error("CommitStatus.GetUniqueSizeOk")
	}
	if _, ok := cs.GetReadyOk(); !ok {
		t.Error("CommitStatus.GetReadyOk")
	}
	if _, ok := cs.GetErrorOk(); ok {
		t.Error("CommitStatus.GetErrorOk should be ok=false on fresh status")
	}
	cs.SetLogicalSize(99)
	cs.SetActualSize(88)
	cs.SetUniqueSize(77)
	if cs.GetLogicalSize() != 99 || cs.GetActualSize() != 88 || cs.GetUniqueSize() != 77 {
		t.Error("CommitStatus setters")
	}

	// Context required-field GetXOk
	ctx := NewContext("p", map[string]interface{}{})
	if _, ok := ctx.GetProviderOk(); !ok {
		t.Error("Context.GetProviderOk")
	}
	if _, ok := ctx.GetPropertiesOk(); !ok {
		t.Error("Context.GetPropertiesOk")
	}
	ctx.SetProperties(map[string]interface{}{"k": "v"})
	if ctx.GetProperties()["k"] != "v" {
		t.Error("Context.SetProperties")
	}

	// NullableCommit.Set on existing wrapper
	c1 := NewCommit("a", map[string]interface{}{})
	c2 := NewCommit("b", map[string]interface{}{})
	n := NewNullableCommit(c1)
	n.Set(c2)
	if n.Get().Id != "b" {
		t.Errorf("NullableCommit.Set didn't update value")
	}
}

// Exercise every remaining Get*Ok / Set* on Operation, ProgressEntry,
// Remote, RemoteParameters, Repository, Volume, VolumeStatus so the
// accessor boilerplate gets covered.

func TestModel_OperationAccessors(t *testing.T) {
	o := NewOperation("op", "PUSH", "RUNNING", "remote", "c1")
	if _, ok := o.GetIdOk(); !ok {
		t.Error("Operation.GetIdOk")
	}
	if _, ok := o.GetTypeOk(); !ok {
		t.Error("Operation.GetTypeOk")
	}
	if _, ok := o.GetStateOk(); !ok {
		t.Error("Operation.GetStateOk")
	}
	if _, ok := o.GetRemoteOk(); !ok {
		t.Error("Operation.GetRemoteOk")
	}
	if _, ok := o.GetCommitIdOk(); !ok {
		t.Error("Operation.GetCommitIdOk")
	}
	o.SetId("op2")
	o.SetType("PULL")
	o.SetRemote("origin2")
	o.SetCommitId("c2")
	if o.GetId() != "op2" || o.GetType() != "PULL" || o.GetRemote() != "origin2" || o.GetCommitId() != "c2" {
		t.Errorf("Operation setters not reflected in getters")
	}
	n := NewNullableOperation(o)
	n.Set(NewOperation("op3", "PUSH", "RUNNING", "remote", "c3"))
	if n.Get().GetId() != "op3" {
		t.Error("NullableOperation.Set/Get")
	}
}

func TestModel_ProgressEntryAccessors(t *testing.T) {
	p := NewProgressEntry(7, "PROGRESS")
	if _, ok := p.GetIdOk(); !ok {
		t.Error("ProgressEntry.GetIdOk")
	}
	if _, ok := p.GetTypeOk(); !ok {
		t.Error("ProgressEntry.GetTypeOk")
	}
	p.SetId(8)
	p.SetType("END")
	p.SetMessage("done")
	p.SetPercent(100)
	if p.GetId() != 8 || p.GetType() != "END" {
		t.Error("ProgressEntry setters")
	}
	if _, ok := p.GetMessageOk(); !ok {
		t.Error("ProgressEntry.GetMessageOk after SetMessage")
	}
	if _, ok := p.GetPercentOk(); !ok {
		t.Error("ProgressEntry.GetPercentOk after SetPercent")
	}
	n := NewNullableProgressEntry(p)
	n.Set(NewProgressEntry(9, "MESSAGE"))
	if n.Get().GetId() != 9 {
		t.Error("NullableProgressEntry.Set/Get")
	}
}

func TestModel_RemoteAccessors(t *testing.T) {
	r := NewRemote("s3", "origin", map[string]interface{}{"k": "v"})
	if _, ok := r.GetProviderOk(); !ok {
		t.Error("Remote.GetProviderOk")
	}
	if _, ok := r.GetNameOk(); !ok {
		t.Error("Remote.GetNameOk")
	}
	if got := r.GetProperties(); got["k"] != "v" {
		t.Error("Remote.GetProperties")
	}
	if _, ok := r.GetPropertiesOk(); !ok {
		t.Error("Remote.GetPropertiesOk")
	}
	r.SetProvider("ssh")
	r.SetProperties(map[string]interface{}{"updated": true})
	if r.GetProvider() != "ssh" || r.GetProperties()["updated"] != true {
		t.Error("Remote setters")
	}
	n := NewNullableRemote(r)
	n.Set(NewRemote("ssh", "backup", map[string]interface{}{}))
	if n.Get().GetName() != "backup" {
		t.Error("NullableRemote.Set/Get")
	}
}

func TestModel_RemoteParametersAccessors(t *testing.T) {
	rp := NewRemoteParameters("s3", map[string]interface{}{"region": "us-west-2"})
	if _, ok := rp.GetProviderOk(); !ok {
		t.Error("RemoteParameters.GetProviderOk")
	}
	if got := rp.GetProperties()["region"]; got != "us-west-2" {
		t.Error("RemoteParameters.GetProperties")
	}
	if _, ok := rp.GetPropertiesOk(); !ok {
		t.Error("RemoteParameters.GetPropertiesOk")
	}
	rp.SetProperties(map[string]interface{}{"bucket": "demo"})
	if rp.GetProperties()["bucket"] != "demo" {
		t.Error("RemoteParameters.SetProperties")
	}
	n := NewNullableRemoteParameters(rp)
	n.Set(NewRemoteParameters("ssh", map[string]interface{}{}))
	if n.Get().GetProvider() != "ssh" {
		t.Error("NullableRemoteParameters.Set/Get")
	}
}

func TestModel_RepositoryAccessors(t *testing.T) {
	r := NewRepository("alpha", map[string]interface{}{"owner": "team"})
	if _, ok := r.GetNameOk(); !ok {
		t.Error("Repository.GetNameOk")
	}
	if got := r.GetProperties()["owner"]; got != "team" {
		t.Error("Repository.GetProperties")
	}
	if _, ok := r.GetPropertiesOk(); !ok {
		t.Error("Repository.GetPropertiesOk")
	}
	r.SetProperties(map[string]interface{}{})
	if len(r.GetProperties()) != 0 {
		t.Error("Repository.SetProperties")
	}
	n := NewNullableRepository(r)
	n.Set(NewRepository("beta", map[string]interface{}{}))
	if n.Get().GetName() != "beta" {
		t.Error("NullableRepository.Set/Get")
	}
}

func TestModel_RepositoryStatusAccessors(t *testing.T) {
	s := NewRepositoryStatus()
	s.SetLastCommit("abc")
	s.SetSourceCommit("xyz")
	if _, ok := s.GetLastCommitOk(); !ok {
		t.Error("RepositoryStatus.GetLastCommitOk")
	}
	if _, ok := s.GetSourceCommitOk(); !ok {
		t.Error("RepositoryStatus.GetSourceCommitOk")
	}
	n := NewNullableRepositoryStatus(s)
	other := NewRepositoryStatus()
	other.SetLastCommit("other")
	n.Set(other)
	if n.Get().GetLastCommit() != "other" {
		t.Error("NullableRepositoryStatus.Set/Get")
	}
}

func TestModel_VolumeAccessors(t *testing.T) {
	v := NewVolume("data", map[string]interface{}{"mount": "/data"})
	if _, ok := v.GetNameOk(); !ok {
		t.Error("Volume.GetNameOk")
	}
	if _, ok := v.GetPropertiesOk(); !ok {
		t.Error("Volume.GetPropertiesOk")
	}
	v.SetConfig(map[string]interface{}{"cfg": "value"})
	if _, ok := v.GetConfigOk(); !ok {
		t.Error("Volume.GetConfigOk after SetConfig")
	}
	v.SetProperties(map[string]interface{}{"updated": true})
	if v.GetProperties()["updated"] != true {
		t.Error("Volume.SetProperties")
	}
	n := NewNullableVolume(v)
	n.Set(NewVolume("logs", map[string]interface{}{}))
	if n.Get().GetName() != "logs" {
		t.Error("NullableVolume.Set/Get")
	}
}

func TestModel_VolumeStatusAccessors(t *testing.T) {
	s := NewVolumeStatus("data", 1024, 800, map[string]interface{}{}, true)
	if _, ok := s.GetNameOk(); !ok {
		t.Error("VolumeStatus.GetNameOk")
	}
	if _, ok := s.GetLogicalSizeOk(); !ok {
		t.Error("VolumeStatus.GetLogicalSizeOk")
	}
	if _, ok := s.GetActualSizeOk(); !ok {
		t.Error("VolumeStatus.GetActualSizeOk")
	}
	if _, ok := s.GetPropertiesOk(); !ok {
		t.Error("VolumeStatus.GetPropertiesOk")
	}
	if _, ok := s.GetReadyOk(); !ok {
		t.Error("VolumeStatus.GetReadyOk")
	}
	if _, ok := s.GetErrorOk(); ok {
		t.Error("VolumeStatus.GetErrorOk should be ok=false initially")
	}
	s.SetName("data2")
	s.SetLogicalSize(2048)
	s.SetActualSize(1024)
	s.SetProperties(map[string]interface{}{"k": "v"})
	s.SetReady(false)
	s.SetError("issue")
	if s.GetName() != "data2" || s.GetLogicalSize() != 2048 || s.GetReady() != false || s.GetError() != "issue" {
		t.Errorf("VolumeStatus setters: %+v", s)
	}
	if _, ok := s.GetErrorOk(); !ok {
		t.Error("VolumeStatus.GetErrorOk after SetError")
	}
	n := NewNullableVolumeStatus(s)
	n.Set(NewVolumeStatus("logs", 0, 0, map[string]interface{}{}, false))
	if n.Get().GetName() != "logs" {
		t.Error("NullableVolumeStatus.Set/Get")
	}
}

func TestModel_CommitStatusNullableSet(t *testing.T) {
	a := NewCommitStatus(1, 1, 1, true)
	b := NewCommitStatus(2, 2, 2, false)
	n := NewNullableCommitStatus(a)
	n.Set(b)
	if n.Get().GetLogicalSize() != 2 {
		t.Error("NullableCommitStatus.Set/Get")
	}
}

func TestModel_ContextNullableSet(t *testing.T) {
	a := NewContext("p1", map[string]interface{}{})
	b := NewContext("p2", map[string]interface{}{})
	n := NewNullableContext(a)
	n.Set(b)
	if n.Get().GetProvider() != "p2" {
		t.Error("NullableContext.Set/Get")
	}
}
