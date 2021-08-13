# EbsVolumeExclusionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeIds** | Pointer to **[]string** | Array of volume IDs that are to be excluded. This is only for object level exclusion. | [optional] 
**MaxVolumeSizeBytes** | Pointer to **NullableInt64** | Any volume larger than this size will be excluded. | [optional] 
**VolumeTypes** | Pointer to **[]string** | Array of volume types to exclude. Eg - gp2, gp3. | [optional] 
**DeviceNames** | Pointer to **[]string** | Array of device names to exclude. Eg - /dev/sda. | [optional] 

## Methods

### NewEbsVolumeExclusionParams

`func NewEbsVolumeExclusionParams() *EbsVolumeExclusionParams`

NewEbsVolumeExclusionParams instantiates a new EbsVolumeExclusionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbsVolumeExclusionParamsWithDefaults

`func NewEbsVolumeExclusionParamsWithDefaults() *EbsVolumeExclusionParams`

NewEbsVolumeExclusionParamsWithDefaults instantiates a new EbsVolumeExclusionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeIds

`func (o *EbsVolumeExclusionParams) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *EbsVolumeExclusionParams) GetVolumeIdsOk() (*[]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIds

`func (o *EbsVolumeExclusionParams) SetVolumeIds(v []string)`

SetVolumeIds sets VolumeIds field to given value.

### HasVolumeIds

`func (o *EbsVolumeExclusionParams) HasVolumeIds() bool`

HasVolumeIds returns a boolean if a field has been set.

### SetVolumeIdsNil

`func (o *EbsVolumeExclusionParams) SetVolumeIdsNil(b bool)`

 SetVolumeIdsNil sets the value for VolumeIds to be an explicit nil

### UnsetVolumeIds
`func (o *EbsVolumeExclusionParams) UnsetVolumeIds()`

UnsetVolumeIds ensures that no value is present for VolumeIds, not even an explicit nil
### GetMaxVolumeSizeBytes

`func (o *EbsVolumeExclusionParams) GetMaxVolumeSizeBytes() int64`

GetMaxVolumeSizeBytes returns the MaxVolumeSizeBytes field if non-nil, zero value otherwise.

### GetMaxVolumeSizeBytesOk

`func (o *EbsVolumeExclusionParams) GetMaxVolumeSizeBytesOk() (*int64, bool)`

GetMaxVolumeSizeBytesOk returns a tuple with the MaxVolumeSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumeSizeBytes

`func (o *EbsVolumeExclusionParams) SetMaxVolumeSizeBytes(v int64)`

SetMaxVolumeSizeBytes sets MaxVolumeSizeBytes field to given value.

### HasMaxVolumeSizeBytes

`func (o *EbsVolumeExclusionParams) HasMaxVolumeSizeBytes() bool`

HasMaxVolumeSizeBytes returns a boolean if a field has been set.

### SetMaxVolumeSizeBytesNil

`func (o *EbsVolumeExclusionParams) SetMaxVolumeSizeBytesNil(b bool)`

 SetMaxVolumeSizeBytesNil sets the value for MaxVolumeSizeBytes to be an explicit nil

### UnsetMaxVolumeSizeBytes
`func (o *EbsVolumeExclusionParams) UnsetMaxVolumeSizeBytes()`

UnsetMaxVolumeSizeBytes ensures that no value is present for MaxVolumeSizeBytes, not even an explicit nil
### GetVolumeTypes

`func (o *EbsVolumeExclusionParams) GetVolumeTypes() []string`

GetVolumeTypes returns the VolumeTypes field if non-nil, zero value otherwise.

### GetVolumeTypesOk

`func (o *EbsVolumeExclusionParams) GetVolumeTypesOk() (*[]string, bool)`

GetVolumeTypesOk returns a tuple with the VolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypes

`func (o *EbsVolumeExclusionParams) SetVolumeTypes(v []string)`

SetVolumeTypes sets VolumeTypes field to given value.

### HasVolumeTypes

`func (o *EbsVolumeExclusionParams) HasVolumeTypes() bool`

HasVolumeTypes returns a boolean if a field has been set.

### SetVolumeTypesNil

`func (o *EbsVolumeExclusionParams) SetVolumeTypesNil(b bool)`

 SetVolumeTypesNil sets the value for VolumeTypes to be an explicit nil

### UnsetVolumeTypes
`func (o *EbsVolumeExclusionParams) UnsetVolumeTypes()`

UnsetVolumeTypes ensures that no value is present for VolumeTypes, not even an explicit nil
### GetDeviceNames

`func (o *EbsVolumeExclusionParams) GetDeviceNames() []string`

GetDeviceNames returns the DeviceNames field if non-nil, zero value otherwise.

### GetDeviceNamesOk

`func (o *EbsVolumeExclusionParams) GetDeviceNamesOk() (*[]string, bool)`

GetDeviceNamesOk returns a tuple with the DeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNames

`func (o *EbsVolumeExclusionParams) SetDeviceNames(v []string)`

SetDeviceNames sets DeviceNames field to given value.

### HasDeviceNames

`func (o *EbsVolumeExclusionParams) HasDeviceNames() bool`

HasDeviceNames returns a boolean if a field has been set.

### SetDeviceNamesNil

`func (o *EbsVolumeExclusionParams) SetDeviceNamesNil(b bool)`

 SetDeviceNamesNil sets the value for DeviceNames to be an explicit nil

### UnsetDeviceNames
`func (o *EbsVolumeExclusionParams) UnsetDeviceNames()`

UnsetDeviceNames ensures that no value is present for DeviceNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


