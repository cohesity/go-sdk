# TenantGroupUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sids** | Pointer to **[]string** | Specifies the array of Sid of the groups. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantGroupUpdateParameters

`func NewTenantGroupUpdateParameters() *TenantGroupUpdateParameters`

NewTenantGroupUpdateParameters instantiates a new TenantGroupUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantGroupUpdateParametersWithDefaults

`func NewTenantGroupUpdateParametersWithDefaults() *TenantGroupUpdateParameters`

NewTenantGroupUpdateParametersWithDefaults instantiates a new TenantGroupUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSids

`func (o *TenantGroupUpdateParameters) GetSids() []string`

GetSids returns the Sids field if non-nil, zero value otherwise.

### GetSidsOk

`func (o *TenantGroupUpdateParameters) GetSidsOk() (*[]string, bool)`

GetSidsOk returns a tuple with the Sids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSids

`func (o *TenantGroupUpdateParameters) SetSids(v []string)`

SetSids sets Sids field to given value.

### HasSids

`func (o *TenantGroupUpdateParameters) HasSids() bool`

HasSids returns a boolean if a field has been set.

### SetSidsNil

`func (o *TenantGroupUpdateParameters) SetSidsNil(b bool)`

 SetSidsNil sets the value for Sids to be an explicit nil

### UnsetSids
`func (o *TenantGroupUpdateParameters) UnsetSids()`

UnsetSids ensures that no value is present for Sids, not even an explicit nil
### GetTenantId

`func (o *TenantGroupUpdateParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantGroupUpdateParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantGroupUpdateParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantGroupUpdateParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantGroupUpdateParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantGroupUpdateParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


