# ViewOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPoint** | Pointer to **NullableString** | The path of the SMB share. | [optional] 
**ViewName** | Pointer to **NullableString** | The name of the view. | [optional] 
**WhitelistRestoreViewForAll** | Pointer to **NullableBool** | Whether to white-list the Exchange restore view for all the IP addresses | [optional] 

## Methods

### NewViewOptions

`func NewViewOptions() *ViewOptions`

NewViewOptions instantiates a new ViewOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewOptionsWithDefaults

`func NewViewOptionsWithDefaults() *ViewOptions`

NewViewOptionsWithDefaults instantiates a new ViewOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPoint

`func (o *ViewOptions) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ViewOptions) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ViewOptions) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *ViewOptions) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### SetMountPointNil

`func (o *ViewOptions) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *ViewOptions) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetViewName

`func (o *ViewOptions) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ViewOptions) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ViewOptions) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ViewOptions) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ViewOptions) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ViewOptions) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetWhitelistRestoreViewForAll

`func (o *ViewOptions) GetWhitelistRestoreViewForAll() bool`

GetWhitelistRestoreViewForAll returns the WhitelistRestoreViewForAll field if non-nil, zero value otherwise.

### GetWhitelistRestoreViewForAllOk

`func (o *ViewOptions) GetWhitelistRestoreViewForAllOk() (*bool, bool)`

GetWhitelistRestoreViewForAllOk returns a tuple with the WhitelistRestoreViewForAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistRestoreViewForAll

`func (o *ViewOptions) SetWhitelistRestoreViewForAll(v bool)`

SetWhitelistRestoreViewForAll sets WhitelistRestoreViewForAll field to given value.

### HasWhitelistRestoreViewForAll

`func (o *ViewOptions) HasWhitelistRestoreViewForAll() bool`

HasWhitelistRestoreViewForAll returns a boolean if a field has been set.

### SetWhitelistRestoreViewForAllNil

`func (o *ViewOptions) SetWhitelistRestoreViewForAllNil(b bool)`

 SetWhitelistRestoreViewForAllNil sets the value for WhitelistRestoreViewForAll to be an explicit nil

### UnsetWhitelistRestoreViewForAll
`func (o *ViewOptions) UnsetWhitelistRestoreViewForAll()`

UnsetWhitelistRestoreViewForAll ensures that no value is present for WhitelistRestoreViewForAll, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


