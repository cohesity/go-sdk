# ClusterTenantConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | **NullableString** | List of Clusters Identifiers to filter from. The format is clusterId:clusterIncarnationId. | 
**OrganizationsEnabled** | **NullableBool** | Wether organizations is enabled on the cluster. | 
**OrganizationsStorageDomainSharingEnabled** | **NullableBool** | Wether storage domain sharing is enabled for organizations on the cluster. | 

## Methods

### NewClusterTenantConfig

`func NewClusterTenantConfig(clusterIdentifier NullableString, organizationsEnabled NullableBool, organizationsStorageDomainSharingEnabled NullableBool, ) *ClusterTenantConfig`

NewClusterTenantConfig instantiates a new ClusterTenantConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTenantConfigWithDefaults

`func NewClusterTenantConfigWithDefaults() *ClusterTenantConfig`

NewClusterTenantConfigWithDefaults instantiates a new ClusterTenantConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *ClusterTenantConfig) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *ClusterTenantConfig) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *ClusterTenantConfig) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.


### SetClusterIdentifierNil

`func (o *ClusterTenantConfig) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *ClusterTenantConfig) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetOrganizationsEnabled

`func (o *ClusterTenantConfig) GetOrganizationsEnabled() bool`

GetOrganizationsEnabled returns the OrganizationsEnabled field if non-nil, zero value otherwise.

### GetOrganizationsEnabledOk

`func (o *ClusterTenantConfig) GetOrganizationsEnabledOk() (*bool, bool)`

GetOrganizationsEnabledOk returns a tuple with the OrganizationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsEnabled

`func (o *ClusterTenantConfig) SetOrganizationsEnabled(v bool)`

SetOrganizationsEnabled sets OrganizationsEnabled field to given value.


### SetOrganizationsEnabledNil

`func (o *ClusterTenantConfig) SetOrganizationsEnabledNil(b bool)`

 SetOrganizationsEnabledNil sets the value for OrganizationsEnabled to be an explicit nil

### UnsetOrganizationsEnabled
`func (o *ClusterTenantConfig) UnsetOrganizationsEnabled()`

UnsetOrganizationsEnabled ensures that no value is present for OrganizationsEnabled, not even an explicit nil
### GetOrganizationsStorageDomainSharingEnabled

`func (o *ClusterTenantConfig) GetOrganizationsStorageDomainSharingEnabled() bool`

GetOrganizationsStorageDomainSharingEnabled returns the OrganizationsStorageDomainSharingEnabled field if non-nil, zero value otherwise.

### GetOrganizationsStorageDomainSharingEnabledOk

`func (o *ClusterTenantConfig) GetOrganizationsStorageDomainSharingEnabledOk() (*bool, bool)`

GetOrganizationsStorageDomainSharingEnabledOk returns a tuple with the OrganizationsStorageDomainSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsStorageDomainSharingEnabled

`func (o *ClusterTenantConfig) SetOrganizationsStorageDomainSharingEnabled(v bool)`

SetOrganizationsStorageDomainSharingEnabled sets OrganizationsStorageDomainSharingEnabled field to given value.


### SetOrganizationsStorageDomainSharingEnabledNil

`func (o *ClusterTenantConfig) SetOrganizationsStorageDomainSharingEnabledNil(b bool)`

 SetOrganizationsStorageDomainSharingEnabledNil sets the value for OrganizationsStorageDomainSharingEnabled to be an explicit nil

### UnsetOrganizationsStorageDomainSharingEnabled
`func (o *ClusterTenantConfig) UnsetOrganizationsStorageDomainSharingEnabled()`

UnsetOrganizationsStorageDomainSharingEnabled ensures that no value is present for OrganizationsStorageDomainSharingEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


