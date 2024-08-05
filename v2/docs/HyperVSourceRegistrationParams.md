# HyperVSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScvmmParams** | Pointer to [**ScvmmRegistrationParams**](ScvmmRegistrationParams.md) |  | [optional] 
**StandaloneClusterParams** | Pointer to [**StandaloneClusterRegistrationParams**](StandaloneClusterRegistrationParams.md) |  | [optional] 
**StandaloneHostParams** | Pointer to [**StandaloneHostRegistrationParams**](StandaloneHostRegistrationParams.md) |  | [optional] 
**Type** | **NullableString** | Specifies the HyperV Source type. | 

## Methods

### NewHyperVSourceRegistrationParams

`func NewHyperVSourceRegistrationParams(type_ NullableString, ) *HyperVSourceRegistrationParams`

NewHyperVSourceRegistrationParams instantiates a new HyperVSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVSourceRegistrationParamsWithDefaults

`func NewHyperVSourceRegistrationParamsWithDefaults() *HyperVSourceRegistrationParams`

NewHyperVSourceRegistrationParamsWithDefaults instantiates a new HyperVSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScvmmParams

`func (o *HyperVSourceRegistrationParams) GetScvmmParams() ScvmmRegistrationParams`

GetScvmmParams returns the ScvmmParams field if non-nil, zero value otherwise.

### GetScvmmParamsOk

`func (o *HyperVSourceRegistrationParams) GetScvmmParamsOk() (*ScvmmRegistrationParams, bool)`

GetScvmmParamsOk returns a tuple with the ScvmmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScvmmParams

`func (o *HyperVSourceRegistrationParams) SetScvmmParams(v ScvmmRegistrationParams)`

SetScvmmParams sets ScvmmParams field to given value.

### HasScvmmParams

`func (o *HyperVSourceRegistrationParams) HasScvmmParams() bool`

HasScvmmParams returns a boolean if a field has been set.

### GetStandaloneClusterParams

`func (o *HyperVSourceRegistrationParams) GetStandaloneClusterParams() StandaloneClusterRegistrationParams`

GetStandaloneClusterParams returns the StandaloneClusterParams field if non-nil, zero value otherwise.

### GetStandaloneClusterParamsOk

`func (o *HyperVSourceRegistrationParams) GetStandaloneClusterParamsOk() (*StandaloneClusterRegistrationParams, bool)`

GetStandaloneClusterParamsOk returns a tuple with the StandaloneClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneClusterParams

`func (o *HyperVSourceRegistrationParams) SetStandaloneClusterParams(v StandaloneClusterRegistrationParams)`

SetStandaloneClusterParams sets StandaloneClusterParams field to given value.

### HasStandaloneClusterParams

`func (o *HyperVSourceRegistrationParams) HasStandaloneClusterParams() bool`

HasStandaloneClusterParams returns a boolean if a field has been set.

### GetStandaloneHostParams

`func (o *HyperVSourceRegistrationParams) GetStandaloneHostParams() StandaloneHostRegistrationParams`

GetStandaloneHostParams returns the StandaloneHostParams field if non-nil, zero value otherwise.

### GetStandaloneHostParamsOk

`func (o *HyperVSourceRegistrationParams) GetStandaloneHostParamsOk() (*StandaloneHostRegistrationParams, bool)`

GetStandaloneHostParamsOk returns a tuple with the StandaloneHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneHostParams

`func (o *HyperVSourceRegistrationParams) SetStandaloneHostParams(v StandaloneHostRegistrationParams)`

SetStandaloneHostParams sets StandaloneHostParams field to given value.

### HasStandaloneHostParams

`func (o *HyperVSourceRegistrationParams) HasStandaloneHostParams() bool`

HasStandaloneHostParams returns a boolean if a field has been set.

### GetType

`func (o *HyperVSourceRegistrationParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperVSourceRegistrationParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperVSourceRegistrationParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *HyperVSourceRegistrationParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HyperVSourceRegistrationParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


