# VolumeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Volume name | 
**LogicalSize** | **int64** | Logical size consumed by the volume | 
**ActualSize** | **int64** | Actual (compressed) size used by the volume | 
**Properties** | **map[string]interface{}** | Client-specific properties | 
**Ready** | **bool** | True if the volume is ready for use in a runtime environmemnt | 
**Error** | **string** | Optional error message if volume asynchronously failed to be created | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


