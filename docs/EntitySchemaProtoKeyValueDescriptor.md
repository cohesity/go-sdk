# EntitySchemaProtoKeyValueDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | Pointer to **NullableString** | Specifies the name of a key. | [optional] 
**ValueType** | Pointer to **NullableInt32** | Specifies the type of the value that is associated with the key. 0 specifies a value type of Int64. 1 specifies a value type of Double. 2 specifies a value type of String. 3 specifies a value type of Bytes. | [optional] 

## Methods

### NewEntitySchemaProtoKeyValueDescriptor

`func NewEntitySchemaProtoKeyValueDescriptor() *EntitySchemaProtoKeyValueDescriptor`

NewEntitySchemaProtoKeyValueDescriptor instantiates a new EntitySchemaProtoKeyValueDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaProtoKeyValueDescriptorWithDefaults

`func NewEntitySchemaProtoKeyValueDescriptorWithDefaults() *EntitySchemaProtoKeyValueDescriptor`

NewEntitySchemaProtoKeyValueDescriptorWithDefaults instantiates a new EntitySchemaProtoKeyValueDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *EntitySchemaProtoKeyValueDescriptor) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *EntitySchemaProtoKeyValueDescriptor) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *EntitySchemaProtoKeyValueDescriptor) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *EntitySchemaProtoKeyValueDescriptor) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *EntitySchemaProtoKeyValueDescriptor) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *EntitySchemaProtoKeyValueDescriptor) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetValueType

`func (o *EntitySchemaProtoKeyValueDescriptor) GetValueType() int32`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *EntitySchemaProtoKeyValueDescriptor) GetValueTypeOk() (*int32, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *EntitySchemaProtoKeyValueDescriptor) SetValueType(v int32)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *EntitySchemaProtoKeyValueDescriptor) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *EntitySchemaProtoKeyValueDescriptor) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *EntitySchemaProtoKeyValueDescriptor) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


