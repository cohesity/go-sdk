# RecoverIsilonParamsRecoverNasVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElastifileTargetParams** | Pointer to [**NullableRecoverIsilonNasVolumeParamsElastifileTargetParams**](RecoverIsilonNasVolumeParamsElastifileTargetParams.md) |  | [optional] 
**FlashbladeTargetParams** | Pointer to [**NullableRecoverGenericNasNasVolumeParamsFlashbladeTargetParams**](RecoverGenericNasNasVolumeParamsFlashbladeTargetParams.md) |  | [optional] 
**GenericNasTargetParams** | Pointer to [**NullableRecoverGpfsNasVolumeParamsGenericNasTargetParams**](RecoverGpfsNasVolumeParamsGenericNasTargetParams.md) |  | [optional] 
**GpfsTargetParams** | Pointer to [**NullableRecoverIsilonNasVolumeParamsGpfsTargetParams**](RecoverIsilonNasVolumeParamsGpfsTargetParams.md) |  | [optional] 
**IsilonTargetParams** | Pointer to [**NullableRecoverIsilonNasVolumeParamsIsilonTargetParams**](RecoverIsilonNasVolumeParamsIsilonTargetParams.md) |  | [optional] 
**NetappTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsNetappTargetParams**](RecoverElastifileNasVolumeParamsNetappTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**ViewTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsViewTargetParams**](RecoverElastifileNasVolumeParamsViewTargetParams.md) |  | [optional] 

## Methods

### NewRecoverIsilonParamsRecoverNasVolumeParams

`func NewRecoverIsilonParamsRecoverNasVolumeParams(targetEnvironment string, ) *RecoverIsilonParamsRecoverNasVolumeParams`

NewRecoverIsilonParamsRecoverNasVolumeParams instantiates a new RecoverIsilonParamsRecoverNasVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverIsilonParamsRecoverNasVolumeParamsWithDefaults

`func NewRecoverIsilonParamsRecoverNasVolumeParamsWithDefaults() *RecoverIsilonParamsRecoverNasVolumeParams`

NewRecoverIsilonParamsRecoverNasVolumeParamsWithDefaults instantiates a new RecoverIsilonParamsRecoverNasVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetElastifileTargetParams() RecoverIsilonNasVolumeParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverIsilonNasVolumeParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetElastifileTargetParams(v RecoverIsilonNasVolumeParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverIsilonParamsRecoverNasVolumeParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetFlashbladeTargetParams() RecoverGenericNasNasVolumeParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverGenericNasNasVolumeParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetFlashbladeTargetParams(v RecoverGenericNasNasVolumeParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverIsilonParamsRecoverNasVolumeParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetGenericNasTargetParams() RecoverGpfsNasVolumeParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverGpfsNasVolumeParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetGenericNasTargetParams(v RecoverGpfsNasVolumeParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverIsilonParamsRecoverNasVolumeParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetGpfsTargetParams() RecoverIsilonNasVolumeParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverIsilonNasVolumeParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetGpfsTargetParams(v RecoverIsilonNasVolumeParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverIsilonParamsRecoverNasVolumeParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetIsilonTargetParams() RecoverIsilonNasVolumeParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverIsilonNasVolumeParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetIsilonTargetParams(v RecoverIsilonNasVolumeParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverIsilonParamsRecoverNasVolumeParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetNetappTargetParams() RecoverElastifileNasVolumeParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetNetappTargetParamsOk() (*RecoverElastifileNasVolumeParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetNetappTargetParams(v RecoverElastifileNasVolumeParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverIsilonParamsRecoverNasVolumeParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetViewTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetViewTargetParams() RecoverElastifileNasVolumeParamsViewTargetParams`

GetViewTargetParams returns the ViewTargetParams field if non-nil, zero value otherwise.

### GetViewTargetParamsOk

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) GetViewTargetParamsOk() (*RecoverElastifileNasVolumeParamsViewTargetParams, bool)`

GetViewTargetParamsOk returns a tuple with the ViewTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetViewTargetParams(v RecoverElastifileNasVolumeParamsViewTargetParams)`

SetViewTargetParams sets ViewTargetParams field to given value.

### HasViewTargetParams

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) HasViewTargetParams() bool`

HasViewTargetParams returns a boolean if a field has been set.

### SetViewTargetParamsNil

`func (o *RecoverIsilonParamsRecoverNasVolumeParams) SetViewTargetParamsNil(b bool)`

 SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil

### UnsetViewTargetParams
`func (o *RecoverIsilonParamsRecoverNasVolumeParams) UnsetViewTargetParams()`

UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


