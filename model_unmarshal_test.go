package ditclient

import (
	"encoding/json"
	"testing"
)

// Each generated UnmarshalJSON has three error paths:
//   1. json.Unmarshal of the raw bytes into a map fails (e.g. not JSON)
//   2. a required property is missing from the map
//   3. the strict decoder rejects an unknown field
//
// The first and third paths are not exercised by the standard "missing
// required field" tests in model_test.go. This file fills those gaps.

func testUnmarshal_BadJSON[T any](t *testing.T, name string) {
	t.Helper()
	var v T
	if err := json.Unmarshal([]byte(`not-json`), &v); err == nil {
		t.Errorf("%s: expected unmarshal error on bad JSON", name)
	}
}

func TestUnmarshal_BadJSON_AllModels(t *testing.T) {
	testUnmarshal_BadJSON[ApiError](t, "ApiError")
	testUnmarshal_BadJSON[Commit](t, "Commit")
	testUnmarshal_BadJSON[CommitStatus](t, "CommitStatus")
	testUnmarshal_BadJSON[Context](t, "Context")
	testUnmarshal_BadJSON[Operation](t, "Operation")
	testUnmarshal_BadJSON[ProgressEntry](t, "ProgressEntry")
	testUnmarshal_BadJSON[Remote](t, "Remote")
	testUnmarshal_BadJSON[RemoteParameters](t, "RemoteParameters")
	testUnmarshal_BadJSON[Repository](t, "Repository")
	testUnmarshal_BadJSON[Volume](t, "Volume")
	testUnmarshal_BadJSON[VolumeStatus](t, "VolumeStatus")
}

func TestUnmarshal_UnknownField_AllModels(t *testing.T) {
	// Each payload contains all required fields plus one unexpected field
	// "_unknown" — DisallowUnknownFields() in the second-pass decoder must
	// reject this and propagate the error.
	cases := []struct {
		name      string
		body      string
		unmarshal func([]byte) error
	}{
		{
			name:      "ApiError",
			body:      `{"message":"m","_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v ApiError; return v.UnmarshalJSON(b) },
		},
		{
			name:      "Commit",
			body:      `{"id":"a","properties":{},"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v Commit; return v.UnmarshalJSON(b) },
		},
		{
			name:      "CommitStatus",
			body:      `{"logicalSize":1,"actualSize":1,"uniqueSize":1,"ready":true,"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v CommitStatus; return v.UnmarshalJSON(b) },
		},
		{
			name:      "Context",
			body:      `{"provider":"p","properties":{},"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v Context; return v.UnmarshalJSON(b) },
		},
		{
			name:      "Operation",
			body:      `{"id":"a","type":"PUSH","state":"R","remote":"r","commitId":"c","_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v Operation; return v.UnmarshalJSON(b) },
		},
		{
			name:      "ProgressEntry",
			body:      `{"id":1,"type":"T","_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v ProgressEntry; return v.UnmarshalJSON(b) },
		},
		{
			name:      "Remote",
			body:      `{"provider":"p","name":"n","properties":{},"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v Remote; return v.UnmarshalJSON(b) },
		},
		{
			name:      "RemoteParameters",
			body:      `{"provider":"p","properties":{},"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v RemoteParameters; return v.UnmarshalJSON(b) },
		},
		{
			name:      "Repository",
			body:      `{"name":"n","properties":{},"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v Repository; return v.UnmarshalJSON(b) },
		},
		{
			name:      "Volume",
			body:      `{"name":"n","properties":{},"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v Volume; return v.UnmarshalJSON(b) },
		},
		{
			name:      "VolumeStatus",
			body:      `{"name":"n","logicalSize":1,"actualSize":1,"properties":{},"ready":true,"_unknown":"x"}`,
			unmarshal: func(b []byte) error { var v VolumeStatus; return v.UnmarshalJSON(b) },
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if err := c.unmarshal([]byte(c.body)); err == nil {
				t.Error("expected error from strict decoder for unknown field")
			}
		})
	}
}

// Required-field missing covered for non-ApiError models, since
// TestApiError_UnmarshalRequiredMissing already covers ApiError.
func TestUnmarshal_RequiredMissing_AllModels(t *testing.T) {
	type unmarshalFn func([]byte) error
	cases := []struct {
		name string
		fn   unmarshalFn
	}{
		{"CommitStatus", func(b []byte) error { var v CommitStatus; return v.UnmarshalJSON(b) }},
		{"Context", func(b []byte) error { var v Context; return v.UnmarshalJSON(b) }},
		{"Operation", func(b []byte) error { var v Operation; return v.UnmarshalJSON(b) }},
		{"ProgressEntry", func(b []byte) error { var v ProgressEntry; return v.UnmarshalJSON(b) }},
		{"Remote", func(b []byte) error { var v Remote; return v.UnmarshalJSON(b) }},
		{"RemoteParameters", func(b []byte) error { var v RemoteParameters; return v.UnmarshalJSON(b) }},
		{"Repository", func(b []byte) error { var v Repository; return v.UnmarshalJSON(b) }},
		{"Volume", func(b []byte) error { var v Volume; return v.UnmarshalJSON(b) }},
		{"VolumeStatus", func(b []byte) error { var v VolumeStatus; return v.UnmarshalJSON(b) }},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if err := c.fn([]byte(`{}`)); err == nil {
				t.Error("expected error when required fields missing")
			}
		})
	}
}
