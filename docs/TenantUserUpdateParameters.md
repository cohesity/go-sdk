# TenantUserUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sids** | Pointer to **[]string** | Specifies the array of Sid of the users. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantUserUpdateParameters

`func NewTenantUserUpdateParameters() *TenantUserUpdateParameters`

NewTenantUserUpdateParameters instantiates a new TenantUserUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUserUpdateParametersWithDefaults

`func NewTenantUserUpdateParametersWithDefaults() *TenantUserUpdateParameters`

NewTenantUserUpdateParametersWithDefaults instantiates a new TenantUserUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSids

`func (o *TenantUserUpdateParameters) GetSids() []string`

GetSids returns the Sids field if non-nil, zero value otherwise.

### GetSidsOk

`func (o *TenantUserUpdateParameters) GetSidsOk() (*[]string, bool)`

GetSidsOk returns a tuple with the Sids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSids

`func (o *TenantUserUpdateParameters) SetSids(v []string)`

SetSids sets Sids field to given value.

### HasSids

`func (o *TenantUserUpdateParameters) HasSids() bool`

HasSids returns a boolean if a field has been set.

### SetSidsNil

`func (o *TenantUserUpdateParameters) SetSidsNil(b bool)`

 SetSidsNil sets the value for Sids to be an explicit nil

### UnsetSids
`func (o *TenantUserUpdateParameters) UnsetSids()`

UnsetSids ensures that no value is present for Sids, not even an explicit nil
### GetTenantId

`func (o *TenantUserUpdateParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantUserUpdateParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantUserUpdateParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantUserUpdateParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantUserUpdateParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantUserUpdateParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


