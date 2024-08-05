# OnedriveBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePathFilter** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**ShouldBackupOnedrive** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewOnedriveBackupParams

`func NewOnedriveBackupParams() *OnedriveBackupParams`

NewOnedriveBackupParams instantiates a new OnedriveBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnedriveBackupParamsWithDefaults

`func NewOnedriveBackupParamsWithDefaults() *OnedriveBackupParams`

NewOnedriveBackupParamsWithDefaults instantiates a new OnedriveBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePathFilter

`func (o *OnedriveBackupParams) GetFilePathFilter() FileFilteringPolicy`

GetFilePathFilter returns the FilePathFilter field if non-nil, zero value otherwise.

### GetFilePathFilterOk

`func (o *OnedriveBackupParams) GetFilePathFilterOk() (*FileFilteringPolicy, bool)`

GetFilePathFilterOk returns a tuple with the FilePathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilter

`func (o *OnedriveBackupParams) SetFilePathFilter(v FileFilteringPolicy)`

SetFilePathFilter sets FilePathFilter field to given value.

### HasFilePathFilter

`func (o *OnedriveBackupParams) HasFilePathFilter() bool`

HasFilePathFilter returns a boolean if a field has been set.

### GetShouldBackupOnedrive

`func (o *OnedriveBackupParams) GetShouldBackupOnedrive() bool`

GetShouldBackupOnedrive returns the ShouldBackupOnedrive field if non-nil, zero value otherwise.

### GetShouldBackupOnedriveOk

`func (o *OnedriveBackupParams) GetShouldBackupOnedriveOk() (*bool, bool)`

GetShouldBackupOnedriveOk returns a tuple with the ShouldBackupOnedrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBackupOnedrive

`func (o *OnedriveBackupParams) SetShouldBackupOnedrive(v bool)`

SetShouldBackupOnedrive sets ShouldBackupOnedrive field to given value.

### HasShouldBackupOnedrive

`func (o *OnedriveBackupParams) HasShouldBackupOnedrive() bool`

HasShouldBackupOnedrive returns a boolean if a field has been set.

### SetShouldBackupOnedriveNil

`func (o *OnedriveBackupParams) SetShouldBackupOnedriveNil(b bool)`

 SetShouldBackupOnedriveNil sets the value for ShouldBackupOnedrive to be an explicit nil

### UnsetShouldBackupOnedrive
`func (o *OnedriveBackupParams) UnsetShouldBackupOnedrive()`

UnsetShouldBackupOnedrive ensures that no value is present for ShouldBackupOnedrive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


