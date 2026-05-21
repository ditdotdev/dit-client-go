# ProgressEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Sequenced entry identifier | 
**Type** | **string** | Entry type | 
**Message** | Pointer to **string** | Optional message for progress step | [optional] 
**Percent** | Pointer to **int32** | Optional percent for step | [optional] 

## Methods

### NewProgressEntry

`func NewProgressEntry(id int32, type_ string, ) *ProgressEntry`

NewProgressEntry instantiates a new ProgressEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressEntryWithDefaults

`func NewProgressEntryWithDefaults() *ProgressEntry`

NewProgressEntryWithDefaults instantiates a new ProgressEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProgressEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProgressEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProgressEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *ProgressEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProgressEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProgressEntry) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *ProgressEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProgressEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProgressEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProgressEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPercent

`func (o *ProgressEntry) GetPercent() int32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ProgressEntry) GetPercentOk() (*int32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ProgressEntry) SetPercent(v int32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ProgressEntry) HasPercent() bool`

HasPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


