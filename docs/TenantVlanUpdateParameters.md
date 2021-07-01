# TenantVlanUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**VlanIfaceNames** | Pointer to **[]string** | Specifies the VlanIfaceNames for respective tenant, in the format of bond1.200. | [optional] 

## Methods

### NewTenantVlanUpdateParameters

`func NewTenantVlanUpdateParameters() *TenantVlanUpdateParameters`

NewTenantVlanUpdateParameters instantiates a new TenantVlanUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantVlanUpdateParametersWithDefaults

`func NewTenantVlanUpdateParametersWithDefaults() *TenantVlanUpdateParameters`

NewTenantVlanUpdateParametersWithDefaults instantiates a new TenantVlanUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantVlanUpdateParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantVlanUpdateParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantVlanUpdateParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantVlanUpdateParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantVlanUpdateParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantVlanUpdateParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetVlanIfaceNames

`func (o *TenantVlanUpdateParameters) GetVlanIfaceNames() []string`

GetVlanIfaceNames returns the VlanIfaceNames field if non-nil, zero value otherwise.

### GetVlanIfaceNamesOk

`func (o *TenantVlanUpdateParameters) GetVlanIfaceNamesOk() (*[]string, bool)`

GetVlanIfaceNamesOk returns a tuple with the VlanIfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIfaceNames

`func (o *TenantVlanUpdateParameters) SetVlanIfaceNames(v []string)`

SetVlanIfaceNames sets VlanIfaceNames field to given value.

### HasVlanIfaceNames

`func (o *TenantVlanUpdateParameters) HasVlanIfaceNames() bool`

HasVlanIfaceNames returns a boolean if a field has been set.

### SetVlanIfaceNamesNil

`func (o *TenantVlanUpdateParameters) SetVlanIfaceNamesNil(b bool)`

 SetVlanIfaceNamesNil sets the value for VlanIfaceNames to be an explicit nil

### UnsetVlanIfaceNames
`func (o *TenantVlanUpdateParameters) UnsetVlanIfaceNames()`

UnsetVlanIfaceNames ensures that no value is present for VlanIfaceNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


