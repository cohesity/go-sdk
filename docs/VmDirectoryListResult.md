# VmDirectoryListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | Cookie is used for paginating results. If ReadVMDirResult is returning partial results, this field will be set. Supplying this cookie will resume listing from where this result left off. | [optional] 
**Entries** | Pointer to [**[]VmDirEntry**](VmDirEntry.md) | Entries is the array of files and folders that are immediate children of the parent directory specified in the request. | [optional] 

## Methods

### NewVmDirectoryListResult

`func NewVmDirectoryListResult() *VmDirectoryListResult`

NewVmDirectoryListResult instantiates a new VmDirectoryListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDirectoryListResultWithDefaults

`func NewVmDirectoryListResultWithDefaults() *VmDirectoryListResult`

NewVmDirectoryListResultWithDefaults instantiates a new VmDirectoryListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *VmDirectoryListResult) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *VmDirectoryListResult) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *VmDirectoryListResult) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *VmDirectoryListResult) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *VmDirectoryListResult) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *VmDirectoryListResult) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetEntries

`func (o *VmDirectoryListResult) GetEntries() []VmDirEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *VmDirectoryListResult) GetEntriesOk() (*[]VmDirEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *VmDirectoryListResult) SetEntries(v []VmDirEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *VmDirectoryListResult) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### SetEntriesNil

`func (o *VmDirectoryListResult) SetEntriesNil(b bool)`

 SetEntriesNil sets the value for Entries to be an explicit nil

### UnsetEntries
`func (o *VmDirectoryListResult) UnsetEntries()`

UnsetEntries ensures that no value is present for Entries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


