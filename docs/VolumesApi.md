# \VolumesApi

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateVolume**](VolumesApi.md#ActivateVolume) | **Post** /v1/repositories/{repositoryName}/volumes/{volumeName}/activate | Activate a volume for use by a repository (e.g. mount)
[**CreateVolume**](VolumesApi.md#CreateVolume) | **Post** /v1/repositories/{repositoryName}/volumes | Create new volume
[**DeactivateVolume**](VolumesApi.md#DeactivateVolume) | **Post** /v1/repositories/{repositoryName}/volumes/{volumeName}/deactivate | Deactivate a volume prior to its deletion (e.g. unmount)
[**DeleteVolume**](VolumesApi.md#DeleteVolume) | **Delete** /v1/repositories/{repositoryName}/volumes/{volumeName} | Remove a volume
[**GetVolume**](VolumesApi.md#GetVolume) | **Get** /v1/repositories/{repositoryName}/volumes/{volumeName} | Get info for a volume
[**GetVolumeStatus**](VolumesApi.md#GetVolumeStatus) | **Get** /v1/repositories/{repositoryName}/volumes/{volumeName}/status | Get status of a volume
[**ListVolumes**](VolumesApi.md#ListVolumes) | **Get** /v1/repositories/{repositoryName}/volumes | List volumes



## ActivateVolume

> ActivateVolume(ctx, repositoryName, volumeName).Execute()

Activate a volume for use by a repository (e.g. mount)

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
	volumeName := "volumeName_example" // string | Name of the volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumesApi.ActivateVolume(context.Background(), repositoryName, volumeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.ActivateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**volumeName** | **string** | Name of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateVolumeRequest struct via the builder pattern


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


## CreateVolume

> Volume CreateVolume(ctx, repositoryName).Volume(volume).Execute()

Create new volume

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
	volume := *openapiclient.NewVolume("Name_example", map[string]interface{}(123)) // Volume | New volume to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesApi.CreateVolume(context.Background(), repositoryName).Volume(volume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesApi.CreateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volume** | [**Volume**](Volume.md) | New volume to create | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateVolume

> DeactivateVolume(ctx, repositoryName, volumeName).Execute()

Deactivate a volume prior to its deletion (e.g. unmount)

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
	volumeName := "volumeName_example" // string | Name of the volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumesApi.DeactivateVolume(context.Background(), repositoryName, volumeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeactivateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**volumeName** | **string** | Name of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateVolumeRequest struct via the builder pattern


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


## DeleteVolume

> DeleteVolume(ctx, repositoryName, volumeName).Execute()

Remove a volume

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
	volumeName := "volumeName_example" // string | Name of the volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumesApi.DeleteVolume(context.Background(), repositoryName, volumeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**volumeName** | **string** | Name of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## GetVolume

> Volume GetVolume(ctx, repositoryName, volumeName).Execute()

Get info for a volume

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
	volumeName := "volumeName_example" // string | Name of the volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesApi.GetVolume(context.Background(), repositoryName, volumeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**volumeName** | **string** | Name of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeStatus

> VolumeStatus GetVolumeStatus(ctx, repositoryName, volumeName).Execute()

Get status of a volume

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
	volumeName := "volumeName_example" // string | Name of the volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesApi.GetVolumeStatus(context.Background(), repositoryName, volumeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetVolumeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumeStatus`: VolumeStatus
	fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetVolumeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 
**volumeName** | **string** | Name of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VolumeStatus**](VolumeStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> []Volume ListVolumes(ctx, repositoryName).Execute()

List volumes

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
	resp, r, err := apiClient.VolumesApi.ListVolumes(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesApi.ListVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

