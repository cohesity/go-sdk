# RecoverMailboxParamsPstParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatePst** | Pointer to **NullableBool** | Specifies if create a PST or MSG for input items. | [optional] 
**Password** | **NullableString** | Specifies Password to be set for generated PSTs. | 
**SizeThresholdBytes** | Pointer to **NullableInt64** | Specifies PST size threshold in bytes. | [optional] 

## Methods

### NewRecoverMailboxParamsPstParams

`func NewRecoverMailboxParamsPstParams(password NullableString, ) *RecoverMailboxParamsPstParams`

NewRecoverMailboxParamsPstParams instantiates a new RecoverMailboxParamsPstParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverMailboxParamsPstParamsWithDefaults

`func NewRecoverMailboxParamsPstParamsWithDefaults() *RecoverMailboxParamsPstParams`

NewRecoverMailboxParamsPstParamsWithDefaults instantiates a new RecoverMailboxParamsPstParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatePst

`func (o *RecoverMailboxParamsPstParams) GetCreatePst() bool`

GetCreatePst returns the CreatePst field if non-nil, zero value otherwise.

### GetCreatePstOk

`func (o *RecoverMailboxParamsPstParams) GetCreatePstOk() (*bool, bool)`

GetCreatePstOk returns a tuple with the CreatePst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePst

`func (o *RecoverMailboxParamsPstParams) SetCreatePst(v bool)`

SetCreatePst sets CreatePst field to given value.

### HasCreatePst

`func (o *RecoverMailboxParamsPstParams) HasCreatePst() bool`

HasCreatePst returns a boolean if a field has been set.

### SetCreatePstNil

`func (o *RecoverMailboxParamsPstParams) SetCreatePstNil(b bool)`

 SetCreatePstNil sets the value for CreatePst to be an explicit nil

### UnsetCreatePst
`func (o *RecoverMailboxParamsPstParams) UnsetCreatePst()`

UnsetCreatePst ensures that no value is present for CreatePst, not even an explicit nil
### GetPassword

`func (o *RecoverMailboxParamsPstParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RecoverMailboxParamsPstParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RecoverMailboxParamsPstParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *RecoverMailboxParamsPstParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RecoverMailboxParamsPstParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSizeThresholdBytes

`func (o *RecoverMailboxParamsPstParams) GetSizeThresholdBytes() int64`

GetSizeThresholdBytes returns the SizeThresholdBytes field if non-nil, zero value otherwise.

### GetSizeThresholdBytesOk

`func (o *RecoverMailboxParamsPstParams) GetSizeThresholdBytesOk() (*int64, bool)`

GetSizeThresholdBytesOk returns a tuple with the SizeThresholdBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeThresholdBytes

`func (o *RecoverMailboxParamsPstParams) SetSizeThresholdBytes(v int64)`

SetSizeThresholdBytes sets SizeThresholdBytes field to given value.

### HasSizeThresholdBytes

`func (o *RecoverMailboxParamsPstParams) HasSizeThresholdBytes() bool`

HasSizeThresholdBytes returns a boolean if a field has been set.

### SetSizeThresholdBytesNil

`func (o *RecoverMailboxParamsPstParams) SetSizeThresholdBytesNil(b bool)`

 SetSizeThresholdBytesNil sets the value for SizeThresholdBytes to be an explicit nil

### UnsetSizeThresholdBytes
`func (o *RecoverMailboxParamsPstParams) UnsetSizeThresholdBytes()`

UnsetSizeThresholdBytes ensures that no value is present for SizeThresholdBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


