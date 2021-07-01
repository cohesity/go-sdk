# TenantProtectionJobUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionJobIds** | Pointer to **[]int64** | Specifies the ProtectionJobIds vec for respective tenant. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantProtectionJobUpdateParameters

`func NewTenantProtectionJobUpdateParameters() *TenantProtectionJobUpdateParameters`

NewTenantProtectionJobUpdateParameters instantiates a new TenantProtectionJobUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantProtectionJobUpdateParametersWithDefaults

`func NewTenantProtectionJobUpdateParametersWithDefaults() *TenantProtectionJobUpdateParameters`

NewTenantProtectionJobUpdateParametersWithDefaults instantiates a new TenantProtectionJobUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionJobIds

`func (o *TenantProtectionJobUpdateParameters) GetProtectionJobIds() []int64`

GetProtectionJobIds returns the ProtectionJobIds field if non-nil, zero value otherwise.

### GetProtectionJobIdsOk

`func (o *TenantProtectionJobUpdateParameters) GetProtectionJobIdsOk() (*[]int64, bool)`

GetProtectionJobIdsOk returns a tuple with the ProtectionJobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobIds

`func (o *TenantProtectionJobUpdateParameters) SetProtectionJobIds(v []int64)`

SetProtectionJobIds sets ProtectionJobIds field to given value.

### HasProtectionJobIds

`func (o *TenantProtectionJobUpdateParameters) HasProtectionJobIds() bool`

HasProtectionJobIds returns a boolean if a field has been set.

### SetProtectionJobIdsNil

`func (o *TenantProtectionJobUpdateParameters) SetProtectionJobIdsNil(b bool)`

 SetProtectionJobIdsNil sets the value for ProtectionJobIds to be an explicit nil

### UnsetProtectionJobIds
`func (o *TenantProtectionJobUpdateParameters) UnsetProtectionJobIds()`

UnsetProtectionJobIds ensures that no value is present for ProtectionJobIds, not even an explicit nil
### GetTenantId

`func (o *TenantProtectionJobUpdateParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantProtectionJobUpdateParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantProtectionJobUpdateParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantProtectionJobUpdateParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantProtectionJobUpdateParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantProtectionJobUpdateParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


