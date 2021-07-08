# ProtectedObjectsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumObjectsProtected** | Pointer to **NullableInt64** | Specifies the total number of protected objects. | [optional] 
**NumObjectsUnprotected** | Pointer to **NullableInt64** | Specifies the total number of unprotected objects. | [optional] 
**ProtectedSizeBytes** | Pointer to **NullableInt64** | Specifies the total size of protected objects in bytes. | [optional] 
**StatsByEnv** | Pointer to [**[]ProtectedObjectsSummaryByEnv**](ProtectedObjectsSummaryByEnv.md) | Specifies the stats of Protected objects by environment. | [optional] 
**UnprotectedSizeBytes** | Pointer to **NullableInt64** | Specifies the total size of unprotected objects in bytes. | [optional] 

## Methods

### NewProtectedObjectsSummary

`func NewProtectedObjectsSummary() *ProtectedObjectsSummary`

NewProtectedObjectsSummary instantiates a new ProtectedObjectsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectsSummaryWithDefaults

`func NewProtectedObjectsSummaryWithDefaults() *ProtectedObjectsSummary`

NewProtectedObjectsSummaryWithDefaults instantiates a new ProtectedObjectsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumObjectsProtected

`func (o *ProtectedObjectsSummary) GetNumObjectsProtected() int64`

GetNumObjectsProtected returns the NumObjectsProtected field if non-nil, zero value otherwise.

### GetNumObjectsProtectedOk

`func (o *ProtectedObjectsSummary) GetNumObjectsProtectedOk() (*int64, bool)`

GetNumObjectsProtectedOk returns a tuple with the NumObjectsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsProtected

`func (o *ProtectedObjectsSummary) SetNumObjectsProtected(v int64)`

SetNumObjectsProtected sets NumObjectsProtected field to given value.

### HasNumObjectsProtected

`func (o *ProtectedObjectsSummary) HasNumObjectsProtected() bool`

HasNumObjectsProtected returns a boolean if a field has been set.

### SetNumObjectsProtectedNil

`func (o *ProtectedObjectsSummary) SetNumObjectsProtectedNil(b bool)`

 SetNumObjectsProtectedNil sets the value for NumObjectsProtected to be an explicit nil

### UnsetNumObjectsProtected
`func (o *ProtectedObjectsSummary) UnsetNumObjectsProtected()`

UnsetNumObjectsProtected ensures that no value is present for NumObjectsProtected, not even an explicit nil
### GetNumObjectsUnprotected

`func (o *ProtectedObjectsSummary) GetNumObjectsUnprotected() int64`

GetNumObjectsUnprotected returns the NumObjectsUnprotected field if non-nil, zero value otherwise.

### GetNumObjectsUnprotectedOk

`func (o *ProtectedObjectsSummary) GetNumObjectsUnprotectedOk() (*int64, bool)`

GetNumObjectsUnprotectedOk returns a tuple with the NumObjectsUnprotected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsUnprotected

`func (o *ProtectedObjectsSummary) SetNumObjectsUnprotected(v int64)`

SetNumObjectsUnprotected sets NumObjectsUnprotected field to given value.

### HasNumObjectsUnprotected

`func (o *ProtectedObjectsSummary) HasNumObjectsUnprotected() bool`

HasNumObjectsUnprotected returns a boolean if a field has been set.

### SetNumObjectsUnprotectedNil

`func (o *ProtectedObjectsSummary) SetNumObjectsUnprotectedNil(b bool)`

 SetNumObjectsUnprotectedNil sets the value for NumObjectsUnprotected to be an explicit nil

### UnsetNumObjectsUnprotected
`func (o *ProtectedObjectsSummary) UnsetNumObjectsUnprotected()`

UnsetNumObjectsUnprotected ensures that no value is present for NumObjectsUnprotected, not even an explicit nil
### GetProtectedSizeBytes

`func (o *ProtectedObjectsSummary) GetProtectedSizeBytes() int64`

GetProtectedSizeBytes returns the ProtectedSizeBytes field if non-nil, zero value otherwise.

### GetProtectedSizeBytesOk

`func (o *ProtectedObjectsSummary) GetProtectedSizeBytesOk() (*int64, bool)`

GetProtectedSizeBytesOk returns a tuple with the ProtectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSizeBytes

`func (o *ProtectedObjectsSummary) SetProtectedSizeBytes(v int64)`

SetProtectedSizeBytes sets ProtectedSizeBytes field to given value.

### HasProtectedSizeBytes

`func (o *ProtectedObjectsSummary) HasProtectedSizeBytes() bool`

HasProtectedSizeBytes returns a boolean if a field has been set.

### SetProtectedSizeBytesNil

`func (o *ProtectedObjectsSummary) SetProtectedSizeBytesNil(b bool)`

 SetProtectedSizeBytesNil sets the value for ProtectedSizeBytes to be an explicit nil

### UnsetProtectedSizeBytes
`func (o *ProtectedObjectsSummary) UnsetProtectedSizeBytes()`

UnsetProtectedSizeBytes ensures that no value is present for ProtectedSizeBytes, not even an explicit nil
### GetStatsByEnv

`func (o *ProtectedObjectsSummary) GetStatsByEnv() []ProtectedObjectsSummaryByEnv`

GetStatsByEnv returns the StatsByEnv field if non-nil, zero value otherwise.

### GetStatsByEnvOk

`func (o *ProtectedObjectsSummary) GetStatsByEnvOk() (*[]ProtectedObjectsSummaryByEnv, bool)`

GetStatsByEnvOk returns a tuple with the StatsByEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByEnv

`func (o *ProtectedObjectsSummary) SetStatsByEnv(v []ProtectedObjectsSummaryByEnv)`

SetStatsByEnv sets StatsByEnv field to given value.

### HasStatsByEnv

`func (o *ProtectedObjectsSummary) HasStatsByEnv() bool`

HasStatsByEnv returns a boolean if a field has been set.

### GetUnprotectedSizeBytes

`func (o *ProtectedObjectsSummary) GetUnprotectedSizeBytes() int64`

GetUnprotectedSizeBytes returns the UnprotectedSizeBytes field if non-nil, zero value otherwise.

### GetUnprotectedSizeBytesOk

`func (o *ProtectedObjectsSummary) GetUnprotectedSizeBytesOk() (*int64, bool)`

GetUnprotectedSizeBytesOk returns a tuple with the UnprotectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedSizeBytes

`func (o *ProtectedObjectsSummary) SetUnprotectedSizeBytes(v int64)`

SetUnprotectedSizeBytes sets UnprotectedSizeBytes field to given value.

### HasUnprotectedSizeBytes

`func (o *ProtectedObjectsSummary) HasUnprotectedSizeBytes() bool`

HasUnprotectedSizeBytes returns a boolean if a field has been set.

### SetUnprotectedSizeBytesNil

`func (o *ProtectedObjectsSummary) SetUnprotectedSizeBytesNil(b bool)`

 SetUnprotectedSizeBytesNil sets the value for UnprotectedSizeBytes to be an explicit nil

### UnsetUnprotectedSizeBytes
`func (o *ProtectedObjectsSummary) UnsetUnprotectedSizeBytes()`

UnsetUnprotectedSizeBytes ensures that no value is present for UnprotectedSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


