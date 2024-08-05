# VmwareObjectEnableSqlProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**VmwareSQLCredentialParams**](VmwareSQLCredentialParams.md) |  | [optional] 
**UseInstalledAgent** | Pointer to **NullableBool** | Specifies if agent is already installed. | [optional] 

## Methods

### NewVmwareObjectEnableSqlProtectionParams

`func NewVmwareObjectEnableSqlProtectionParams() *VmwareObjectEnableSqlProtectionParams`

NewVmwareObjectEnableSqlProtectionParams instantiates a new VmwareObjectEnableSqlProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectEnableSqlProtectionParamsWithDefaults

`func NewVmwareObjectEnableSqlProtectionParamsWithDefaults() *VmwareObjectEnableSqlProtectionParams`

NewVmwareObjectEnableSqlProtectionParamsWithDefaults instantiates a new VmwareObjectEnableSqlProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *VmwareObjectEnableSqlProtectionParams) GetCredentials() VmwareSQLCredentialParams`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *VmwareObjectEnableSqlProtectionParams) GetCredentialsOk() (*VmwareSQLCredentialParams, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *VmwareObjectEnableSqlProtectionParams) SetCredentials(v VmwareSQLCredentialParams)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *VmwareObjectEnableSqlProtectionParams) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetUseInstalledAgent

`func (o *VmwareObjectEnableSqlProtectionParams) GetUseInstalledAgent() bool`

GetUseInstalledAgent returns the UseInstalledAgent field if non-nil, zero value otherwise.

### GetUseInstalledAgentOk

`func (o *VmwareObjectEnableSqlProtectionParams) GetUseInstalledAgentOk() (*bool, bool)`

GetUseInstalledAgentOk returns a tuple with the UseInstalledAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInstalledAgent

`func (o *VmwareObjectEnableSqlProtectionParams) SetUseInstalledAgent(v bool)`

SetUseInstalledAgent sets UseInstalledAgent field to given value.

### HasUseInstalledAgent

`func (o *VmwareObjectEnableSqlProtectionParams) HasUseInstalledAgent() bool`

HasUseInstalledAgent returns a boolean if a field has been set.

### SetUseInstalledAgentNil

`func (o *VmwareObjectEnableSqlProtectionParams) SetUseInstalledAgentNil(b bool)`

 SetUseInstalledAgentNil sets the value for UseInstalledAgent to be an explicit nil

### UnsetUseInstalledAgent
`func (o *VmwareObjectEnableSqlProtectionParams) UnsetUseInstalledAgent()`

UnsetUseInstalledAgent ensures that no value is present for UseInstalledAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


