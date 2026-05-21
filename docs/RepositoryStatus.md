# RepositoryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCommit** | Pointer to **string** | The latest commit ID for the repository | [optional] 
**SourceCommit** | Pointer to **string** | The source commit for the current state (last checkout or commit) | [optional] 

## Methods

### NewRepositoryStatus

`func NewRepositoryStatus() *RepositoryStatus`

NewRepositoryStatus instantiates a new RepositoryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryStatusWithDefaults

`func NewRepositoryStatusWithDefaults() *RepositoryStatus`

NewRepositoryStatusWithDefaults instantiates a new RepositoryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCommit

`func (o *RepositoryStatus) GetLastCommit() string`

GetLastCommit returns the LastCommit field if non-nil, zero value otherwise.

### GetLastCommitOk

`func (o *RepositoryStatus) GetLastCommitOk() (*string, bool)`

GetLastCommitOk returns a tuple with the LastCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommit

`func (o *RepositoryStatus) SetLastCommit(v string)`

SetLastCommit sets LastCommit field to given value.

### HasLastCommit

`func (o *RepositoryStatus) HasLastCommit() bool`

HasLastCommit returns a boolean if a field has been set.

### GetSourceCommit

`func (o *RepositoryStatus) GetSourceCommit() string`

GetSourceCommit returns the SourceCommit field if non-nil, zero value otherwise.

### GetSourceCommitOk

`func (o *RepositoryStatus) GetSourceCommitOk() (*string, bool)`

GetSourceCommitOk returns a tuple with the SourceCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCommit

`func (o *RepositoryStatus) SetSourceCommit(v string)`

SetSourceCommit sets SourceCommit field to given value.

### HasSourceCommit

`func (o *RepositoryStatus) HasSourceCommit() bool`

HasSourceCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


