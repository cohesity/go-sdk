# OnPremDeployTargetResultVmwareParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Moref** | Pointer to [**MOref**](MOref.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** | UUID of the recovered VMware VM | [optional] 

## Methods

### NewOnPremDeployTargetResultVmwareParams

`func NewOnPremDeployTargetResultVmwareParams() *OnPremDeployTargetResultVmwareParams`

NewOnPremDeployTargetResultVmwareParams instantiates a new OnPremDeployTargetResultVmwareParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremDeployTargetResultVmwareParamsWithDefaults

`func NewOnPremDeployTargetResultVmwareParamsWithDefaults() *OnPremDeployTargetResultVmwareParams`

NewOnPremDeployTargetResultVmwareParamsWithDefaults instantiates a new OnPremDeployTargetResultVmwareParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoref

`func (o *OnPremDeployTargetResultVmwareParams) GetMoref() MOref`

GetMoref returns the Moref field if non-nil, zero value otherwise.

### GetMorefOk

`func (o *OnPremDeployTargetResultVmwareParams) GetMorefOk() (*MOref, bool)`

GetMorefOk returns a tuple with the Moref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoref

`func (o *OnPremDeployTargetResultVmwareParams) SetMoref(v MOref)`

SetMoref sets Moref field to given value.

### HasMoref

`func (o *OnPremDeployTargetResultVmwareParams) HasMoref() bool`

HasMoref returns a boolean if a field has been set.

### GetUuid

`func (o *OnPremDeployTargetResultVmwareParams) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OnPremDeployTargetResultVmwareParams) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OnPremDeployTargetResultVmwareParams) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OnPremDeployTargetResultVmwareParams) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *OnPremDeployTargetResultVmwareParams) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *OnPremDeployTargetResultVmwareParams) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


