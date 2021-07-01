# TenantActiveDirectoryUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryDomains** | Pointer to **[]string** | Specifies the ActiveDirectoryDomain vec for respective tenant. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantActiveDirectoryUpdateParameters

`func NewTenantActiveDirectoryUpdateParameters() *TenantActiveDirectoryUpdateParameters`

NewTenantActiveDirectoryUpdateParameters instantiates a new TenantActiveDirectoryUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantActiveDirectoryUpdateParametersWithDefaults

`func NewTenantActiveDirectoryUpdateParametersWithDefaults() *TenantActiveDirectoryUpdateParameters`

NewTenantActiveDirectoryUpdateParametersWithDefaults instantiates a new TenantActiveDirectoryUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryDomains

`func (o *TenantActiveDirectoryUpdateParameters) GetActiveDirectoryDomains() []string`

GetActiveDirectoryDomains returns the ActiveDirectoryDomains field if non-nil, zero value otherwise.

### GetActiveDirectoryDomainsOk

`func (o *TenantActiveDirectoryUpdateParameters) GetActiveDirectoryDomainsOk() (*[]string, bool)`

GetActiveDirectoryDomainsOk returns a tuple with the ActiveDirectoryDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryDomains

`func (o *TenantActiveDirectoryUpdateParameters) SetActiveDirectoryDomains(v []string)`

SetActiveDirectoryDomains sets ActiveDirectoryDomains field to given value.

### HasActiveDirectoryDomains

`func (o *TenantActiveDirectoryUpdateParameters) HasActiveDirectoryDomains() bool`

HasActiveDirectoryDomains returns a boolean if a field has been set.

### SetActiveDirectoryDomainsNil

`func (o *TenantActiveDirectoryUpdateParameters) SetActiveDirectoryDomainsNil(b bool)`

 SetActiveDirectoryDomainsNil sets the value for ActiveDirectoryDomains to be an explicit nil

### UnsetActiveDirectoryDomains
`func (o *TenantActiveDirectoryUpdateParameters) UnsetActiveDirectoryDomains()`

UnsetActiveDirectoryDomains ensures that no value is present for ActiveDirectoryDomains, not even an explicit nil
### GetTenantId

`func (o *TenantActiveDirectoryUpdateParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantActiveDirectoryUpdateParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantActiveDirectoryUpdateParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantActiveDirectoryUpdateParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantActiveDirectoryUpdateParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantActiveDirectoryUpdateParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


