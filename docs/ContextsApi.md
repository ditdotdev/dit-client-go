# \ContextsApi

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContext**](ContextsApi.md#GetContext) | **Get** /v1/context | Get current context



## GetContext

> Context GetContext(ctx).Execute()

Get current context

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
	resp, r, err := apiClient.ContextsApi.GetContext(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContext`: Context
	fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetContext`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextRequest struct via the builder pattern


### Return type

[**Context**](Context.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

