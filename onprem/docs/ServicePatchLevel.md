# ServicePatchLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **NullableString** | Specifies the name of the service. | [optional] 
**PatchLevel** | Pointer to **NullableInt64** | Specifies patch level of the service after the patch operation. | [optional] 
**PatchVersion** | Pointer to **NullableString** | Specifies the version of the service patch after the patch operation. | [optional] 
**StartLevel** | Pointer to **NullableInt64** | Specifies patch level of the service before the patch operation. | [optional] 
**StartVersion** | Pointer to **NullableString** | Specifies the version of the service running on the cluster before the patch operation. | [optional] 

## Methods

### NewServicePatchLevel

`func NewServicePatchLevel() *ServicePatchLevel`

NewServicePatchLevel instantiates a new ServicePatchLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePatchLevelWithDefaults

`func NewServicePatchLevelWithDefaults() *ServicePatchLevel`

NewServicePatchLevelWithDefaults instantiates a new ServicePatchLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServicePatchLevel) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServicePatchLevel) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServicePatchLevel) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServicePatchLevel) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ServicePatchLevel) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ServicePatchLevel) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetPatchLevel

`func (o *ServicePatchLevel) GetPatchLevel() int64`

GetPatchLevel returns the PatchLevel field if non-nil, zero value otherwise.

### GetPatchLevelOk

`func (o *ServicePatchLevel) GetPatchLevelOk() (*int64, bool)`

GetPatchLevelOk returns a tuple with the PatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchLevel

`func (o *ServicePatchLevel) SetPatchLevel(v int64)`

SetPatchLevel sets PatchLevel field to given value.

### HasPatchLevel

`func (o *ServicePatchLevel) HasPatchLevel() bool`

HasPatchLevel returns a boolean if a field has been set.

### SetPatchLevelNil

`func (o *ServicePatchLevel) SetPatchLevelNil(b bool)`

 SetPatchLevelNil sets the value for PatchLevel to be an explicit nil

### UnsetPatchLevel
`func (o *ServicePatchLevel) UnsetPatchLevel()`

UnsetPatchLevel ensures that no value is present for PatchLevel, not even an explicit nil
### GetPatchVersion

`func (o *ServicePatchLevel) GetPatchVersion() string`

GetPatchVersion returns the PatchVersion field if non-nil, zero value otherwise.

### GetPatchVersionOk

`func (o *ServicePatchLevel) GetPatchVersionOk() (*string, bool)`

GetPatchVersionOk returns a tuple with the PatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchVersion

`func (o *ServicePatchLevel) SetPatchVersion(v string)`

SetPatchVersion sets PatchVersion field to given value.

### HasPatchVersion

`func (o *ServicePatchLevel) HasPatchVersion() bool`

HasPatchVersion returns a boolean if a field has been set.

### SetPatchVersionNil

`func (o *ServicePatchLevel) SetPatchVersionNil(b bool)`

 SetPatchVersionNil sets the value for PatchVersion to be an explicit nil

### UnsetPatchVersion
`func (o *ServicePatchLevel) UnsetPatchVersion()`

UnsetPatchVersion ensures that no value is present for PatchVersion, not even an explicit nil
### GetStartLevel

`func (o *ServicePatchLevel) GetStartLevel() int64`

GetStartLevel returns the StartLevel field if non-nil, zero value otherwise.

### GetStartLevelOk

`func (o *ServicePatchLevel) GetStartLevelOk() (*int64, bool)`

GetStartLevelOk returns a tuple with the StartLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLevel

`func (o *ServicePatchLevel) SetStartLevel(v int64)`

SetStartLevel sets StartLevel field to given value.

### HasStartLevel

`func (o *ServicePatchLevel) HasStartLevel() bool`

HasStartLevel returns a boolean if a field has been set.

### SetStartLevelNil

`func (o *ServicePatchLevel) SetStartLevelNil(b bool)`

 SetStartLevelNil sets the value for StartLevel to be an explicit nil

### UnsetStartLevel
`func (o *ServicePatchLevel) UnsetStartLevel()`

UnsetStartLevel ensures that no value is present for StartLevel, not even an explicit nil
### GetStartVersion

`func (o *ServicePatchLevel) GetStartVersion() string`

GetStartVersion returns the StartVersion field if non-nil, zero value otherwise.

### GetStartVersionOk

`func (o *ServicePatchLevel) GetStartVersionOk() (*string, bool)`

GetStartVersionOk returns a tuple with the StartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVersion

`func (o *ServicePatchLevel) SetStartVersion(v string)`

SetStartVersion sets StartVersion field to given value.

### HasStartVersion

`func (o *ServicePatchLevel) HasStartVersion() bool`

HasStartVersion returns a boolean if a field has been set.

### SetStartVersionNil

`func (o *ServicePatchLevel) SetStartVersionNil(b bool)`

 SetStartVersionNil sets the value for StartVersion to be an explicit nil

### UnsetStartVersion
`func (o *ServicePatchLevel) UnsetStartVersion()`

UnsetStartVersion ensures that no value is present for StartVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


