# VmwareObjectEntityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdpInfo** | Pointer to [**VmwareCdpObject**](VmwareCdpObject.md) |  | [optional] 
**IsTemplate** | Pointer to **NullableBool** | Specifies if the object is a VM template. | [optional] 
**MoRef** | Pointer to [**MOref**](MOref.md) |  | [optional] 
**Type** | Pointer to **NullableString** | VMware Object type. | [optional] 

## Methods

### NewVmwareObjectEntityParams

`func NewVmwareObjectEntityParams() *VmwareObjectEntityParams`

NewVmwareObjectEntityParams instantiates a new VmwareObjectEntityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectEntityParamsWithDefaults

`func NewVmwareObjectEntityParamsWithDefaults() *VmwareObjectEntityParams`

NewVmwareObjectEntityParamsWithDefaults instantiates a new VmwareObjectEntityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdpInfo

`func (o *VmwareObjectEntityParams) GetCdpInfo() VmwareCdpObject`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *VmwareObjectEntityParams) GetCdpInfoOk() (*VmwareCdpObject, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *VmwareObjectEntityParams) SetCdpInfo(v VmwareCdpObject)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *VmwareObjectEntityParams) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.

### GetIsTemplate

`func (o *VmwareObjectEntityParams) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *VmwareObjectEntityParams) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *VmwareObjectEntityParams) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *VmwareObjectEntityParams) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### SetIsTemplateNil

`func (o *VmwareObjectEntityParams) SetIsTemplateNil(b bool)`

 SetIsTemplateNil sets the value for IsTemplate to be an explicit nil

### UnsetIsTemplate
`func (o *VmwareObjectEntityParams) UnsetIsTemplate()`

UnsetIsTemplate ensures that no value is present for IsTemplate, not even an explicit nil
### GetMoRef

`func (o *VmwareObjectEntityParams) GetMoRef() MOref`

GetMoRef returns the MoRef field if non-nil, zero value otherwise.

### GetMoRefOk

`func (o *VmwareObjectEntityParams) GetMoRefOk() (*MOref, bool)`

GetMoRefOk returns a tuple with the MoRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoRef

`func (o *VmwareObjectEntityParams) SetMoRef(v MOref)`

SetMoRef sets MoRef field to given value.

### HasMoRef

`func (o *VmwareObjectEntityParams) HasMoRef() bool`

HasMoRef returns a boolean if a field has been set.

### GetType

`func (o *VmwareObjectEntityParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmwareObjectEntityParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmwareObjectEntityParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VmwareObjectEntityParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VmwareObjectEntityParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VmwareObjectEntityParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


