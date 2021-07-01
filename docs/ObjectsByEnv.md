# ObjectsByEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvType** | Pointer to **NullableString** | Environment Type. | [optional] 
**NumObjects** | Pointer to **NullableInt32** | Number of Objects. | [optional] 

## Methods

### NewObjectsByEnv

`func NewObjectsByEnv() *ObjectsByEnv`

NewObjectsByEnv instantiates a new ObjectsByEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsByEnvWithDefaults

`func NewObjectsByEnvWithDefaults() *ObjectsByEnv`

NewObjectsByEnvWithDefaults instantiates a new ObjectsByEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvType

`func (o *ObjectsByEnv) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *ObjectsByEnv) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *ObjectsByEnv) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *ObjectsByEnv) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *ObjectsByEnv) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *ObjectsByEnv) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetNumObjects

`func (o *ObjectsByEnv) GetNumObjects() int32`

GetNumObjects returns the NumObjects field if non-nil, zero value otherwise.

### GetNumObjectsOk

`func (o *ObjectsByEnv) GetNumObjectsOk() (*int32, bool)`

GetNumObjectsOk returns a tuple with the NumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjects

`func (o *ObjectsByEnv) SetNumObjects(v int32)`

SetNumObjects sets NumObjects field to given value.

### HasNumObjects

`func (o *ObjectsByEnv) HasNumObjects() bool`

HasNumObjects returns a boolean if a field has been set.

### SetNumObjectsNil

`func (o *ObjectsByEnv) SetNumObjectsNil(b bool)`

 SetNumObjectsNil sets the value for NumObjects to be an explicit nil

### UnsetNumObjects
`func (o *ObjectsByEnv) UnsetNumObjects()`

UnsetNumObjects ensures that no value is present for NumObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


