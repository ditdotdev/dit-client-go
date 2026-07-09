// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package ditclient

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// ---------------------------------------------------------------------------
// atoi
// ---------------------------------------------------------------------------

func TestAtoi_Valid(t *testing.T) {
	got, err := atoi("42")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
}

func TestAtoi_Invalid(t *testing.T) {
	if _, err := atoi("not-a-number"); err == nil {
		t.Error("expected error for non-numeric input")
	}
}

// ---------------------------------------------------------------------------
// strlen
// ---------------------------------------------------------------------------

func TestStrlen_ASCII(t *testing.T) {
	if got := strlen("hello"); got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
}

func TestStrlen_UTF8(t *testing.T) {
	// "héllo" has 5 runes but 6 bytes
	if got := strlen("héllo"); got != 5 {
		t.Errorf("expected 5 runes, got %d", got)
	}
}

func TestStrlen_Empty(t *testing.T) {
	if got := strlen(""); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

// ---------------------------------------------------------------------------
// typeCheckParameter
// ---------------------------------------------------------------------------

func TestTypeCheckParameter_NilObject(t *testing.T) {
	if err := typeCheckParameter(nil, "string", "name"); err != nil {
		t.Errorf("expected nil error for nil input, got %v", err)
	}
}

func TestTypeCheckParameter_Matching(t *testing.T) {
	if err := typeCheckParameter("hello", "string", "name"); err != nil {
		t.Errorf("expected nil for matching type, got %v", err)
	}
}

func TestTypeCheckParameter_Mismatch(t *testing.T) {
	err := typeCheckParameter(42, "string", "name")
	if err == nil {
		t.Fatal("expected error for type mismatch")
	}
	if !strings.Contains(err.Error(), "name") || !strings.Contains(err.Error(), "string") {
		t.Errorf("error should describe expected/actual type: %v", err)
	}
}

// ---------------------------------------------------------------------------
// addFile
// ---------------------------------------------------------------------------

func TestAddFile_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	if err := os.WriteFile(path, []byte("hello file"), 0o600); err != nil {
		t.Fatalf("write tempfile: %v", err)
	}

	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := addFile(w, "file", path); err != nil {
		t.Fatalf("addFile failed: %v", err)
	}
	_ = w.Close()

	if !strings.Contains(buf.String(), "hello file") {
		t.Errorf("expected file content in multipart body, got %q", buf.String())
	}
}

func TestAddFile_MissingPath(t *testing.T) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	if err := addFile(w, "file", filepath.Join(t.TempDir(), "no-such-file")); err == nil {
		t.Error("expected error for missing file")
	}
}

// ---------------------------------------------------------------------------
// newStrictDecoder
// ---------------------------------------------------------------------------

func TestNewStrictDecoder_RejectsUnknownFields(t *testing.T) {
	type sample struct {
		Known string `json:"known"`
	}
	dec := newStrictDecoder([]byte(`{"known":"ok","extra":"nope"}`))
	var s sample
	if err := dec.Decode(&s); err == nil {
		t.Error("expected error for unknown field")
	}
}

func TestNewStrictDecoder_AcceptsKnownFields(t *testing.T) {
	type sample struct {
		Known string `json:"known"`
	}
	dec := newStrictDecoder([]byte(`{"known":"ok"}`))
	var s sample
	if err := dec.Decode(&s); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.Known != "ok" {
		t.Errorf("expected ok, got %q", s.Known)
	}
}

// ---------------------------------------------------------------------------
// parameterAddToHeaderOrQuery – exercise remaining type branches.
// ---------------------------------------------------------------------------

func TestParameterAddToHeaderOrQuery_StringIntoQuery(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "key", "value", "form", "")
	if got := q.Get("key"); got != "value" {
		t.Errorf("expected key=value, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_IntIntoQuery(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "n", int64(42), "form", "")
	if got := q.Get("n"); got != "42" {
		t.Errorf("expected n=42, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_UintIntoQuery(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "u", uint64(7), "form", "")
	if got := q.Get("u"); got != "7" {
		t.Errorf("expected u=7, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_FloatIntoQuery(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "f", 3.14, "form", "")
	if got := q.Get("f"); !strings.HasPrefix(got, "3.14") {
		t.Errorf("expected ~3.14, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_BoolIntoQuery(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "b", true, "form", "")
	if got := q.Get("b"); got != "true" {
		t.Errorf("expected b=true, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_SliceIntoQuery(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "tag", []string{"v1", "v2"}, "form", "multi")
	got := q["tag"]
	if len(got) != 2 || got[0] != "v1" || got[1] != "v2" {
		t.Errorf("expected [v1 v2], got %v", got)
	}
}

func TestParameterAddToHeaderOrQuery_SliceCSV(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "tag", []string{"a", "b", "c"}, "form", "csv")
	if got := q.Get("tag"); got != "a,b,c" {
		t.Errorf("expected csv joined a,b,c, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_PointerIntoQuery(t *testing.T) {
	q := url.Values{}
	s := "ptr"
	parameterAddToHeaderOrQuery(q, "p", &s, "form", "")
	if got := q.Get("p"); got != "ptr" {
		t.Errorf("expected p=ptr, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_MapIntoHeader(t *testing.T) {
	h := map[string]string{}
	parameterAddToHeaderOrQuery(h, "x", "y", "form", "")
	if h["x"] != "y" {
		t.Errorf("expected x=y, got %v", h)
	}
}

func TestParameterAddToHeaderOrQuery_TimeStruct(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "when", time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC), "form", "")
	if got := q.Get("when"); !strings.HasPrefix(got, "2024-01-02T03:04:05") {
		t.Errorf("expected RFC3339Nano, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_MapValueIntoQuery(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "m", map[string]string{"k": "v"}, "deepObject", "")
	if got := q.Get("m[k]"); got != "v" {
		t.Errorf("expected m[k]=v, got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_DeepObjectSlice(t *testing.T) {
	q := url.Values{}
	parameterAddToHeaderOrQuery(q, "k", []string{"x", "y"}, "deepObject", "multi")
	if got := q.Get("k[0]"); got != "x" {
		t.Errorf("expected k[0]=x, got %q", got)
	}
	if got := q.Get("k[1]"); got != "y" {
		t.Errorf("expected k[1]=y, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// parameterValueToString – pointer + MappedNullable branch
// ---------------------------------------------------------------------------

func TestParameterValueToString_NonMappedPointer(t *testing.T) {
	s := "hello"
	// *string does not implement MappedNullable; the function should return ""
	if got := parameterValueToString(&s, "key"); got != "" {
		t.Errorf("expected empty for non-MappedNullable pointer, got %q", got)
	}
}

func TestParameterValueToString_MappedNullablePointer(t *testing.T) {
	c := NewCommit("abc", map[string]interface{}{"k": "v"})
	if got := parameterValueToString(c, "id"); got != "abc" {
		t.Errorf("expected abc, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// callAPI – Debug mode wraps the request/response in DumpRequestOut/DumpResponse.
// Exercise the Debug=true path; assertions are minimal because log.Printf
// goes to the default logger.
// ---------------------------------------------------------------------------

func TestCallAPI_DebugMode(t *testing.T) {
	logBuf := &bytes.Buffer{}
	origOut := log.Writer()
	log.SetOutput(logBuf)
	defer log.SetOutput(origOut)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: ts.URL}}
	cfg.Debug = true
	client := NewAPIClient(cfg)

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}
	resp, err := client.callAPI(req)
	if err != nil {
		t.Fatalf("callAPI: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	logged := logBuf.String()
	if !strings.Contains(logged, "GET") {
		t.Errorf("expected request dump containing GET in debug log, got: %s", logged)
	}
}

// ---------------------------------------------------------------------------
// prepareRequest – multipart/form-data and x-www-form-urlencoded branches.
// ---------------------------------------------------------------------------

func TestPrepareRequest_MultipartFormData(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	client := NewAPIClient(cfg)

	headers := map[string]string{"Content-Type": "multipart/form-data"}
	form := url.Values{}
	form.Add("foo", "bar")

	dir := t.TempDir()
	fp := filepath.Join(dir, "in.txt")
	if err := os.WriteFile(fp, []byte("contents"), 0o600); err != nil {
		t.Fatalf("write tempfile: %v", err)
	}

	files := []formFile{
		{fileBytes: []byte("payload"), fileName: filepath.Base(fp), formFileName: "upload"},
	}

	req, err := client.prepareRequest(
		context.Background(),
		"http://example.com/path",
		http.MethodPost,
		nil,
		headers,
		url.Values{},
		form,
		files,
	)
	if err != nil {
		t.Fatalf("prepareRequest: %v", err)
	}

	got := req.Header.Get("Content-Type")
	if !strings.HasPrefix(got, "multipart/form-data") {
		t.Errorf("expected multipart/form-data Content-Type, got %q", got)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "payload") || !strings.Contains(string(body), "bar") {
		t.Errorf("expected payload and form value in body, got %q", string(body))
	}
}

func TestPrepareRequest_FormURLEncoded(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	client := NewAPIClient(cfg)

	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	form := url.Values{}
	form.Add("a", "1")
	form.Add("b", "2")

	req, err := client.prepareRequest(
		context.Background(),
		"http://example.com/path",
		http.MethodPost,
		nil,
		headers,
		url.Values{},
		form,
		nil,
	)
	if err != nil {
		t.Fatalf("prepareRequest: %v", err)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "a=1") || !strings.Contains(string(body), "b=2") {
		t.Errorf("expected a=1 and b=2 in body, got %q", string(body))
	}
	if req.Header.Get("Content-Length") == "" {
		t.Error("expected Content-Length header to be set")
	}
}

func TestPrepareRequest_HostAndSchemeOverride(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	cfg.Host = "override.example"
	cfg.Scheme = "https"
	client := NewAPIClient(cfg)

	req, err := client.prepareRequest(
		context.Background(),
		"http://example.com/some/path",
		http.MethodGet,
		nil,
		map[string]string{},
		url.Values{},
		url.Values{},
		nil,
	)
	if err != nil {
		t.Fatalf("prepareRequest: %v", err)
	}
	if req.URL.Host != "override.example" {
		t.Errorf("expected host override, got %q", req.URL.Host)
	}
	if req.URL.Scheme != "https" {
		t.Errorf("expected https scheme, got %q", req.URL.Scheme)
	}
}

func TestPrepareRequest_DefaultHeader(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	cfg.AddDefaultHeader("X-Default", "yes")
	client := NewAPIClient(cfg)

	req, err := client.prepareRequest(
		context.Background(),
		"http://example.com/path",
		http.MethodGet,
		nil,
		map[string]string{},
		url.Values{},
		url.Values{},
		nil,
	)
	if err != nil {
		t.Fatalf("prepareRequest: %v", err)
	}
	if req.Header.Get("X-Default") != "yes" {
		t.Errorf("expected X-Default header, got %q", req.Header.Get("X-Default"))
	}
}

func TestPrepareRequest_BadURL(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	client := NewAPIClient(cfg)

	_, err := client.prepareRequest(
		context.Background(),
		"://bad-url",
		http.MethodGet,
		nil,
		map[string]string{},
		url.Values{},
		url.Values{},
		nil,
	)
	if err == nil {
		t.Fatal("expected error parsing bad URL")
	}
}

func TestPrepareRequest_PostBodyAndMultipart_Conflict(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	client := NewAPIClient(cfg)

	form := url.Values{}
	form.Add("k", "v")

	_, err := client.prepareRequest(
		context.Background(),
		"http://example.com/path",
		http.MethodPost,
		"json body",
		map[string]string{"Content-Type": "multipart/form-data"},
		url.Values{},
		form,
		nil,
	)
	if err == nil {
		t.Fatal("expected error combining postBody with multipart form")
	}
}

func TestPrepareRequest_PostBodyAndFormURLEncoded_Conflict(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	client := NewAPIClient(cfg)

	form := url.Values{}
	form.Add("k", "v")

	_, err := client.prepareRequest(
		context.Background(),
		"http://example.com/path",
		http.MethodPost,
		"json body",
		map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
		url.Values{},
		form,
		nil,
	)
	if err == nil {
		t.Fatal("expected error combining postBody with form-urlencoded body")
	}
}

func TestPrepareRequest_QueryParams(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://example.com"}}
	client := NewAPIClient(cfg)

	q := url.Values{}
	q.Add("a", "1")
	q.Add("b", "2")

	req, err := client.prepareRequest(
		context.Background(),
		"http://example.com/path",
		http.MethodGet,
		nil,
		map[string]string{},
		q,
		url.Values{},
		nil,
	)
	if err != nil {
		t.Fatalf("prepareRequest: %v", err)
	}
	if got := req.URL.Query().Get("a"); got != "1" {
		t.Errorf("expected a=1, got %q", got)
	}
	if got := req.URL.Query().Get("b"); got != "2" {
		t.Errorf("expected b=2, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// decode – exercise the *os.File and **os.File branches.
// ---------------------------------------------------------------------------

func TestDecode_FilePointer(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var f *os.File
	if err := client.decode(&f, []byte("payload"), "application/octet-stream"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected *os.File to be assigned")
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()
	contents, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("read tempfile: %v", err)
	}
	if string(contents) != "payload" {
		t.Errorf("expected payload, got %q", string(contents))
	}
}

func TestDecode_JSONMalformed(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var out map[string]string
	if err := client.decode(&out, []byte(`{not-json`), "application/json"); err == nil {
		t.Error("expected error for malformed JSON")
	}
}

func TestDecode_XMLMalformed(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	type sample struct {
		Name string `xml:"name"`
	}
	var s sample
	if err := client.decode(&s, []byte(`<bad`), "application/xml"); err == nil {
		t.Error("expected error for malformed XML")
	}
}

// ---------------------------------------------------------------------------
// setBody – additional branches
// ---------------------------------------------------------------------------

func TestSetBody_FilePointer(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "in.txt")
	if err := os.WriteFile(path, []byte("file-bytes"), 0o600); err != nil {
		t.Fatalf("write tempfile: %v", err)
	}
	fp, err := os.Open(path) //nolint:gosec // test fixture, path is controlled by t.TempDir
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer func() { _ = fp.Close() }()

	buf, err := setBody(fp, "application/octet-stream")
	if err != nil {
		t.Fatalf("setBody: %v", err)
	}
	if buf.String() != "file-bytes" {
		t.Errorf("expected file-bytes, got %q", buf.String())
	}
}

// ---------------------------------------------------------------------------
// formatErrorMessage – non-struct branch (no Title/Detail fields)
// ---------------------------------------------------------------------------

func TestFormatErrorMessage_NoTitleOrDetail(t *testing.T) {
	type empty struct{}
	got := formatErrorMessage("500 Internal Server Error", &empty{})
	if !strings.Contains(got, "500") {
		t.Errorf("expected status prefix, got %q", got)
	}
}

func TestFormatErrorMessage_WithTitleAndDetail(t *testing.T) {
	type rfc7807 struct {
		Title  string
		Detail string
	}
	got := formatErrorMessage("400 Bad Request", &rfc7807{Title: "Invalid", Detail: "something broke"})
	if !strings.Contains(got, "Invalid") || !strings.Contains(got, "something broke") {
		t.Errorf("expected title and detail in message, got %q", got)
	}
}
