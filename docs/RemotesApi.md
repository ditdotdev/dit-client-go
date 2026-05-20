# \RemotesApi

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemote**](RemotesApi.md#CreateRemote) | **Post** /v1/repositories/{repositoryName}/remotes | Create new remote
[**DeleteRemote**](RemotesApi.md#DeleteRemote) | **Delete** /v1/repositories/{repositoryName}/remotes/{remoteName} | Delete remote
[**GetRemote**](RemotesApi.md#GetRemote) | **Get** /v1/repositories/{repositoryName}/remotes/{remoteName} | Get information about a particular remote
[**GetRemoteCommit**](RemotesApi.md#GetRemoteCommit) | **Get** /v1/repositories/{repositoryName}/remotes/{remoteName}/commits/{commitId} | Get a remote commit
[**ListRemoteCommits**](RemotesApi.md#ListRemoteCommits) | **Get** /v1/repositories/{repositoryName}/remotes/{remoteName}/commits | List remote commits
[**ListRemotes**](RemotesApi.md#ListRemotes) | **Get** /v1/repositories/{repositoryName}/remotes | Get list of remotes
[**UpdateRemote**](RemotesApi.md#UpdateRemote) | **Post** /v1/repositories/{repositoryName}/remotes/{remoteName} | Update remote information



## CreateRemote

> Remote CreateRemote(ctx, repositoryName, remote)

Create new remote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remote** | [**Remote**](Remote.md)| Remote to create | 

### Return type

[**Remote**](Remote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemote

> DeleteRemote(ctx, repositoryName, remoteName)

Delete remote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remoteName** | **string**| Name of the remote | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemote

> Remote GetRemote(ctx, repositoryName, remoteName)

Get information about a particular remote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remoteName** | **string**| Name of the remote | 

### Return type

[**Remote**](Remote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteCommit

> Commit GetRemoteCommit(ctx, repositoryName, remoteName, commitId, datadatdatRemoteParameters)

Get a remote commit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remoteName** | **string**| Name of the remote | 
**commitId** | **string**| Commit identifier | 
**datadatdatRemoteParameters** | [**RemoteParameters**](.md)| Remote-specific parameters | 

### Return type

[**Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteCommits

> []Commit ListRemoteCommits(ctx, repositoryName, remoteName, datadatdatRemoteParameters, optional)

List remote commits

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remoteName** | **string**| Name of the remote | 
**datadatdatRemoteParameters** | [**RemoteParameters**](.md)| Remote-specific parameters | 
 **optional** | ***ListRemoteCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRemoteCommitsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | ***[]string** | Tags (name or name&#x3D;value) to search for | 

### Return type

[**[]Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemotes

> []Remote ListRemotes(ctx, repositoryName)

Get list of remotes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 

### Return type

[**[]Remote**](Remote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemote

> Remote UpdateRemote(ctx, repositoryName, remoteName, remote)

Update remote information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remoteName** | **string**| Name of the remote | 
**remote** | [**Remote**](Remote.md)| Remote information to update | 

### Return type

[**Remote**](Remote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

