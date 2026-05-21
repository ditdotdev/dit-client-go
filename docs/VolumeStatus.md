# VolumeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Volume name | 
**LogicalSize** | **int64** | Logical size consumed by the volume | 
**ActualSize** | **int64** | Actual (compressed) size used by the volume | 
**Properties** | **map[string]interface{}** | Client-specific properties | 
**Ready** | **bool** | True if the volume is ready for use in a runtime environment | 
**Error** | Pointer to **string** | Optional error message if volume asynchronously failed to be created | [optional] 

## Methods

### NewVolumeStatus

`func NewVolumeStatus(name string, logicalSize int64, actualSize int64, properties map[string]interface{}, ready bool, ) *VolumeStatus`

NewVolumeStatus instantiates a new VolumeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeStatusWithDefaults

`func NewVolumeStatusWithDefaults() *VolumeStatus`

NewVolumeStatusWithDefaults instantiates a new VolumeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeStatus) SetName(v string)`

SetName sets Name field to given value.


### GetLogicalSize

`func (o *VolumeStatus) GetLogicalSize() int64`

GetLogicalSize returns the LogicalSize field if non-nil, zero value otherwise.

### GetLogicalSizeOk

`func (o *VolumeStatus) GetLogicalSizeOk() (*int64, bool)`

GetLogicalSizeOk returns a tuple with the LogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSize

`func (o *VolumeStatus) SetLogicalSize(v int64)`

SetLogicalSize sets LogicalSize field to given value.


### GetActualSize

`func (o *VolumeStatus) GetActualSize() int64`

GetActualSize returns the ActualSize field if non-nil, zero value otherwise.

### GetActualSizeOk

`func (o *VolumeStatus) GetActualSizeOk() (*int64, bool)`

GetActualSizeOk returns a tuple with the ActualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSize

`func (o *VolumeStatus) SetActualSize(v int64)`

SetActualSize sets ActualSize field to given value.


### GetProperties

`func (o *VolumeStatus) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VolumeStatus) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VolumeStatus) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.


### GetReady

`func (o *VolumeStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *VolumeStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *VolumeStatus) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetError

`func (o *VolumeStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VolumeStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VolumeStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VolumeStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


