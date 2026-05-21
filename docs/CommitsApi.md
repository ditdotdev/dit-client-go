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

> CheckoutCommit(ctx, repositoryName, commitId).Execute()

Checkout the given commit

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
	commitId := "commitId_example" // string | Commit identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommitsApi.CheckoutCommit(context.Background(), repositoryName, commitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.CheckoutCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**commitId** | **string** | Commit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutCommitRequest struct via the builder pattern


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


## CreateCommit

> Commit CreateCommit(ctx, repositoryName).Commit(commit).Execute()

Create new commit

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
	commit := *openapiclient.NewCommit("Id_example", map[string]interface{}(123)) // Commit | New commit to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommitsApi.CreateCommit(context.Background(), repositoryName).Commit(commit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.CreateCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommit`: Commit
	fmt.Fprintf(os.Stdout, "Response from `CommitsApi.CreateCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commit** | [**Commit**](Commit.md) | New commit to create | 

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

> DeleteCommit(ctx, repositoryName, commitId).Execute()

Discard a past commit

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
	commitId := "commitId_example" // string | Commit identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommitsApi.DeleteCommit(context.Background(), repositoryName, commitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.DeleteCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**commitId** | **string** | Commit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommitRequest struct via the builder pattern


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


## GetCommit

> Commit GetCommit(ctx, repositoryName, commitId).Execute()

Get information for a specific commit

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
	commitId := "commitId_example" // string | Commit identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommitsApi.GetCommit(context.Background(), repositoryName, commitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommit`: Commit
	fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**commitId** | **string** | Commit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> CommitStatus GetCommitStatus(ctx, repositoryName, commitId).Execute()

Get commit status

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
	commitId := "commitId_example" // string | Commit identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommitsApi.GetCommitStatus(context.Background(), repositoryName, commitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetCommitStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommitStatus`: CommitStatus
	fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetCommitStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**commitId** | **string** | Commit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> []Commit ListCommits(ctx, repositoryName).Tag(tag).Execute()

Get commit history for a repository

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
	tag := []string{"Inner_example"} // []string | Tags (name or name=value) to search for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommitsApi.ListCommits(context.Background(), repositoryName).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.ListCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommits`: []Commit
	fmt.Fprintf(os.Stdout, "Response from `CommitsApi.ListCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **[]string** | Tags (name or name&#x3D;value) to search for | 

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

> Commit UpdateCommit(ctx, repositoryName, commitId).Commit(commit).Execute()

Update tags for a previous commit

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
	commitId := "commitId_example" // string | Commit identifier
	commit := *openapiclient.NewCommit("Id_example", map[string]interface{}(123)) // Commit | Commit contents to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommitsApi.UpdateCommit(context.Background(), repositoryName, commitId).Commit(commit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.UpdateCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCommit`: Commit
	fmt.Fprintf(os.Stdout, "Response from `CommitsApi.UpdateCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**commitId** | **string** | Commit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commit** | [**Commit**](Commit.md) | Commit contents to update | 

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

