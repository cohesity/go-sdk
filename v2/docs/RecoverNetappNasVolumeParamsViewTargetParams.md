# RecoverNetappNasVolumeParamsViewTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosPolicy** | Pointer to [**NasQosPolicy**](NasQosPolicy.md) |  | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the view. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverNetappNasVolumeParamsViewTargetParams

`func NewRecoverNetappNasVolumeParamsViewTargetParams() *RecoverNetappNasVolumeParamsViewTargetParams`

NewRecoverNetappNasVolumeParamsViewTargetParams instantiates a new RecoverNetappNasVolumeParamsViewTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappNasVolumeParamsViewTargetParamsWithDefaults

`func NewRecoverNetappNasVolumeParamsViewTargetParamsWithDefaults() *RecoverNetappNasVolumeParamsViewTargetParams`

NewRecoverNetappNasVolumeParamsViewTargetParamsWithDefaults instantiates a new RecoverNetappNasVolumeParamsViewTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosPolicy

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) GetQosPolicy() NasQosPolicy`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) GetQosPolicyOk() (*NasQosPolicy, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) SetQosPolicy(v NasQosPolicy)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### GetViewName

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RecoverNetappNasVolumeParamsViewTargetParams) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVlanConfig

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverNetappNasVolumeParamsViewTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


