# OnPremTenantConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationsEnabled** | **NullableBool** | Wether organizations is enabled on the cluster. | 
**OrganizationsStorageDomainSharingEnabled** | **NullableBool** | Wether storage domain sharing is enabled for organizations on the cluster. | 

## Methods

### NewOnPremTenantConfig

`func NewOnPremTenantConfig(organizationsEnabled NullableBool, organizationsStorageDomainSharingEnabled NullableBool, ) *OnPremTenantConfig`

NewOnPremTenantConfig instantiates a new OnPremTenantConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremTenantConfigWithDefaults

`func NewOnPremTenantConfigWithDefaults() *OnPremTenantConfig`

NewOnPremTenantConfigWithDefaults instantiates a new OnPremTenantConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationsEnabled

`func (o *OnPremTenantConfig) GetOrganizationsEnabled() bool`

GetOrganizationsEnabled returns the OrganizationsEnabled field if non-nil, zero value otherwise.

### GetOrganizationsEnabledOk

`func (o *OnPremTenantConfig) GetOrganizationsEnabledOk() (*bool, bool)`

GetOrganizationsEnabledOk returns a tuple with the OrganizationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsEnabled

`func (o *OnPremTenantConfig) SetOrganizationsEnabled(v bool)`

SetOrganizationsEnabled sets OrganizationsEnabled field to given value.


### SetOrganizationsEnabledNil

`func (o *OnPremTenantConfig) SetOrganizationsEnabledNil(b bool)`

 SetOrganizationsEnabledNil sets the value for OrganizationsEnabled to be an explicit nil

### UnsetOrganizationsEnabled
`func (o *OnPremTenantConfig) UnsetOrganizationsEnabled()`

UnsetOrganizationsEnabled ensures that no value is present for OrganizationsEnabled, not even an explicit nil
### GetOrganizationsStorageDomainSharingEnabled

`func (o *OnPremTenantConfig) GetOrganizationsStorageDomainSharingEnabled() bool`

GetOrganizationsStorageDomainSharingEnabled returns the OrganizationsStorageDomainSharingEnabled field if non-nil, zero value otherwise.

### GetOrganizationsStorageDomainSharingEnabledOk

`func (o *OnPremTenantConfig) GetOrganizationsStorageDomainSharingEnabledOk() (*bool, bool)`

GetOrganizationsStorageDomainSharingEnabledOk returns a tuple with the OrganizationsStorageDomainSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsStorageDomainSharingEnabled

`func (o *OnPremTenantConfig) SetOrganizationsStorageDomainSharingEnabled(v bool)`

SetOrganizationsStorageDomainSharingEnabled sets OrganizationsStorageDomainSharingEnabled field to given value.


### SetOrganizationsStorageDomainSharingEnabledNil

`func (o *OnPremTenantConfig) SetOrganizationsStorageDomainSharingEnabledNil(b bool)`

 SetOrganizationsStorageDomainSharingEnabledNil sets the value for OrganizationsStorageDomainSharingEnabled to be an explicit nil

### UnsetOrganizationsStorageDomainSharingEnabled
`func (o *OnPremTenantConfig) UnsetOrganizationsStorageDomainSharingEnabled()`

UnsetOrganizationsStorageDomainSharingEnabled ensures that no value is present for OrganizationsStorageDomainSharingEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


