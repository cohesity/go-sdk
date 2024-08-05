# SupportTotpKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** | Specifies the TOTP account name to be configured for support user. | [optional] 
**TotpSecretKey** | Pointer to **string** | Specifies the TOTP secret key. | [optional] 
**TotpUri** | Pointer to **string** | Specifies the TOTP key URI for generating MFA QR code. | [optional] 

## Methods

### NewSupportTotpKeyInfo

`func NewSupportTotpKeyInfo() *SupportTotpKeyInfo`

NewSupportTotpKeyInfo instantiates a new SupportTotpKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTotpKeyInfoWithDefaults

`func NewSupportTotpKeyInfoWithDefaults() *SupportTotpKeyInfo`

NewSupportTotpKeyInfoWithDefaults instantiates a new SupportTotpKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *SupportTotpKeyInfo) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SupportTotpKeyInfo) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SupportTotpKeyInfo) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SupportTotpKeyInfo) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetTotpSecretKey

`func (o *SupportTotpKeyInfo) GetTotpSecretKey() string`

GetTotpSecretKey returns the TotpSecretKey field if non-nil, zero value otherwise.

### GetTotpSecretKeyOk

`func (o *SupportTotpKeyInfo) GetTotpSecretKeyOk() (*string, bool)`

GetTotpSecretKeyOk returns a tuple with the TotpSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpSecretKey

`func (o *SupportTotpKeyInfo) SetTotpSecretKey(v string)`

SetTotpSecretKey sets TotpSecretKey field to given value.

### HasTotpSecretKey

`func (o *SupportTotpKeyInfo) HasTotpSecretKey() bool`

HasTotpSecretKey returns a boolean if a field has been set.

### GetTotpUri

`func (o *SupportTotpKeyInfo) GetTotpUri() string`

GetTotpUri returns the TotpUri field if non-nil, zero value otherwise.

### GetTotpUriOk

`func (o *SupportTotpKeyInfo) GetTotpUriOk() (*string, bool)`

GetTotpUriOk returns a tuple with the TotpUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUri

`func (o *SupportTotpKeyInfo) SetTotpUri(v string)`

SetTotpUri sets TotpUri field to given value.

### HasTotpUri

`func (o *SupportTotpKeyInfo) HasTotpUri() bool`

HasTotpUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


