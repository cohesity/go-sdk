# RestoreExchangeParamsViewOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPoint** | Pointer to **NullableString** | The path to access the SMB share. | [optional] 
**ViewName** | Pointer to **NullableString** | View to which the files of an Exchange database have to be cloned. | [optional] 
**WhitelistRestoreViewForAll** | Pointer to **NullableBool** | If set to true then when restore view is cloned then white-list all IPs not just the agent IP. | [optional] 

## Methods

### NewRestoreExchangeParamsViewOptions

`func NewRestoreExchangeParamsViewOptions() *RestoreExchangeParamsViewOptions`

NewRestoreExchangeParamsViewOptions instantiates a new RestoreExchangeParamsViewOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreExchangeParamsViewOptionsWithDefaults

`func NewRestoreExchangeParamsViewOptionsWithDefaults() *RestoreExchangeParamsViewOptions`

NewRestoreExchangeParamsViewOptionsWithDefaults instantiates a new RestoreExchangeParamsViewOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPoint

`func (o *RestoreExchangeParamsViewOptions) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *RestoreExchangeParamsViewOptions) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *RestoreExchangeParamsViewOptions) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *RestoreExchangeParamsViewOptions) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### SetMountPointNil

`func (o *RestoreExchangeParamsViewOptions) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *RestoreExchangeParamsViewOptions) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetViewName

`func (o *RestoreExchangeParamsViewOptions) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RestoreExchangeParamsViewOptions) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RestoreExchangeParamsViewOptions) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RestoreExchangeParamsViewOptions) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RestoreExchangeParamsViewOptions) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RestoreExchangeParamsViewOptions) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetWhitelistRestoreViewForAll

`func (o *RestoreExchangeParamsViewOptions) GetWhitelistRestoreViewForAll() bool`

GetWhitelistRestoreViewForAll returns the WhitelistRestoreViewForAll field if non-nil, zero value otherwise.

### GetWhitelistRestoreViewForAllOk

`func (o *RestoreExchangeParamsViewOptions) GetWhitelistRestoreViewForAllOk() (*bool, bool)`

GetWhitelistRestoreViewForAllOk returns a tuple with the WhitelistRestoreViewForAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistRestoreViewForAll

`func (o *RestoreExchangeParamsViewOptions) SetWhitelistRestoreViewForAll(v bool)`

SetWhitelistRestoreViewForAll sets WhitelistRestoreViewForAll field to given value.

### HasWhitelistRestoreViewForAll

`func (o *RestoreExchangeParamsViewOptions) HasWhitelistRestoreViewForAll() bool`

HasWhitelistRestoreViewForAll returns a boolean if a field has been set.

### SetWhitelistRestoreViewForAllNil

`func (o *RestoreExchangeParamsViewOptions) SetWhitelistRestoreViewForAllNil(b bool)`

 SetWhitelistRestoreViewForAllNil sets the value for WhitelistRestoreViewForAll to be an explicit nil

### UnsetWhitelistRestoreViewForAll
`func (o *RestoreExchangeParamsViewOptions) UnsetWhitelistRestoreViewForAll()`

UnsetWhitelistRestoreViewForAll ensures that no value is present for WhitelistRestoreViewForAll, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


