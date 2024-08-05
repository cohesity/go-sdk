# PowerOffVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmIds** | Pointer to **[]float32** | Specifies the Entity IDs of the VMs to be powered off. | [optional] 

## Methods

### NewPowerOffVmParams

`func NewPowerOffVmParams() *PowerOffVmParams`

NewPowerOffVmParams instantiates a new PowerOffVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerOffVmParamsWithDefaults

`func NewPowerOffVmParamsWithDefaults() *PowerOffVmParams`

NewPowerOffVmParamsWithDefaults instantiates a new PowerOffVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmIds

`func (o *PowerOffVmParams) GetVmIds() []float32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *PowerOffVmParams) GetVmIdsOk() (*[]float32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *PowerOffVmParams) SetVmIds(v []float32)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *PowerOffVmParams) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIdsNil

`func (o *PowerOffVmParams) SetVmIdsNil(b bool)`

 SetVmIdsNil sets the value for VmIds to be an explicit nil

### UnsetVmIds
`func (o *PowerOffVmParams) UnsetVmIds()`

UnsetVmIds ensures that no value is present for VmIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


