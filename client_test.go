package Datadatdatclient

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

// ---------------------------------------------------------------------------
// selectHeaderContentType
// ---------------------------------------------------------------------------

func TestSelectHeaderContentType_Empty(t *testing.T) {
	got := selectHeaderContentType([]string{})
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestSelectHeaderContentType_PrefersJSON(t *testing.T) {
	got := selectHeaderContentType([]string{"text/plain", "application/json", "application/xml"})
	if got != "application/json" {
		t.Errorf("expected application/json, got %q", got)
	}
}

func TestSelectHeaderContentType_FallsBackToFirst(t *testing.T) {
	got := selectHeaderContentType([]string{"application/xml", "text/html"})
	if got != "application/xml" {
		t.Errorf("expected application/xml, got %q", got)
	}
}

func TestSelectHeaderContentType_SingleJSON(t *testing.T) {
	got := selectHeaderContentType([]string{"application/json"})
	if got != "application/json" {
		t.Errorf("expected application/json, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// selectHeaderAccept
// ---------------------------------------------------------------------------

func TestSelectHeaderAccept_Empty(t *testing.T) {
	got := selectHeaderAccept([]string{})
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestSelectHeaderAccept_PrefersJSON(t *testing.T) {
	got := selectHeaderAccept([]string{"text/plain", "application/json"})
	if got != "application/json" {
		t.Errorf("expected application/json, got %q", got)
	}
}

func TestSelectHeaderAccept_JoinsWhenNoJSON(t *testing.T) {
	got := selectHeaderAccept([]string{"text/plain", "application/xml"})
	if got != "text/plain,application/xml" {
		t.Errorf("expected joined string, got %q", got)
	}
}

func TestSelectHeaderAccept_CaseInsensitiveJSON(t *testing.T) {
	got := selectHeaderAccept([]string{"Application/JSON"})
	if got != "application/json" {
		t.Errorf("expected application/json, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// contains
// ---------------------------------------------------------------------------

func TestContains_Found(t *testing.T) {
	if !contains([]string{"a", "b", "c"}, "b") {
		t.Error("expected true for element in slice")
	}
}

func TestContains_NotFound(t *testing.T) {
	if contains([]string{"a", "b", "c"}, "d") {
		t.Error("expected false for element not in slice")
	}
}

func TestContains_CaseInsensitive(t *testing.T) {
	if !contains([]string{"Application/JSON"}, "application/json") {
		t.Error("expected case insensitive match")
	}
}

func TestContains_EmptySlice(t *testing.T) {
	if contains([]string{}, "anything") {
		t.Error("expected false for empty slice")
	}
}

// ---------------------------------------------------------------------------
// parameterToString
// ---------------------------------------------------------------------------

func TestParameterToString_StringValue(t *testing.T) {
	got := parameterToString("hello", "")
	if got != "hello" {
		t.Errorf("expected hello, got %q", got)
	}
}

func TestParameterToString_IntValue(t *testing.T) {
	got := parameterToString(42, "")
	if got != "42" {
		t.Errorf("expected 42, got %q", got)
	}
}

func TestParameterToString_BoolValue(t *testing.T) {
	got := parameterToString(true, "")
	if got != "true" {
		t.Errorf("expected true, got %q", got)
	}
}

func TestParameterToString_TimeValue(t *testing.T) {
	ts := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)
	got := parameterToString(ts, "")
	if got != "2025-01-15T12:00:00Z" {
		t.Errorf("expected RFC3339 time, got %q", got)
	}
}

func TestParameterToString_SliceCSV(t *testing.T) {
	got := parameterToString([]string{"a", "b", "c"}, "csv")
	if got != "a,b,c" {
		t.Errorf("expected a,b,c, got %q", got)
	}
}

func TestParameterToString_SliceSSV(t *testing.T) {
	got := parameterToString([]string{"a", "b", "c"}, "ssv")
	if got != "a b c" {
		t.Errorf("expected 'a b c', got %q", got)
	}
}

func TestParameterToString_SliceTSV(t *testing.T) {
	got := parameterToString([]string{"a", "b", "c"}, "tsv")
	if got != "a\tb\tc" {
		t.Errorf("expected tab-separated, got %q", got)
	}
}

func TestParameterToString_SlicePipes(t *testing.T) {
	got := parameterToString([]string{"a", "b", "c"}, "pipes")
	if got != "a|b|c" {
		t.Errorf("expected pipe-separated, got %q", got)
	}
}

func TestParameterToString_SliceDefaultDelimiter(t *testing.T) {
	// Without a recognized collection format, the delimiter is empty string.
	// ReplaceAll replaces spaces with "", so "[a b]" becomes "[ab]", trimmed to "ab".
	got := parameterToString([]string{"a", "b"}, "")
	if got != "ab" {
		t.Errorf("expected 'ab', got %q", got)
	}
}

// ---------------------------------------------------------------------------
// parameterToJson
// ---------------------------------------------------------------------------

func TestParameterToJson_SimpleStruct(t *testing.T) {
	input := map[string]string{"key": "value"}
	got, err := parameterToJson(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, `"key":"value"`) {
		t.Errorf("expected JSON with key:value, got %q", got)
	}
}

func TestParameterToJson_Nil(t *testing.T) {
	got, err := parameterToJson(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(got) != "null" {
		t.Errorf("expected null, got %q", got)
	}
}

func TestParameterToJson_SliceOfInts(t *testing.T) {
	got, err := parameterToJson([]int{1, 2, 3})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(got) != "[1,2,3]" {
		t.Errorf("expected [1,2,3], got %q", got)
	}
}

// ---------------------------------------------------------------------------
// detectContentType
// ---------------------------------------------------------------------------

func TestDetectContentType_Struct(t *testing.T) {
	type sample struct{ Name string }
	got := detectContentType(sample{Name: "test"})
	if got != "application/json; charset=utf-8" {
		t.Errorf("expected application/json, got %q", got)
	}
}

func TestDetectContentType_Map(t *testing.T) {
	got := detectContentType(map[string]int{"a": 1})
	if got != "application/json; charset=utf-8" {
		t.Errorf("expected application/json, got %q", got)
	}
}

func TestDetectContentType_Pointer(t *testing.T) {
	s := "hello"
	got := detectContentType(&s)
	if got != "application/json; charset=utf-8" {
		t.Errorf("expected application/json, got %q", got)
	}
}

func TestDetectContentType_String(t *testing.T) {
	got := detectContentType("hello")
	if got != "text/plain; charset=utf-8" {
		t.Errorf("expected text/plain, got %q", got)
	}
}

func TestDetectContentType_ByteSlice(t *testing.T) {
	got := detectContentType([]byte("<html></html>"))
	// http.DetectContentType will return text/html
	if !strings.Contains(got, "text/html") {
		t.Errorf("expected text/html for HTML bytes, got %q", got)
	}
}

func TestDetectContentType_SliceOfStructs(t *testing.T) {
	type item struct{ ID int }
	got := detectContentType([]item{{ID: 1}})
	if got != "application/json; charset=utf-8" {
		t.Errorf("expected application/json for slice of structs, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// parseCacheControl
// ---------------------------------------------------------------------------

func TestParseCacheControl_Empty(t *testing.T) {
	h := http.Header{}
	cc := parseCacheControl(h)
	if len(cc) != 0 {
		t.Errorf("expected empty cache control, got %v", cc)
	}
}

func TestParseCacheControl_MaxAge(t *testing.T) {
	h := http.Header{}
	h.Set("Cache-Control", "max-age=3600")
	cc := parseCacheControl(h)
	if cc["max-age"] != "3600" {
		t.Errorf("expected max-age=3600, got %q", cc["max-age"])
	}
}

func TestParseCacheControl_MultipleDirectives(t *testing.T) {
	h := http.Header{}
	h.Set("Cache-Control", "public, max-age=600, no-transform")
	cc := parseCacheControl(h)
	if cc["max-age"] != "600" {
		t.Errorf("expected max-age=600, got %q", cc["max-age"])
	}
	if _, ok := cc["public"]; !ok {
		t.Error("expected public directive")
	}
	if _, ok := cc["no-transform"]; !ok {
		t.Error("expected no-transform directive")
	}
}

func TestParseCacheControl_NoCache(t *testing.T) {
	h := http.Header{}
	h.Set("Cache-Control", "no-cache, no-store")
	cc := parseCacheControl(h)
	if _, ok := cc["no-cache"]; !ok {
		t.Error("expected no-cache directive")
	}
	if _, ok := cc["no-store"]; !ok {
		t.Error("expected no-store directive")
	}
}

// ---------------------------------------------------------------------------
// CacheExpires
// ---------------------------------------------------------------------------

func TestCacheExpires_WithMaxAge(t *testing.T) {
	h := http.Header{}
	now := time.Now().UTC()
	h.Set("Date", now.Format(time.RFC1123))
	h.Set("Cache-Control", "max-age=60")

	resp := &http.Response{Header: h}
	expires := CacheExpires(resp)

	// The expires time should be approximately now + 60 seconds
	expected := now.Add(60 * time.Second)
	diff := expires.Sub(expected)
	if diff < -2*time.Second || diff > 2*time.Second {
		t.Errorf("expected expires ~%v, got %v (diff: %v)", expected, expires, diff)
	}
}

func TestCacheExpires_WithExpiresHeader(t *testing.T) {
	h := http.Header{}
	now := time.Now().UTC()
	expiresTime := now.Add(120 * time.Second)
	h.Set("Date", now.Format(time.RFC1123))
	h.Set("Expires", expiresTime.Format(time.RFC1123))

	resp := &http.Response{Header: h}
	expires := CacheExpires(resp)

	diff := expires.Sub(expiresTime)
	if diff < -2*time.Second || diff > 2*time.Second {
		t.Errorf("expected expires ~%v, got %v (diff: %v)", expiresTime, expires, diff)
	}
}

func TestCacheExpires_NoDateHeader(t *testing.T) {
	h := http.Header{}
	resp := &http.Response{Header: h}

	before := time.Now()
	expires := CacheExpires(resp)
	after := time.Now()

	// When Date header is missing, should return approximately time.Now()
	if expires.Before(before.Add(-time.Second)) || expires.After(after.Add(time.Second)) {
		t.Errorf("expected expires near now, got %v", expires)
	}
}

func TestCacheExpires_InvalidMaxAge(t *testing.T) {
	h := http.Header{}
	now := time.Now().UTC()
	h.Set("Date", now.Format(time.RFC1123))
	h.Set("Cache-Control", "max-age=not-a-number")

	resp := &http.Response{Header: h}
	expires := CacheExpires(resp)

	// With invalid max-age, expires should be set to the parsed "now" from Date header
	diff := expires.Sub(now)
	if diff < -2*time.Second || diff > 2*time.Second {
		t.Errorf("expected expires near date header time, got %v (diff: %v)", expires, diff)
	}
}

// ---------------------------------------------------------------------------
// shouldUseMultipart
// ---------------------------------------------------------------------------

func TestShouldUseMultipart_WithMultipartAndFormParams(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	headers := map[string]string{"Content-Type": "multipart/form-data"}
	formParams := url.Values{"field": {"value"}}

	if !client.shouldUseMultipart(headers, formParams, nil, "") {
		t.Error("expected true when Content-Type is multipart and form params present")
	}
}

func TestShouldUseMultipart_WithFileBytes(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	headers := map[string]string{}
	formParams := url.Values{}

	if !client.shouldUseMultipart(headers, formParams, []byte("data"), "file.txt") {
		t.Error("expected true when file bytes and filename are present")
	}
}

func TestShouldUseMultipart_False(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	headers := map[string]string{"Content-Type": "application/json"}
	formParams := url.Values{}

	if client.shouldUseMultipart(headers, formParams, nil, "") {
		t.Error("expected false when not multipart and no file")
	}
}

func TestShouldUseMultipart_MultipartButNoFormParams(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	headers := map[string]string{"Content-Type": "multipart/form-data"}
	formParams := url.Values{}

	if client.shouldUseMultipart(headers, formParams, nil, "") {
		t.Error("expected false when multipart but no form params and no file")
	}
}

func TestShouldUseMultipart_FileBytesButNoName(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	headers := map[string]string{}
	formParams := url.Values{}

	if client.shouldUseMultipart(headers, formParams, []byte("data"), "") {
		t.Error("expected false when file bytes present but no filename")
	}
}

// ---------------------------------------------------------------------------
// setBody
// ---------------------------------------------------------------------------

func TestSetBody_StringContent(t *testing.T) {
	buf, err := setBody("hello world", "text/plain")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.String() != "hello world" {
		t.Errorf("expected 'hello world', got %q", buf.String())
	}
}

func TestSetBody_StringPointer(t *testing.T) {
	s := "pointer content"
	buf, err := setBody(&s, "text/plain")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.String() != "pointer content" {
		t.Errorf("expected 'pointer content', got %q", buf.String())
	}
}

func TestSetBody_ByteSlice(t *testing.T) {
	buf, err := setBody([]byte("byte data"), "application/octet-stream")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.String() != "byte data" {
		t.Errorf("expected 'byte data', got %q", buf.String())
	}
}

func TestSetBody_Reader(t *testing.T) {
	reader := bytes.NewBufferString("reader content")
	buf, err := setBody(reader, "text/plain")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.String() != "reader content" {
		t.Errorf("expected 'reader content', got %q", buf.String())
	}
}

func TestSetBody_JSONStruct(t *testing.T) {
	type sample struct {
		Name string `json:"name"`
	}
	buf, err := setBody(sample{Name: "test"}, "application/json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), `"name":"test"`) {
		t.Errorf("expected JSON with name:test, got %q", buf.String())
	}
}

func TestSetBody_XMLStruct(t *testing.T) {
	type sample struct {
		Name string `xml:"name"`
	}
	buf, err := setBody(sample{Name: "test"}, "application/xml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "<name>test</name>") {
		t.Errorf("expected XML with <name>test</name>, got %q", buf.String())
	}
}

func TestSetBody_UnsupportedContentType(t *testing.T) {
	// An integer with an unsupported content type should produce an error
	_, err := setBody(42, "application/octet-stream")
	if err == nil {
		t.Error("expected error for unsupported body type, got nil")
	}
}

// ---------------------------------------------------------------------------
// GenericOpenAPIError
// ---------------------------------------------------------------------------

func TestGenericOpenAPIError_Error(t *testing.T) {
	e := GenericOpenAPIError{error: "something went wrong"}
	if e.Error() != "something went wrong" {
		t.Errorf("expected error message, got %q", e.Error())
	}
}

func TestGenericOpenAPIError_Body(t *testing.T) {
	body := []byte(`{"error": "not found"}`)
	e := GenericOpenAPIError{body: body}
	if !bytes.Equal(e.Body(), body) {
		t.Errorf("expected body %q, got %q", body, e.Body())
	}
}

func TestGenericOpenAPIError_BodyNil(t *testing.T) {
	e := GenericOpenAPIError{}
	if e.Body() != nil {
		t.Errorf("expected nil body, got %q", e.Body())
	}
}

func TestGenericOpenAPIError_Model(t *testing.T) {
	model := ApiError{Code: "NOT_FOUND", Message: "not found"}
	e := GenericOpenAPIError{model: model}
	got, ok := e.Model().(ApiError)
	if !ok {
		t.Fatal("expected model to be ApiError")
	}
	if got.Code != "NOT_FOUND" {
		t.Errorf("expected code NOT_FOUND, got %q", got.Code)
	}
}

func TestGenericOpenAPIError_ModelNil(t *testing.T) {
	e := GenericOpenAPIError{}
	if e.Model() != nil {
		t.Error("expected nil model")
	}
}

func TestGenericOpenAPIError_ImplementsErrorInterface(t *testing.T) {
	var err error = GenericOpenAPIError{error: "test"}
	if err.Error() != "test" {
		t.Errorf("expected GenericOpenAPIError to satisfy error interface")
	}
}

// ---------------------------------------------------------------------------
// NewAPIClient
// ---------------------------------------------------------------------------

func TestNewAPIClient_DefaultHTTPClient(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	if client.cfg.HTTPClient == nil {
		t.Error("expected default HTTP client to be set")
	}
}

func TestNewAPIClient_CustomHTTPClient(t *testing.T) {
	cfg := NewConfiguration()
	custom := &http.Client{Timeout: 30 * time.Second}
	cfg.HTTPClient = custom
	client := NewAPIClient(cfg)
	if client.cfg.HTTPClient != custom {
		t.Error("expected custom HTTP client to be preserved")
	}
}

func TestNewAPIClient_ServicesInitialized(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	if client.CommitsApi == nil {
		t.Error("CommitsApi should be initialized")
	}
	if client.ContextsApi == nil {
		t.Error("ContextsApi should be initialized")
	}
	if client.OperationsApi == nil {
		t.Error("OperationsApi should be initialized")
	}
	if client.RemotesApi == nil {
		t.Error("RemotesApi should be initialized")
	}
	if client.RepositoriesApi == nil {
		t.Error("RepositoriesApi should be initialized")
	}
	if client.VolumesApi == nil {
		t.Error("VolumesApi should be initialized")
	}
}

func TestNewAPIClient_ChangeBasePath(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	client.ChangeBasePath("http://example.com:9090")
	if client.GetConfig().BasePath != "http://example.com:9090" {
		t.Errorf("expected changed base path, got %q", client.GetConfig().BasePath)
	}
}

func TestNewAPIClient_GetConfig(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	if client.GetConfig() != cfg {
		t.Error("GetConfig should return the same configuration")
	}
}

// ---------------------------------------------------------------------------
// decode
// ---------------------------------------------------------------------------

func TestDecode_JSONContent(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var result map[string]string
	err := client.decode(&result, []byte(`{"key":"value"}`), "application/json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["key"] != "value" {
		t.Errorf("expected key=value, got %q", result["key"])
	}
}

func TestDecode_JSONVndContent(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var result map[string]string
	err := client.decode(&result, []byte(`{"key":"value"}`), "application/vnd.api+json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["key"] != "value" {
		t.Errorf("expected key=value, got %q", result["key"])
	}
}

func TestDecode_XMLContent(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	type Item struct {
		Name string `xml:"name"`
	}
	var result Item
	err := client.decode(&result, []byte(`<Item><name>test</name></Item>`), "application/xml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "test" {
		t.Errorf("expected name=test, got %q", result.Name)
	}
}

func TestDecode_StringTarget(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var result string
	err := client.decode(&result, []byte("hello"), "text/plain")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "hello" {
		t.Errorf("expected hello, got %q", result)
	}
}

func TestDecode_EmptyBody(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var result map[string]string
	err := client.decode(&result, []byte{}, "application/json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDecode_UnsupportedContentType(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	var result map[string]string
	err := client.decode(&result, []byte("data"), "application/octet-stream")
	if err == nil {
		t.Error("expected error for unsupported content type")
	}
}
