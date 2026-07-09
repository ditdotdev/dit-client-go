// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package ditclient

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

// utils.go ships with primitive NullableX wrappers (Bool, Int, Int32, Int64,
// Float32, Float64, String, Time) plus the PtrX helpers. Our generated
// models do not reference these wrappers (our spec has no nullable
// primitives), so they end up as 0%-covered dead-ish code unless we
// exercise them explicitly.

func TestPtrTime(t *testing.T) {
	now := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	if got := *PtrTime(now); !got.Equal(now) {
		t.Errorf("PtrTime: got %v, want %v", got, now)
	}
}

func TestNullableBool_Lifecycle(t *testing.T) {
	exerciseNullablePrimitive(t,
		NewNullableBool(PtrBool(true)),
		func() ([]byte, error) { return NewNullableBool(PtrBool(true)).MarshalJSON() },
		func(data []byte) error { return new(NullableBool).UnmarshalJSON(data) },
	)
}

func TestNullableInt_Lifecycle(t *testing.T) {
	exerciseNullablePrimitive(t,
		NewNullableInt(PtrInt(7)),
		func() ([]byte, error) { return NewNullableInt(PtrInt(7)).MarshalJSON() },
		func(data []byte) error { return new(NullableInt).UnmarshalJSON(data) },
	)
}

func TestNullableInt32_Lifecycle(t *testing.T) {
	exerciseNullablePrimitive(t,
		NewNullableInt32(PtrInt32(7)),
		func() ([]byte, error) { return NewNullableInt32(PtrInt32(7)).MarshalJSON() },
		func(data []byte) error { return new(NullableInt32).UnmarshalJSON(data) },
	)
}

func TestNullableInt64_Lifecycle(t *testing.T) {
	exerciseNullablePrimitive(t,
		NewNullableInt64(PtrInt64(7)),
		func() ([]byte, error) { return NewNullableInt64(PtrInt64(7)).MarshalJSON() },
		func(data []byte) error { return new(NullableInt64).UnmarshalJSON(data) },
	)
}

func TestNullableFloat32_Lifecycle(t *testing.T) {
	exerciseNullablePrimitive(t,
		NewNullableFloat32(PtrFloat32(1.5)),
		func() ([]byte, error) { return NewNullableFloat32(PtrFloat32(1.5)).MarshalJSON() },
		func(data []byte) error { return new(NullableFloat32).UnmarshalJSON(data) },
	)
}

func TestNullableFloat64_Lifecycle(t *testing.T) {
	exerciseNullablePrimitive(t,
		NewNullableFloat64(PtrFloat64(1.5)),
		func() ([]byte, error) { return NewNullableFloat64(PtrFloat64(1.5)).MarshalJSON() },
		func(data []byte) error { return new(NullableFloat64).UnmarshalJSON(data) },
	)
}

func TestNullableString_Lifecycle(t *testing.T) {
	exerciseNullablePrimitive(t,
		NewNullableString(PtrString("hi")),
		func() ([]byte, error) { return NewNullableString(PtrString("hi")).MarshalJSON() },
		func(data []byte) error { return new(NullableString).UnmarshalJSON(data) },
	)
}

func TestNullableTime_Lifecycle(t *testing.T) {
	now := time.Date(2026, 5, 21, 12, 0, 0, 0, time.UTC)
	n := NewNullableTime(&now)
	if !n.IsSet() {
		t.Error("expected IsSet=true")
	}
	if got := n.Get(); got == nil || !got.Equal(now) {
		t.Errorf("NullableTime.Get: %v", got)
	}
	other := now.Add(time.Hour)
	n.Set(&other)
	if got := n.Get(); !got.Equal(other) {
		t.Errorf("NullableTime.Set: %v", got)
	}
	data, err := n.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if !bytes.Contains(data, []byte("2026")) {
		t.Errorf("MarshalJSON: %s", data)
	}
	var n2 NullableTime
	if err := n2.UnmarshalJSON(data); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	if !n2.IsSet() {
		t.Error("expected IsSet=true after unmarshal")
	}
	n2.Unset()
	if n2.IsSet() {
		t.Error("expected IsSet=false after Unset")
	}
}

// ---------------------------------------------------------------------------

// nullablePrimitive is the shared interface for the primitive NullableX types.
type nullablePrimitive interface {
	IsSet() bool
	Unset()
}

func exerciseNullablePrimitive(
	t *testing.T,
	n nullablePrimitive,
	marshal func() ([]byte, error),
	unmarshal func([]byte) error,
) {
	t.Helper()
	if !n.IsSet() {
		t.Error("expected IsSet=true after construction")
	}
	data, err := marshal()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty marshalled output")
	}
	if err := unmarshal(data); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	n.Unset()
	if n.IsSet() {
		t.Error("expected IsSet=false after Unset")
	}
}

// ---------------------------------------------------------------------------
// Nullable*.Set / Get / Get-while-unset
// ---------------------------------------------------------------------------

func TestNullablePrimitive_GetSetAfterConstruction(t *testing.T) {
	n := NewNullableInt(PtrInt(5))
	if got := *n.Get(); got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
	v := 10
	n.Set(&v)
	if got := *n.Get(); got != 10 {
		t.Errorf("after Set: %d", got)
	}

	ns := NewNullableString(PtrString("a"))
	if got := *ns.Get(); got != "a" {
		t.Errorf("expected a, got %q", got)
	}
	s := "b"
	ns.Set(&s)
	if got := *ns.Get(); got != "b" {
		t.Errorf("after Set: %q", got)
	}

	nb := NewNullableBool(PtrBool(true))
	if got := *nb.Get(); !got {
		t.Errorf("expected true")
	}
	f := false
	nb.Set(&f)
	if got := *nb.Get(); got {
		t.Errorf("after Set: %v", got)
	}
}

func TestNullableInt32_GetSet(t *testing.T) {
	n := NewNullableInt32(PtrInt32(1))
	if got := *n.Get(); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
	v := int32(2)
	n.Set(&v)
	if got := *n.Get(); got != 2 {
		t.Errorf("after Set: %d", got)
	}
}

func TestNullableInt64_GetSet(t *testing.T) {
	n := NewNullableInt64(PtrInt64(int64(1)))
	if got := *n.Get(); got != int64(1) {
		t.Errorf("expected 1, got %d", got)
	}
	v := int64(2)
	n.Set(&v)
	if got := *n.Get(); got != int64(2) {
		t.Errorf("after Set: %d", got)
	}
}

func TestNullableFloat32_GetSet(t *testing.T) {
	n := NewNullableFloat32(PtrFloat32(1.0))
	v := float32(2.0)
	n.Set(&v)
	if got := *n.Get(); got != float32(2.0) {
		t.Errorf("after Set: %f", got)
	}
}

func TestNullableFloat64_GetSet(t *testing.T) {
	n := NewNullableFloat64(PtrFloat64(1.0))
	v := float64(2.0)
	n.Set(&v)
	if got := *n.Get(); got != float64(2.0) {
		t.Errorf("after Set: %f", got)
	}
}

func TestNullablePrimitive_UnmarshalJSON_Null(t *testing.T) {
	// Per the openapi-generator's nullable contract, "null" unmarshals to
	// isSet=true with a zero pointer (the field was present in JSON, but
	// explicitly null).
	var n NullableString
	if err := json.Unmarshal([]byte("null"), &n); err != nil {
		t.Fatalf("Unmarshal null: %v", err)
	}
	if !n.IsSet() {
		t.Error("expected IsSet=true after Unmarshal(null)")
	}
	if n.Get() != nil {
		t.Errorf("expected nil after Unmarshal(null), got %v", n.Get())
	}
}
