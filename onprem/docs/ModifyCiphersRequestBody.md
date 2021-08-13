# ModifyCiphersRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | If true, the ciphers passed in will be enabled on the cluster and all other ciphers will be disabled. If false, the ciphers specified will be disabled and all other ciphers on the cluster will be enabled. | 
**Ciphers** | **[]string** | Specifies a list of ciphers to enable/disable on the cluster. | 

## Methods

### NewModifyCiphersRequestBody

`func NewModifyCiphersRequestBody(enable bool, ciphers []string, ) *ModifyCiphersRequestBody`

NewModifyCiphersRequestBody instantiates a new ModifyCiphersRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCiphersRequestBodyWithDefaults

`func NewModifyCiphersRequestBodyWithDefaults() *ModifyCiphersRequestBody`

NewModifyCiphersRequestBodyWithDefaults instantiates a new ModifyCiphersRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ModifyCiphersRequestBody) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ModifyCiphersRequestBody) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ModifyCiphersRequestBody) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetCiphers

`func (o *ModifyCiphersRequestBody) GetCiphers() []string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *ModifyCiphersRequestBody) GetCiphersOk() (*[]string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *ModifyCiphersRequestBody) SetCiphers(v []string)`

SetCiphers sets Ciphers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


