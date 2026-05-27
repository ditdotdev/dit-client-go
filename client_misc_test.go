package datadatdatclient

import (
	"context"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

// ---------------------------------------------------------------------------
// decode: pointer-to-pointer-to-os.File branch
// ---------------------------------------------------------------------------

func TestDecode_DoublePointerToFile(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var f *os.File
	if err := client.decode(&f, []byte("payload-pp"), "application/octet-stream"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected *os.File to be allocated")
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()
	got, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(got) != "payload-pp" {
		t.Errorf("expected payload-pp, got %q", string(got))
	}
}

// ---------------------------------------------------------------------------
// parameterAddToHeaderOrQuery: MappedNullable struct value (non-pointer) and
// a struct that doesn't implement MappedNullable.
// ---------------------------------------------------------------------------

func TestParameterAddToHeaderOrQuery_MappedNullableStruct(t *testing.T) {
	// A struct value (not pointer) implementing MappedNullable triggers the
	// "case reflect.Struct" -> "if t, ok := obj.(MappedNullable); ok" branch.
	q := url.Values{}
	c := Commit{Id: "abc", Properties: map[string]interface{}{}}
	parameterAddToHeaderOrQuery(q, "commit", c, "form", "")
	// The value will be flattened from ToMap; assertion is on the side-effect
	// existing — keys may be derived from properties' shape.
	if len(q) == 0 {
		t.Error("expected MappedNullable struct to populate query")
	}
}

func TestParameterAddToHeaderOrQuery_PlainStruct(t *testing.T) {
	type plain struct{ X int }
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "p", plain{X: 1}, "form", "")
	// Plain struct falls through to default `value = v.Type().String() + " value"`.
	if got := q.Get("p"); got == "" {
		t.Error("expected non-empty value for plain struct")
	}
}

// ---------------------------------------------------------------------------
// parameterValueToString – ToMap-error branch is hard to hit because all
// our models' ToMap returns nil; cover the "non-Ptr GetActualInstanceValue"
// branch via a value type lacking it (which just falls back to %v).
// ---------------------------------------------------------------------------

type actualValuer struct{}

func (actualValuer) GetActualInstanceValue() interface{} { return "deep" }

func TestParameterValueToString_GetActualInstanceValue(t *testing.T) {
	if got := parameterValueToString(actualValuer{}, "key"); got != "deep" {
		t.Errorf("expected deep, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// addFile via real file copy: covers the io.Copy success branch.
// ---------------------------------------------------------------------------

func TestAddFile_LargerContent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "big.txt")
	payload := make([]byte, 64*1024)
	for i := range payload {
		payload[i] = byte(i % 256)
	}
	if err := os.WriteFile(path, payload, 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	// Just verify that addFile completes without error on a non-trivial file.
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	client := NewAPIClient(cfg)
	headers := map[string]string{"Content-Type": "multipart/form-data"}
	form := url.Values{}
	form.Add("@upload", path) // form key prefixed with @ takes the file path

	req, err := client.prepareRequest(
		context.Background(),
		"http://example.com/path",
		"POST",
		nil,
		headers,
		url.Values{},
		form,
		nil,
	)
	if err != nil {
		t.Fatalf("prepareRequest: %v", err)
	}
	if req == nil {
		t.Fatal("expected request")
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if len(body) < len(payload) {
		t.Errorf("expected body to include file content; len=%d payload=%d", len(body), len(payload))
	}
}

// ---------------------------------------------------------------------------
// Cover IsNil for additional types.
// ---------------------------------------------------------------------------

func TestIsNil_Slice(t *testing.T) {
	var s []int
	if !IsNil(s) {
		t.Error("nil slice should be nil")
	}
	if IsNil([]int{1}) {
		t.Error("non-empty slice should not be nil")
	}
}

func TestIsNil_Map(t *testing.T) {
	var m map[string]int
	if !IsNil(m) {
		t.Error("nil map should be nil")
	}
}

func TestIsNil_Func(t *testing.T) {
	var f func()
	if !IsNil(f) {
		t.Error("nil func should be nil")
	}
}

func TestIsNil_Chan(t *testing.T) {
	var c chan int
	if !IsNil(c) {
		t.Error("nil chan should be nil")
	}
}

func TestIsNil_Array(t *testing.T) {
	if !IsNil([0]int{}) {
		t.Error("zero-length array should be IsZero / treated as nil")
	}
}

// ---------------------------------------------------------------------------
// parameterToJson – marshal error path. Channels can't be marshaled.
// ---------------------------------------------------------------------------

func TestParameterToJson_MarshalError(t *testing.T) {
	ch := make(chan int)
	_, err := parameterToJson(ch)
	if err == nil {
		t.Error("expected error marshaling a channel")
	}
}

// ---------------------------------------------------------------------------
// ToMap on Volume with Config populated covers the !IsNil(o.Config) branch.
// ---------------------------------------------------------------------------

func TestVolume_ToMap_WithConfig(t *testing.T) {
	v := NewVolume("data", map[string]interface{}{})
	v.SetConfig(map[string]interface{}{"a": 1})
	m, err := v.ToMap()
	if err != nil {
		t.Fatalf("ToMap: %v", err)
	}
	if _, ok := m["config"]; !ok {
		t.Error("expected config key in map when Config is set")
	}
}
