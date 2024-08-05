# OutlookBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePathFilter** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**ShouldBackupMailbox** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewOutlookBackupParams

`func NewOutlookBackupParams() *OutlookBackupParams`

NewOutlookBackupParams instantiates a new OutlookBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookBackupParamsWithDefaults

`func NewOutlookBackupParamsWithDefaults() *OutlookBackupParams`

NewOutlookBackupParamsWithDefaults instantiates a new OutlookBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePathFilter

`func (o *OutlookBackupParams) GetFilePathFilter() FileFilteringPolicy`

GetFilePathFilter returns the FilePathFilter field if non-nil, zero value otherwise.

### GetFilePathFilterOk

`func (o *OutlookBackupParams) GetFilePathFilterOk() (*FileFilteringPolicy, bool)`

GetFilePathFilterOk returns a tuple with the FilePathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilter

`func (o *OutlookBackupParams) SetFilePathFilter(v FileFilteringPolicy)`

SetFilePathFilter sets FilePathFilter field to given value.

### HasFilePathFilter

`func (o *OutlookBackupParams) HasFilePathFilter() bool`

HasFilePathFilter returns a boolean if a field has been set.

### GetShouldBackupMailbox

`func (o *OutlookBackupParams) GetShouldBackupMailbox() bool`

GetShouldBackupMailbox returns the ShouldBackupMailbox field if non-nil, zero value otherwise.

### GetShouldBackupMailboxOk

`func (o *OutlookBackupParams) GetShouldBackupMailboxOk() (*bool, bool)`

GetShouldBackupMailboxOk returns a tuple with the ShouldBackupMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBackupMailbox

`func (o *OutlookBackupParams) SetShouldBackupMailbox(v bool)`

SetShouldBackupMailbox sets ShouldBackupMailbox field to given value.

### HasShouldBackupMailbox

`func (o *OutlookBackupParams) HasShouldBackupMailbox() bool`

HasShouldBackupMailbox returns a boolean if a field has been set.

### SetShouldBackupMailboxNil

`func (o *OutlookBackupParams) SetShouldBackupMailboxNil(b bool)`

 SetShouldBackupMailboxNil sets the value for ShouldBackupMailbox to be an explicit nil

### UnsetShouldBackupMailbox
`func (o *OutlookBackupParams) UnsetShouldBackupMailbox()`

UnsetShouldBackupMailbox ensures that no value is present for ShouldBackupMailbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


