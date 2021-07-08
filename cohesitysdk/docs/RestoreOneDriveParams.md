# RestoreOneDriveParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveOwnerVec** | Pointer to [**[]RestoreOneDriveParamsDriveOwner**](RestoreOneDriveParamsDriveOwner.md) | The list of users/groups whose drives are being restored. | [optional] 
**RestoreToOriginal** | Pointer to **NullableBool** | Whether or not all drive items are restored to original location. | [optional] 
**TargetDriveId** | Pointer to **NullableString** | The id of the drive in which items will be restored. | [optional] 
**TargetFolderPath** | Pointer to **NullableString** | All drives part of various users listed in drive_owner_vec will be restored to the drive belonging to target_user having id target_drive_id. Let&#39;s say drive_owner_vec is A and B; drive_vec of A and B is 111 and 222 respectively; target_user is C; target_drive_id is 333. The final folder-hierarchy after restore job is finished will look like this : C:333: {target_folder_path}/| |A/111/{whatever is there in restore_item_vec of 111} |B/222/{whatever is there in restore_item_vec of 222} | [optional] 
**TargetUser** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreOneDriveParams

`func NewRestoreOneDriveParams() *RestoreOneDriveParams`

NewRestoreOneDriveParams instantiates a new RestoreOneDriveParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOneDriveParamsWithDefaults

`func NewRestoreOneDriveParamsWithDefaults() *RestoreOneDriveParams`

NewRestoreOneDriveParamsWithDefaults instantiates a new RestoreOneDriveParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveOwnerVec

`func (o *RestoreOneDriveParams) GetDriveOwnerVec() []RestoreOneDriveParamsDriveOwner`

GetDriveOwnerVec returns the DriveOwnerVec field if non-nil, zero value otherwise.

### GetDriveOwnerVecOk

`func (o *RestoreOneDriveParams) GetDriveOwnerVecOk() (*[]RestoreOneDriveParamsDriveOwner, bool)`

GetDriveOwnerVecOk returns a tuple with the DriveOwnerVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveOwnerVec

`func (o *RestoreOneDriveParams) SetDriveOwnerVec(v []RestoreOneDriveParamsDriveOwner)`

SetDriveOwnerVec sets DriveOwnerVec field to given value.

### HasDriveOwnerVec

`func (o *RestoreOneDriveParams) HasDriveOwnerVec() bool`

HasDriveOwnerVec returns a boolean if a field has been set.

### SetDriveOwnerVecNil

`func (o *RestoreOneDriveParams) SetDriveOwnerVecNil(b bool)`

 SetDriveOwnerVecNil sets the value for DriveOwnerVec to be an explicit nil

### UnsetDriveOwnerVec
`func (o *RestoreOneDriveParams) UnsetDriveOwnerVec()`

UnsetDriveOwnerVec ensures that no value is present for DriveOwnerVec, not even an explicit nil
### GetRestoreToOriginal

`func (o *RestoreOneDriveParams) GetRestoreToOriginal() bool`

GetRestoreToOriginal returns the RestoreToOriginal field if non-nil, zero value otherwise.

### GetRestoreToOriginalOk

`func (o *RestoreOneDriveParams) GetRestoreToOriginalOk() (*bool, bool)`

GetRestoreToOriginalOk returns a tuple with the RestoreToOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginal

`func (o *RestoreOneDriveParams) SetRestoreToOriginal(v bool)`

SetRestoreToOriginal sets RestoreToOriginal field to given value.

### HasRestoreToOriginal

`func (o *RestoreOneDriveParams) HasRestoreToOriginal() bool`

HasRestoreToOriginal returns a boolean if a field has been set.

### SetRestoreToOriginalNil

`func (o *RestoreOneDriveParams) SetRestoreToOriginalNil(b bool)`

 SetRestoreToOriginalNil sets the value for RestoreToOriginal to be an explicit nil

### UnsetRestoreToOriginal
`func (o *RestoreOneDriveParams) UnsetRestoreToOriginal()`

UnsetRestoreToOriginal ensures that no value is present for RestoreToOriginal, not even an explicit nil
### GetTargetDriveId

`func (o *RestoreOneDriveParams) GetTargetDriveId() string`

GetTargetDriveId returns the TargetDriveId field if non-nil, zero value otherwise.

### GetTargetDriveIdOk

`func (o *RestoreOneDriveParams) GetTargetDriveIdOk() (*string, bool)`

GetTargetDriveIdOk returns a tuple with the TargetDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDriveId

`func (o *RestoreOneDriveParams) SetTargetDriveId(v string)`

SetTargetDriveId sets TargetDriveId field to given value.

### HasTargetDriveId

`func (o *RestoreOneDriveParams) HasTargetDriveId() bool`

HasTargetDriveId returns a boolean if a field has been set.

### SetTargetDriveIdNil

`func (o *RestoreOneDriveParams) SetTargetDriveIdNil(b bool)`

 SetTargetDriveIdNil sets the value for TargetDriveId to be an explicit nil

### UnsetTargetDriveId
`func (o *RestoreOneDriveParams) UnsetTargetDriveId()`

UnsetTargetDriveId ensures that no value is present for TargetDriveId, not even an explicit nil
### GetTargetFolderPath

`func (o *RestoreOneDriveParams) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *RestoreOneDriveParams) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *RestoreOneDriveParams) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *RestoreOneDriveParams) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *RestoreOneDriveParams) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *RestoreOneDriveParams) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetUser

`func (o *RestoreOneDriveParams) GetTargetUser() EntityProto`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *RestoreOneDriveParams) GetTargetUserOk() (*EntityProto, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *RestoreOneDriveParams) SetTargetUser(v EntityProto)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *RestoreOneDriveParams) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


