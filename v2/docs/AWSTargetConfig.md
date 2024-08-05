# AWSTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the AWS Replication target. | [optional] [readonly] 
**Region** | **NullableInt64** | Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target. | 
**RegionName** | Pointer to **NullableString** | Specifies name of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target. | [optional] [readonly] 
**SourceId** | **NullableInt64** | Specifies the source id of the AWS protection source registered on Cohesity cluster. | 

## Methods

### NewAWSTargetConfig

`func NewAWSTargetConfig(region NullableInt64, sourceId NullableInt64, ) *AWSTargetConfig`

NewAWSTargetConfig instantiates a new AWSTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSTargetConfigWithDefaults

`func NewAWSTargetConfigWithDefaults() *AWSTargetConfig`

NewAWSTargetConfigWithDefaults instantiates a new AWSTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AWSTargetConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AWSTargetConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AWSTargetConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AWSTargetConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AWSTargetConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AWSTargetConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegion

`func (o *AWSTargetConfig) GetRegion() int64`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSTargetConfig) GetRegionOk() (*int64, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSTargetConfig) SetRegion(v int64)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AWSTargetConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AWSTargetConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRegionName

`func (o *AWSTargetConfig) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AWSTargetConfig) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AWSTargetConfig) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AWSTargetConfig) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### SetRegionNameNil

`func (o *AWSTargetConfig) SetRegionNameNil(b bool)`

 SetRegionNameNil sets the value for RegionName to be an explicit nil

### UnsetRegionName
`func (o *AWSTargetConfig) UnsetRegionName()`

UnsetRegionName ensures that no value is present for RegionName, not even an explicit nil
### GetSourceId

`func (o *AWSTargetConfig) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AWSTargetConfig) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AWSTargetConfig) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *AWSTargetConfig) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AWSTargetConfig) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


