# HeliosTenantClusterResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | Pointer to **NullableString** | Specifies the list of cluster identifiers. The format is clusterId:clusterIncarnationId. | [optional] 
**StorageDomainNames** | Pointer to **[]string** | List of assigned Storage Domains. | [optional] 
**SourceNames** | Pointer to **[]string** | List of completely assigned Sources. | [optional] 
**SourceStats** | Pointer to [**[]HeliosSourceObjectsStats**](HeliosSourceObjectsStats.md) | Number of assigned objects grouped by source names. | [optional] 

## Methods

### NewHeliosTenantClusterResources

`func NewHeliosTenantClusterResources() *HeliosTenantClusterResources`

NewHeliosTenantClusterResources instantiates a new HeliosTenantClusterResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTenantClusterResourcesWithDefaults

`func NewHeliosTenantClusterResourcesWithDefaults() *HeliosTenantClusterResources`

NewHeliosTenantClusterResourcesWithDefaults instantiates a new HeliosTenantClusterResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *HeliosTenantClusterResources) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosTenantClusterResources) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosTenantClusterResources) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosTenantClusterResources) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosTenantClusterResources) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosTenantClusterResources) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetStorageDomainNames

`func (o *HeliosTenantClusterResources) GetStorageDomainNames() []string`

GetStorageDomainNames returns the StorageDomainNames field if non-nil, zero value otherwise.

### GetStorageDomainNamesOk

`func (o *HeliosTenantClusterResources) GetStorageDomainNamesOk() (*[]string, bool)`

GetStorageDomainNamesOk returns a tuple with the StorageDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainNames

`func (o *HeliosTenantClusterResources) SetStorageDomainNames(v []string)`

SetStorageDomainNames sets StorageDomainNames field to given value.

### HasStorageDomainNames

`func (o *HeliosTenantClusterResources) HasStorageDomainNames() bool`

HasStorageDomainNames returns a boolean if a field has been set.

### SetStorageDomainNamesNil

`func (o *HeliosTenantClusterResources) SetStorageDomainNamesNil(b bool)`

 SetStorageDomainNamesNil sets the value for StorageDomainNames to be an explicit nil

### UnsetStorageDomainNames
`func (o *HeliosTenantClusterResources) UnsetStorageDomainNames()`

UnsetStorageDomainNames ensures that no value is present for StorageDomainNames, not even an explicit nil
### GetSourceNames

`func (o *HeliosTenantClusterResources) GetSourceNames() []string`

GetSourceNames returns the SourceNames field if non-nil, zero value otherwise.

### GetSourceNamesOk

`func (o *HeliosTenantClusterResources) GetSourceNamesOk() (*[]string, bool)`

GetSourceNamesOk returns a tuple with the SourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNames

`func (o *HeliosTenantClusterResources) SetSourceNames(v []string)`

SetSourceNames sets SourceNames field to given value.

### HasSourceNames

`func (o *HeliosTenantClusterResources) HasSourceNames() bool`

HasSourceNames returns a boolean if a field has been set.

### SetSourceNamesNil

`func (o *HeliosTenantClusterResources) SetSourceNamesNil(b bool)`

 SetSourceNamesNil sets the value for SourceNames to be an explicit nil

### UnsetSourceNames
`func (o *HeliosTenantClusterResources) UnsetSourceNames()`

UnsetSourceNames ensures that no value is present for SourceNames, not even an explicit nil
### GetSourceStats

`func (o *HeliosTenantClusterResources) GetSourceStats() []HeliosSourceObjectsStats`

GetSourceStats returns the SourceStats field if non-nil, zero value otherwise.

### GetSourceStatsOk

`func (o *HeliosTenantClusterResources) GetSourceStatsOk() (*[]HeliosSourceObjectsStats, bool)`

GetSourceStatsOk returns a tuple with the SourceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStats

`func (o *HeliosTenantClusterResources) SetSourceStats(v []HeliosSourceObjectsStats)`

SetSourceStats sets SourceStats field to given value.

### HasSourceStats

`func (o *HeliosTenantClusterResources) HasSourceStats() bool`

HasSourceStats returns a boolean if a field has been set.

### SetSourceStatsNil

`func (o *HeliosTenantClusterResources) SetSourceStatsNil(b bool)`

 SetSourceStatsNil sets the value for SourceStats to be an explicit nil

### UnsetSourceStats
`func (o *HeliosTenantClusterResources) UnsetSourceStats()`

UnsetSourceStats ensures that no value is present for SourceStats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


