# RecoverNetappParamsRecoverNasVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElastifileTargetParams** | Pointer to [**NullableRecoverNetappNasVolumeParamsElastifileTargetParams**](RecoverNetappNasVolumeParamsElastifileTargetParams.md) |  | [optional] 
**FlashbladeTargetParams** | Pointer to [**NullableRecoverNetappNasVolumeParamsFlashbladeTargetParams**](RecoverNetappNasVolumeParamsFlashbladeTargetParams.md) |  | [optional] 
**GenericNasTargetParams** | Pointer to [**NullableRecoverNetappNasVolumeParamsGenericNasTargetParams**](RecoverNetappNasVolumeParamsGenericNasTargetParams.md) |  | [optional] 
**GpfsTargetParams** | Pointer to [**NullableRecoverNetappNasVolumeParamsGpfsTargetParams**](RecoverNetappNasVolumeParamsGpfsTargetParams.md) |  | [optional] 
**IsFromSourceInitiatedProtection** | Pointer to **NullableBool** | Specifies if the snapshot trying to recover is from a source initiated protection. | [optional] 
**IsilonTargetParams** | Pointer to [**NullableRecoverNetappNasVolumeParamsIsilonTargetParams**](RecoverNetappNasVolumeParamsIsilonTargetParams.md) |  | [optional] 
**NetappTargetParams** | Pointer to [**NullableRecoverNetappNasVolumeParamsNetappTargetParams**](RecoverNetappNasVolumeParamsNetappTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**ViewTargetParams** | Pointer to [**NullableRecoverNetappNasVolumeParamsViewTargetParams**](RecoverNetappNasVolumeParamsViewTargetParams.md) |  | [optional] 

## Methods

### NewRecoverNetappParamsRecoverNasVolumeParams

`func NewRecoverNetappParamsRecoverNasVolumeParams(targetEnvironment string, ) *RecoverNetappParamsRecoverNasVolumeParams`

NewRecoverNetappParamsRecoverNasVolumeParams instantiates a new RecoverNetappParamsRecoverNasVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappParamsRecoverNasVolumeParamsWithDefaults

`func NewRecoverNetappParamsRecoverNasVolumeParamsWithDefaults() *RecoverNetappParamsRecoverNasVolumeParams`

NewRecoverNetappParamsRecoverNasVolumeParamsWithDefaults instantiates a new RecoverNetappParamsRecoverNasVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetElastifileTargetParams() RecoverNetappNasVolumeParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverNetappNasVolumeParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetElastifileTargetParams(v RecoverNetappNasVolumeParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetFlashbladeTargetParams() RecoverNetappNasVolumeParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverNetappNasVolumeParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetFlashbladeTargetParams(v RecoverNetappNasVolumeParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetGenericNasTargetParams() RecoverNetappNasVolumeParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverNetappNasVolumeParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetGenericNasTargetParams(v RecoverNetappNasVolumeParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetGpfsTargetParams() RecoverNetappNasVolumeParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverNetappNasVolumeParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetGpfsTargetParams(v RecoverNetappNasVolumeParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsFromSourceInitiatedProtection

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetIsFromSourceInitiatedProtection() bool`

GetIsFromSourceInitiatedProtection returns the IsFromSourceInitiatedProtection field if non-nil, zero value otherwise.

### GetIsFromSourceInitiatedProtectionOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetIsFromSourceInitiatedProtectionOk() (*bool, bool)`

GetIsFromSourceInitiatedProtectionOk returns a tuple with the IsFromSourceInitiatedProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSourceInitiatedProtection

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetIsFromSourceInitiatedProtection(v bool)`

SetIsFromSourceInitiatedProtection sets IsFromSourceInitiatedProtection field to given value.

### HasIsFromSourceInitiatedProtection

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasIsFromSourceInitiatedProtection() bool`

HasIsFromSourceInitiatedProtection returns a boolean if a field has been set.

### SetIsFromSourceInitiatedProtectionNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetIsFromSourceInitiatedProtectionNil(b bool)`

 SetIsFromSourceInitiatedProtectionNil sets the value for IsFromSourceInitiatedProtection to be an explicit nil

### UnsetIsFromSourceInitiatedProtection
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetIsFromSourceInitiatedProtection()`

UnsetIsFromSourceInitiatedProtection ensures that no value is present for IsFromSourceInitiatedProtection, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetIsilonTargetParams() RecoverNetappNasVolumeParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverNetappNasVolumeParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetIsilonTargetParams(v RecoverNetappNasVolumeParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetNetappTargetParams() RecoverNetappNasVolumeParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetNetappTargetParamsOk() (*RecoverNetappNasVolumeParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetNetappTargetParams(v RecoverNetappNasVolumeParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetViewTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetViewTargetParams() RecoverNetappNasVolumeParamsViewTargetParams`

GetViewTargetParams returns the ViewTargetParams field if non-nil, zero value otherwise.

### GetViewTargetParamsOk

`func (o *RecoverNetappParamsRecoverNasVolumeParams) GetViewTargetParamsOk() (*RecoverNetappNasVolumeParamsViewTargetParams, bool)`

GetViewTargetParamsOk returns a tuple with the ViewTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetViewTargetParams(v RecoverNetappNasVolumeParamsViewTargetParams)`

SetViewTargetParams sets ViewTargetParams field to given value.

### HasViewTargetParams

`func (o *RecoverNetappParamsRecoverNasVolumeParams) HasViewTargetParams() bool`

HasViewTargetParams returns a boolean if a field has been set.

### SetViewTargetParamsNil

`func (o *RecoverNetappParamsRecoverNasVolumeParams) SetViewTargetParamsNil(b bool)`

 SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil

### UnsetViewTargetParams
`func (o *RecoverNetappParamsRecoverNasVolumeParams) UnsetViewTargetParams()`

UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


