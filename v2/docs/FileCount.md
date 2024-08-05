# FileCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** | Specifies the number of files with size in this range. | [optional] 
**LowerSizeBytes** | Pointer to **NullableInt64** | Specifies the lower bound of file size in bytes. This value is inclusive. | [optional] 
**UpperSizeBytes** | Pointer to **NullableInt64** | Specifies the upper bound of file size in bytes. This value is exclusive. | [optional] 

## Methods

### NewFileCount

`func NewFileCount() *FileCount`

NewFileCount instantiates a new FileCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCountWithDefaults

`func NewFileCountWithDefaults() *FileCount`

NewFileCountWithDefaults instantiates a new FileCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FileCount) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FileCount) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FileCount) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *FileCount) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *FileCount) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *FileCount) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLowerSizeBytes

`func (o *FileCount) GetLowerSizeBytes() int64`

GetLowerSizeBytes returns the LowerSizeBytes field if non-nil, zero value otherwise.

### GetLowerSizeBytesOk

`func (o *FileCount) GetLowerSizeBytesOk() (*int64, bool)`

GetLowerSizeBytesOk returns a tuple with the LowerSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerSizeBytes

`func (o *FileCount) SetLowerSizeBytes(v int64)`

SetLowerSizeBytes sets LowerSizeBytes field to given value.

### HasLowerSizeBytes

`func (o *FileCount) HasLowerSizeBytes() bool`

HasLowerSizeBytes returns a boolean if a field has been set.

### SetLowerSizeBytesNil

`func (o *FileCount) SetLowerSizeBytesNil(b bool)`

 SetLowerSizeBytesNil sets the value for LowerSizeBytes to be an explicit nil

### UnsetLowerSizeBytes
`func (o *FileCount) UnsetLowerSizeBytes()`

UnsetLowerSizeBytes ensures that no value is present for LowerSizeBytes, not even an explicit nil
### GetUpperSizeBytes

`func (o *FileCount) GetUpperSizeBytes() int64`

GetUpperSizeBytes returns the UpperSizeBytes field if non-nil, zero value otherwise.

### GetUpperSizeBytesOk

`func (o *FileCount) GetUpperSizeBytesOk() (*int64, bool)`

GetUpperSizeBytesOk returns a tuple with the UpperSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperSizeBytes

`func (o *FileCount) SetUpperSizeBytes(v int64)`

SetUpperSizeBytes sets UpperSizeBytes field to given value.

### HasUpperSizeBytes

`func (o *FileCount) HasUpperSizeBytes() bool`

HasUpperSizeBytes returns a boolean if a field has been set.

### SetUpperSizeBytesNil

`func (o *FileCount) SetUpperSizeBytesNil(b bool)`

 SetUpperSizeBytesNil sets the value for UpperSizeBytes to be an explicit nil

### UnsetUpperSizeBytes
`func (o *FileCount) UnsetUpperSizeBytes()`

UnsetUpperSizeBytes ensures that no value is present for UpperSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


