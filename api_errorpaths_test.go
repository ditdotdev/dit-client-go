// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package ditclient

import (
	"context"
	"net/http"
	"testing"
)

// This file adds the 500-status (default branch) and malformed-body (decode
// failure inside the 400/404/default error handlers) tests for every API
// Execute method that lacks them. Each test exercises:
//   - the "default" error branch (status != 400 / != 404), and
//   - the decode-error path inside an error branch (when the server returns
//     malformed JSON), which causes formatErrorMessage to be bypassed and
//     newErr.error to be replaced by the decode error message.
//
// These are the highest-yield missing coverage targets across the api_*.go
// generated files.

// Body helper for "malformed JSON" responses. The server claims JSON in the
// Content-Type but returns something that json.Unmarshal will reject.
const malformedJSON = `{not valid json`

func writeMalformed(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(malformedJSON))
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(`{"message":"` + msg + `"}`))
}

// ---------------------------------------------------------------------------
// Commits
// ---------------------------------------------------------------------------

func TestErrPath_CheckoutCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, err := client.CommitsApi.CheckoutCommit(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CheckoutCommit_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	if _, err := client.CommitsApi.CheckoutCommit(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.CommitsApi.CreateCommit(context.Background(), "r").
		Commit(Commit{Id: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateCommit_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.CommitsApi.CreateCommit(context.Background(), "r").
		Commit(Commit{Id: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateCommit_MalformedSuccess(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(malformedJSON))
	})
	defer ts.Close()
	_, _, err := client.CommitsApi.CreateCommit(context.Background(), "r").
		Commit(Commit{Id: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error from malformed success body")
	}
}

func TestErrPath_DeleteCommit_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	if _, err := client.CommitsApi.DeleteCommit(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetCommit_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	if _, _, err := client.CommitsApi.GetCommit(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetCommitStatus_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.CommitsApi.GetCommitStatus(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetCommitStatus_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	if _, _, err := client.CommitsApi.GetCommitStatus(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListCommits_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.CommitsApi.ListCommits(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.CommitsApi.UpdateCommit(context.Background(), "r", "c").
		Commit(Commit{Id: "c2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// Remotes
// ---------------------------------------------------------------------------

func TestErrPath_CreateRemote_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.CreateRemote(context.Background(), "r").
		Remote(Remote{Provider: "s3", Name: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateRemote_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.CreateRemote(context.Background(), "r").
		Remote(Remote{Provider: "s3", Name: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeleteRemote_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, err := client.RemotesApi.DeleteRemote(context.Background(), "r", "x").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRemote_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.RemotesApi.GetRemote(context.Background(), "r", "x").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRemote_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	if _, _, err := client.RemotesApi.GetRemote(context.Background(), "r", "x").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRemoteCommit_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.GetRemoteCommit(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRemoteCommit_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusBadRequest, "bad")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.GetRemoteCommit(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListRemoteCommits_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.ListRemoteCommits(context.Background(), "r", "x").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListRemoteCommits_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusNotFound, "missing")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.ListRemoteCommits(context.Background(), "r", "x").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListRemotes_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.RemotesApi.ListRemotes(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRemote_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.UpdateRemote(context.Background(), "r", "x").
		Remote(Remote{Provider: "s3", Name: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRemote_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusNotFound, "missing")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.UpdateRemote(context.Background(), "r", "x").
		Remote(Remote{Provider: "s3", Name: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRemote_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusBadRequest, "bad")
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.UpdateRemote(context.Background(), "r", "x").
		Remote(Remote{Provider: "s3", Name: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// Operations
// ---------------------------------------------------------------------------

func TestErrPath_AbortOperation_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, err := client.OperationsApi.AbortOperation(context.Background(), "op-1").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetOperation_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.OperationsApi.GetOperation(context.Background(), "op-1").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetOperationProgress_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.OperationsApi.GetOperationProgress(context.Background(), "op-1").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Pull_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Pull(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Pull_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusBadRequest, "bad")
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Pull(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Pull_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusNotFound, "missing")
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Pull(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Push_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Push(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Push_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusNotFound, "missing")
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Push(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// Repositories
// ---------------------------------------------------------------------------

func TestErrPath_CreateRepository_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.CreateRepository(context.Background()).
		Repository(Repository{Name: "alpha", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeleteRepository_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, err := client.RepositoriesApi.DeleteRepository(context.Background(), "alpha").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRepository_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.RepositoriesApi.GetRepository(context.Background(), "alpha").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRepositoryStatus_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.RepositoriesApi.GetRepositoryStatus(context.Background(), "alpha").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRepository_BadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusBadRequest, "bad")
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "alpha").
		Repository(Repository{Name: "beta", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRepository_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusNotFound, "missing")
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "alpha").
		Repository(Repository{Name: "beta", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRepository_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "alpha").
		Repository(Repository{Name: "beta", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListRepositories_MalformedSuccess(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(malformedJSON))
	})
	defer ts.Close()
	if _, _, err := client.RepositoriesApi.ListRepositories(context.Background()).Execute(); err == nil {
		t.Fatal("expected error from malformed body")
	}
}

// ---------------------------------------------------------------------------
// Volumes
// ---------------------------------------------------------------------------

func TestErrPath_ActivateVolume_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, err := client.VolumesApi.ActivateVolume(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateVolume_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	_, _, err := client.VolumesApi.CreateVolume(context.Background(), "r").
		Volume(Volume{Name: "data", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeactivateVolume_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, err := client.VolumesApi.DeactivateVolume(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeleteVolume_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, err := client.VolumesApi.DeleteVolume(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetVolume_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.VolumesApi.GetVolume(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetVolumeStatus_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.VolumesApi.GetVolumeStatus(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListVolumes_ServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusInternalServerError, "boom")
	})
	defer ts.Close()
	if _, _, err := client.VolumesApi.ListVolumes(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// Contexts
// ---------------------------------------------------------------------------

func TestErrPath_GetContext_NotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusNotFound, "missing")
	})
	defer ts.Close()
	if _, _, err := client.ContextsApi.GetContext(context.Background()).Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetContext_MalformedSuccess(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(malformedJSON))
	})
	defer ts.Close()
	if _, _, err := client.ContextsApi.GetContext(context.Background()).Execute(); err == nil {
		t.Fatal("expected error from malformed body")
	}
}

// ---------------------------------------------------------------------------
// Additional malformed-error-status tests: cover the decode-error branch
// inside specific 400/404 status handlers across various methods.
// ---------------------------------------------------------------------------

func TestErrPath_CreateVolume_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.VolumesApi.CreateVolume(context.Background(), "r").
		Volume(Volume{Name: "v", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateVolume_MalformedServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusInternalServerError)
	})
	defer ts.Close()
	_, _, err := client.VolumesApi.CreateVolume(context.Background(), "r").
		Volume(Volume{Name: "v", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateRepository_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.CreateRepository(context.Background()).
		Repository(Repository{Name: "r", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRepository_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "r").
		Repository(Repository{Name: "r2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRepository_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.UpdateRepository(context.Background(), "r").
		Repository(Repository{Name: "r2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRemote_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.UpdateRemote(context.Background(), "r", "x").
		Remote(Remote{Provider: "s3", Name: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateRemote_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.UpdateRemote(context.Background(), "r", "x").
		Remote(Remote{Provider: "s3", Name: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Pull_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Pull(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Pull_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Pull(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Push_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Push(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_Push_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.Push(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRemoteCommit_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.GetRemoteCommit(context.Background(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListRemoteCommits_MalformedBadRequest(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusBadRequest)
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.ListRemoteCommits(context.Background(), "r", "x").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_CreateRemote_MalformedServerError(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusInternalServerError)
	})
	defer ts.Close()
	_, _, err := client.RemotesApi.CreateRemote(context.Background(), "r").
		Remote(Remote{Provider: "s3", Name: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeleteVolume_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, err := client.VolumesApi.DeleteVolume(context.Background(), "r", "v").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ActivateVolume_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, err := client.VolumesApi.ActivateVolume(context.Background(), "r", "v").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeactivateVolume_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, err := client.VolumesApi.DeactivateVolume(context.Background(), "r", "v").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeleteRemote_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, err := client.RemotesApi.DeleteRemote(context.Background(), "r", "x").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_DeleteRepository_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, err := client.RepositoriesApi.DeleteRepository(context.Background(), "r").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_AbortOperation_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, err := client.OperationsApi.AbortOperation(context.Background(), "op").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetOperation_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.GetOperation(context.Background(), "op").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetOperationProgress_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.OperationsApi.GetOperationProgress(context.Background(), "op").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetVolume_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.VolumesApi.GetVolume(context.Background(), "r", "v").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetVolumeStatus_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.VolumesApi.GetVolumeStatus(context.Background(), "r", "v").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRepository_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.GetRepository(context.Background(), "r").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetRepositoryStatus_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.RepositoriesApi.GetRepositoryStatus(context.Background(), "r").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListVolumes_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.VolumesApi.ListVolumes(context.Background(), "r").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_ListCommits_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.CommitsApi.ListCommits(context.Background(), "r").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_GetCommitStatus_MalformedNotFound2(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.CommitsApi.GetCommitStatus(context.Background(), "r", "c").Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestErrPath_UpdateCommit_MalformedNotFound(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		writeMalformed(w, http.StatusNotFound)
	})
	defer ts.Close()
	_, _, err := client.CommitsApi.UpdateCommit(context.Background(), "r", "c").
		Commit(Commit{Id: "c2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---------------------------------------------------------------------------
// Network error path: callAPI returns err when the request can't be made.
// We point at a closed listener to force a connection refused, which then
// flows through the "if err != nil || localVarHTTPResponse == nil" branch
// in every Execute method.
// ---------------------------------------------------------------------------

func TestErrPath_ListRemotes_NetworkError(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:1"}} // port 1 not listening
	client := NewAPIClient(cfg)
	if _, _, err := client.RemotesApi.ListRemotes(context.Background(), "r").Execute(); err == nil {
		t.Fatal("expected network error")
	}
}

func TestErrPath_GetCommit_NetworkError(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:1"}}
	client := NewAPIClient(cfg)
	if _, _, err := client.CommitsApi.GetCommit(context.Background(), "r", "c").Execute(); err == nil {
		t.Fatal("expected network error")
	}
}

func TestErrPath_DeleteVolume_NetworkError(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://127.0.0.1:1"}}
	client := NewAPIClient(cfg)
	if _, err := client.VolumesApi.DeleteVolume(context.Background(), "r", "v").Execute(); err == nil {
		t.Fatal("expected network error")
	}
}

// ---------------------------------------------------------------------------
// Context-canceled path through callAPI.
// ---------------------------------------------------------------------------

func TestErrPath_ListRemotes_ContextCanceled(t *testing.T) {
	ts, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	})
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel before the request runs
	if _, _, err := client.RemotesApi.ListRemotes(ctx, "r").Execute(); err == nil {
		t.Fatal("expected error from canceled context")
	}
}
