# RestoreOneDriveParamsDriveOwnerDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEntireDriveRequired** | Pointer to **NullableBool** | Specify if the entire drive is to be restored. This field should be false if restore_item_vec size &gt; 0. | [optional] 
**RestoreDriveId** | Pointer to **NullableString** | Id of the drive whose items are being restored. | [optional] 
**RestoreItemVec** | Pointer to [**[]RestoreOneDriveParamsDriveItem**](RestoreOneDriveParamsDriveItem.md) | List of drive paths that need to be restored. | [optional] 

## Methods

### NewRestoreOneDriveParamsDriveOwnerDrive

`func NewRestoreOneDriveParamsDriveOwnerDrive() *RestoreOneDriveParamsDriveOwnerDrive`

NewRestoreOneDriveParamsDriveOwnerDrive instantiates a new RestoreOneDriveParamsDriveOwnerDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOneDriveParamsDriveOwnerDriveWithDefaults

`func NewRestoreOneDriveParamsDriveOwnerDriveWithDefaults() *RestoreOneDriveParamsDriveOwnerDrive`

NewRestoreOneDriveParamsDriveOwnerDriveWithDefaults instantiates a new RestoreOneDriveParamsDriveOwnerDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEntireDriveRequired

`func (o *RestoreOneDriveParamsDriveOwnerDrive) GetIsEntireDriveRequired() bool`

GetIsEntireDriveRequired returns the IsEntireDriveRequired field if non-nil, zero value otherwise.

### GetIsEntireDriveRequiredOk

`func (o *RestoreOneDriveParamsDriveOwnerDrive) GetIsEntireDriveRequiredOk() (*bool, bool)`

GetIsEntireDriveRequiredOk returns a tuple with the IsEntireDriveRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntireDriveRequired

`func (o *RestoreOneDriveParamsDriveOwnerDrive) SetIsEntireDriveRequired(v bool)`

SetIsEntireDriveRequired sets IsEntireDriveRequired field to given value.

### HasIsEntireDriveRequired

`func (o *RestoreOneDriveParamsDriveOwnerDrive) HasIsEntireDriveRequired() bool`

HasIsEntireDriveRequired returns a boolean if a field has been set.

### SetIsEntireDriveRequiredNil

`func (o *RestoreOneDriveParamsDriveOwnerDrive) SetIsEntireDriveRequiredNil(b bool)`

 SetIsEntireDriveRequiredNil sets the value for IsEntireDriveRequired to be an explicit nil

### UnsetIsEntireDriveRequired
`func (o *RestoreOneDriveParamsDriveOwnerDrive) UnsetIsEntireDriveRequired()`

UnsetIsEntireDriveRequired ensures that no value is present for IsEntireDriveRequired, not even an explicit nil
### GetRestoreDriveId

`func (o *RestoreOneDriveParamsDriveOwnerDrive) GetRestoreDriveId() string`

GetRestoreDriveId returns the RestoreDriveId field if non-nil, zero value otherwise.

### GetRestoreDriveIdOk

`func (o *RestoreOneDriveParamsDriveOwnerDrive) GetRestoreDriveIdOk() (*string, bool)`

GetRestoreDriveIdOk returns a tuple with the RestoreDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDriveId

`func (o *RestoreOneDriveParamsDriveOwnerDrive) SetRestoreDriveId(v string)`

SetRestoreDriveId sets RestoreDriveId field to given value.

### HasRestoreDriveId

`func (o *RestoreOneDriveParamsDriveOwnerDrive) HasRestoreDriveId() bool`

HasRestoreDriveId returns a boolean if a field has been set.

### SetRestoreDriveIdNil

`func (o *RestoreOneDriveParamsDriveOwnerDrive) SetRestoreDriveIdNil(b bool)`

 SetRestoreDriveIdNil sets the value for RestoreDriveId to be an explicit nil

### UnsetRestoreDriveId
`func (o *RestoreOneDriveParamsDriveOwnerDrive) UnsetRestoreDriveId()`

UnsetRestoreDriveId ensures that no value is present for RestoreDriveId, not even an explicit nil
### GetRestoreItemVec

`func (o *RestoreOneDriveParamsDriveOwnerDrive) GetRestoreItemVec() []RestoreOneDriveParamsDriveItem`

GetRestoreItemVec returns the RestoreItemVec field if non-nil, zero value otherwise.

### GetRestoreItemVecOk

`func (o *RestoreOneDriveParamsDriveOwnerDrive) GetRestoreItemVecOk() (*[]RestoreOneDriveParamsDriveItem, bool)`

GetRestoreItemVecOk returns a tuple with the RestoreItemVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreItemVec

`func (o *RestoreOneDriveParamsDriveOwnerDrive) SetRestoreItemVec(v []RestoreOneDriveParamsDriveItem)`

SetRestoreItemVec sets RestoreItemVec field to given value.

### HasRestoreItemVec

`func (o *RestoreOneDriveParamsDriveOwnerDrive) HasRestoreItemVec() bool`

HasRestoreItemVec returns a boolean if a field has been set.

### SetRestoreItemVecNil

`func (o *RestoreOneDriveParamsDriveOwnerDrive) SetRestoreItemVecNil(b bool)`

 SetRestoreItemVecNil sets the value for RestoreItemVec to be an explicit nil

### UnsetRestoreItemVec
`func (o *RestoreOneDriveParamsDriveOwnerDrive) UnsetRestoreItemVec()`

UnsetRestoreItemVec ensures that no value is present for RestoreItemVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


