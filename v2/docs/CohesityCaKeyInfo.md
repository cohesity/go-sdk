# CohesityCaKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | Certificate chain in pem format | [optional] 
**PrivateKey** | Pointer to **NullableString** | Private key of the Cohesity CA. | [optional] 

## Methods

### NewCohesityCaKeyInfo

`func NewCohesityCaKeyInfo() *CohesityCaKeyInfo`

NewCohesityCaKeyInfo instantiates a new CohesityCaKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCohesityCaKeyInfoWithDefaults

`func NewCohesityCaKeyInfoWithDefaults() *CohesityCaKeyInfo`

NewCohesityCaKeyInfoWithDefaults instantiates a new CohesityCaKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaChain

`func (o *CohesityCaKeyInfo) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *CohesityCaKeyInfo) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *CohesityCaKeyInfo) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.

### HasCaChain

`func (o *CohesityCaKeyInfo) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CohesityCaKeyInfo) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CohesityCaKeyInfo) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CohesityCaKeyInfo) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CohesityCaKeyInfo) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CohesityCaKeyInfo) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CohesityCaKeyInfo) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


