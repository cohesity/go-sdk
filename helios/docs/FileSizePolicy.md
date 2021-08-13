# FileSizePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **NullableString** | Specifies condition for the file selection. | [optional] 
**NBytes** | Pointer to **NullableInt64** | Specifies the number of bytes. | [optional] 

## Methods

### NewFileSizePolicy

`func NewFileSizePolicy() *FileSizePolicy`

NewFileSizePolicy instantiates a new FileSizePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSizePolicyWithDefaults

`func NewFileSizePolicyWithDefaults() *FileSizePolicy`

NewFileSizePolicyWithDefaults instantiates a new FileSizePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *FileSizePolicy) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *FileSizePolicy) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *FileSizePolicy) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *FileSizePolicy) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *FileSizePolicy) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *FileSizePolicy) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetNBytes

`func (o *FileSizePolicy) GetNBytes() int64`

GetNBytes returns the NBytes field if non-nil, zero value otherwise.

### GetNBytesOk

`func (o *FileSizePolicy) GetNBytesOk() (*int64, bool)`

GetNBytesOk returns a tuple with the NBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNBytes

`func (o *FileSizePolicy) SetNBytes(v int64)`

SetNBytes sets NBytes field to given value.

### HasNBytes

`func (o *FileSizePolicy) HasNBytes() bool`

HasNBytes returns a boolean if a field has been set.

### SetNBytesNil

`func (o *FileSizePolicy) SetNBytesNil(b bool)`

 SetNBytesNil sets the value for NBytes to be an explicit nil

### UnsetNBytes
`func (o *FileSizePolicy) UnsetNBytes()`

UnsetNBytes ensures that no value is present for NBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


