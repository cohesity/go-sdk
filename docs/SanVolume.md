# SanVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **NullableString** | Specifies the created time (e.g., \&quot;2015-07-21T17:59:41Z\&quot;) of the volume. | [optional] 
**ParentVolume** | Pointer to **NullableString** | Specifies the name of the source volume, if this volume was copied or cloned from it. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Specifies the serial number of the volume. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the provisioned size in bytes of the volume. | [optional] 
**UsedBytes** | Pointer to **NullableInt64** | Specifies the total space actually used by the volume. | [optional] 

## Methods

### NewSanVolume

`func NewSanVolume() *SanVolume`

NewSanVolume instantiates a new SanVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSanVolumeWithDefaults

`func NewSanVolumeWithDefaults() *SanVolume`

NewSanVolumeWithDefaults instantiates a new SanVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *SanVolume) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SanVolume) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SanVolume) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SanVolume) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *SanVolume) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *SanVolume) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetParentVolume

`func (o *SanVolume) GetParentVolume() string`

GetParentVolume returns the ParentVolume field if non-nil, zero value otherwise.

### GetParentVolumeOk

`func (o *SanVolume) GetParentVolumeOk() (*string, bool)`

GetParentVolumeOk returns a tuple with the ParentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVolume

`func (o *SanVolume) SetParentVolume(v string)`

SetParentVolume sets ParentVolume field to given value.

### HasParentVolume

`func (o *SanVolume) HasParentVolume() bool`

HasParentVolume returns a boolean if a field has been set.

### SetParentVolumeNil

`func (o *SanVolume) SetParentVolumeNil(b bool)`

 SetParentVolumeNil sets the value for ParentVolume to be an explicit nil

### UnsetParentVolume
`func (o *SanVolume) UnsetParentVolume()`

UnsetParentVolume ensures that no value is present for ParentVolume, not even an explicit nil
### GetSerialNumber

`func (o *SanVolume) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SanVolume) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SanVolume) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SanVolume) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *SanVolume) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *SanVolume) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetSizeBytes

`func (o *SanVolume) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SanVolume) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SanVolume) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *SanVolume) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *SanVolume) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *SanVolume) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
### GetUsedBytes

`func (o *SanVolume) GetUsedBytes() int64`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *SanVolume) GetUsedBytesOk() (*int64, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *SanVolume) SetUsedBytes(v int64)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *SanVolume) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### SetUsedBytesNil

`func (o *SanVolume) SetUsedBytesNil(b bool)`

 SetUsedBytesNil sets the value for UsedBytes to be an explicit nil

### UnsetUsedBytes
`func (o *SanVolume) UnsetUsedBytes()`

UnsetUsedBytes ensures that no value is present for UsedBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


