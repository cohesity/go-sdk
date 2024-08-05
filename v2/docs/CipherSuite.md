# CipherSuite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cipher** | Pointer to **string** | Specifies the cipher suite used for TLS handshake. | [optional] 

## Methods

### NewCipherSuite

`func NewCipherSuite() *CipherSuite`

NewCipherSuite instantiates a new CipherSuite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCipherSuiteWithDefaults

`func NewCipherSuiteWithDefaults() *CipherSuite`

NewCipherSuiteWithDefaults instantiates a new CipherSuite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipher

`func (o *CipherSuite) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *CipherSuite) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *CipherSuite) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *CipherSuite) HasCipher() bool`

HasCipher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


