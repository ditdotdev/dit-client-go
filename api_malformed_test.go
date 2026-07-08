// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package ditclient

import (
	"context"
	"net/http"
	"testing"
)

// This file covers the success-path decode-error branch in every Execute
// method that decodes a response body. The server returns a 2xx status with
// a Content-Type of application/json but a body that fails json.Unmarshal,
// which forces the trailing `if err := decode(...); err != nil { ... }`
// branch (the only uncovered code path remaining in many Execute funcs
// after the error-status tests in api_errorpaths_test.go).

func mfWriteOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(malformedJSON))
}

// ---------------------------------------------------------------------------
// Commits
// ---------------------------------------------------------------------------

func TestMalformed_GetCommit(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.CommitsApi.GetCommit(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetCommitStatus(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.CommitsApi.GetCommitStatus(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_ListCommits(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.CommitsApi.ListCommits(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_UpdateCommit(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.CommitsApi.UpdateCommit(context.Background(), "r", "c").
		Commit(Commit{Id: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

// ---------------------------------------------------------------------------
// Remotes
// ---------------------------------------------------------------------------

func TestMalformed_ListRemotes(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.RemotesApi.ListRemotes(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetRemote(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.RemotesApi.GetRemote(context.Background(), "r", "x").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_CreateRemote(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.RemotesApi.CreateRemote(context.Background(), "r").
		Remote(Remote{Provider: "s3", Name: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_UpdateRemote(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.RemotesApi.UpdateRemote(context.Background(), "r", "x").
		Remote(Remote{Provider: "s3", Name: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetRemoteCommit(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.RemotesApi.GetRemoteCommit(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_ListRemoteCommits(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.RemotesApi.ListRemoteCommits(context.Background(), "r", "x").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

// ---------------------------------------------------------------------------
// Operations
// ---------------------------------------------------------------------------

func TestMalformed_ListOperations(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.OperationsApi.ListOperations(context.Background()).Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetOperation(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.OperationsApi.GetOperation(context.Background(), "op").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetOperationProgress(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.OperationsApi.GetOperationProgress(context.Background(), "op").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_Pull(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.OperationsApi.Pull(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_Push(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.OperationsApi.Push(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

// ---------------------------------------------------------------------------
// Repositories
// ---------------------------------------------------------------------------

func TestMalformed_GetRepository(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.RepositoriesApi.GetRepository(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetRepositoryStatus(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.RepositoriesApi.GetRepositoryStatus(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_CreateRepository(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.RepositoriesApi.CreateRepository(context.Background()).
		Repository(Repository{Name: "r", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_UpdateRepository(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "r").
		Repository(Repository{Name: "r2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}

// ---------------------------------------------------------------------------
// Volumes
// ---------------------------------------------------------------------------

func TestMalformed_ListVolumes(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.VolumesApi.ListVolumes(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetVolume(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.VolumesApi.GetVolume(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_GetVolumeStatus(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	if _, _, err := client.VolumesApi.GetVolumeStatus(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestMalformed_CreateVolume(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) { mfWriteOK(w) })
	defer ts.Close()
	_, _, err := client.VolumesApi.CreateVolume(context.Background(), "r").
		Volume(Volume{Name: "v", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected decode error")
	}
}
