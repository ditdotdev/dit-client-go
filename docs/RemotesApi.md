# \RemotesApi

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemote**](RemotesApi.md#CreateRemote) | **Post** /v1/repositories/{repositoryName}/remotes | Create new remote
[**DeleteRemote**](RemotesApi.md#DeleteRemote) | **Delete** /v1/repositories/{repositoryName}/remotes/{remoteName} | Delete remote
[**GetRemote**](RemotesApi.md#GetRemote) | **Get** /v1/repositories/{repositoryName}/remotes/{remoteName} | Get information about a particular remote
[**GetRemoteCommit**](RemotesApi.md#GetRemoteCommit) | **Post** /v1/repositories/{repositoryName}/remotes/{remoteName}/commits/{commitId} | Get a remote commit
[**ListRemoteCommits**](RemotesApi.md#ListRemoteCommits) | **Post** /v1/repositories/{repositoryName}/remotes/{remoteName}/commits | List remote commits
[**ListRemotes**](RemotesApi.md#ListRemotes) | **Get** /v1/repositories/{repositoryName}/remotes | Get list of remotes
[**UpdateRemote**](RemotesApi.md#UpdateRemote) | **Post** /v1/repositories/{repositoryName}/remotes/{remoteName} | Update remote information



## CreateRemote

> Remote CreateRemote(ctx, repositoryName).Remote(remote).Execute()

Create new remote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ditdotdev/dit-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository
	remote := *openapiclient.NewRemote("Provider_example", "Name_example", map[string]interface{}(123)) // Remote | Remote to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesApi.CreateRemote(context.Background(), repositoryName).Remote(remote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.CreateRemote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRemote`: Remote
	fmt.Fprintf(os.Stdout, "Response from `RemotesApi.CreateRemote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remote** | [**Remote**](Remote.md) | Remote to create | 

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

> DeleteRemote(ctx, repositoryName, remoteName).Execute()

Delete remote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ditdotdev/dit-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository
	remoteName := "remoteName_example" // string | Name of the remote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemotesApi.DeleteRemote(context.Background(), repositoryName, remoteName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.DeleteRemote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**remoteName** | **string** | Name of the remote | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteRequest struct via the builder pattern


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


## GetRemote

> Remote GetRemote(ctx, repositoryName, remoteName).Execute()

Get information about a particular remote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ditdotdev/dit-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository
	remoteName := "remoteName_example" // string | Name of the remote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesApi.GetRemote(context.Background(), repositoryName, remoteName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.GetRemote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemote`: Remote
	fmt.Fprintf(os.Stdout, "Response from `RemotesApi.GetRemote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**remoteName** | **string** | Name of the remote | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> Commit GetRemoteCommit(ctx, repositoryName, remoteName, commitId).RemoteParameters(remoteParameters).Execute()

Get a remote commit



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ditdotdev/dit-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository
	remoteName := "remoteName_example" // string | Name of the remote
	commitId := "commitId_example" // string | Commit identifier
	remoteParameters := *openapiclient.NewRemoteParameters("Provider_example", map[string]interface{}(123)) // RemoteParameters | Provider-specific parameters used to reach the remote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesApi.GetRemoteCommit(context.Background(), repositoryName, remoteName, commitId).RemoteParameters(remoteParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.GetRemoteCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteCommit`: Commit
	fmt.Fprintf(os.Stdout, "Response from `RemotesApi.GetRemoteCommit`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetRemoteCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **remoteParameters** | [**RemoteParameters**](RemoteParameters.md) | Provider-specific parameters used to reach the remote | 

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


## ListRemoteCommits

> []Commit ListRemoteCommits(ctx, repositoryName, remoteName).RemoteParameters(remoteParameters).Tag(tag).Execute()

List remote commits



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ditdotdev/dit-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository
	remoteName := "remoteName_example" // string | Name of the remote
	remoteParameters := *openapiclient.NewRemoteParameters("Provider_example", map[string]interface{}(123)) // RemoteParameters | Provider-specific parameters used to reach the remote
	tag := []string{"Inner_example"} // []string | Tags (name or name=value) to search for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesApi.ListRemoteCommits(context.Background(), repositoryName, remoteName).RemoteParameters(remoteParameters).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.ListRemoteCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRemoteCommits`: []Commit
	fmt.Fprintf(os.Stdout, "Response from `RemotesApi.ListRemoteCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**remoteName** | **string** | Name of the remote | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remoteParameters** | [**RemoteParameters**](RemoteParameters.md) | Provider-specific parameters used to reach the remote | 
 **tag** | **[]string** | Tags (name or name&#x3D;value) to search for | 

### Return type

[**[]Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemotes

> []Remote ListRemotes(ctx, repositoryName).Execute()

Get list of remotes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ditdotdev/dit-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesApi.ListRemotes(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.ListRemotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRemotes`: []Remote
	fmt.Fprintf(os.Stdout, "Response from `RemotesApi.ListRemotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRemotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> Remote UpdateRemote(ctx, repositoryName, remoteName).Remote(remote).Execute()

Update remote information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ditdotdev/dit-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository
	remoteName := "remoteName_example" // string | Name of the remote
	remote := *openapiclient.NewRemote("Provider_example", "Name_example", map[string]interface{}(123)) // Remote | Remote information to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesApi.UpdateRemote(context.Background(), repositoryName, remoteName).Remote(remote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.UpdateRemote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemote`: Remote
	fmt.Fprintf(os.Stdout, "Response from `RemotesApi.UpdateRemote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**remoteName** | **string** | Name of the remote | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remote** | [**Remote**](Remote.md) | Remote information to update | 

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

