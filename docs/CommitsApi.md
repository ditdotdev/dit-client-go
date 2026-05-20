# \CommitsApi

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckoutCommit**](CommitsApi.md#CheckoutCommit) | **Post** /v1/repositories/{repositoryName}/commits/{commitId}/checkout | Checkout the given commit
[**CreateCommit**](CommitsApi.md#CreateCommit) | **Post** /v1/repositories/{repositoryName}/commits | Create new commit
[**DeleteCommit**](CommitsApi.md#DeleteCommit) | **Delete** /v1/repositories/{repositoryName}/commits/{commitId} | Discard a past commit
[**GetCommit**](CommitsApi.md#GetCommit) | **Get** /v1/repositories/{repositoryName}/commits/{commitId} | Get information for a specific commit
[**GetCommitStatus**](CommitsApi.md#GetCommitStatus) | **Get** /v1/repositories/{repositoryName}/commits/{commitId}/status | Get commit status
[**ListCommits**](CommitsApi.md#ListCommits) | **Get** /v1/repositories/{repositoryName}/commits | Get commit history for a repository
[**UpdateCommit**](CommitsApi.md#UpdateCommit) | **Post** /v1/repositories/{repositoryName}/commits/{commitId} | Update tags for a previous commit



## CheckoutCommit

> CheckoutCommit(ctx, repositoryName, commitId)

Checkout the given commit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**commitId** | **string**| Commit identifier | 

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


## CreateCommit

> Commit CreateCommit(ctx, repositoryName, commit)

Create new commit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**commit** | [**Commit**](Commit.md)| New commit to create | 

### Return type

[**Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommit

> DeleteCommit(ctx, repositoryName, commitId)

Discard a past commit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**commitId** | **string**| Commit identifier | 

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


## GetCommit

> Commit GetCommit(ctx, repositoryName, commitId)

Get information for a specific commit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**commitId** | **string**| Commit identifier | 

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


## GetCommitStatus

> CommitStatus GetCommitStatus(ctx, repositoryName, commitId)

Get commit status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**commitId** | **string**| Commit identifier | 

### Return type

[**CommitStatus**](CommitStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommits

> []Commit ListCommits(ctx, repositoryName, optional)

Get commit history for a repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
 **optional** | ***ListCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCommitsOpts struct


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


## UpdateCommit

> Commit UpdateCommit(ctx, repositoryName, commitId, commit)

Update tags for a previous commit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**commitId** | **string**| Commit identifier | 
**commit** | [**Commit**](Commit.md)| Commit contents to update | 

### Return type

[**Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

