# TotpKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpKeyName** | Pointer to **string** | Specifies the TOTP key name. | [optional] 
**TotpKeyId** | Pointer to **string** | Specifies the TOTP key ID. | [optional] 
**TotpSecretKey** | Pointer to **string** | Specifies the TOTP secret key. | [optional] 
**TotpUri** | Pointer to **string** | Specifies the TOTP key URI for generating MFA QR code. | [optional] 

## Methods

### NewTotpKeyInfo

`func NewTotpKeyInfo() *TotpKeyInfo`

NewTotpKeyInfo instantiates a new TotpKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotpKeyInfoWithDefaults

`func NewTotpKeyInfoWithDefaults() *TotpKeyInfo`

NewTotpKeyInfoWithDefaults instantiates a new TotpKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpKeyName

`func (o *TotpKeyInfo) GetTotpKeyName() string`

GetTotpKeyName returns the TotpKeyName field if non-nil, zero value otherwise.

### GetTotpKeyNameOk

`func (o *TotpKeyInfo) GetTotpKeyNameOk() (*string, bool)`

GetTotpKeyNameOk returns a tuple with the TotpKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpKeyName

`func (o *TotpKeyInfo) SetTotpKeyName(v string)`

SetTotpKeyName sets TotpKeyName field to given value.

### HasTotpKeyName

`func (o *TotpKeyInfo) HasTotpKeyName() bool`

HasTotpKeyName returns a boolean if a field has been set.

### GetTotpKeyId

`func (o *TotpKeyInfo) GetTotpKeyId() string`

GetTotpKeyId returns the TotpKeyId field if non-nil, zero value otherwise.

### GetTotpKeyIdOk

`func (o *TotpKeyInfo) GetTotpKeyIdOk() (*string, bool)`

GetTotpKeyIdOk returns a tuple with the TotpKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpKeyId

`func (o *TotpKeyInfo) SetTotpKeyId(v string)`

SetTotpKeyId sets TotpKeyId field to given value.

### HasTotpKeyId

`func (o *TotpKeyInfo) HasTotpKeyId() bool`

HasTotpKeyId returns a boolean if a field has been set.

### GetTotpSecretKey

`func (o *TotpKeyInfo) GetTotpSecretKey() string`

GetTotpSecretKey returns the TotpSecretKey field if non-nil, zero value otherwise.

### GetTotpSecretKeyOk

`func (o *TotpKeyInfo) GetTotpSecretKeyOk() (*string, bool)`

GetTotpSecretKeyOk returns a tuple with the TotpSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpSecretKey

`func (o *TotpKeyInfo) SetTotpSecretKey(v string)`

SetTotpSecretKey sets TotpSecretKey field to given value.

### HasTotpSecretKey

`func (o *TotpKeyInfo) HasTotpSecretKey() bool`

HasTotpSecretKey returns a boolean if a field has been set.

### GetTotpUri

`func (o *TotpKeyInfo) GetTotpUri() string`

GetTotpUri returns the TotpUri field if non-nil, zero value otherwise.

### GetTotpUriOk

`func (o *TotpKeyInfo) GetTotpUriOk() (*string, bool)`

GetTotpUriOk returns a tuple with the TotpUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUri

`func (o *TotpKeyInfo) SetTotpUri(v string)`

SetTotpUri sets TotpUri field to given value.

### HasTotpUri

`func (o *TotpKeyInfo) HasTotpUri() bool`

HasTotpUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


