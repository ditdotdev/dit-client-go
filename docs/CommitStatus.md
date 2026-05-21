# CommitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalSize** | **int64** | Logical size of data referenced by commit | 
**ActualSize** | **int64** | Actual size of data referenced by commit | 
**UniqueSize** | **int64** | Amount of data uniquely held by this commit | 
**Ready** | **bool** | Whether this commit can be used as the source of an operation or whether it&#39;s still being created | 
**Error** | Pointer to **string** | If commit failed to be created, error string explaining why | [optional] 

## Methods

### NewCommitStatus

`func NewCommitStatus(logicalSize int64, actualSize int64, uniqueSize int64, ready bool, ) *CommitStatus`

NewCommitStatus instantiates a new CommitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitStatusWithDefaults

`func NewCommitStatusWithDefaults() *CommitStatus`

NewCommitStatusWithDefaults instantiates a new CommitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalSize

`func (o *CommitStatus) GetLogicalSize() int64`

GetLogicalSize returns the LogicalSize field if non-nil, zero value otherwise.

### GetLogicalSizeOk

`func (o *CommitStatus) GetLogicalSizeOk() (*int64, bool)`

GetLogicalSizeOk returns a tuple with the LogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSize

`func (o *CommitStatus) SetLogicalSize(v int64)`

SetLogicalSize sets LogicalSize field to given value.


### GetActualSize

`func (o *CommitStatus) GetActualSize() int64`

GetActualSize returns the ActualSize field if non-nil, zero value otherwise.

### GetActualSizeOk

`func (o *CommitStatus) GetActualSizeOk() (*int64, bool)`

GetActualSizeOk returns a tuple with the ActualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSize

`func (o *CommitStatus) SetActualSize(v int64)`

SetActualSize sets ActualSize field to given value.


### GetUniqueSize

`func (o *CommitStatus) GetUniqueSize() int64`

GetUniqueSize returns the UniqueSize field if non-nil, zero value otherwise.

### GetUniqueSizeOk

`func (o *CommitStatus) GetUniqueSizeOk() (*int64, bool)`

GetUniqueSizeOk returns a tuple with the UniqueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSize

`func (o *CommitStatus) SetUniqueSize(v int64)`

SetUniqueSize sets UniqueSize field to given value.


### GetReady

`func (o *CommitStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *CommitStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *CommitStatus) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetError

`func (o *CommitStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CommitStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CommitStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CommitStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


