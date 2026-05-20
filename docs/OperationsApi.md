# \OperationsApi

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortOperation**](OperationsApi.md#AbortOperation) | **Delete** /v1/operations/{operationId} | Abort operation
[**GetOperation**](OperationsApi.md#GetOperation) | **Get** /v1/operations/{operationId} | Get operation
[**GetOperationProgress**](OperationsApi.md#GetOperationProgress) | **Get** /v1/operations/{operationId}/progress | Get operation progress
[**ListOperations**](OperationsApi.md#ListOperations) | **Get** /v1/operations | List operations
[**Pull**](OperationsApi.md#Pull) | **Post** /v1/repositories/{repositoryName}/remotes/{remoteName}/commits/{commitId}/pull | Start a pull operation
[**Push**](OperationsApi.md#Push) | **Post** /v1/repositories/{repositoryName}/remotes/{remoteName}/commits/{commitId}/push | Start a push operation



## AbortOperation

> AbortOperation(ctx, operationId)

Abort operation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**| Operation identifier | 

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


## GetOperation

> Operation GetOperation(ctx, operationId)

Get operation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**| Operation identifier | 

### Return type

[**Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperationProgress

> []ProgressEntry GetOperationProgress(ctx, operationId, optional)

Get operation progress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**| Operation identifier | 
 **optional** | ***GetOperationProgressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOperationProgressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastId** | ***int32** |  | 

### Return type

[**[]ProgressEntry**](ProgressEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperations

> []Operation ListOperations(ctx, optional)

List operations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOperationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | ***string** | Limit to the given repository | 

### Return type

[**[]Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Pull

> Operation Pull(ctx, repositoryName, remoteName, commitId, remoteParameters, optional)

Start a pull operation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remoteName** | **string**| Name of the remote | 
**commitId** | **string**| Commit identifier | 
**remoteParameters** | [**RemoteParameters**](RemoteParameters.md)| Provider specific parameters | 
 **optional** | ***PullOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PullOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **metadataOnly** | ***bool** | Transfer only tag metadata | 

### Return type

[**Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Push

> Operation Push(ctx, repositoryName, remoteName, commitId, remoteParameters, optional)

Start a push operation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**remoteName** | **string**| Name of the remote | 
**commitId** | **string**| Commit identifier | 
**remoteParameters** | [**RemoteParameters**](RemoteParameters.md)| Provider specific parameters | 
 **optional** | ***PushOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PushOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **metadataOnly** | ***bool** | Transfer only tag metadata | 

### Return type

[**Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

