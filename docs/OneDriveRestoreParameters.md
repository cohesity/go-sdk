# OneDriveRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveOwnerList** | Pointer to [**[]OneDriveOwner**](OneDriveOwner.md) | Specifies the list of Drive owners which are to be restored along with the details of their drives. | [optional] 
**RestoreToOriginalDrive** | Pointer to **NullableBool** | Specifies whether the objects are to be restored to the original drive. | [optional] 
**TargetDriveId** | Pointer to **NullableString** | Specifies the Drive Id of the target user where the OneDrive items are to be recovered. | [optional] 
**TargetFolderPath** | Pointer to **NullableString** | Specifies the target folder path within the drive where recovery has to be done. | [optional] 
**TargetUser** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 

## Methods

### NewOneDriveRestoreParameters

`func NewOneDriveRestoreParameters() *OneDriveRestoreParameters`

NewOneDriveRestoreParameters instantiates a new OneDriveRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveRestoreParametersWithDefaults

`func NewOneDriveRestoreParametersWithDefaults() *OneDriveRestoreParameters`

NewOneDriveRestoreParametersWithDefaults instantiates a new OneDriveRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveOwnerList

`func (o *OneDriveRestoreParameters) GetDriveOwnerList() []OneDriveOwner`

GetDriveOwnerList returns the DriveOwnerList field if non-nil, zero value otherwise.

### GetDriveOwnerListOk

`func (o *OneDriveRestoreParameters) GetDriveOwnerListOk() (*[]OneDriveOwner, bool)`

GetDriveOwnerListOk returns a tuple with the DriveOwnerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveOwnerList

`func (o *OneDriveRestoreParameters) SetDriveOwnerList(v []OneDriveOwner)`

SetDriveOwnerList sets DriveOwnerList field to given value.

### HasDriveOwnerList

`func (o *OneDriveRestoreParameters) HasDriveOwnerList() bool`

HasDriveOwnerList returns a boolean if a field has been set.

### SetDriveOwnerListNil

`func (o *OneDriveRestoreParameters) SetDriveOwnerListNil(b bool)`

 SetDriveOwnerListNil sets the value for DriveOwnerList to be an explicit nil

### UnsetDriveOwnerList
`func (o *OneDriveRestoreParameters) UnsetDriveOwnerList()`

UnsetDriveOwnerList ensures that no value is present for DriveOwnerList, not even an explicit nil
### GetRestoreToOriginalDrive

`func (o *OneDriveRestoreParameters) GetRestoreToOriginalDrive() bool`

GetRestoreToOriginalDrive returns the RestoreToOriginalDrive field if non-nil, zero value otherwise.

### GetRestoreToOriginalDriveOk

`func (o *OneDriveRestoreParameters) GetRestoreToOriginalDriveOk() (*bool, bool)`

GetRestoreToOriginalDriveOk returns a tuple with the RestoreToOriginalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginalDrive

`func (o *OneDriveRestoreParameters) SetRestoreToOriginalDrive(v bool)`

SetRestoreToOriginalDrive sets RestoreToOriginalDrive field to given value.

### HasRestoreToOriginalDrive

`func (o *OneDriveRestoreParameters) HasRestoreToOriginalDrive() bool`

HasRestoreToOriginalDrive returns a boolean if a field has been set.

### SetRestoreToOriginalDriveNil

`func (o *OneDriveRestoreParameters) SetRestoreToOriginalDriveNil(b bool)`

 SetRestoreToOriginalDriveNil sets the value for RestoreToOriginalDrive to be an explicit nil

### UnsetRestoreToOriginalDrive
`func (o *OneDriveRestoreParameters) UnsetRestoreToOriginalDrive()`

UnsetRestoreToOriginalDrive ensures that no value is present for RestoreToOriginalDrive, not even an explicit nil
### GetTargetDriveId

`func (o *OneDriveRestoreParameters) GetTargetDriveId() string`

GetTargetDriveId returns the TargetDriveId field if non-nil, zero value otherwise.

### GetTargetDriveIdOk

`func (o *OneDriveRestoreParameters) GetTargetDriveIdOk() (*string, bool)`

GetTargetDriveIdOk returns a tuple with the TargetDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDriveId

`func (o *OneDriveRestoreParameters) SetTargetDriveId(v string)`

SetTargetDriveId sets TargetDriveId field to given value.

### HasTargetDriveId

`func (o *OneDriveRestoreParameters) HasTargetDriveId() bool`

HasTargetDriveId returns a boolean if a field has been set.

### SetTargetDriveIdNil

`func (o *OneDriveRestoreParameters) SetTargetDriveIdNil(b bool)`

 SetTargetDriveIdNil sets the value for TargetDriveId to be an explicit nil

### UnsetTargetDriveId
`func (o *OneDriveRestoreParameters) UnsetTargetDriveId()`

UnsetTargetDriveId ensures that no value is present for TargetDriveId, not even an explicit nil
### GetTargetFolderPath

`func (o *OneDriveRestoreParameters) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *OneDriveRestoreParameters) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *OneDriveRestoreParameters) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *OneDriveRestoreParameters) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *OneDriveRestoreParameters) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *OneDriveRestoreParameters) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetUser

`func (o *OneDriveRestoreParameters) GetTargetUser() ProtectionSource`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *OneDriveRestoreParameters) GetTargetUserOk() (*ProtectionSource, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *OneDriveRestoreParameters) SetTargetUser(v ProtectionSource)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *OneDriveRestoreParameters) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


