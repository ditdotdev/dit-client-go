# \RepositoriesApi

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepository**](RepositoriesApi.md#CreateRepository) | **Post** /v1/repositories | Create new repository
[**DeleteRepository**](RepositoriesApi.md#DeleteRepository) | **Delete** /v1/repositories/{repositoryName} | Remove a repository
[**GetRepository**](RepositoriesApi.md#GetRepository) | **Get** /v1/repositories/{repositoryName} | Get info for a repository
[**GetRepositoryStatus**](RepositoriesApi.md#GetRepositoryStatus) | **Get** /v1/repositories/{repositoryName}/status | Get current status of a repository
[**ListRepositories**](RepositoriesApi.md#ListRepositories) | **Get** /v1/repositories | List repositories
[**UpdateRepository**](RepositoriesApi.md#UpdateRepository) | **Post** /v1/repositories/{repositoryName} | Update or rename a repository



## CreateRepository

> Repository CreateRepository(ctx, repository)

Create new repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | [**Repository**](Repository.md)| New repository to create | 

### Return type

[**Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepository

> DeleteRepository(ctx, repositoryName)

Remove a repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 

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


## GetRepository

> Repository GetRepository(ctx, repositoryName)

Get info for a repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 

### Return type

[**Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryStatus

> RepositoryStatus GetRepositoryStatus(ctx, repositoryName)

Get current status of a repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 

### Return type

[**RepositoryStatus**](RepositoryStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepositories

> []Repository ListRepositories(ctx, )

List repositories

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepository

> Repository UpdateRepository(ctx, repositoryName, repository)

Update or rename a repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**repository** | [**Repository**](Repository.md)| New repository | 

### Return type

[**Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

