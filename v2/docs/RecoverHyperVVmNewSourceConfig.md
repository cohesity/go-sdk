# RecoverHyperVVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScvmmServerParams** | Pointer to [**RecoverHyperVVmSCVMMSourceConfig**](RecoverHyperVVmSCVMMSourceConfig.md) |  | [optional] 
**SourceType** | **NullableString** | Specifies the type of HyperV source to which the VMs are being restored. | 
**StandaloneClusterParams** | Pointer to [**RecoverHyperVVmStandaloneClusterSourceConfig**](RecoverHyperVVmStandaloneClusterSourceConfig.md) |  | [optional] 
**StandaloneHostParams** | Pointer to [**RecoverHyperVVmStandaloneHostSourceConfig**](RecoverHyperVVmStandaloneHostSourceConfig.md) |  | [optional] 

## Methods

### NewRecoverHyperVVmNewSourceConfig

`func NewRecoverHyperVVmNewSourceConfig(sourceType NullableString, ) *RecoverHyperVVmNewSourceConfig`

NewRecoverHyperVVmNewSourceConfig instantiates a new RecoverHyperVVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmNewSourceConfigWithDefaults

`func NewRecoverHyperVVmNewSourceConfigWithDefaults() *RecoverHyperVVmNewSourceConfig`

NewRecoverHyperVVmNewSourceConfigWithDefaults instantiates a new RecoverHyperVVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScvmmServerParams

`func (o *RecoverHyperVVmNewSourceConfig) GetScvmmServerParams() RecoverHyperVVmSCVMMSourceConfig`

GetScvmmServerParams returns the ScvmmServerParams field if non-nil, zero value otherwise.

### GetScvmmServerParamsOk

`func (o *RecoverHyperVVmNewSourceConfig) GetScvmmServerParamsOk() (*RecoverHyperVVmSCVMMSourceConfig, bool)`

GetScvmmServerParamsOk returns a tuple with the ScvmmServerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScvmmServerParams

`func (o *RecoverHyperVVmNewSourceConfig) SetScvmmServerParams(v RecoverHyperVVmSCVMMSourceConfig)`

SetScvmmServerParams sets ScvmmServerParams field to given value.

### HasScvmmServerParams

`func (o *RecoverHyperVVmNewSourceConfig) HasScvmmServerParams() bool`

HasScvmmServerParams returns a boolean if a field has been set.

### GetSourceType

`func (o *RecoverHyperVVmNewSourceConfig) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RecoverHyperVVmNewSourceConfig) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RecoverHyperVVmNewSourceConfig) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *RecoverHyperVVmNewSourceConfig) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *RecoverHyperVVmNewSourceConfig) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetStandaloneClusterParams

`func (o *RecoverHyperVVmNewSourceConfig) GetStandaloneClusterParams() RecoverHyperVVmStandaloneClusterSourceConfig`

GetStandaloneClusterParams returns the StandaloneClusterParams field if non-nil, zero value otherwise.

### GetStandaloneClusterParamsOk

`func (o *RecoverHyperVVmNewSourceConfig) GetStandaloneClusterParamsOk() (*RecoverHyperVVmStandaloneClusterSourceConfig, bool)`

GetStandaloneClusterParamsOk returns a tuple with the StandaloneClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneClusterParams

`func (o *RecoverHyperVVmNewSourceConfig) SetStandaloneClusterParams(v RecoverHyperVVmStandaloneClusterSourceConfig)`

SetStandaloneClusterParams sets StandaloneClusterParams field to given value.

### HasStandaloneClusterParams

`func (o *RecoverHyperVVmNewSourceConfig) HasStandaloneClusterParams() bool`

HasStandaloneClusterParams returns a boolean if a field has been set.

### GetStandaloneHostParams

`func (o *RecoverHyperVVmNewSourceConfig) GetStandaloneHostParams() RecoverHyperVVmStandaloneHostSourceConfig`

GetStandaloneHostParams returns the StandaloneHostParams field if non-nil, zero value otherwise.

### GetStandaloneHostParamsOk

`func (o *RecoverHyperVVmNewSourceConfig) GetStandaloneHostParamsOk() (*RecoverHyperVVmStandaloneHostSourceConfig, bool)`

GetStandaloneHostParamsOk returns a tuple with the StandaloneHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneHostParams

`func (o *RecoverHyperVVmNewSourceConfig) SetStandaloneHostParams(v RecoverHyperVVmStandaloneHostSourceConfig)`

SetStandaloneHostParams sets StandaloneHostParams field to given value.

### HasStandaloneHostParams

`func (o *RecoverHyperVVmNewSourceConfig) HasStandaloneHostParams() bool`

HasStandaloneHostParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


