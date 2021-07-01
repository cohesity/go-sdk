# VolumeInfoDiskInfoPhysicalRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **NullableInt64** | Length of this range in bytes. | [optional] 
**Offset** | Pointer to **NullableInt64** | Offset of this range in disk file from beginning of file. | [optional] 

## Methods

### NewVolumeInfoDiskInfoPhysicalRange

`func NewVolumeInfoDiskInfoPhysicalRange() *VolumeInfoDiskInfoPhysicalRange`

NewVolumeInfoDiskInfoPhysicalRange instantiates a new VolumeInfoDiskInfoPhysicalRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeInfoDiskInfoPhysicalRangeWithDefaults

`func NewVolumeInfoDiskInfoPhysicalRangeWithDefaults() *VolumeInfoDiskInfoPhysicalRange`

NewVolumeInfoDiskInfoPhysicalRangeWithDefaults instantiates a new VolumeInfoDiskInfoPhysicalRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *VolumeInfoDiskInfoPhysicalRange) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VolumeInfoDiskInfoPhysicalRange) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VolumeInfoDiskInfoPhysicalRange) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *VolumeInfoDiskInfoPhysicalRange) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *VolumeInfoDiskInfoPhysicalRange) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *VolumeInfoDiskInfoPhysicalRange) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetOffset

`func (o *VolumeInfoDiskInfoPhysicalRange) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VolumeInfoDiskInfoPhysicalRange) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VolumeInfoDiskInfoPhysicalRange) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VolumeInfoDiskInfoPhysicalRange) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *VolumeInfoDiskInfoPhysicalRange) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *VolumeInfoDiskInfoPhysicalRange) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


