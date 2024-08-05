# RecoverElastifileNasVolumeParamsViewTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosPolicy** | Pointer to [**NasQosPolicy**](NasQosPolicy.md) |  | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the view. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverElastifileNasVolumeParamsViewTargetParams

`func NewRecoverElastifileNasVolumeParamsViewTargetParams() *RecoverElastifileNasVolumeParamsViewTargetParams`

NewRecoverElastifileNasVolumeParamsViewTargetParams instantiates a new RecoverElastifileNasVolumeParamsViewTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileNasVolumeParamsViewTargetParamsWithDefaults

`func NewRecoverElastifileNasVolumeParamsViewTargetParamsWithDefaults() *RecoverElastifileNasVolumeParamsViewTargetParams`

NewRecoverElastifileNasVolumeParamsViewTargetParamsWithDefaults instantiates a new RecoverElastifileNasVolumeParamsViewTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosPolicy

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) GetQosPolicy() NasQosPolicy`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) GetQosPolicyOk() (*NasQosPolicy, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) SetQosPolicy(v NasQosPolicy)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### GetViewName

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVlanConfig

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverElastifileNasVolumeParamsViewTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


