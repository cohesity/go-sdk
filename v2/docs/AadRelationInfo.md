# AadRelationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestNodeId** | Pointer to **string** | Specifies Unique ID of the destination node. | [optional] [readonly] 
**RelationAttributes** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the list of node relation attributes provided in key/value pair. | [optional] 
**RelationType** | Pointer to **NullableString** | Specified type of the aad node relation. | [optional] 
**SrcNodeId** | Pointer to **string** | Specifies Unique ID of the source node. | [optional] [readonly] 

## Methods

### NewAadRelationInfo

`func NewAadRelationInfo() *AadRelationInfo`

NewAadRelationInfo instantiates a new AadRelationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadRelationInfoWithDefaults

`func NewAadRelationInfoWithDefaults() *AadRelationInfo`

NewAadRelationInfoWithDefaults instantiates a new AadRelationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestNodeId

`func (o *AadRelationInfo) GetDestNodeId() string`

GetDestNodeId returns the DestNodeId field if non-nil, zero value otherwise.

### GetDestNodeIdOk

`func (o *AadRelationInfo) GetDestNodeIdOk() (*string, bool)`

GetDestNodeIdOk returns a tuple with the DestNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestNodeId

`func (o *AadRelationInfo) SetDestNodeId(v string)`

SetDestNodeId sets DestNodeId field to given value.

### HasDestNodeId

`func (o *AadRelationInfo) HasDestNodeId() bool`

HasDestNodeId returns a boolean if a field has been set.

### GetRelationAttributes

`func (o *AadRelationInfo) GetRelationAttributes() []KeyValuePair`

GetRelationAttributes returns the RelationAttributes field if non-nil, zero value otherwise.

### GetRelationAttributesOk

`func (o *AadRelationInfo) GetRelationAttributesOk() (*[]KeyValuePair, bool)`

GetRelationAttributesOk returns a tuple with the RelationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationAttributes

`func (o *AadRelationInfo) SetRelationAttributes(v []KeyValuePair)`

SetRelationAttributes sets RelationAttributes field to given value.

### HasRelationAttributes

`func (o *AadRelationInfo) HasRelationAttributes() bool`

HasRelationAttributes returns a boolean if a field has been set.

### SetRelationAttributesNil

`func (o *AadRelationInfo) SetRelationAttributesNil(b bool)`

 SetRelationAttributesNil sets the value for RelationAttributes to be an explicit nil

### UnsetRelationAttributes
`func (o *AadRelationInfo) UnsetRelationAttributes()`

UnsetRelationAttributes ensures that no value is present for RelationAttributes, not even an explicit nil
### GetRelationType

`func (o *AadRelationInfo) GetRelationType() string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *AadRelationInfo) GetRelationTypeOk() (*string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *AadRelationInfo) SetRelationType(v string)`

SetRelationType sets RelationType field to given value.

### HasRelationType

`func (o *AadRelationInfo) HasRelationType() bool`

HasRelationType returns a boolean if a field has been set.

### SetRelationTypeNil

`func (o *AadRelationInfo) SetRelationTypeNil(b bool)`

 SetRelationTypeNil sets the value for RelationType to be an explicit nil

### UnsetRelationType
`func (o *AadRelationInfo) UnsetRelationType()`

UnsetRelationType ensures that no value is present for RelationType, not even an explicit nil
### GetSrcNodeId

`func (o *AadRelationInfo) GetSrcNodeId() string`

GetSrcNodeId returns the SrcNodeId field if non-nil, zero value otherwise.

### GetSrcNodeIdOk

`func (o *AadRelationInfo) GetSrcNodeIdOk() (*string, bool)`

GetSrcNodeIdOk returns a tuple with the SrcNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcNodeId

`func (o *AadRelationInfo) SetSrcNodeId(v string)`

SetSrcNodeId sets SrcNodeId field to given value.

### HasSrcNodeId

`func (o *AadRelationInfo) HasSrcNodeId() bool`

HasSrcNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


