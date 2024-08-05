# RecoverElastifileParamsRecoverNasVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElastifileTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsElastifileTargetParams**](RecoverElastifileNasVolumeParamsElastifileTargetParams.md) |  | [optional] 
**FlashbladeTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsFlashbladeTargetParams**](RecoverElastifileNasVolumeParamsFlashbladeTargetParams.md) |  | [optional] 
**GenericNasTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsGenericNasTargetParams**](RecoverElastifileNasVolumeParamsGenericNasTargetParams.md) |  | [optional] 
**GpfsTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsGpfsTargetParams**](RecoverElastifileNasVolumeParamsGpfsTargetParams.md) |  | [optional] 
**IsilonTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsIsilonTargetParams**](RecoverElastifileNasVolumeParamsIsilonTargetParams.md) |  | [optional] 
**NetappTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsNetappTargetParams**](RecoverElastifileNasVolumeParamsNetappTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**ViewTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsViewTargetParams**](RecoverElastifileNasVolumeParamsViewTargetParams.md) |  | [optional] 

## Methods

### NewRecoverElastifileParamsRecoverNasVolumeParams

`func NewRecoverElastifileParamsRecoverNasVolumeParams(targetEnvironment string, ) *RecoverElastifileParamsRecoverNasVolumeParams`

NewRecoverElastifileParamsRecoverNasVolumeParams instantiates a new RecoverElastifileParamsRecoverNasVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileParamsRecoverNasVolumeParamsWithDefaults

`func NewRecoverElastifileParamsRecoverNasVolumeParamsWithDefaults() *RecoverElastifileParamsRecoverNasVolumeParams`

NewRecoverElastifileParamsRecoverNasVolumeParamsWithDefaults instantiates a new RecoverElastifileParamsRecoverNasVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetElastifileTargetParams() RecoverElastifileNasVolumeParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverElastifileNasVolumeParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetElastifileTargetParams(v RecoverElastifileNasVolumeParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetFlashbladeTargetParams() RecoverElastifileNasVolumeParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverElastifileNasVolumeParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetFlashbladeTargetParams(v RecoverElastifileNasVolumeParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGenericNasTargetParams() RecoverElastifileNasVolumeParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverElastifileNasVolumeParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGenericNasTargetParams(v RecoverElastifileNasVolumeParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGpfsTargetParams() RecoverElastifileNasVolumeParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverElastifileNasVolumeParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGpfsTargetParams(v RecoverElastifileNasVolumeParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetIsilonTargetParams() RecoverElastifileNasVolumeParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverElastifileNasVolumeParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetIsilonTargetParams(v RecoverElastifileNasVolumeParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetNetappTargetParams() RecoverElastifileNasVolumeParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetNetappTargetParamsOk() (*RecoverElastifileNasVolumeParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetNetappTargetParams(v RecoverElastifileNasVolumeParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetViewTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetViewTargetParams() RecoverElastifileNasVolumeParamsViewTargetParams`

GetViewTargetParams returns the ViewTargetParams field if non-nil, zero value otherwise.

### GetViewTargetParamsOk

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetViewTargetParamsOk() (*RecoverElastifileNasVolumeParamsViewTargetParams, bool)`

GetViewTargetParamsOk returns a tuple with the ViewTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetViewTargetParams(v RecoverElastifileNasVolumeParamsViewTargetParams)`

SetViewTargetParams sets ViewTargetParams field to given value.

### HasViewTargetParams

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasViewTargetParams() bool`

HasViewTargetParams returns a boolean if a field has been set.

### SetViewTargetParamsNil

`func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetViewTargetParamsNil(b bool)`

 SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil

### UnsetViewTargetParams
`func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetViewTargetParams()`

UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


