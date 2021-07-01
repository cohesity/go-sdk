# ProtectedObjectsByEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvType** | Pointer to **NullableString** | Environment Type. | [optional] 
**ProtectedCount** | Pointer to **NullableInt32** | Number of Protected Objects. | [optional] 
**ProtectedSizeBytes** | Pointer to **NullableInt64** | Size of Protected Objects. | [optional] 
**UnprotectedCount** | Pointer to **NullableInt32** | Number of Unprotected Objects. | [optional] 
**UnprotectedSizeBytes** | Pointer to **NullableInt64** | Size of Unprotected Objects. | [optional] 

## Methods

### NewProtectedObjectsByEnv

`func NewProtectedObjectsByEnv() *ProtectedObjectsByEnv`

NewProtectedObjectsByEnv instantiates a new ProtectedObjectsByEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectsByEnvWithDefaults

`func NewProtectedObjectsByEnvWithDefaults() *ProtectedObjectsByEnv`

NewProtectedObjectsByEnvWithDefaults instantiates a new ProtectedObjectsByEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvType

`func (o *ProtectedObjectsByEnv) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *ProtectedObjectsByEnv) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *ProtectedObjectsByEnv) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *ProtectedObjectsByEnv) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *ProtectedObjectsByEnv) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *ProtectedObjectsByEnv) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetProtectedCount

`func (o *ProtectedObjectsByEnv) GetProtectedCount() int32`

GetProtectedCount returns the ProtectedCount field if non-nil, zero value otherwise.

### GetProtectedCountOk

`func (o *ProtectedObjectsByEnv) GetProtectedCountOk() (*int32, bool)`

GetProtectedCountOk returns a tuple with the ProtectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedCount

`func (o *ProtectedObjectsByEnv) SetProtectedCount(v int32)`

SetProtectedCount sets ProtectedCount field to given value.

### HasProtectedCount

`func (o *ProtectedObjectsByEnv) HasProtectedCount() bool`

HasProtectedCount returns a boolean if a field has been set.

### SetProtectedCountNil

`func (o *ProtectedObjectsByEnv) SetProtectedCountNil(b bool)`

 SetProtectedCountNil sets the value for ProtectedCount to be an explicit nil

### UnsetProtectedCount
`func (o *ProtectedObjectsByEnv) UnsetProtectedCount()`

UnsetProtectedCount ensures that no value is present for ProtectedCount, not even an explicit nil
### GetProtectedSizeBytes

`func (o *ProtectedObjectsByEnv) GetProtectedSizeBytes() int64`

GetProtectedSizeBytes returns the ProtectedSizeBytes field if non-nil, zero value otherwise.

### GetProtectedSizeBytesOk

`func (o *ProtectedObjectsByEnv) GetProtectedSizeBytesOk() (*int64, bool)`

GetProtectedSizeBytesOk returns a tuple with the ProtectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSizeBytes

`func (o *ProtectedObjectsByEnv) SetProtectedSizeBytes(v int64)`

SetProtectedSizeBytes sets ProtectedSizeBytes field to given value.

### HasProtectedSizeBytes

`func (o *ProtectedObjectsByEnv) HasProtectedSizeBytes() bool`

HasProtectedSizeBytes returns a boolean if a field has been set.

### SetProtectedSizeBytesNil

`func (o *ProtectedObjectsByEnv) SetProtectedSizeBytesNil(b bool)`

 SetProtectedSizeBytesNil sets the value for ProtectedSizeBytes to be an explicit nil

### UnsetProtectedSizeBytes
`func (o *ProtectedObjectsByEnv) UnsetProtectedSizeBytes()`

UnsetProtectedSizeBytes ensures that no value is present for ProtectedSizeBytes, not even an explicit nil
### GetUnprotectedCount

`func (o *ProtectedObjectsByEnv) GetUnprotectedCount() int32`

GetUnprotectedCount returns the UnprotectedCount field if non-nil, zero value otherwise.

### GetUnprotectedCountOk

`func (o *ProtectedObjectsByEnv) GetUnprotectedCountOk() (*int32, bool)`

GetUnprotectedCountOk returns a tuple with the UnprotectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedCount

`func (o *ProtectedObjectsByEnv) SetUnprotectedCount(v int32)`

SetUnprotectedCount sets UnprotectedCount field to given value.

### HasUnprotectedCount

`func (o *ProtectedObjectsByEnv) HasUnprotectedCount() bool`

HasUnprotectedCount returns a boolean if a field has been set.

### SetUnprotectedCountNil

`func (o *ProtectedObjectsByEnv) SetUnprotectedCountNil(b bool)`

 SetUnprotectedCountNil sets the value for UnprotectedCount to be an explicit nil

### UnsetUnprotectedCount
`func (o *ProtectedObjectsByEnv) UnsetUnprotectedCount()`

UnsetUnprotectedCount ensures that no value is present for UnprotectedCount, not even an explicit nil
### GetUnprotectedSizeBytes

`func (o *ProtectedObjectsByEnv) GetUnprotectedSizeBytes() int64`

GetUnprotectedSizeBytes returns the UnprotectedSizeBytes field if non-nil, zero value otherwise.

### GetUnprotectedSizeBytesOk

`func (o *ProtectedObjectsByEnv) GetUnprotectedSizeBytesOk() (*int64, bool)`

GetUnprotectedSizeBytesOk returns a tuple with the UnprotectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedSizeBytes

`func (o *ProtectedObjectsByEnv) SetUnprotectedSizeBytes(v int64)`

SetUnprotectedSizeBytes sets UnprotectedSizeBytes field to given value.

### HasUnprotectedSizeBytes

`func (o *ProtectedObjectsByEnv) HasUnprotectedSizeBytes() bool`

HasUnprotectedSizeBytes returns a boolean if a field has been set.

### SetUnprotectedSizeBytesNil

`func (o *ProtectedObjectsByEnv) SetUnprotectedSizeBytesNil(b bool)`

 SetUnprotectedSizeBytesNil sets the value for UnprotectedSizeBytes to be an explicit nil

### UnsetUnprotectedSizeBytes
`func (o *ProtectedObjectsByEnv) UnsetUnprotectedSizeBytes()`

UnsetUnprotectedSizeBytes ensures that no value is present for UnprotectedSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


