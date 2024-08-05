# MfaParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpCode** | Pointer to **NullableString** | Specifies OTP code for MFA verification. | [optional] 
**OtpType** | Pointer to **NullableString** | Specifies the list of mechanism to receive the OTP code. Supported types are: TOTP (Helios OnPrem Only) -&gt; Time based OTP. Email OTP (Helios OnPrem Only) -&gt; OTP via Email. | [optional] 

## Methods

### NewMfaParams

`func NewMfaParams() *MfaParams`

NewMfaParams instantiates a new MfaParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaParamsWithDefaults

`func NewMfaParamsWithDefaults() *MfaParams`

NewMfaParamsWithDefaults instantiates a new MfaParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpCode

`func (o *MfaParams) GetOtpCode() string`

GetOtpCode returns the OtpCode field if non-nil, zero value otherwise.

### GetOtpCodeOk

`func (o *MfaParams) GetOtpCodeOk() (*string, bool)`

GetOtpCodeOk returns a tuple with the OtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpCode

`func (o *MfaParams) SetOtpCode(v string)`

SetOtpCode sets OtpCode field to given value.

### HasOtpCode

`func (o *MfaParams) HasOtpCode() bool`

HasOtpCode returns a boolean if a field has been set.

### SetOtpCodeNil

`func (o *MfaParams) SetOtpCodeNil(b bool)`

 SetOtpCodeNil sets the value for OtpCode to be an explicit nil

### UnsetOtpCode
`func (o *MfaParams) UnsetOtpCode()`

UnsetOtpCode ensures that no value is present for OtpCode, not even an explicit nil
### GetOtpType

`func (o *MfaParams) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *MfaParams) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *MfaParams) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.

### HasOtpType

`func (o *MfaParams) HasOtpType() bool`

HasOtpType returns a boolean if a field has been set.

### SetOtpTypeNil

`func (o *MfaParams) SetOtpTypeNil(b bool)`

 SetOtpTypeNil sets the value for OtpType to be an explicit nil

### UnsetOtpType
`func (o *MfaParams) UnsetOtpType()`

UnsetOtpType ensures that no value is present for OtpType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


