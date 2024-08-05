# RecoverNetappNasVolumeParams

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

### NewRecoverNetappNasVolumeParams

`func NewRecoverNetappNasVolumeParams(targetEnvironment string, ) *RecoverNetappNasVolumeParams`

NewRecoverNetappNasVolumeParams instantiates a new RecoverNetappNasVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappNasVolumeParamsWithDefaults

`func NewRecoverNetappNasVolumeParamsWithDefaults() *RecoverNetappNasVolumeParams`

NewRecoverNetappNasVolumeParamsWithDefaults instantiates a new RecoverNetappNasVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverNetappNasVolumeParams) GetElastifileTargetParams() RecoverNetappNasVolumeParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverNetappNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverNetappNasVolumeParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverNetappNasVolumeParams) SetElastifileTargetParams(v RecoverNetappNasVolumeParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverNetappNasVolumeParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverNetappNasVolumeParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverNetappNasVolumeParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverNetappNasVolumeParams) GetFlashbladeTargetParams() RecoverNetappNasVolumeParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverNetappNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverNetappNasVolumeParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverNetappNasVolumeParams) SetFlashbladeTargetParams(v RecoverNetappNasVolumeParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverNetappNasVolumeParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverNetappNasVolumeParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverNetappNasVolumeParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverNetappNasVolumeParams) GetGenericNasTargetParams() RecoverNetappNasVolumeParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverNetappNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverNetappNasVolumeParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverNetappNasVolumeParams) SetGenericNasTargetParams(v RecoverNetappNasVolumeParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverNetappNasVolumeParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverNetappNasVolumeParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverNetappNasVolumeParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverNetappNasVolumeParams) GetGpfsTargetParams() RecoverNetappNasVolumeParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverNetappNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverNetappNasVolumeParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverNetappNasVolumeParams) SetGpfsTargetParams(v RecoverNetappNasVolumeParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverNetappNasVolumeParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverNetappNasVolumeParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverNetappNasVolumeParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsFromSourceInitiatedProtection

`func (o *RecoverNetappNasVolumeParams) GetIsFromSourceInitiatedProtection() bool`

GetIsFromSourceInitiatedProtection returns the IsFromSourceInitiatedProtection field if non-nil, zero value otherwise.

### GetIsFromSourceInitiatedProtectionOk

`func (o *RecoverNetappNasVolumeParams) GetIsFromSourceInitiatedProtectionOk() (*bool, bool)`

GetIsFromSourceInitiatedProtectionOk returns a tuple with the IsFromSourceInitiatedProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSourceInitiatedProtection

`func (o *RecoverNetappNasVolumeParams) SetIsFromSourceInitiatedProtection(v bool)`

SetIsFromSourceInitiatedProtection sets IsFromSourceInitiatedProtection field to given value.

### HasIsFromSourceInitiatedProtection

`func (o *RecoverNetappNasVolumeParams) HasIsFromSourceInitiatedProtection() bool`

HasIsFromSourceInitiatedProtection returns a boolean if a field has been set.

### SetIsFromSourceInitiatedProtectionNil

`func (o *RecoverNetappNasVolumeParams) SetIsFromSourceInitiatedProtectionNil(b bool)`

 SetIsFromSourceInitiatedProtectionNil sets the value for IsFromSourceInitiatedProtection to be an explicit nil

### UnsetIsFromSourceInitiatedProtection
`func (o *RecoverNetappNasVolumeParams) UnsetIsFromSourceInitiatedProtection()`

UnsetIsFromSourceInitiatedProtection ensures that no value is present for IsFromSourceInitiatedProtection, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverNetappNasVolumeParams) GetIsilonTargetParams() RecoverNetappNasVolumeParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverNetappNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverNetappNasVolumeParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverNetappNasVolumeParams) SetIsilonTargetParams(v RecoverNetappNasVolumeParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverNetappNasVolumeParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverNetappNasVolumeParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverNetappNasVolumeParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverNetappNasVolumeParams) GetNetappTargetParams() RecoverNetappNasVolumeParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverNetappNasVolumeParams) GetNetappTargetParamsOk() (*RecoverNetappNasVolumeParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverNetappNasVolumeParams) SetNetappTargetParams(v RecoverNetappNasVolumeParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverNetappNasVolumeParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverNetappNasVolumeParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverNetappNasVolumeParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverNetappNasVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverNetappNasVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverNetappNasVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetViewTargetParams

`func (o *RecoverNetappNasVolumeParams) GetViewTargetParams() RecoverNetappNasVolumeParamsViewTargetParams`

GetViewTargetParams returns the ViewTargetParams field if non-nil, zero value otherwise.

### GetViewTargetParamsOk

`func (o *RecoverNetappNasVolumeParams) GetViewTargetParamsOk() (*RecoverNetappNasVolumeParamsViewTargetParams, bool)`

GetViewTargetParamsOk returns a tuple with the ViewTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTargetParams

`func (o *RecoverNetappNasVolumeParams) SetViewTargetParams(v RecoverNetappNasVolumeParamsViewTargetParams)`

SetViewTargetParams sets ViewTargetParams field to given value.

### HasViewTargetParams

`func (o *RecoverNetappNasVolumeParams) HasViewTargetParams() bool`

HasViewTargetParams returns a boolean if a field has been set.

### SetViewTargetParamsNil

`func (o *RecoverNetappNasVolumeParams) SetViewTargetParamsNil(b bool)`

 SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil

### UnsetViewTargetParams
`func (o *RecoverNetappNasVolumeParams) UnsetViewTargetParams()`

UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


