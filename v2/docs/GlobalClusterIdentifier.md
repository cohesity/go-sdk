# GlobalClusterIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | Pointer to **NullableString** | List of Clusters Identifiers to filter from. The format is clusterId:clusterIncarnationId. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id of the cluster. Only valid for DMaaS clusters. | [optional] 

## Methods

### NewGlobalClusterIdentifier

`func NewGlobalClusterIdentifier() *GlobalClusterIdentifier`

NewGlobalClusterIdentifier instantiates a new GlobalClusterIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalClusterIdentifierWithDefaults

`func NewGlobalClusterIdentifierWithDefaults() *GlobalClusterIdentifier`

NewGlobalClusterIdentifierWithDefaults instantiates a new GlobalClusterIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *GlobalClusterIdentifier) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *GlobalClusterIdentifier) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *GlobalClusterIdentifier) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *GlobalClusterIdentifier) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *GlobalClusterIdentifier) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *GlobalClusterIdentifier) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetRegionId

`func (o *GlobalClusterIdentifier) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GlobalClusterIdentifier) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GlobalClusterIdentifier) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GlobalClusterIdentifier) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *GlobalClusterIdentifier) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *GlobalClusterIdentifier) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


