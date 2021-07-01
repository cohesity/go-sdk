# TenantActiveDirectoryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryDomains** | Pointer to **[]string** | Specifies the ActiveDirectoryDomain vec for respective tenant. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantActiveDirectoryUpdate

`func NewTenantActiveDirectoryUpdate() *TenantActiveDirectoryUpdate`

NewTenantActiveDirectoryUpdate instantiates a new TenantActiveDirectoryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantActiveDirectoryUpdateWithDefaults

`func NewTenantActiveDirectoryUpdateWithDefaults() *TenantActiveDirectoryUpdate`

NewTenantActiveDirectoryUpdateWithDefaults instantiates a new TenantActiveDirectoryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryDomains

`func (o *TenantActiveDirectoryUpdate) GetActiveDirectoryDomains() []string`

GetActiveDirectoryDomains returns the ActiveDirectoryDomains field if non-nil, zero value otherwise.

### GetActiveDirectoryDomainsOk

`func (o *TenantActiveDirectoryUpdate) GetActiveDirectoryDomainsOk() (*[]string, bool)`

GetActiveDirectoryDomainsOk returns a tuple with the ActiveDirectoryDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryDomains

`func (o *TenantActiveDirectoryUpdate) SetActiveDirectoryDomains(v []string)`

SetActiveDirectoryDomains sets ActiveDirectoryDomains field to given value.

### HasActiveDirectoryDomains

`func (o *TenantActiveDirectoryUpdate) HasActiveDirectoryDomains() bool`

HasActiveDirectoryDomains returns a boolean if a field has been set.

### SetActiveDirectoryDomainsNil

`func (o *TenantActiveDirectoryUpdate) SetActiveDirectoryDomainsNil(b bool)`

 SetActiveDirectoryDomainsNil sets the value for ActiveDirectoryDomains to be an explicit nil

### UnsetActiveDirectoryDomains
`func (o *TenantActiveDirectoryUpdate) UnsetActiveDirectoryDomains()`

UnsetActiveDirectoryDomains ensures that no value is present for ActiveDirectoryDomains, not even an explicit nil
### GetTenantId

`func (o *TenantActiveDirectoryUpdate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantActiveDirectoryUpdate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantActiveDirectoryUpdate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantActiveDirectoryUpdate) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantActiveDirectoryUpdate) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantActiveDirectoryUpdate) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


