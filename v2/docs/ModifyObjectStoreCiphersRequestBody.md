# ModifyObjectStoreCiphersRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphers** | **[]string** | Specifies a list of object store ciphers to enable/disable on the cluster. | 
**Enable** | **NullableBool** | If true, the ciphers passed in will be enabled on the cluster and all other ciphers will be disabled. If false, the ciphers specified will be disabled and all other ciphers on the cluster will be enabled. | 

## Methods

### NewModifyObjectStoreCiphersRequestBody

`func NewModifyObjectStoreCiphersRequestBody(ciphers []string, enable NullableBool, ) *ModifyObjectStoreCiphersRequestBody`

NewModifyObjectStoreCiphersRequestBody instantiates a new ModifyObjectStoreCiphersRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyObjectStoreCiphersRequestBodyWithDefaults

`func NewModifyObjectStoreCiphersRequestBodyWithDefaults() *ModifyObjectStoreCiphersRequestBody`

NewModifyObjectStoreCiphersRequestBodyWithDefaults instantiates a new ModifyObjectStoreCiphersRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphers

`func (o *ModifyObjectStoreCiphersRequestBody) GetCiphers() []string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *ModifyObjectStoreCiphersRequestBody) GetCiphersOk() (*[]string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *ModifyObjectStoreCiphersRequestBody) SetCiphers(v []string)`

SetCiphers sets Ciphers field to given value.


### SetCiphersNil

`func (o *ModifyObjectStoreCiphersRequestBody) SetCiphersNil(b bool)`

 SetCiphersNil sets the value for Ciphers to be an explicit nil

### UnsetCiphers
`func (o *ModifyObjectStoreCiphersRequestBody) UnsetCiphers()`

UnsetCiphers ensures that no value is present for Ciphers, not even an explicit nil
### GetEnable

`func (o *ModifyObjectStoreCiphersRequestBody) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ModifyObjectStoreCiphersRequestBody) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ModifyObjectStoreCiphersRequestBody) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### SetEnableNil

`func (o *ModifyObjectStoreCiphersRequestBody) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *ModifyObjectStoreCiphersRequestBody) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


