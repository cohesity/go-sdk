# ViewAliasInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasName** | Pointer to **NullableString** | Alias name. | [optional] 
**ViewPath** | Pointer to **NullableString** | View path for the alias. | [optional] 
**SmbConfig** | Pointer to [**AliasSmbConfig**](AliasSmbConfig.md) | SMB config for the alias (share). | [optional] 
**ClientSubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | List of external client subnet IPs that are allowed to access the share. | [optional] 

## Methods

### NewViewAliasInfo

`func NewViewAliasInfo() *ViewAliasInfo`

NewViewAliasInfo instantiates a new ViewAliasInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewAliasInfoWithDefaults

`func NewViewAliasInfoWithDefaults() *ViewAliasInfo`

NewViewAliasInfoWithDefaults instantiates a new ViewAliasInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasName

`func (o *ViewAliasInfo) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *ViewAliasInfo) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *ViewAliasInfo) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *ViewAliasInfo) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### SetAliasNameNil

`func (o *ViewAliasInfo) SetAliasNameNil(b bool)`

 SetAliasNameNil sets the value for AliasName to be an explicit nil

### UnsetAliasName
`func (o *ViewAliasInfo) UnsetAliasName()`

UnsetAliasName ensures that no value is present for AliasName, not even an explicit nil
### GetViewPath

`func (o *ViewAliasInfo) GetViewPath() string`

GetViewPath returns the ViewPath field if non-nil, zero value otherwise.

### GetViewPathOk

`func (o *ViewAliasInfo) GetViewPathOk() (*string, bool)`

GetViewPathOk returns a tuple with the ViewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPath

`func (o *ViewAliasInfo) SetViewPath(v string)`

SetViewPath sets ViewPath field to given value.

### HasViewPath

`func (o *ViewAliasInfo) HasViewPath() bool`

HasViewPath returns a boolean if a field has been set.

### SetViewPathNil

`func (o *ViewAliasInfo) SetViewPathNil(b bool)`

 SetViewPathNil sets the value for ViewPath to be an explicit nil

### UnsetViewPath
`func (o *ViewAliasInfo) UnsetViewPath()`

UnsetViewPath ensures that no value is present for ViewPath, not even an explicit nil
### GetSmbConfig

`func (o *ViewAliasInfo) GetSmbConfig() AliasSmbConfig`

GetSmbConfig returns the SmbConfig field if non-nil, zero value otherwise.

### GetSmbConfigOk

`func (o *ViewAliasInfo) GetSmbConfigOk() (*AliasSmbConfig, bool)`

GetSmbConfigOk returns a tuple with the SmbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbConfig

`func (o *ViewAliasInfo) SetSmbConfig(v AliasSmbConfig)`

SetSmbConfig sets SmbConfig field to given value.

### HasSmbConfig

`func (o *ViewAliasInfo) HasSmbConfig() bool`

HasSmbConfig returns a boolean if a field has been set.

### GetClientSubnetWhitelist

`func (o *ViewAliasInfo) GetClientSubnetWhitelist() []Subnet`

GetClientSubnetWhitelist returns the ClientSubnetWhitelist field if non-nil, zero value otherwise.

### GetClientSubnetWhitelistOk

`func (o *ViewAliasInfo) GetClientSubnetWhitelistOk() (*[]Subnet, bool)`

GetClientSubnetWhitelistOk returns a tuple with the ClientSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetWhitelist

`func (o *ViewAliasInfo) SetClientSubnetWhitelist(v []Subnet)`

SetClientSubnetWhitelist sets ClientSubnetWhitelist field to given value.

### HasClientSubnetWhitelist

`func (o *ViewAliasInfo) HasClientSubnetWhitelist() bool`

HasClientSubnetWhitelist returns a boolean if a field has been set.

### SetClientSubnetWhitelistNil

`func (o *ViewAliasInfo) SetClientSubnetWhitelistNil(b bool)`

 SetClientSubnetWhitelistNil sets the value for ClientSubnetWhitelist to be an explicit nil

### UnsetClientSubnetWhitelist
`func (o *ViewAliasInfo) UnsetClientSubnetWhitelist()`

UnsetClientSubnetWhitelist ensures that no value is present for ClientSubnetWhitelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


