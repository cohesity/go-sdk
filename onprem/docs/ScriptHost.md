# ScriptHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **NullableString** | Specifies the Hostname or IP address of the host where the pre and post script will be run. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username for the host. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the Operating system type of the host. | [optional] 

## Methods

### NewScriptHost

`func NewScriptHost() *ScriptHost`

NewScriptHost instantiates a new ScriptHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptHostWithDefaults

`func NewScriptHostWithDefaults() *ScriptHost`

NewScriptHostWithDefaults instantiates a new ScriptHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *ScriptHost) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ScriptHost) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ScriptHost) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ScriptHost) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *ScriptHost) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *ScriptHost) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetUsername

`func (o *ScriptHost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ScriptHost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ScriptHost) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ScriptHost) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ScriptHost) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ScriptHost) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetHostType

`func (o *ScriptHost) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *ScriptHost) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *ScriptHost) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *ScriptHost) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *ScriptHost) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *ScriptHost) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


