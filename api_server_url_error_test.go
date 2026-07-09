// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package ditclient

import (
	"context"
	"testing"
)

// Each generated Execute method begins with:
//   localBasePath, err := a.client.cfg.ServerURLWithContext(...)
//   if err != nil { return ..., &GenericOpenAPIError{...} }
//
// Passing a context whose ContextServerIndex is the wrong type causes
// getServerOperationIndex to return an error, which propagates through
// ServerURLWithContext. This file exercises that early-return branch
// for every Execute path, which is otherwise unreached.

func badIndexCtx() context.Context {
	return context.WithValue(context.Background(), ContextServerIndex, "not-an-int")
}

func TestServerURLErr_CheckoutCommit(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.CommitsApi.CheckoutCommit(badIndexCtx(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_CreateCommit(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.CommitsApi.CreateCommit(badIndexCtx(), "r").
		Commit(Commit{Id: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_DeleteCommit(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.CommitsApi.DeleteCommit(badIndexCtx(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetCommit(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.CommitsApi.GetCommit(badIndexCtx(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetCommitStatus(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.CommitsApi.GetCommitStatus(badIndexCtx(), "r", "c").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_ListCommits(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.CommitsApi.ListCommits(badIndexCtx(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_UpdateCommit(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.CommitsApi.UpdateCommit(badIndexCtx(), "r", "c").
		Commit(Commit{Id: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetContext(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.ContextsApi.GetContext(badIndexCtx()).Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_AbortOperation(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.OperationsApi.AbortOperation(badIndexCtx(), "op").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetOperation(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.OperationsApi.GetOperation(badIndexCtx(), "op").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetOperationProgress(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.OperationsApi.GetOperationProgress(badIndexCtx(), "op").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_ListOperations(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.OperationsApi.ListOperations(badIndexCtx()).Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_Pull(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.OperationsApi.Pull(badIndexCtx(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_Push(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.OperationsApi.Push(badIndexCtx(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_CreateRemote(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.RemotesApi.CreateRemote(badIndexCtx(), "r").
		Remote(Remote{Provider: "s3", Name: "x", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_DeleteRemote(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.RemotesApi.DeleteRemote(badIndexCtx(), "r", "x").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetRemote(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.RemotesApi.GetRemote(badIndexCtx(), "r", "x").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetRemoteCommit(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.RemotesApi.GetRemoteCommit(badIndexCtx(), "r", "x", "c").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_ListRemoteCommits(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.RemotesApi.ListRemoteCommits(badIndexCtx(), "r", "x").
		RemoteParameters(RemoteParameters{Provider: "s3", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_ListRemotes(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.RemotesApi.ListRemotes(badIndexCtx(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_UpdateRemote(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.RemotesApi.UpdateRemote(badIndexCtx(), "r", "x").
		Remote(Remote{Provider: "s3", Name: "x2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_CreateRepository(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.RepositoriesApi.CreateRepository(badIndexCtx()).
		Repository(Repository{Name: "r", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_DeleteRepository(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.RepositoriesApi.DeleteRepository(badIndexCtx(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetRepository(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.RepositoriesApi.GetRepository(badIndexCtx(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetRepositoryStatus(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.RepositoriesApi.GetRepositoryStatus(badIndexCtx(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_ListRepositories(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.RepositoriesApi.ListRepositories(badIndexCtx()).Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_UpdateRepository(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.RepositoriesApi.UpdateRepository(badIndexCtx(), "r").
		Repository(Repository{Name: "r2", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_ActivateVolume(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.VolumesApi.ActivateVolume(badIndexCtx(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_CreateVolume(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	_, _, err := c.VolumesApi.CreateVolume(badIndexCtx(), "r").
		Volume(Volume{Name: "v", Properties: map[string]interface{}{}}).Execute()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_DeactivateVolume(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.VolumesApi.DeactivateVolume(badIndexCtx(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_DeleteVolume(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, err := c.VolumesApi.DeleteVolume(badIndexCtx(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetVolume(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.VolumesApi.GetVolume(badIndexCtx(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_GetVolumeStatus(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.VolumesApi.GetVolumeStatus(badIndexCtx(), "r", "v").Execute(); err == nil {
		t.Fatal("expected error")
	}
}

func TestServerURLErr_ListVolumes(t *testing.T) {
	cfg := NewConfiguration()
	c := NewAPIClient(cfg)
	if _, _, err := c.VolumesApi.ListVolumes(badIndexCtx(), "r").Execute(); err == nil {
		t.Fatal("expected error")
	}
}
