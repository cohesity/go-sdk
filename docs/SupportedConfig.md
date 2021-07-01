# SupportedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinNodesAllowed** | Pointer to **NullableInt32** | Specifies the minimum number of Nodes supported for this Cluster type. For example, a Cohesity Cluster hosted directly on hardware must have at least 3 Nodes. | [optional] 
**SupportedErasureCoding** | Pointer to **[]string** | Array of Supported Erasure Coding Options.  List the supported Erasure Coding options for the current number of Nodes (nodeCount) in this Cluster. Each string in the array is in the following format: \&quot;NumDataStripes:NumCodedStripes\&quot; For example if there are 3 nodes in the Cluster, the following Erasure Coding mode is returned: 2:1. See the Cohesity Dashboard help documentation for details. | [optional] 

## Methods

### NewSupportedConfig

`func NewSupportedConfig() *SupportedConfig`

NewSupportedConfig instantiates a new SupportedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedConfigWithDefaults

`func NewSupportedConfigWithDefaults() *SupportedConfig`

NewSupportedConfigWithDefaults instantiates a new SupportedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinNodesAllowed

`func (o *SupportedConfig) GetMinNodesAllowed() int32`

GetMinNodesAllowed returns the MinNodesAllowed field if non-nil, zero value otherwise.

### GetMinNodesAllowedOk

`func (o *SupportedConfig) GetMinNodesAllowedOk() (*int32, bool)`

GetMinNodesAllowedOk returns a tuple with the MinNodesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodesAllowed

`func (o *SupportedConfig) SetMinNodesAllowed(v int32)`

SetMinNodesAllowed sets MinNodesAllowed field to given value.

### HasMinNodesAllowed

`func (o *SupportedConfig) HasMinNodesAllowed() bool`

HasMinNodesAllowed returns a boolean if a field has been set.

### SetMinNodesAllowedNil

`func (o *SupportedConfig) SetMinNodesAllowedNil(b bool)`

 SetMinNodesAllowedNil sets the value for MinNodesAllowed to be an explicit nil

### UnsetMinNodesAllowed
`func (o *SupportedConfig) UnsetMinNodesAllowed()`

UnsetMinNodesAllowed ensures that no value is present for MinNodesAllowed, not even an explicit nil
### GetSupportedErasureCoding

`func (o *SupportedConfig) GetSupportedErasureCoding() []string`

GetSupportedErasureCoding returns the SupportedErasureCoding field if non-nil, zero value otherwise.

### GetSupportedErasureCodingOk

`func (o *SupportedConfig) GetSupportedErasureCodingOk() (*[]string, bool)`

GetSupportedErasureCodingOk returns a tuple with the SupportedErasureCoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedErasureCoding

`func (o *SupportedConfig) SetSupportedErasureCoding(v []string)`

SetSupportedErasureCoding sets SupportedErasureCoding field to given value.

### HasSupportedErasureCoding

`func (o *SupportedConfig) HasSupportedErasureCoding() bool`

HasSupportedErasureCoding returns a boolean if a field has been set.

### SetSupportedErasureCodingNil

`func (o *SupportedConfig) SetSupportedErasureCodingNil(b bool)`

 SetSupportedErasureCodingNil sets the value for SupportedErasureCoding to be an explicit nil

### UnsetSupportedErasureCoding
`func (o *SupportedConfig) UnsetSupportedErasureCoding()`

UnsetSupportedErasureCoding ensures that no value is present for SupportedErasureCoding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


