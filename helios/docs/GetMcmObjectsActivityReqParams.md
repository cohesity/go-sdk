# GetMcmObjectsActivityReqParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIdentifiers** | Pointer to [**[]McmObjectIdentifier**](McmObjectIdentifier.md) | Specifies the list of object identifiers to filter the activity. | [optional] 
**Environments** | Pointer to **[]string** | Specifies the list of environments. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the activity event. | [optional] 
**FromTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in Unix timestamp epoch in microsecond which filters all the activity started after this value. | [optional] 
**ToTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in Unix timestamp epoch in microsecond which filters all the activity started before this value. | [optional] 
**ActivityTypes** | Pointer to **[]string** | Specifies the activity types. | [optional] 
**BackupRunParams** | Pointer to [**NullableBackupRunFilterParams**](BackupRunFilterParams.md) | Specifies the additional filters in case activity type is set to &#39;BackupRun&#39;. | [optional] 
**ArchivalRunParams** | Pointer to [**ArchivalRunFilterParams**](ArchivalRunFilterParams.md) |  | [optional] 
**RestoreParams** | Pointer to [**NullableRestoreFilterParams**](RestoreFilterParams.md) | Specifies the additional filters in case activity type is set to &#39;Restore&#39;. | [optional] 

## Methods

### NewGetMcmObjectsActivityReqParams

`func NewGetMcmObjectsActivityReqParams() *GetMcmObjectsActivityReqParams`

NewGetMcmObjectsActivityReqParams instantiates a new GetMcmObjectsActivityReqParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMcmObjectsActivityReqParamsWithDefaults

`func NewGetMcmObjectsActivityReqParamsWithDefaults() *GetMcmObjectsActivityReqParams`

NewGetMcmObjectsActivityReqParamsWithDefaults instantiates a new GetMcmObjectsActivityReqParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectIdentifiers

`func (o *GetMcmObjectsActivityReqParams) GetObjectIdentifiers() []McmObjectIdentifier`

GetObjectIdentifiers returns the ObjectIdentifiers field if non-nil, zero value otherwise.

### GetObjectIdentifiersOk

`func (o *GetMcmObjectsActivityReqParams) GetObjectIdentifiersOk() (*[]McmObjectIdentifier, bool)`

GetObjectIdentifiersOk returns a tuple with the ObjectIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdentifiers

`func (o *GetMcmObjectsActivityReqParams) SetObjectIdentifiers(v []McmObjectIdentifier)`

SetObjectIdentifiers sets ObjectIdentifiers field to given value.

### HasObjectIdentifiers

`func (o *GetMcmObjectsActivityReqParams) HasObjectIdentifiers() bool`

HasObjectIdentifiers returns a boolean if a field has been set.

### SetObjectIdentifiersNil

`func (o *GetMcmObjectsActivityReqParams) SetObjectIdentifiersNil(b bool)`

 SetObjectIdentifiersNil sets the value for ObjectIdentifiers to be an explicit nil

### UnsetObjectIdentifiers
`func (o *GetMcmObjectsActivityReqParams) UnsetObjectIdentifiers()`

UnsetObjectIdentifiers ensures that no value is present for ObjectIdentifiers, not even an explicit nil
### GetEnvironments

`func (o *GetMcmObjectsActivityReqParams) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *GetMcmObjectsActivityReqParams) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *GetMcmObjectsActivityReqParams) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *GetMcmObjectsActivityReqParams) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### SetEnvironmentsNil

`func (o *GetMcmObjectsActivityReqParams) SetEnvironmentsNil(b bool)`

 SetEnvironmentsNil sets the value for Environments to be an explicit nil

### UnsetEnvironments
`func (o *GetMcmObjectsActivityReqParams) UnsetEnvironments()`

UnsetEnvironments ensures that no value is present for Environments, not even an explicit nil
### GetStatus

`func (o *GetMcmObjectsActivityReqParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMcmObjectsActivityReqParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMcmObjectsActivityReqParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMcmObjectsActivityReqParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetMcmObjectsActivityReqParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetMcmObjectsActivityReqParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFromTimeUsecs

`func (o *GetMcmObjectsActivityReqParams) GetFromTimeUsecs() int64`

GetFromTimeUsecs returns the FromTimeUsecs field if non-nil, zero value otherwise.

### GetFromTimeUsecsOk

`func (o *GetMcmObjectsActivityReqParams) GetFromTimeUsecsOk() (*int64, bool)`

GetFromTimeUsecsOk returns a tuple with the FromTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTimeUsecs

`func (o *GetMcmObjectsActivityReqParams) SetFromTimeUsecs(v int64)`

SetFromTimeUsecs sets FromTimeUsecs field to given value.

### HasFromTimeUsecs

`func (o *GetMcmObjectsActivityReqParams) HasFromTimeUsecs() bool`

HasFromTimeUsecs returns a boolean if a field has been set.

### SetFromTimeUsecsNil

`func (o *GetMcmObjectsActivityReqParams) SetFromTimeUsecsNil(b bool)`

 SetFromTimeUsecsNil sets the value for FromTimeUsecs to be an explicit nil

### UnsetFromTimeUsecs
`func (o *GetMcmObjectsActivityReqParams) UnsetFromTimeUsecs()`

UnsetFromTimeUsecs ensures that no value is present for FromTimeUsecs, not even an explicit nil
### GetToTimeUsecs

`func (o *GetMcmObjectsActivityReqParams) GetToTimeUsecs() int64`

GetToTimeUsecs returns the ToTimeUsecs field if non-nil, zero value otherwise.

### GetToTimeUsecsOk

`func (o *GetMcmObjectsActivityReqParams) GetToTimeUsecsOk() (*int64, bool)`

GetToTimeUsecsOk returns a tuple with the ToTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTimeUsecs

`func (o *GetMcmObjectsActivityReqParams) SetToTimeUsecs(v int64)`

SetToTimeUsecs sets ToTimeUsecs field to given value.

### HasToTimeUsecs

`func (o *GetMcmObjectsActivityReqParams) HasToTimeUsecs() bool`

HasToTimeUsecs returns a boolean if a field has been set.

### SetToTimeUsecsNil

`func (o *GetMcmObjectsActivityReqParams) SetToTimeUsecsNil(b bool)`

 SetToTimeUsecsNil sets the value for ToTimeUsecs to be an explicit nil

### UnsetToTimeUsecs
`func (o *GetMcmObjectsActivityReqParams) UnsetToTimeUsecs()`

UnsetToTimeUsecs ensures that no value is present for ToTimeUsecs, not even an explicit nil
### GetActivityTypes

`func (o *GetMcmObjectsActivityReqParams) GetActivityTypes() []string`

GetActivityTypes returns the ActivityTypes field if non-nil, zero value otherwise.

### GetActivityTypesOk

`func (o *GetMcmObjectsActivityReqParams) GetActivityTypesOk() (*[]string, bool)`

GetActivityTypesOk returns a tuple with the ActivityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypes

`func (o *GetMcmObjectsActivityReqParams) SetActivityTypes(v []string)`

SetActivityTypes sets ActivityTypes field to given value.

### HasActivityTypes

`func (o *GetMcmObjectsActivityReqParams) HasActivityTypes() bool`

HasActivityTypes returns a boolean if a field has been set.

### SetActivityTypesNil

`func (o *GetMcmObjectsActivityReqParams) SetActivityTypesNil(b bool)`

 SetActivityTypesNil sets the value for ActivityTypes to be an explicit nil

### UnsetActivityTypes
`func (o *GetMcmObjectsActivityReqParams) UnsetActivityTypes()`

UnsetActivityTypes ensures that no value is present for ActivityTypes, not even an explicit nil
### GetBackupRunParams

`func (o *GetMcmObjectsActivityReqParams) GetBackupRunParams() BackupRunFilterParams`

GetBackupRunParams returns the BackupRunParams field if non-nil, zero value otherwise.

### GetBackupRunParamsOk

`func (o *GetMcmObjectsActivityReqParams) GetBackupRunParamsOk() (*BackupRunFilterParams, bool)`

GetBackupRunParamsOk returns a tuple with the BackupRunParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunParams

`func (o *GetMcmObjectsActivityReqParams) SetBackupRunParams(v BackupRunFilterParams)`

SetBackupRunParams sets BackupRunParams field to given value.

### HasBackupRunParams

`func (o *GetMcmObjectsActivityReqParams) HasBackupRunParams() bool`

HasBackupRunParams returns a boolean if a field has been set.

### SetBackupRunParamsNil

`func (o *GetMcmObjectsActivityReqParams) SetBackupRunParamsNil(b bool)`

 SetBackupRunParamsNil sets the value for BackupRunParams to be an explicit nil

### UnsetBackupRunParams
`func (o *GetMcmObjectsActivityReqParams) UnsetBackupRunParams()`

UnsetBackupRunParams ensures that no value is present for BackupRunParams, not even an explicit nil
### GetArchivalRunParams

`func (o *GetMcmObjectsActivityReqParams) GetArchivalRunParams() ArchivalRunFilterParams`

GetArchivalRunParams returns the ArchivalRunParams field if non-nil, zero value otherwise.

### GetArchivalRunParamsOk

`func (o *GetMcmObjectsActivityReqParams) GetArchivalRunParamsOk() (*ArchivalRunFilterParams, bool)`

GetArchivalRunParamsOk returns a tuple with the ArchivalRunParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalRunParams

`func (o *GetMcmObjectsActivityReqParams) SetArchivalRunParams(v ArchivalRunFilterParams)`

SetArchivalRunParams sets ArchivalRunParams field to given value.

### HasArchivalRunParams

`func (o *GetMcmObjectsActivityReqParams) HasArchivalRunParams() bool`

HasArchivalRunParams returns a boolean if a field has been set.

### GetRestoreParams

`func (o *GetMcmObjectsActivityReqParams) GetRestoreParams() RestoreFilterParams`

GetRestoreParams returns the RestoreParams field if non-nil, zero value otherwise.

### GetRestoreParamsOk

`func (o *GetMcmObjectsActivityReqParams) GetRestoreParamsOk() (*RestoreFilterParams, bool)`

GetRestoreParamsOk returns a tuple with the RestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreParams

`func (o *GetMcmObjectsActivityReqParams) SetRestoreParams(v RestoreFilterParams)`

SetRestoreParams sets RestoreParams field to given value.

### HasRestoreParams

`func (o *GetMcmObjectsActivityReqParams) HasRestoreParams() bool`

HasRestoreParams returns a boolean if a field has been set.

### SetRestoreParamsNil

`func (o *GetMcmObjectsActivityReqParams) SetRestoreParamsNil(b bool)`

 SetRestoreParamsNil sets the value for RestoreParams to be an explicit nil

### UnsetRestoreParams
`func (o *GetMcmObjectsActivityReqParams) UnsetRestoreParams()`

UnsetRestoreParams ensures that no value is present for RestoreParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


