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

> ActivateVolume(ctx, repositoryName, volumeName)

Activate a volume for use by a repository (e.g. mount)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**volumeName** | **string**| Name of the volume | 

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

> Volume CreateVolume(ctx, repositoryName, volume)

Create new volume

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**volume** | [**Volume**](Volume.md)| New volume to create | 

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

> DeactivateVolume(ctx, repositoryName, volumeName)

Deactivate a volume prior to its deletion (e.g. unmount)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**volumeName** | **string**| Name of the volume | 

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

> DeleteVolume(ctx, repositoryName, volumeName)

Remove a volume

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**volumeName** | **string**| Name of the volume | 

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

> Volume GetVolume(ctx, repositoryName, volumeName)

Get info for a volume

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**volumeName** | **string**| Name of the volume | 

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

> VolumeStatus GetVolumeStatus(ctx, repositoryName, volumeName)

Get status of a volume

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 
**volumeName** | **string**| Name of the volume | 

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

> []Volume ListVolumes(ctx, repositoryName)

List volumes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string**| Name of the repository | 

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

