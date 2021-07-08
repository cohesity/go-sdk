# VaultProviderStatsByEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** | Specifies the count of the objects of the specified environment. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type. | [optional] 
**Size** | Pointer to **NullableInt64** | Specifies the size of the entities of the specified environment. | [optional] 

## Methods

### NewVaultProviderStatsByEnv

`func NewVaultProviderStatsByEnv() *VaultProviderStatsByEnv`

NewVaultProviderStatsByEnv instantiates a new VaultProviderStatsByEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultProviderStatsByEnvWithDefaults

`func NewVaultProviderStatsByEnvWithDefaults() *VaultProviderStatsByEnv`

NewVaultProviderStatsByEnvWithDefaults instantiates a new VaultProviderStatsByEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VaultProviderStatsByEnv) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VaultProviderStatsByEnv) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VaultProviderStatsByEnv) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VaultProviderStatsByEnv) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *VaultProviderStatsByEnv) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *VaultProviderStatsByEnv) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetEnvironment

`func (o *VaultProviderStatsByEnv) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *VaultProviderStatsByEnv) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *VaultProviderStatsByEnv) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *VaultProviderStatsByEnv) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *VaultProviderStatsByEnv) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *VaultProviderStatsByEnv) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetSize

`func (o *VaultProviderStatsByEnv) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VaultProviderStatsByEnv) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VaultProviderStatsByEnv) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VaultProviderStatsByEnv) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VaultProviderStatsByEnv) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VaultProviderStatsByEnv) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


