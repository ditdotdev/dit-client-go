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

> AbortOperation(ctx, operationId).Execute()

Abort operation

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
	operationId := "operationId_example" // string | Operation identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OperationsApi.AbortOperation(context.Background(), operationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApi.AbortOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortOperationRequest struct via the builder pattern


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


## GetOperation

> Operation GetOperation(ctx, operationId).Execute()

Get operation

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
	operationId := "operationId_example" // string | Operation identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApi.GetOperation(context.Background(), operationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApi.GetOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperation`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsApi.GetOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []ProgressEntry GetOperationProgress(ctx, operationId).LastId(lastId).Execute()

Get operation progress

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
	operationId := "operationId_example" // string | Operation identifier
	lastId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApi.GetOperationProgress(context.Background(), operationId).LastId(lastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApi.GetOperationProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationProgress`: []ProgressEntry
	fmt.Fprintf(os.Stdout, "Response from `OperationsApi.GetOperationProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastId** | **int32** |  | 

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

> []Operation ListOperations(ctx).Repository(repository).Execute()

List operations

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
	repository := "repository_example" // string | Limit to the given repository (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApi.ListOperations(context.Background()).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApi.ListOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperations`: []Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsApi.ListOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | **string** | Limit to the given repository | 

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

> Operation Pull(ctx, repositoryName, remoteName, commitId).RemoteParameters(remoteParameters).MetadataOnly(metadataOnly).Execute()

Start a pull operation

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
	remoteName := "remoteName_example" // string | Name of the remote
	commitId := "commitId_example" // string | Commit identifier
	remoteParameters := *openapiclient.NewRemoteParameters("Provider_example", map[string]interface{}(123)) // RemoteParameters | Provider specific parameters
	metadataOnly := true // bool | Transfer only tag metadata (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApi.Pull(context.Background(), repositoryName, remoteName, commitId).RemoteParameters(remoteParameters).MetadataOnly(metadataOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApi.Pull``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Pull`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsApi.Pull`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**remoteName** | **string** | Name of the remote | 
**commitId** | **string** | Commit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **remoteParameters** | [**RemoteParameters**](RemoteParameters.md) | Provider specific parameters | 
 **metadataOnly** | **bool** | Transfer only tag metadata | 

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

> Operation Push(ctx, repositoryName, remoteName, commitId).RemoteParameters(remoteParameters).MetadataOnly(metadataOnly).Execute()

Start a push operation

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
	remoteName := "remoteName_example" // string | Name of the remote
	commitId := "commitId_example" // string | Commit identifier
	remoteParameters := *openapiclient.NewRemoteParameters("Provider_example", map[string]interface{}(123)) // RemoteParameters | Provider specific parameters
	metadataOnly := true // bool | Transfer only tag metadata (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApi.Push(context.Background(), repositoryName, remoteName, commitId).RemoteParameters(remoteParameters).MetadataOnly(metadataOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApi.Push``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Push`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsApi.Push`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**remoteName** | **string** | Name of the remote | 
**commitId** | **string** | Commit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **remoteParameters** | [**RemoteParameters**](RemoteParameters.md) | Provider specific parameters | 
 **metadataOnly** | **bool** | Transfer only tag metadata | 

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

