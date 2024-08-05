# FreeDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** | Specifies the location of disk. | [optional] 
**Path** | Pointer to **NullableString** | Specifies path of disk. | [optional] 
**SerialNumber** | **NullableString** | Specifies serial number of disk. | 
**SizeInBytes** | Pointer to **NullableInt64** | Size of disk. | [optional] 

## Methods

### NewFreeDisk

`func NewFreeDisk(serialNumber NullableString, ) *FreeDisk`

NewFreeDisk instantiates a new FreeDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeDiskWithDefaults

`func NewFreeDiskWithDefaults() *FreeDisk`

NewFreeDiskWithDefaults instantiates a new FreeDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *FreeDisk) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FreeDisk) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FreeDisk) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FreeDisk) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *FreeDisk) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *FreeDisk) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetPath

`func (o *FreeDisk) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FreeDisk) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FreeDisk) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FreeDisk) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *FreeDisk) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *FreeDisk) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSerialNumber

`func (o *FreeDisk) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *FreeDisk) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *FreeDisk) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### SetSerialNumberNil

`func (o *FreeDisk) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *FreeDisk) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetSizeInBytes

`func (o *FreeDisk) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *FreeDisk) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *FreeDisk) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *FreeDisk) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.

### SetSizeInBytesNil

`func (o *FreeDisk) SetSizeInBytesNil(b bool)`

 SetSizeInBytesNil sets the value for SizeInBytes to be an explicit nil

### UnsetSizeInBytes
`func (o *FreeDisk) UnsetSizeInBytes()`

UnsetSizeInBytes ensures that no value is present for SizeInBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


