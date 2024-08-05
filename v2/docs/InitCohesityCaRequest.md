# InitCohesityCaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | **[]string** | Certificate chain in pem format | 
**PrivateKey** | **string** | Private key (RSA 4096). | 

## Methods

### NewInitCohesityCaRequest

`func NewInitCohesityCaRequest(caChain []string, privateKey string, ) *InitCohesityCaRequest`

NewInitCohesityCaRequest instantiates a new InitCohesityCaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitCohesityCaRequestWithDefaults

`func NewInitCohesityCaRequestWithDefaults() *InitCohesityCaRequest`

NewInitCohesityCaRequestWithDefaults instantiates a new InitCohesityCaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaChain

`func (o *InitCohesityCaRequest) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *InitCohesityCaRequest) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *InitCohesityCaRequest) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### GetPrivateKey

`func (o *InitCohesityCaRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *InitCohesityCaRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *InitCohesityCaRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


