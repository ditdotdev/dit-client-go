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

> Repository CreateRepository(ctx).Repository(repository).Execute()

Create new repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/datadatdat/datadatdat-client-go"
)

func main() {
	repository := *openapiclient.NewRepository("Name_example", map[string]interface{}(123)) // Repository | New repository to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesApi.CreateRepository(context.Background()).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.CreateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRepository`: Repository
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.CreateRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | [**Repository**](Repository.md) | New repository to create | 

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

> DeleteRepository(ctx, repositoryName).Execute()

Remove a repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/datadatdat/datadatdat-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoriesApi.DeleteRepository(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.DeleteRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> Repository GetRepository(ctx, repositoryName).Execute()

Get info for a repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/datadatdat/datadatdat-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesApi.GetRepository(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository`: Repository
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> RepositoryStatus GetRepositoryStatus(ctx, repositoryName).Execute()

Get current status of a repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/datadatdat/datadatdat-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesApi.GetRepositoryStatus(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetRepositoryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryStatus`: RepositoryStatus
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetRepositoryStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Repository ListRepositories(ctx).Execute()

List repositories

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/datadatdat/datadatdat-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesApi.ListRepositories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.ListRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRepositories`: []Repository
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.ListRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRepositoriesRequest struct via the builder pattern


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

> Repository UpdateRepository(ctx, repositoryName).Repository(repository).Execute()

Update or rename a repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/datadatdat/datadatdat-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository
	repository := *openapiclient.NewRepository("Name_example", map[string]interface{}(123)) // Repository | New repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesApi.UpdateRepository(context.Background(), repositoryName).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.UpdateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRepository`: Repository
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.UpdateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | [**Repository**](Repository.md) | New repository | 

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

