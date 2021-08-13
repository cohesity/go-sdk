# VerifyTotpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpCode** | Pointer to **NullableString** | Specifies the Totp code. | [optional] 

## Methods

### NewVerifyTotpRequest

`func NewVerifyTotpRequest() *VerifyTotpRequest`

NewVerifyTotpRequest instantiates a new VerifyTotpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyTotpRequestWithDefaults

`func NewVerifyTotpRequestWithDefaults() *VerifyTotpRequest`

NewVerifyTotpRequestWithDefaults instantiates a new VerifyTotpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpCode

`func (o *VerifyTotpRequest) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *VerifyTotpRequest) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *VerifyTotpRequest) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *VerifyTotpRequest) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.

### SetTotpCodeNil

`func (o *VerifyTotpRequest) SetTotpCodeNil(b bool)`

 SetTotpCodeNil sets the value for TotpCode to be an explicit nil

### UnsetTotpCode
`func (o *VerifyTotpRequest) UnsetTotpCode()`

UnsetTotpCode ensures that no value is present for TotpCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


