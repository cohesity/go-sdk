# ObjectProtectionStatsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**ProtectedCount** | Pointer to **NullableInt64** | Specifies the count of the protected leaf objects. | [optional] 
**UnprotectedCount** | Pointer to **NullableInt64** | Specifies the count of the unprotected leaf objects. | [optional] 
**ProtectedSizeBytes** | Pointer to **NullableInt64** | Specifies the protected logical size in bytes. | [optional] 
**UnprotectedSizeBytes** | Pointer to **NullableInt64** | Specifies the unprotected logical size in bytes. | [optional] 

## Methods

### NewObjectProtectionStatsSummary

`func NewObjectProtectionStatsSummary() *ObjectProtectionStatsSummary`

NewObjectProtectionStatsSummary instantiates a new ObjectProtectionStatsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectProtectionStatsSummaryWithDefaults

`func NewObjectProtectionStatsSummaryWithDefaults() *ObjectProtectionStatsSummary`

NewObjectProtectionStatsSummaryWithDefaults instantiates a new ObjectProtectionStatsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ObjectProtectionStatsSummary) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectProtectionStatsSummary) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectProtectionStatsSummary) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectProtectionStatsSummary) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectProtectionStatsSummary) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectProtectionStatsSummary) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetProtectedCount

`func (o *ObjectProtectionStatsSummary) GetProtectedCount() int64`

GetProtectedCount returns the ProtectedCount field if non-nil, zero value otherwise.

### GetProtectedCountOk

`func (o *ObjectProtectionStatsSummary) GetProtectedCountOk() (*int64, bool)`

GetProtectedCountOk returns a tuple with the ProtectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedCount

`func (o *ObjectProtectionStatsSummary) SetProtectedCount(v int64)`

SetProtectedCount sets ProtectedCount field to given value.

### HasProtectedCount

`func (o *ObjectProtectionStatsSummary) HasProtectedCount() bool`

HasProtectedCount returns a boolean if a field has been set.

### SetProtectedCountNil

`func (o *ObjectProtectionStatsSummary) SetProtectedCountNil(b bool)`

 SetProtectedCountNil sets the value for ProtectedCount to be an explicit nil

### UnsetProtectedCount
`func (o *ObjectProtectionStatsSummary) UnsetProtectedCount()`

UnsetProtectedCount ensures that no value is present for ProtectedCount, not even an explicit nil
### GetUnprotectedCount

`func (o *ObjectProtectionStatsSummary) GetUnprotectedCount() int64`

GetUnprotectedCount returns the UnprotectedCount field if non-nil, zero value otherwise.

### GetUnprotectedCountOk

`func (o *ObjectProtectionStatsSummary) GetUnprotectedCountOk() (*int64, bool)`

GetUnprotectedCountOk returns a tuple with the UnprotectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedCount

`func (o *ObjectProtectionStatsSummary) SetUnprotectedCount(v int64)`

SetUnprotectedCount sets UnprotectedCount field to given value.

### HasUnprotectedCount

`func (o *ObjectProtectionStatsSummary) HasUnprotectedCount() bool`

HasUnprotectedCount returns a boolean if a field has been set.

### SetUnprotectedCountNil

`func (o *ObjectProtectionStatsSummary) SetUnprotectedCountNil(b bool)`

 SetUnprotectedCountNil sets the value for UnprotectedCount to be an explicit nil

### UnsetUnprotectedCount
`func (o *ObjectProtectionStatsSummary) UnsetUnprotectedCount()`

UnsetUnprotectedCount ensures that no value is present for UnprotectedCount, not even an explicit nil
### GetProtectedSizeBytes

`func (o *ObjectProtectionStatsSummary) GetProtectedSizeBytes() int64`

GetProtectedSizeBytes returns the ProtectedSizeBytes field if non-nil, zero value otherwise.

### GetProtectedSizeBytesOk

`func (o *ObjectProtectionStatsSummary) GetProtectedSizeBytesOk() (*int64, bool)`

GetProtectedSizeBytesOk returns a tuple with the ProtectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSizeBytes

`func (o *ObjectProtectionStatsSummary) SetProtectedSizeBytes(v int64)`

SetProtectedSizeBytes sets ProtectedSizeBytes field to given value.

### HasProtectedSizeBytes

`func (o *ObjectProtectionStatsSummary) HasProtectedSizeBytes() bool`

HasProtectedSizeBytes returns a boolean if a field has been set.

### SetProtectedSizeBytesNil

`func (o *ObjectProtectionStatsSummary) SetProtectedSizeBytesNil(b bool)`

 SetProtectedSizeBytesNil sets the value for ProtectedSizeBytes to be an explicit nil

### UnsetProtectedSizeBytes
`func (o *ObjectProtectionStatsSummary) UnsetProtectedSizeBytes()`

UnsetProtectedSizeBytes ensures that no value is present for ProtectedSizeBytes, not even an explicit nil
### GetUnprotectedSizeBytes

`func (o *ObjectProtectionStatsSummary) GetUnprotectedSizeBytes() int64`

GetUnprotectedSizeBytes returns the UnprotectedSizeBytes field if non-nil, zero value otherwise.

### GetUnprotectedSizeBytesOk

`func (o *ObjectProtectionStatsSummary) GetUnprotectedSizeBytesOk() (*int64, bool)`

GetUnprotectedSizeBytesOk returns a tuple with the UnprotectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedSizeBytes

`func (o *ObjectProtectionStatsSummary) SetUnprotectedSizeBytes(v int64)`

SetUnprotectedSizeBytes sets UnprotectedSizeBytes field to given value.

### HasUnprotectedSizeBytes

`func (o *ObjectProtectionStatsSummary) HasUnprotectedSizeBytes() bool`

HasUnprotectedSizeBytes returns a boolean if a field has been set.

### SetUnprotectedSizeBytesNil

`func (o *ObjectProtectionStatsSummary) SetUnprotectedSizeBytesNil(b bool)`

 SetUnprotectedSizeBytesNil sets the value for UnprotectedSizeBytes to be an explicit nil

### UnsetUnprotectedSizeBytes
`func (o *ObjectProtectionStatsSummary) UnsetUnprotectedSizeBytes()`

UnsetUnprotectedSizeBytes ensures that no value is present for UnprotectedSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


