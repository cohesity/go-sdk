# ListTenantData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | Pointer to [**[]HeliosTenant**](HeliosTenant.md) | List of tenants and the clusters to which they belong. | [optional] 

## Methods

### NewListTenantData

`func NewListTenantData() *ListTenantData`

NewListTenantData instantiates a new ListTenantData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenantDataWithDefaults

`func NewListTenantDataWithDefaults() *ListTenantData`

NewListTenantDataWithDefaults instantiates a new ListTenantData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *ListTenantData) GetTenants() []HeliosTenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListTenantData) GetTenantsOk() (*[]HeliosTenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListTenantData) SetTenants(v []HeliosTenant)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListTenantData) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


