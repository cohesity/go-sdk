# DataTieringSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment type of the data tiering source. | [optional] 
**GenericNasParams** | Pointer to [**GenericNasDataTieringParams**](GenericNasDataTieringParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonDataTieringParams**](IsilonDataTieringParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappDataTieringParams**](NetappDataTieringParams.md) |  | [optional] 

## Methods

### NewDataTieringSource

`func NewDataTieringSource() *DataTieringSource`

NewDataTieringSource instantiates a new DataTieringSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringSourceWithDefaults

`func NewDataTieringSourceWithDefaults() *DataTieringSource`

NewDataTieringSourceWithDefaults instantiates a new DataTieringSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DataTieringSource) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DataTieringSource) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DataTieringSource) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DataTieringSource) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *DataTieringSource) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *DataTieringSource) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetGenericNasParams

`func (o *DataTieringSource) GetGenericNasParams() GenericNasDataTieringParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *DataTieringSource) GetGenericNasParamsOk() (*GenericNasDataTieringParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *DataTieringSource) SetGenericNasParams(v GenericNasDataTieringParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *DataTieringSource) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *DataTieringSource) GetIsilonParams() IsilonDataTieringParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *DataTieringSource) GetIsilonParamsOk() (*IsilonDataTieringParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *DataTieringSource) SetIsilonParams(v IsilonDataTieringParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *DataTieringSource) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *DataTieringSource) GetNetappParams() NetappDataTieringParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *DataTieringSource) GetNetappParamsOk() (*NetappDataTieringParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *DataTieringSource) SetNetappParams(v NetappDataTieringParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *DataTieringSource) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


