# TenantProtectionPolicyUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyIds** | Pointer to **[]string** | Specifies the PolicyIds for respective tenant. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantProtectionPolicyUpdateParameters

`func NewTenantProtectionPolicyUpdateParameters() *TenantProtectionPolicyUpdateParameters`

NewTenantProtectionPolicyUpdateParameters instantiates a new TenantProtectionPolicyUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantProtectionPolicyUpdateParametersWithDefaults

`func NewTenantProtectionPolicyUpdateParametersWithDefaults() *TenantProtectionPolicyUpdateParameters`

NewTenantProtectionPolicyUpdateParametersWithDefaults instantiates a new TenantProtectionPolicyUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyIds

`func (o *TenantProtectionPolicyUpdateParameters) GetPolicyIds() []string`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *TenantProtectionPolicyUpdateParameters) GetPolicyIdsOk() (*[]string, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *TenantProtectionPolicyUpdateParameters) SetPolicyIds(v []string)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *TenantProtectionPolicyUpdateParameters) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### SetPolicyIdsNil

`func (o *TenantProtectionPolicyUpdateParameters) SetPolicyIdsNil(b bool)`

 SetPolicyIdsNil sets the value for PolicyIds to be an explicit nil

### UnsetPolicyIds
`func (o *TenantProtectionPolicyUpdateParameters) UnsetPolicyIds()`

UnsetPolicyIds ensures that no value is present for PolicyIds, not even an explicit nil
### GetTenantId

`func (o *TenantProtectionPolicyUpdateParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantProtectionPolicyUpdateParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantProtectionPolicyUpdateParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantProtectionPolicyUpdateParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantProtectionPolicyUpdateParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantProtectionPolicyUpdateParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


