# TeamsFileItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTimeSecs** | Pointer to **NullableInt64** | Specifies the Unix timestamp epoch in seconds at which this item is created. | [optional] 
**DriveName** | Pointer to **NullableString** | Specifies the name of the drive location for this file. | [optional] 
**FileType** | Pointer to **NullableString** | Specifies the file type. | [optional] 
**ItemSize** | Pointer to **NullableInt64** | Specifies the size in bytes for the indexed item. | [optional] 

## Methods

### NewTeamsFileItem

`func NewTeamsFileItem() *TeamsFileItem`

NewTeamsFileItem instantiates a new TeamsFileItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsFileItemWithDefaults

`func NewTeamsFileItemWithDefaults() *TeamsFileItem`

NewTeamsFileItemWithDefaults instantiates a new TeamsFileItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTimeSecs

`func (o *TeamsFileItem) GetCreationTimeSecs() int64`

GetCreationTimeSecs returns the CreationTimeSecs field if non-nil, zero value otherwise.

### GetCreationTimeSecsOk

`func (o *TeamsFileItem) GetCreationTimeSecsOk() (*int64, bool)`

GetCreationTimeSecsOk returns a tuple with the CreationTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeSecs

`func (o *TeamsFileItem) SetCreationTimeSecs(v int64)`

SetCreationTimeSecs sets CreationTimeSecs field to given value.

### HasCreationTimeSecs

`func (o *TeamsFileItem) HasCreationTimeSecs() bool`

HasCreationTimeSecs returns a boolean if a field has been set.

### SetCreationTimeSecsNil

`func (o *TeamsFileItem) SetCreationTimeSecsNil(b bool)`

 SetCreationTimeSecsNil sets the value for CreationTimeSecs to be an explicit nil

### UnsetCreationTimeSecs
`func (o *TeamsFileItem) UnsetCreationTimeSecs()`

UnsetCreationTimeSecs ensures that no value is present for CreationTimeSecs, not even an explicit nil
### GetDriveName

`func (o *TeamsFileItem) GetDriveName() string`

GetDriveName returns the DriveName field if non-nil, zero value otherwise.

### GetDriveNameOk

`func (o *TeamsFileItem) GetDriveNameOk() (*string, bool)`

GetDriveNameOk returns a tuple with the DriveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveName

`func (o *TeamsFileItem) SetDriveName(v string)`

SetDriveName sets DriveName field to given value.

### HasDriveName

`func (o *TeamsFileItem) HasDriveName() bool`

HasDriveName returns a boolean if a field has been set.

### SetDriveNameNil

`func (o *TeamsFileItem) SetDriveNameNil(b bool)`

 SetDriveNameNil sets the value for DriveName to be an explicit nil

### UnsetDriveName
`func (o *TeamsFileItem) UnsetDriveName()`

UnsetDriveName ensures that no value is present for DriveName, not even an explicit nil
### GetFileType

`func (o *TeamsFileItem) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *TeamsFileItem) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *TeamsFileItem) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *TeamsFileItem) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### SetFileTypeNil

`func (o *TeamsFileItem) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *TeamsFileItem) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil
### GetItemSize

`func (o *TeamsFileItem) GetItemSize() int64`

GetItemSize returns the ItemSize field if non-nil, zero value otherwise.

### GetItemSizeOk

`func (o *TeamsFileItem) GetItemSizeOk() (*int64, bool)`

GetItemSizeOk returns a tuple with the ItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSize

`func (o *TeamsFileItem) SetItemSize(v int64)`

SetItemSize sets ItemSize field to given value.

### HasItemSize

`func (o *TeamsFileItem) HasItemSize() bool`

HasItemSize returns a boolean if a field has been set.

### SetItemSizeNil

`func (o *TeamsFileItem) SetItemSizeNil(b bool)`

 SetItemSizeNil sets the value for ItemSize to be an explicit nil

### UnsetItemSize
`func (o *TeamsFileItem) UnsetItemSize()`

UnsetItemSize ensures that no value is present for ItemSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


