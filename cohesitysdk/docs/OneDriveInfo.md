# OneDriveInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveId** | Pointer to **NullableString** | Specifies the Id of the Drive. | [optional] 
**DriveItemList** | Pointer to [**[]OneDriveItem**](OneDriveItem.md) | Specifies the Drive items such as files/folders. | [optional] 
**RestoreEntireDrive** | Pointer to **NullableBool** | Specifies whether entire drive is to be restored. This should be set to false if specific drive items are to be restored within &#39;DriveItemList&#39;. | [optional] 

## Methods

### NewOneDriveInfo

`func NewOneDriveInfo() *OneDriveInfo`

NewOneDriveInfo instantiates a new OneDriveInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveInfoWithDefaults

`func NewOneDriveInfoWithDefaults() *OneDriveInfo`

NewOneDriveInfoWithDefaults instantiates a new OneDriveInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveId

`func (o *OneDriveInfo) GetDriveId() string`

GetDriveId returns the DriveId field if non-nil, zero value otherwise.

### GetDriveIdOk

`func (o *OneDriveInfo) GetDriveIdOk() (*string, bool)`

GetDriveIdOk returns a tuple with the DriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveId

`func (o *OneDriveInfo) SetDriveId(v string)`

SetDriveId sets DriveId field to given value.

### HasDriveId

`func (o *OneDriveInfo) HasDriveId() bool`

HasDriveId returns a boolean if a field has been set.

### SetDriveIdNil

`func (o *OneDriveInfo) SetDriveIdNil(b bool)`

 SetDriveIdNil sets the value for DriveId to be an explicit nil

### UnsetDriveId
`func (o *OneDriveInfo) UnsetDriveId()`

UnsetDriveId ensures that no value is present for DriveId, not even an explicit nil
### GetDriveItemList

`func (o *OneDriveInfo) GetDriveItemList() []OneDriveItem`

GetDriveItemList returns the DriveItemList field if non-nil, zero value otherwise.

### GetDriveItemListOk

`func (o *OneDriveInfo) GetDriveItemListOk() (*[]OneDriveItem, bool)`

GetDriveItemListOk returns a tuple with the DriveItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveItemList

`func (o *OneDriveInfo) SetDriveItemList(v []OneDriveItem)`

SetDriveItemList sets DriveItemList field to given value.

### HasDriveItemList

`func (o *OneDriveInfo) HasDriveItemList() bool`

HasDriveItemList returns a boolean if a field has been set.

### SetDriveItemListNil

`func (o *OneDriveInfo) SetDriveItemListNil(b bool)`

 SetDriveItemListNil sets the value for DriveItemList to be an explicit nil

### UnsetDriveItemList
`func (o *OneDriveInfo) UnsetDriveItemList()`

UnsetDriveItemList ensures that no value is present for DriveItemList, not even an explicit nil
### GetRestoreEntireDrive

`func (o *OneDriveInfo) GetRestoreEntireDrive() bool`

GetRestoreEntireDrive returns the RestoreEntireDrive field if non-nil, zero value otherwise.

### GetRestoreEntireDriveOk

`func (o *OneDriveInfo) GetRestoreEntireDriveOk() (*bool, bool)`

GetRestoreEntireDriveOk returns a tuple with the RestoreEntireDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntireDrive

`func (o *OneDriveInfo) SetRestoreEntireDrive(v bool)`

SetRestoreEntireDrive sets RestoreEntireDrive field to given value.

### HasRestoreEntireDrive

`func (o *OneDriveInfo) HasRestoreEntireDrive() bool`

HasRestoreEntireDrive returns a boolean if a field has been set.

### SetRestoreEntireDriveNil

`func (o *OneDriveInfo) SetRestoreEntireDriveNil(b bool)`

 SetRestoreEntireDriveNil sets the value for RestoreEntireDrive to be an explicit nil

### UnsetRestoreEntireDrive
`func (o *OneDriveInfo) UnsetRestoreEntireDrive()`

UnsetRestoreEntireDrive ensures that no value is present for RestoreEntireDrive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


