# RestoreSiteParamsSiteOwnerDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEntireDriveRequired** | Pointer to **NullableBool** | Specify if the entire drive is to be restored. This field should be false if restore_item_vec size &gt; 0. | [optional] 
**RestoreDriveId** | Pointer to **NullableString** | Id of the drive whose items are being restored. | [optional] 
**RestoreDriveName** | Pointer to **NullableString** | Specifies the name of the drive whos items are being restored. NOTE: For restore either the drive Id or the name must be populated. | [optional] 
**RestorePathVec** | Pointer to [**[]RestoreSiteParamsDriveItem**](RestoreSiteParamsDriveItem.md) | List of drive paths that need to be restored. | [optional] 

## Methods

### NewRestoreSiteParamsSiteOwnerDrive

`func NewRestoreSiteParamsSiteOwnerDrive() *RestoreSiteParamsSiteOwnerDrive`

NewRestoreSiteParamsSiteOwnerDrive instantiates a new RestoreSiteParamsSiteOwnerDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSiteParamsSiteOwnerDriveWithDefaults

`func NewRestoreSiteParamsSiteOwnerDriveWithDefaults() *RestoreSiteParamsSiteOwnerDrive`

NewRestoreSiteParamsSiteOwnerDriveWithDefaults instantiates a new RestoreSiteParamsSiteOwnerDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEntireDriveRequired

`func (o *RestoreSiteParamsSiteOwnerDrive) GetIsEntireDriveRequired() bool`

GetIsEntireDriveRequired returns the IsEntireDriveRequired field if non-nil, zero value otherwise.

### GetIsEntireDriveRequiredOk

`func (o *RestoreSiteParamsSiteOwnerDrive) GetIsEntireDriveRequiredOk() (*bool, bool)`

GetIsEntireDriveRequiredOk returns a tuple with the IsEntireDriveRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntireDriveRequired

`func (o *RestoreSiteParamsSiteOwnerDrive) SetIsEntireDriveRequired(v bool)`

SetIsEntireDriveRequired sets IsEntireDriveRequired field to given value.

### HasIsEntireDriveRequired

`func (o *RestoreSiteParamsSiteOwnerDrive) HasIsEntireDriveRequired() bool`

HasIsEntireDriveRequired returns a boolean if a field has been set.

### SetIsEntireDriveRequiredNil

`func (o *RestoreSiteParamsSiteOwnerDrive) SetIsEntireDriveRequiredNil(b bool)`

 SetIsEntireDriveRequiredNil sets the value for IsEntireDriveRequired to be an explicit nil

### UnsetIsEntireDriveRequired
`func (o *RestoreSiteParamsSiteOwnerDrive) UnsetIsEntireDriveRequired()`

UnsetIsEntireDriveRequired ensures that no value is present for IsEntireDriveRequired, not even an explicit nil
### GetRestoreDriveId

`func (o *RestoreSiteParamsSiteOwnerDrive) GetRestoreDriveId() string`

GetRestoreDriveId returns the RestoreDriveId field if non-nil, zero value otherwise.

### GetRestoreDriveIdOk

`func (o *RestoreSiteParamsSiteOwnerDrive) GetRestoreDriveIdOk() (*string, bool)`

GetRestoreDriveIdOk returns a tuple with the RestoreDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDriveId

`func (o *RestoreSiteParamsSiteOwnerDrive) SetRestoreDriveId(v string)`

SetRestoreDriveId sets RestoreDriveId field to given value.

### HasRestoreDriveId

`func (o *RestoreSiteParamsSiteOwnerDrive) HasRestoreDriveId() bool`

HasRestoreDriveId returns a boolean if a field has been set.

### SetRestoreDriveIdNil

`func (o *RestoreSiteParamsSiteOwnerDrive) SetRestoreDriveIdNil(b bool)`

 SetRestoreDriveIdNil sets the value for RestoreDriveId to be an explicit nil

### UnsetRestoreDriveId
`func (o *RestoreSiteParamsSiteOwnerDrive) UnsetRestoreDriveId()`

UnsetRestoreDriveId ensures that no value is present for RestoreDriveId, not even an explicit nil
### GetRestoreDriveName

`func (o *RestoreSiteParamsSiteOwnerDrive) GetRestoreDriveName() string`

GetRestoreDriveName returns the RestoreDriveName field if non-nil, zero value otherwise.

### GetRestoreDriveNameOk

`func (o *RestoreSiteParamsSiteOwnerDrive) GetRestoreDriveNameOk() (*string, bool)`

GetRestoreDriveNameOk returns a tuple with the RestoreDriveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDriveName

`func (o *RestoreSiteParamsSiteOwnerDrive) SetRestoreDriveName(v string)`

SetRestoreDriveName sets RestoreDriveName field to given value.

### HasRestoreDriveName

`func (o *RestoreSiteParamsSiteOwnerDrive) HasRestoreDriveName() bool`

HasRestoreDriveName returns a boolean if a field has been set.

### SetRestoreDriveNameNil

`func (o *RestoreSiteParamsSiteOwnerDrive) SetRestoreDriveNameNil(b bool)`

 SetRestoreDriveNameNil sets the value for RestoreDriveName to be an explicit nil

### UnsetRestoreDriveName
`func (o *RestoreSiteParamsSiteOwnerDrive) UnsetRestoreDriveName()`

UnsetRestoreDriveName ensures that no value is present for RestoreDriveName, not even an explicit nil
### GetRestorePathVec

`func (o *RestoreSiteParamsSiteOwnerDrive) GetRestorePathVec() []RestoreSiteParamsDriveItem`

GetRestorePathVec returns the RestorePathVec field if non-nil, zero value otherwise.

### GetRestorePathVecOk

`func (o *RestoreSiteParamsSiteOwnerDrive) GetRestorePathVecOk() (*[]RestoreSiteParamsDriveItem, bool)`

GetRestorePathVecOk returns a tuple with the RestorePathVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePathVec

`func (o *RestoreSiteParamsSiteOwnerDrive) SetRestorePathVec(v []RestoreSiteParamsDriveItem)`

SetRestorePathVec sets RestorePathVec field to given value.

### HasRestorePathVec

`func (o *RestoreSiteParamsSiteOwnerDrive) HasRestorePathVec() bool`

HasRestorePathVec returns a boolean if a field has been set.

### SetRestorePathVecNil

`func (o *RestoreSiteParamsSiteOwnerDrive) SetRestorePathVecNil(b bool)`

 SetRestorePathVecNil sets the value for RestorePathVec to be an explicit nil

### UnsetRestorePathVec
`func (o *RestoreSiteParamsSiteOwnerDrive) UnsetRestorePathVec()`

UnsetRestorePathVec ensures that no value is present for RestorePathVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


