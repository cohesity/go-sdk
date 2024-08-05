# AadNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeAttributes** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the list of nodes&#39;s attributes as key/value pair. | [optional] 
**NodeType** | Pointer to **NullableString** | Specifies the type of aad node. | [optional] 

## Methods

### NewAadNodeInfo

`func NewAadNodeInfo() *AadNodeInfo`

NewAadNodeInfo instantiates a new AadNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadNodeInfoWithDefaults

`func NewAadNodeInfoWithDefaults() *AadNodeInfo`

NewAadNodeInfoWithDefaults instantiates a new AadNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeAttributes

`func (o *AadNodeInfo) GetNodeAttributes() []KeyValuePair`

GetNodeAttributes returns the NodeAttributes field if non-nil, zero value otherwise.

### GetNodeAttributesOk

`func (o *AadNodeInfo) GetNodeAttributesOk() (*[]KeyValuePair, bool)`

GetNodeAttributesOk returns a tuple with the NodeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAttributes

`func (o *AadNodeInfo) SetNodeAttributes(v []KeyValuePair)`

SetNodeAttributes sets NodeAttributes field to given value.

### HasNodeAttributes

`func (o *AadNodeInfo) HasNodeAttributes() bool`

HasNodeAttributes returns a boolean if a field has been set.

### SetNodeAttributesNil

`func (o *AadNodeInfo) SetNodeAttributesNil(b bool)`

 SetNodeAttributesNil sets the value for NodeAttributes to be an explicit nil

### UnsetNodeAttributes
`func (o *AadNodeInfo) UnsetNodeAttributes()`

UnsetNodeAttributes ensures that no value is present for NodeAttributes, not even an explicit nil
### GetNodeType

`func (o *AadNodeInfo) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *AadNodeInfo) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *AadNodeInfo) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *AadNodeInfo) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### SetNodeTypeNil

`func (o *AadNodeInfo) SetNodeTypeNil(b bool)`

 SetNodeTypeNil sets the value for NodeType to be an explicit nil

### UnsetNodeType
`func (o *AadNodeInfo) UnsetNodeType()`

UnsetNodeType ensures that no value is present for NodeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


