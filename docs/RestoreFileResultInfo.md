# RestoreFileResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyStats** | Pointer to [**RestoreFileCopyStats**](RestoreFileCopyStats.md) |  | [optional] 
**DestinationDir** | Pointer to **NullableString** | This is set to the destination directory where the file/directory was copied. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**RestoredFileInfo** | Pointer to [**RestoredFileInfo**](RestoredFileInfo.md) |  | [optional] 
**Status** | Pointer to **NullableInt32** | Status of the restore. | [optional] 

## Methods

### NewRestoreFileResultInfo

`func NewRestoreFileResultInfo() *RestoreFileResultInfo`

NewRestoreFileResultInfo instantiates a new RestoreFileResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFileResultInfoWithDefaults

`func NewRestoreFileResultInfoWithDefaults() *RestoreFileResultInfo`

NewRestoreFileResultInfoWithDefaults instantiates a new RestoreFileResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyStats

`func (o *RestoreFileResultInfo) GetCopyStats() RestoreFileCopyStats`

GetCopyStats returns the CopyStats field if non-nil, zero value otherwise.

### GetCopyStatsOk

`func (o *RestoreFileResultInfo) GetCopyStatsOk() (*RestoreFileCopyStats, bool)`

GetCopyStatsOk returns a tuple with the CopyStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyStats

`func (o *RestoreFileResultInfo) SetCopyStats(v RestoreFileCopyStats)`

SetCopyStats sets CopyStats field to given value.

### HasCopyStats

`func (o *RestoreFileResultInfo) HasCopyStats() bool`

HasCopyStats returns a boolean if a field has been set.

### GetDestinationDir

`func (o *RestoreFileResultInfo) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *RestoreFileResultInfo) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *RestoreFileResultInfo) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.

### HasDestinationDir

`func (o *RestoreFileResultInfo) HasDestinationDir() bool`

HasDestinationDir returns a boolean if a field has been set.

### SetDestinationDirNil

`func (o *RestoreFileResultInfo) SetDestinationDirNil(b bool)`

 SetDestinationDirNil sets the value for DestinationDir to be an explicit nil

### UnsetDestinationDir
`func (o *RestoreFileResultInfo) UnsetDestinationDir()`

UnsetDestinationDir ensures that no value is present for DestinationDir, not even an explicit nil
### GetError

`func (o *RestoreFileResultInfo) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestoreFileResultInfo) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestoreFileResultInfo) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RestoreFileResultInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRestoredFileInfo

`func (o *RestoreFileResultInfo) GetRestoredFileInfo() RestoredFileInfo`

GetRestoredFileInfo returns the RestoredFileInfo field if non-nil, zero value otherwise.

### GetRestoredFileInfoOk

`func (o *RestoreFileResultInfo) GetRestoredFileInfoOk() (*RestoredFileInfo, bool)`

GetRestoredFileInfoOk returns a tuple with the RestoredFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredFileInfo

`func (o *RestoreFileResultInfo) SetRestoredFileInfo(v RestoredFileInfo)`

SetRestoredFileInfo sets RestoredFileInfo field to given value.

### HasRestoredFileInfo

`func (o *RestoreFileResultInfo) HasRestoredFileInfo() bool`

HasRestoredFileInfo returns a boolean if a field has been set.

### GetStatus

`func (o *RestoreFileResultInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreFileResultInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreFileResultInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestoreFileResultInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RestoreFileResultInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RestoreFileResultInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


