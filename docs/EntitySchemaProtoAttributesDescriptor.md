# EntitySchemaProtoAttributesDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeVec** | Pointer to [**[]EntitySchemaProtoKeyValueDescriptor**](EntitySchemaProtoKeyValueDescriptor.md) | Array of Attributes.  List of attributes about an entity. | [optional] 
**KeyAttributeNameIndex** | Pointer to **NullableInt32** | Specifies the attribute to use as a unique identifier for the entity. This value is returned in entityId when the GET public/statistics/entities operation is run. | [optional] 

## Methods

### NewEntitySchemaProtoAttributesDescriptor

`func NewEntitySchemaProtoAttributesDescriptor() *EntitySchemaProtoAttributesDescriptor`

NewEntitySchemaProtoAttributesDescriptor instantiates a new EntitySchemaProtoAttributesDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaProtoAttributesDescriptorWithDefaults

`func NewEntitySchemaProtoAttributesDescriptorWithDefaults() *EntitySchemaProtoAttributesDescriptor`

NewEntitySchemaProtoAttributesDescriptorWithDefaults instantiates a new EntitySchemaProtoAttributesDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeVec

`func (o *EntitySchemaProtoAttributesDescriptor) GetAttributeVec() []EntitySchemaProtoKeyValueDescriptor`

GetAttributeVec returns the AttributeVec field if non-nil, zero value otherwise.

### GetAttributeVecOk

`func (o *EntitySchemaProtoAttributesDescriptor) GetAttributeVecOk() (*[]EntitySchemaProtoKeyValueDescriptor, bool)`

GetAttributeVecOk returns a tuple with the AttributeVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeVec

`func (o *EntitySchemaProtoAttributesDescriptor) SetAttributeVec(v []EntitySchemaProtoKeyValueDescriptor)`

SetAttributeVec sets AttributeVec field to given value.

### HasAttributeVec

`func (o *EntitySchemaProtoAttributesDescriptor) HasAttributeVec() bool`

HasAttributeVec returns a boolean if a field has been set.

### SetAttributeVecNil

`func (o *EntitySchemaProtoAttributesDescriptor) SetAttributeVecNil(b bool)`

 SetAttributeVecNil sets the value for AttributeVec to be an explicit nil

### UnsetAttributeVec
`func (o *EntitySchemaProtoAttributesDescriptor) UnsetAttributeVec()`

UnsetAttributeVec ensures that no value is present for AttributeVec, not even an explicit nil
### GetKeyAttributeNameIndex

`func (o *EntitySchemaProtoAttributesDescriptor) GetKeyAttributeNameIndex() int32`

GetKeyAttributeNameIndex returns the KeyAttributeNameIndex field if non-nil, zero value otherwise.

### GetKeyAttributeNameIndexOk

`func (o *EntitySchemaProtoAttributesDescriptor) GetKeyAttributeNameIndexOk() (*int32, bool)`

GetKeyAttributeNameIndexOk returns a tuple with the KeyAttributeNameIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAttributeNameIndex

`func (o *EntitySchemaProtoAttributesDescriptor) SetKeyAttributeNameIndex(v int32)`

SetKeyAttributeNameIndex sets KeyAttributeNameIndex field to given value.

### HasKeyAttributeNameIndex

`func (o *EntitySchemaProtoAttributesDescriptor) HasKeyAttributeNameIndex() bool`

HasKeyAttributeNameIndex returns a boolean if a field has been set.

### SetKeyAttributeNameIndexNil

`func (o *EntitySchemaProtoAttributesDescriptor) SetKeyAttributeNameIndexNil(b bool)`

 SetKeyAttributeNameIndexNil sets the value for KeyAttributeNameIndex to be an explicit nil

### UnsetKeyAttributeNameIndex
`func (o *EntitySchemaProtoAttributesDescriptor) UnsetKeyAttributeNameIndex()`

UnsetKeyAttributeNameIndex ensures that no value is present for KeyAttributeNameIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


