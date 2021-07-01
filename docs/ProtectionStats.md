# ProtectionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumFailed** | Pointer to **NullableInt32** | Number of Failed Objects. | [optional] 
**NumObjects** | Pointer to **NullableInt32** | Number of Objects. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Size in Bytes. | [optional] 

## Methods

### NewProtectionStats

`func NewProtectionStats() *ProtectionStats`

NewProtectionStats instantiates a new ProtectionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionStatsWithDefaults

`func NewProtectionStatsWithDefaults() *ProtectionStats`

NewProtectionStatsWithDefaults instantiates a new ProtectionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumFailed

`func (o *ProtectionStats) GetNumFailed() int32`

GetNumFailed returns the NumFailed field if non-nil, zero value otherwise.

### GetNumFailedOk

`func (o *ProtectionStats) GetNumFailedOk() (*int32, bool)`

GetNumFailedOk returns a tuple with the NumFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailed

`func (o *ProtectionStats) SetNumFailed(v int32)`

SetNumFailed sets NumFailed field to given value.

### HasNumFailed

`func (o *ProtectionStats) HasNumFailed() bool`

HasNumFailed returns a boolean if a field has been set.

### SetNumFailedNil

`func (o *ProtectionStats) SetNumFailedNil(b bool)`

 SetNumFailedNil sets the value for NumFailed to be an explicit nil

### UnsetNumFailed
`func (o *ProtectionStats) UnsetNumFailed()`

UnsetNumFailed ensures that no value is present for NumFailed, not even an explicit nil
### GetNumObjects

`func (o *ProtectionStats) GetNumObjects() int32`

GetNumObjects returns the NumObjects field if non-nil, zero value otherwise.

### GetNumObjectsOk

`func (o *ProtectionStats) GetNumObjectsOk() (*int32, bool)`

GetNumObjectsOk returns a tuple with the NumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjects

`func (o *ProtectionStats) SetNumObjects(v int32)`

SetNumObjects sets NumObjects field to given value.

### HasNumObjects

`func (o *ProtectionStats) HasNumObjects() bool`

HasNumObjects returns a boolean if a field has been set.

### SetNumObjectsNil

`func (o *ProtectionStats) SetNumObjectsNil(b bool)`

 SetNumObjectsNil sets the value for NumObjects to be an explicit nil

### UnsetNumObjects
`func (o *ProtectionStats) UnsetNumObjects()`

UnsetNumObjects ensures that no value is present for NumObjects, not even an explicit nil
### GetSizeBytes

`func (o *ProtectionStats) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ProtectionStats) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ProtectionStats) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *ProtectionStats) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *ProtectionStats) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *ProtectionStats) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


