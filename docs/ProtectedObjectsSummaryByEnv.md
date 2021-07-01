# ProtectedObjectsSummaryByEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment. | [optional] 
**NumObjectsProtected** | Pointer to **NullableInt64** | Specifies the total number of protected objects. | [optional] 
**NumObjectsUnprotected** | Pointer to **NullableInt64** | Specifies the total number of unprotected objects. | [optional] 
**ProtectedSizeBytes** | Pointer to **NullableInt64** | Specifies the total size of protected objects in bytes. | [optional] 
**UnprotectedSizeBytes** | Pointer to **NullableInt64** | Specifies the total size of unprotected objects in bytes. | [optional] 

## Methods

### NewProtectedObjectsSummaryByEnv

`func NewProtectedObjectsSummaryByEnv() *ProtectedObjectsSummaryByEnv`

NewProtectedObjectsSummaryByEnv instantiates a new ProtectedObjectsSummaryByEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectsSummaryByEnvWithDefaults

`func NewProtectedObjectsSummaryByEnvWithDefaults() *ProtectedObjectsSummaryByEnv`

NewProtectedObjectsSummaryByEnvWithDefaults instantiates a new ProtectedObjectsSummaryByEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ProtectedObjectsSummaryByEnv) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectedObjectsSummaryByEnv) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectedObjectsSummaryByEnv) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectedObjectsSummaryByEnv) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectedObjectsSummaryByEnv) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectedObjectsSummaryByEnv) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetNumObjectsProtected

`func (o *ProtectedObjectsSummaryByEnv) GetNumObjectsProtected() int64`

GetNumObjectsProtected returns the NumObjectsProtected field if non-nil, zero value otherwise.

### GetNumObjectsProtectedOk

`func (o *ProtectedObjectsSummaryByEnv) GetNumObjectsProtectedOk() (*int64, bool)`

GetNumObjectsProtectedOk returns a tuple with the NumObjectsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsProtected

`func (o *ProtectedObjectsSummaryByEnv) SetNumObjectsProtected(v int64)`

SetNumObjectsProtected sets NumObjectsProtected field to given value.

### HasNumObjectsProtected

`func (o *ProtectedObjectsSummaryByEnv) HasNumObjectsProtected() bool`

HasNumObjectsProtected returns a boolean if a field has been set.

### SetNumObjectsProtectedNil

`func (o *ProtectedObjectsSummaryByEnv) SetNumObjectsProtectedNil(b bool)`

 SetNumObjectsProtectedNil sets the value for NumObjectsProtected to be an explicit nil

### UnsetNumObjectsProtected
`func (o *ProtectedObjectsSummaryByEnv) UnsetNumObjectsProtected()`

UnsetNumObjectsProtected ensures that no value is present for NumObjectsProtected, not even an explicit nil
### GetNumObjectsUnprotected

`func (o *ProtectedObjectsSummaryByEnv) GetNumObjectsUnprotected() int64`

GetNumObjectsUnprotected returns the NumObjectsUnprotected field if non-nil, zero value otherwise.

### GetNumObjectsUnprotectedOk

`func (o *ProtectedObjectsSummaryByEnv) GetNumObjectsUnprotectedOk() (*int64, bool)`

GetNumObjectsUnprotectedOk returns a tuple with the NumObjectsUnprotected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsUnprotected

`func (o *ProtectedObjectsSummaryByEnv) SetNumObjectsUnprotected(v int64)`

SetNumObjectsUnprotected sets NumObjectsUnprotected field to given value.

### HasNumObjectsUnprotected

`func (o *ProtectedObjectsSummaryByEnv) HasNumObjectsUnprotected() bool`

HasNumObjectsUnprotected returns a boolean if a field has been set.

### SetNumObjectsUnprotectedNil

`func (o *ProtectedObjectsSummaryByEnv) SetNumObjectsUnprotectedNil(b bool)`

 SetNumObjectsUnprotectedNil sets the value for NumObjectsUnprotected to be an explicit nil

### UnsetNumObjectsUnprotected
`func (o *ProtectedObjectsSummaryByEnv) UnsetNumObjectsUnprotected()`

UnsetNumObjectsUnprotected ensures that no value is present for NumObjectsUnprotected, not even an explicit nil
### GetProtectedSizeBytes

`func (o *ProtectedObjectsSummaryByEnv) GetProtectedSizeBytes() int64`

GetProtectedSizeBytes returns the ProtectedSizeBytes field if non-nil, zero value otherwise.

### GetProtectedSizeBytesOk

`func (o *ProtectedObjectsSummaryByEnv) GetProtectedSizeBytesOk() (*int64, bool)`

GetProtectedSizeBytesOk returns a tuple with the ProtectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSizeBytes

`func (o *ProtectedObjectsSummaryByEnv) SetProtectedSizeBytes(v int64)`

SetProtectedSizeBytes sets ProtectedSizeBytes field to given value.

### HasProtectedSizeBytes

`func (o *ProtectedObjectsSummaryByEnv) HasProtectedSizeBytes() bool`

HasProtectedSizeBytes returns a boolean if a field has been set.

### SetProtectedSizeBytesNil

`func (o *ProtectedObjectsSummaryByEnv) SetProtectedSizeBytesNil(b bool)`

 SetProtectedSizeBytesNil sets the value for ProtectedSizeBytes to be an explicit nil

### UnsetProtectedSizeBytes
`func (o *ProtectedObjectsSummaryByEnv) UnsetProtectedSizeBytes()`

UnsetProtectedSizeBytes ensures that no value is present for ProtectedSizeBytes, not even an explicit nil
### GetUnprotectedSizeBytes

`func (o *ProtectedObjectsSummaryByEnv) GetUnprotectedSizeBytes() int64`

GetUnprotectedSizeBytes returns the UnprotectedSizeBytes field if non-nil, zero value otherwise.

### GetUnprotectedSizeBytesOk

`func (o *ProtectedObjectsSummaryByEnv) GetUnprotectedSizeBytesOk() (*int64, bool)`

GetUnprotectedSizeBytesOk returns a tuple with the UnprotectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedSizeBytes

`func (o *ProtectedObjectsSummaryByEnv) SetUnprotectedSizeBytes(v int64)`

SetUnprotectedSizeBytes sets UnprotectedSizeBytes field to given value.

### HasUnprotectedSizeBytes

`func (o *ProtectedObjectsSummaryByEnv) HasUnprotectedSizeBytes() bool`

HasUnprotectedSizeBytes returns a boolean if a field has been set.

### SetUnprotectedSizeBytesNil

`func (o *ProtectedObjectsSummaryByEnv) SetUnprotectedSizeBytesNil(b bool)`

 SetUnprotectedSizeBytesNil sets the value for UnprotectedSizeBytes to be an explicit nil

### UnsetUnprotectedSizeBytes
`func (o *ProtectedObjectsSummaryByEnv) UnsetUnprotectedSizeBytes()`

UnsetUnprotectedSizeBytes ensures that no value is present for UnprotectedSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


