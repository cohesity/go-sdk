# TenantConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationsEnabled** | **NullableBool** | Specifies whether organizations is enabled for the account. | 

## Methods

### NewTenantConfig

`func NewTenantConfig(organizationsEnabled NullableBool, ) *TenantConfig`

NewTenantConfig instantiates a new TenantConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantConfigWithDefaults

`func NewTenantConfigWithDefaults() *TenantConfig`

NewTenantConfigWithDefaults instantiates a new TenantConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationsEnabled

`func (o *TenantConfig) GetOrganizationsEnabled() bool`

GetOrganizationsEnabled returns the OrganizationsEnabled field if non-nil, zero value otherwise.

### GetOrganizationsEnabledOk

`func (o *TenantConfig) GetOrganizationsEnabledOk() (*bool, bool)`

GetOrganizationsEnabledOk returns a tuple with the OrganizationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsEnabled

`func (o *TenantConfig) SetOrganizationsEnabled(v bool)`

SetOrganizationsEnabled sets OrganizationsEnabled field to given value.


### SetOrganizationsEnabledNil

`func (o *TenantConfig) SetOrganizationsEnabledNil(b bool)`

 SetOrganizationsEnabledNil sets the value for OrganizationsEnabled to be an explicit nil

### UnsetOrganizationsEnabled
`func (o *TenantConfig) UnsetOrganizationsEnabled()`

UnsetOrganizationsEnabled ensures that no value is present for OrganizationsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


