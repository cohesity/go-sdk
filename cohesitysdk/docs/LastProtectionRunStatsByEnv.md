# LastProtectionRunStatsByEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment. | [optional] 
**NumObjectsFailed** | Pointer to **NullableInt64** | Specifies the count of objects that failed last Protection Run. | [optional] 
**NumObjectsFailedSla** | Pointer to **NullableInt64** | Specifies the count of objects that failed sla in the last Run. | [optional] 
**NumObjectsMetSla** | Pointer to **NullableInt64** | Specifies the count of objects that met sla in the last Run. | [optional] 

## Methods

### NewLastProtectionRunStatsByEnv

`func NewLastProtectionRunStatsByEnv() *LastProtectionRunStatsByEnv`

NewLastProtectionRunStatsByEnv instantiates a new LastProtectionRunStatsByEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastProtectionRunStatsByEnvWithDefaults

`func NewLastProtectionRunStatsByEnvWithDefaults() *LastProtectionRunStatsByEnv`

NewLastProtectionRunStatsByEnvWithDefaults instantiates a new LastProtectionRunStatsByEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *LastProtectionRunStatsByEnv) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *LastProtectionRunStatsByEnv) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *LastProtectionRunStatsByEnv) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *LastProtectionRunStatsByEnv) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *LastProtectionRunStatsByEnv) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *LastProtectionRunStatsByEnv) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetNumObjectsFailed

`func (o *LastProtectionRunStatsByEnv) GetNumObjectsFailed() int64`

GetNumObjectsFailed returns the NumObjectsFailed field if non-nil, zero value otherwise.

### GetNumObjectsFailedOk

`func (o *LastProtectionRunStatsByEnv) GetNumObjectsFailedOk() (*int64, bool)`

GetNumObjectsFailedOk returns a tuple with the NumObjectsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsFailed

`func (o *LastProtectionRunStatsByEnv) SetNumObjectsFailed(v int64)`

SetNumObjectsFailed sets NumObjectsFailed field to given value.

### HasNumObjectsFailed

`func (o *LastProtectionRunStatsByEnv) HasNumObjectsFailed() bool`

HasNumObjectsFailed returns a boolean if a field has been set.

### SetNumObjectsFailedNil

`func (o *LastProtectionRunStatsByEnv) SetNumObjectsFailedNil(b bool)`

 SetNumObjectsFailedNil sets the value for NumObjectsFailed to be an explicit nil

### UnsetNumObjectsFailed
`func (o *LastProtectionRunStatsByEnv) UnsetNumObjectsFailed()`

UnsetNumObjectsFailed ensures that no value is present for NumObjectsFailed, not even an explicit nil
### GetNumObjectsFailedSla

`func (o *LastProtectionRunStatsByEnv) GetNumObjectsFailedSla() int64`

GetNumObjectsFailedSla returns the NumObjectsFailedSla field if non-nil, zero value otherwise.

### GetNumObjectsFailedSlaOk

`func (o *LastProtectionRunStatsByEnv) GetNumObjectsFailedSlaOk() (*int64, bool)`

GetNumObjectsFailedSlaOk returns a tuple with the NumObjectsFailedSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsFailedSla

`func (o *LastProtectionRunStatsByEnv) SetNumObjectsFailedSla(v int64)`

SetNumObjectsFailedSla sets NumObjectsFailedSla field to given value.

### HasNumObjectsFailedSla

`func (o *LastProtectionRunStatsByEnv) HasNumObjectsFailedSla() bool`

HasNumObjectsFailedSla returns a boolean if a field has been set.

### SetNumObjectsFailedSlaNil

`func (o *LastProtectionRunStatsByEnv) SetNumObjectsFailedSlaNil(b bool)`

 SetNumObjectsFailedSlaNil sets the value for NumObjectsFailedSla to be an explicit nil

### UnsetNumObjectsFailedSla
`func (o *LastProtectionRunStatsByEnv) UnsetNumObjectsFailedSla()`

UnsetNumObjectsFailedSla ensures that no value is present for NumObjectsFailedSla, not even an explicit nil
### GetNumObjectsMetSla

`func (o *LastProtectionRunStatsByEnv) GetNumObjectsMetSla() int64`

GetNumObjectsMetSla returns the NumObjectsMetSla field if non-nil, zero value otherwise.

### GetNumObjectsMetSlaOk

`func (o *LastProtectionRunStatsByEnv) GetNumObjectsMetSlaOk() (*int64, bool)`

GetNumObjectsMetSlaOk returns a tuple with the NumObjectsMetSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsMetSla

`func (o *LastProtectionRunStatsByEnv) SetNumObjectsMetSla(v int64)`

SetNumObjectsMetSla sets NumObjectsMetSla field to given value.

### HasNumObjectsMetSla

`func (o *LastProtectionRunStatsByEnv) HasNumObjectsMetSla() bool`

HasNumObjectsMetSla returns a boolean if a field has been set.

### SetNumObjectsMetSlaNil

`func (o *LastProtectionRunStatsByEnv) SetNumObjectsMetSlaNil(b bool)`

 SetNumObjectsMetSlaNil sets the value for NumObjectsMetSla to be an explicit nil

### UnsetNumObjectsMetSla
`func (o *LastProtectionRunStatsByEnv) UnsetNumObjectsMetSla()`

UnsetNumObjectsMetSla ensures that no value is present for NumObjectsMetSla, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


