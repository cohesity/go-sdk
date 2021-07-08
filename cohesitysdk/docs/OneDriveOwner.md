# OneDriveOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveInfoList** | Pointer to [**[]OneDriveInfo**](OneDriveInfo.md) | Specifies the Drives that a user owns which are to be restored. | [optional] 
**UserDetailObject** | Pointer to [**RestoreObjectDetails**](RestoreObjectDetails.md) |  | [optional] 

## Methods

### NewOneDriveOwner

`func NewOneDriveOwner() *OneDriveOwner`

NewOneDriveOwner instantiates a new OneDriveOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveOwnerWithDefaults

`func NewOneDriveOwnerWithDefaults() *OneDriveOwner`

NewOneDriveOwnerWithDefaults instantiates a new OneDriveOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveInfoList

`func (o *OneDriveOwner) GetDriveInfoList() []OneDriveInfo`

GetDriveInfoList returns the DriveInfoList field if non-nil, zero value otherwise.

### GetDriveInfoListOk

`func (o *OneDriveOwner) GetDriveInfoListOk() (*[]OneDriveInfo, bool)`

GetDriveInfoListOk returns a tuple with the DriveInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveInfoList

`func (o *OneDriveOwner) SetDriveInfoList(v []OneDriveInfo)`

SetDriveInfoList sets DriveInfoList field to given value.

### HasDriveInfoList

`func (o *OneDriveOwner) HasDriveInfoList() bool`

HasDriveInfoList returns a boolean if a field has been set.

### SetDriveInfoListNil

`func (o *OneDriveOwner) SetDriveInfoListNil(b bool)`

 SetDriveInfoListNil sets the value for DriveInfoList to be an explicit nil

### UnsetDriveInfoList
`func (o *OneDriveOwner) UnsetDriveInfoList()`

UnsetDriveInfoList ensures that no value is present for DriveInfoList, not even an explicit nil
### GetUserDetailObject

`func (o *OneDriveOwner) GetUserDetailObject() RestoreObjectDetails`

GetUserDetailObject returns the UserDetailObject field if non-nil, zero value otherwise.

### GetUserDetailObjectOk

`func (o *OneDriveOwner) GetUserDetailObjectOk() (*RestoreObjectDetails, bool)`

GetUserDetailObjectOk returns a tuple with the UserDetailObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetailObject

`func (o *OneDriveOwner) SetUserDetailObject(v RestoreObjectDetails)`

SetUserDetailObject sets UserDetailObject field to given value.

### HasUserDetailObject

`func (o *OneDriveOwner) HasUserDetailObject() bool`

HasUserDetailObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


