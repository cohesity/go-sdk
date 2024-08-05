# ClusterSLAMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSLA** | Pointer to **NullableInt32** | Specifies the default SLA in minutes at cluster level which will be applied as default across all protection groups. This value is only used by UI to populate the default value. | [optional] 
**MinimumSLA** | Pointer to **NullableInt32** | Specifies the minimum SLA in minutes at cluster level which will be validated against all protection groups SLA configuration. If the provided SLA at protection group creation or update is less than this value then protection group request will be invalidated. | [optional] 

## Methods

### NewClusterSLAMetadata

`func NewClusterSLAMetadata() *ClusterSLAMetadata`

NewClusterSLAMetadata instantiates a new ClusterSLAMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSLAMetadataWithDefaults

`func NewClusterSLAMetadataWithDefaults() *ClusterSLAMetadata`

NewClusterSLAMetadataWithDefaults instantiates a new ClusterSLAMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSLA

`func (o *ClusterSLAMetadata) GetDefaultSLA() int32`

GetDefaultSLA returns the DefaultSLA field if non-nil, zero value otherwise.

### GetDefaultSLAOk

`func (o *ClusterSLAMetadata) GetDefaultSLAOk() (*int32, bool)`

GetDefaultSLAOk returns a tuple with the DefaultSLA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSLA

`func (o *ClusterSLAMetadata) SetDefaultSLA(v int32)`

SetDefaultSLA sets DefaultSLA field to given value.

### HasDefaultSLA

`func (o *ClusterSLAMetadata) HasDefaultSLA() bool`

HasDefaultSLA returns a boolean if a field has been set.

### SetDefaultSLANil

`func (o *ClusterSLAMetadata) SetDefaultSLANil(b bool)`

 SetDefaultSLANil sets the value for DefaultSLA to be an explicit nil

### UnsetDefaultSLA
`func (o *ClusterSLAMetadata) UnsetDefaultSLA()`

UnsetDefaultSLA ensures that no value is present for DefaultSLA, not even an explicit nil
### GetMinimumSLA

`func (o *ClusterSLAMetadata) GetMinimumSLA() int32`

GetMinimumSLA returns the MinimumSLA field if non-nil, zero value otherwise.

### GetMinimumSLAOk

`func (o *ClusterSLAMetadata) GetMinimumSLAOk() (*int32, bool)`

GetMinimumSLAOk returns a tuple with the MinimumSLA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSLA

`func (o *ClusterSLAMetadata) SetMinimumSLA(v int32)`

SetMinimumSLA sets MinimumSLA field to given value.

### HasMinimumSLA

`func (o *ClusterSLAMetadata) HasMinimumSLA() bool`

HasMinimumSLA returns a boolean if a field has been set.

### SetMinimumSLANil

`func (o *ClusterSLAMetadata) SetMinimumSLANil(b bool)`

 SetMinimumSLANil sets the value for MinimumSLA to be an explicit nil

### UnsetMinimumSLA
`func (o *ClusterSLAMetadata) UnsetMinimumSLA()`

UnsetMinimumSLA ensures that no value is present for MinimumSLA, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


