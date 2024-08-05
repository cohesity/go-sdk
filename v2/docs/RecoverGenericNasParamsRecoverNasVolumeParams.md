# RecoverGenericNasParamsRecoverNasVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElastifileTargetParams** | Pointer to [**NullableRecoverFlashbladeNasVolumeParamsElastifileTargetParams**](RecoverFlashbladeNasVolumeParamsElastifileTargetParams.md) |  | [optional] 
**FlashbladeTargetParams** | Pointer to [**NullableRecoverGenericNasNasVolumeParamsFlashbladeTargetParams**](RecoverGenericNasNasVolumeParamsFlashbladeTargetParams.md) |  | [optional] 
**GenericNasTargetParams** | Pointer to [**NullableRecoverGenericNasNasVolumeParamsGenericNasTargetParams**](RecoverGenericNasNasVolumeParamsGenericNasTargetParams.md) |  | [optional] 
**GpfsTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsGpfsTargetParams**](RecoverElastifileNasVolumeParamsGpfsTargetParams.md) |  | [optional] 
**IsilonTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsIsilonTargetParams**](RecoverElastifileNasVolumeParamsIsilonTargetParams.md) |  | [optional] 
**NetappTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsNetappTargetParams**](RecoverElastifileNasVolumeParamsNetappTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**ViewTargetParams** | Pointer to [**NullableRecoverElastifileNasVolumeParamsViewTargetParams**](RecoverElastifileNasVolumeParamsViewTargetParams.md) |  | [optional] 

## Methods

### NewRecoverGenericNasParamsRecoverNasVolumeParams

`func NewRecoverGenericNasParamsRecoverNasVolumeParams(targetEnvironment string, ) *RecoverGenericNasParamsRecoverNasVolumeParams`

NewRecoverGenericNasParamsRecoverNasVolumeParams instantiates a new RecoverGenericNasParamsRecoverNasVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGenericNasParamsRecoverNasVolumeParamsWithDefaults

`func NewRecoverGenericNasParamsRecoverNasVolumeParamsWithDefaults() *RecoverGenericNasParamsRecoverNasVolumeParams`

NewRecoverGenericNasParamsRecoverNasVolumeParamsWithDefaults instantiates a new RecoverGenericNasParamsRecoverNasVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetElastifileTargetParams() RecoverFlashbladeNasVolumeParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverFlashbladeNasVolumeParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetElastifileTargetParams(v RecoverFlashbladeNasVolumeParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetFlashbladeTargetParams() RecoverGenericNasNasVolumeParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverGenericNasNasVolumeParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetFlashbladeTargetParams(v RecoverGenericNasNasVolumeParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetGenericNasTargetParams() RecoverGenericNasNasVolumeParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverGenericNasNasVolumeParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetGenericNasTargetParams(v RecoverGenericNasNasVolumeParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetGpfsTargetParams() RecoverElastifileNasVolumeParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverElastifileNasVolumeParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetGpfsTargetParams(v RecoverElastifileNasVolumeParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetIsilonTargetParams() RecoverElastifileNasVolumeParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverElastifileNasVolumeParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetIsilonTargetParams(v RecoverElastifileNasVolumeParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetNetappTargetParams() RecoverElastifileNasVolumeParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetNetappTargetParamsOk() (*RecoverElastifileNasVolumeParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetNetappTargetParams(v RecoverElastifileNasVolumeParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetViewTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetViewTargetParams() RecoverElastifileNasVolumeParamsViewTargetParams`

GetViewTargetParams returns the ViewTargetParams field if non-nil, zero value otherwise.

### GetViewTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) GetViewTargetParamsOk() (*RecoverElastifileNasVolumeParamsViewTargetParams, bool)`

GetViewTargetParamsOk returns a tuple with the ViewTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetViewTargetParams(v RecoverElastifileNasVolumeParamsViewTargetParams)`

SetViewTargetParams sets ViewTargetParams field to given value.

### HasViewTargetParams

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) HasViewTargetParams() bool`

HasViewTargetParams returns a boolean if a field has been set.

### SetViewTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) SetViewTargetParamsNil(b bool)`

 SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil

### UnsetViewTargetParams
`func (o *RecoverGenericNasParamsRecoverNasVolumeParams) UnsetViewTargetParams()`

UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


