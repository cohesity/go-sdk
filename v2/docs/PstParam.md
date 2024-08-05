# PstParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatePst** | Pointer to **NullableBool** | Specifies if create a PST or MSG for input items. | [optional] 
**Password** | **NullableString** | Specifies Password to be set for generated PSTs. | 
**SizeThresholdBytes** | Pointer to **NullableInt64** | Specifies PST size threshold in bytes. | [optional] 

## Methods

### NewPstParam

`func NewPstParam(password NullableString, ) *PstParam`

NewPstParam instantiates a new PstParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPstParamWithDefaults

`func NewPstParamWithDefaults() *PstParam`

NewPstParamWithDefaults instantiates a new PstParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatePst

`func (o *PstParam) GetCreatePst() bool`

GetCreatePst returns the CreatePst field if non-nil, zero value otherwise.

### GetCreatePstOk

`func (o *PstParam) GetCreatePstOk() (*bool, bool)`

GetCreatePstOk returns a tuple with the CreatePst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePst

`func (o *PstParam) SetCreatePst(v bool)`

SetCreatePst sets CreatePst field to given value.

### HasCreatePst

`func (o *PstParam) HasCreatePst() bool`

HasCreatePst returns a boolean if a field has been set.

### SetCreatePstNil

`func (o *PstParam) SetCreatePstNil(b bool)`

 SetCreatePstNil sets the value for CreatePst to be an explicit nil

### UnsetCreatePst
`func (o *PstParam) UnsetCreatePst()`

UnsetCreatePst ensures that no value is present for CreatePst, not even an explicit nil
### GetPassword

`func (o *PstParam) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PstParam) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PstParam) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *PstParam) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PstParam) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSizeThresholdBytes

`func (o *PstParam) GetSizeThresholdBytes() int64`

GetSizeThresholdBytes returns the SizeThresholdBytes field if non-nil, zero value otherwise.

### GetSizeThresholdBytesOk

`func (o *PstParam) GetSizeThresholdBytesOk() (*int64, bool)`

GetSizeThresholdBytesOk returns a tuple with the SizeThresholdBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeThresholdBytes

`func (o *PstParam) SetSizeThresholdBytes(v int64)`

SetSizeThresholdBytes sets SizeThresholdBytes field to given value.

### HasSizeThresholdBytes

`func (o *PstParam) HasSizeThresholdBytes() bool`

HasSizeThresholdBytes returns a boolean if a field has been set.

### SetSizeThresholdBytesNil

`func (o *PstParam) SetSizeThresholdBytesNil(b bool)`

 SetSizeThresholdBytesNil sets the value for SizeThresholdBytes to be an explicit nil

### UnsetSizeThresholdBytes
`func (o *PstParam) UnsetSizeThresholdBytes()`

UnsetSizeThresholdBytes ensures that no value is present for SizeThresholdBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


