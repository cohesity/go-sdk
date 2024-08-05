# SecurityConfigSshConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshTimeoutInMins** | Pointer to **NullableInt32** | Specifies the number of minutes of remaining idle before session timeout. | [optional] 

## Methods

### NewSecurityConfigSshConfiguration

`func NewSecurityConfigSshConfiguration() *SecurityConfigSshConfiguration`

NewSecurityConfigSshConfiguration instantiates a new SecurityConfigSshConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigSshConfigurationWithDefaults

`func NewSecurityConfigSshConfigurationWithDefaults() *SecurityConfigSshConfiguration`

NewSecurityConfigSshConfigurationWithDefaults instantiates a new SecurityConfigSshConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshTimeoutInMins

`func (o *SecurityConfigSshConfiguration) GetSshTimeoutInMins() int32`

GetSshTimeoutInMins returns the SshTimeoutInMins field if non-nil, zero value otherwise.

### GetSshTimeoutInMinsOk

`func (o *SecurityConfigSshConfiguration) GetSshTimeoutInMinsOk() (*int32, bool)`

GetSshTimeoutInMinsOk returns a tuple with the SshTimeoutInMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshTimeoutInMins

`func (o *SecurityConfigSshConfiguration) SetSshTimeoutInMins(v int32)`

SetSshTimeoutInMins sets SshTimeoutInMins field to given value.

### HasSshTimeoutInMins

`func (o *SecurityConfigSshConfiguration) HasSshTimeoutInMins() bool`

HasSshTimeoutInMins returns a boolean if a field has been set.

### SetSshTimeoutInMinsNil

`func (o *SecurityConfigSshConfiguration) SetSshTimeoutInMinsNil(b bool)`

 SetSshTimeoutInMinsNil sets the value for SshTimeoutInMins to be an explicit nil

### UnsetSshTimeoutInMins
`func (o *SecurityConfigSshConfiguration) UnsetSshTimeoutInMins()`

UnsetSshTimeoutInMins ensures that no value is present for SshTimeoutInMins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


