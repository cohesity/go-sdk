# SiteDriveInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveId** | Pointer to **NullableString** | Specifies the Id of the Drive. | [optional] 
**DriveItemList** | Pointer to [**[]SiteDriveItem**](SiteDriveItem.md) | Specifies the Drive items such as files/folders. | [optional] 
**DriveName** | Pointer to **NullableString** | Specifies the drive name for the document repsitory. Incase of drive Id not present, this field must be populated to trigger restore. | [optional] 
**RestoreEntireDrive** | Pointer to **NullableBool** | Specifies whether entire drive is to be restored. This should be set to false if specific drive items are to be restored within &#39;DriveItemList&#39;. | [optional] 

## Methods

### NewSiteDriveInfo

`func NewSiteDriveInfo() *SiteDriveInfo`

NewSiteDriveInfo instantiates a new SiteDriveInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteDriveInfoWithDefaults

`func NewSiteDriveInfoWithDefaults() *SiteDriveInfo`

NewSiteDriveInfoWithDefaults instantiates a new SiteDriveInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveId

`func (o *SiteDriveInfo) GetDriveId() string`

GetDriveId returns the DriveId field if non-nil, zero value otherwise.

### GetDriveIdOk

`func (o *SiteDriveInfo) GetDriveIdOk() (*string, bool)`

GetDriveIdOk returns a tuple with the DriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveId

`func (o *SiteDriveInfo) SetDriveId(v string)`

SetDriveId sets DriveId field to given value.

### HasDriveId

`func (o *SiteDriveInfo) HasDriveId() bool`

HasDriveId returns a boolean if a field has been set.

### SetDriveIdNil

`func (o *SiteDriveInfo) SetDriveIdNil(b bool)`

 SetDriveIdNil sets the value for DriveId to be an explicit nil

### UnsetDriveId
`func (o *SiteDriveInfo) UnsetDriveId()`

UnsetDriveId ensures that no value is present for DriveId, not even an explicit nil
### GetDriveItemList

`func (o *SiteDriveInfo) GetDriveItemList() []SiteDriveItem`

GetDriveItemList returns the DriveItemList field if non-nil, zero value otherwise.

### GetDriveItemListOk

`func (o *SiteDriveInfo) GetDriveItemListOk() (*[]SiteDriveItem, bool)`

GetDriveItemListOk returns a tuple with the DriveItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveItemList

`func (o *SiteDriveInfo) SetDriveItemList(v []SiteDriveItem)`

SetDriveItemList sets DriveItemList field to given value.

### HasDriveItemList

`func (o *SiteDriveInfo) HasDriveItemList() bool`

HasDriveItemList returns a boolean if a field has been set.

### SetDriveItemListNil

`func (o *SiteDriveInfo) SetDriveItemListNil(b bool)`

 SetDriveItemListNil sets the value for DriveItemList to be an explicit nil

### UnsetDriveItemList
`func (o *SiteDriveInfo) UnsetDriveItemList()`

UnsetDriveItemList ensures that no value is present for DriveItemList, not even an explicit nil
### GetDriveName

`func (o *SiteDriveInfo) GetDriveName() string`

GetDriveName returns the DriveName field if non-nil, zero value otherwise.

### GetDriveNameOk

`func (o *SiteDriveInfo) GetDriveNameOk() (*string, bool)`

GetDriveNameOk returns a tuple with the DriveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveName

`func (o *SiteDriveInfo) SetDriveName(v string)`

SetDriveName sets DriveName field to given value.

### HasDriveName

`func (o *SiteDriveInfo) HasDriveName() bool`

HasDriveName returns a boolean if a field has been set.

### SetDriveNameNil

`func (o *SiteDriveInfo) SetDriveNameNil(b bool)`

 SetDriveNameNil sets the value for DriveName to be an explicit nil

### UnsetDriveName
`func (o *SiteDriveInfo) UnsetDriveName()`

UnsetDriveName ensures that no value is present for DriveName, not even an explicit nil
### GetRestoreEntireDrive

`func (o *SiteDriveInfo) GetRestoreEntireDrive() bool`

GetRestoreEntireDrive returns the RestoreEntireDrive field if non-nil, zero value otherwise.

### GetRestoreEntireDriveOk

`func (o *SiteDriveInfo) GetRestoreEntireDriveOk() (*bool, bool)`

GetRestoreEntireDriveOk returns a tuple with the RestoreEntireDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntireDrive

`func (o *SiteDriveInfo) SetRestoreEntireDrive(v bool)`

SetRestoreEntireDrive sets RestoreEntireDrive field to given value.

### HasRestoreEntireDrive

`func (o *SiteDriveInfo) HasRestoreEntireDrive() bool`

HasRestoreEntireDrive returns a boolean if a field has been set.

### SetRestoreEntireDriveNil

`func (o *SiteDriveInfo) SetRestoreEntireDriveNil(b bool)`

 SetRestoreEntireDriveNil sets the value for RestoreEntireDrive to be an explicit nil

### UnsetRestoreEntireDrive
`func (o *SiteDriveInfo) UnsetRestoreEntireDrive()`

UnsetRestoreEntireDrive ensures that no value is present for RestoreEntireDrive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


