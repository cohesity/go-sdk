# VmwareSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsxiParams** | Pointer to [**EsxiRegistrationParams**](EsxiRegistrationParams.md) |  | [optional] 
**Type** | **NullableString** | Specifies the VMware Source type. | 
**VCenterParams** | Pointer to [**VcenterRegistrationParams**](VcenterRegistrationParams.md) |  | [optional] 
**VcdParams** | Pointer to [**VcdRegistrationParams**](VcdRegistrationParams.md) |  | [optional] 

## Methods

### NewVmwareSourceRegistrationParams

`func NewVmwareSourceRegistrationParams(type_ NullableString, ) *VmwareSourceRegistrationParams`

NewVmwareSourceRegistrationParams instantiates a new VmwareSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareSourceRegistrationParamsWithDefaults

`func NewVmwareSourceRegistrationParamsWithDefaults() *VmwareSourceRegistrationParams`

NewVmwareSourceRegistrationParamsWithDefaults instantiates a new VmwareSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsxiParams

`func (o *VmwareSourceRegistrationParams) GetEsxiParams() EsxiRegistrationParams`

GetEsxiParams returns the EsxiParams field if non-nil, zero value otherwise.

### GetEsxiParamsOk

`func (o *VmwareSourceRegistrationParams) GetEsxiParamsOk() (*EsxiRegistrationParams, bool)`

GetEsxiParamsOk returns a tuple with the EsxiParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxiParams

`func (o *VmwareSourceRegistrationParams) SetEsxiParams(v EsxiRegistrationParams)`

SetEsxiParams sets EsxiParams field to given value.

### HasEsxiParams

`func (o *VmwareSourceRegistrationParams) HasEsxiParams() bool`

HasEsxiParams returns a boolean if a field has been set.

### GetType

`func (o *VmwareSourceRegistrationParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmwareSourceRegistrationParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmwareSourceRegistrationParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *VmwareSourceRegistrationParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VmwareSourceRegistrationParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVCenterParams

`func (o *VmwareSourceRegistrationParams) GetVCenterParams() VcenterRegistrationParams`

GetVCenterParams returns the VCenterParams field if non-nil, zero value otherwise.

### GetVCenterParamsOk

`func (o *VmwareSourceRegistrationParams) GetVCenterParamsOk() (*VcenterRegistrationParams, bool)`

GetVCenterParamsOk returns a tuple with the VCenterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterParams

`func (o *VmwareSourceRegistrationParams) SetVCenterParams(v VcenterRegistrationParams)`

SetVCenterParams sets VCenterParams field to given value.

### HasVCenterParams

`func (o *VmwareSourceRegistrationParams) HasVCenterParams() bool`

HasVCenterParams returns a boolean if a field has been set.

### GetVcdParams

`func (o *VmwareSourceRegistrationParams) GetVcdParams() VcdRegistrationParams`

GetVcdParams returns the VcdParams field if non-nil, zero value otherwise.

### GetVcdParamsOk

`func (o *VmwareSourceRegistrationParams) GetVcdParamsOk() (*VcdRegistrationParams, bool)`

GetVcdParamsOk returns a tuple with the VcdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdParams

`func (o *VmwareSourceRegistrationParams) SetVcdParams(v VcdRegistrationParams)`

SetVcdParams sets VcdParams field to given value.

### HasVcdParams

`func (o *VmwareSourceRegistrationParams) HasVcdParams() bool`

HasVcdParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


