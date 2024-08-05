# GenerateCsrResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrPem** | Pointer to **NullableString** | Specifies csr in pem format. | [optional] 
**PrivateKey** | Pointer to **NullableString** | Actual private key of the certificate. | [optional] 

## Methods

### NewGenerateCsrResponse

`func NewGenerateCsrResponse() *GenerateCsrResponse`

NewGenerateCsrResponse instantiates a new GenerateCsrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCsrResponseWithDefaults

`func NewGenerateCsrResponseWithDefaults() *GenerateCsrResponse`

NewGenerateCsrResponseWithDefaults instantiates a new GenerateCsrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrPem

`func (o *GenerateCsrResponse) GetCsrPem() string`

GetCsrPem returns the CsrPem field if non-nil, zero value otherwise.

### GetCsrPemOk

`func (o *GenerateCsrResponse) GetCsrPemOk() (*string, bool)`

GetCsrPemOk returns a tuple with the CsrPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrPem

`func (o *GenerateCsrResponse) SetCsrPem(v string)`

SetCsrPem sets CsrPem field to given value.

### HasCsrPem

`func (o *GenerateCsrResponse) HasCsrPem() bool`

HasCsrPem returns a boolean if a field has been set.

### SetCsrPemNil

`func (o *GenerateCsrResponse) SetCsrPemNil(b bool)`

 SetCsrPemNil sets the value for CsrPem to be an explicit nil

### UnsetCsrPem
`func (o *GenerateCsrResponse) UnsetCsrPem()`

UnsetCsrPem ensures that no value is present for CsrPem, not even an explicit nil
### GetPrivateKey

`func (o *GenerateCsrResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GenerateCsrResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GenerateCsrResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GenerateCsrResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *GenerateCsrResponse) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *GenerateCsrResponse) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


