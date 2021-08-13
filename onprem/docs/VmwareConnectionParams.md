# VmwareConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableString** | Specifies the VMware Source type. | 
**VcdParams** | Pointer to [**VcdConnectionParams**](VcdConnectionParams.md) |  | [optional] 

## Methods

### NewVmwareConnectionParams

`func NewVmwareConnectionParams(type_ NullableString, ) *VmwareConnectionParams`

NewVmwareConnectionParams instantiates a new VmwareConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareConnectionParamsWithDefaults

`func NewVmwareConnectionParamsWithDefaults() *VmwareConnectionParams`

NewVmwareConnectionParamsWithDefaults instantiates a new VmwareConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VmwareConnectionParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmwareConnectionParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmwareConnectionParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *VmwareConnectionParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VmwareConnectionParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVcdParams

`func (o *VmwareConnectionParams) GetVcdParams() VcdConnectionParams`

GetVcdParams returns the VcdParams field if non-nil, zero value otherwise.

### GetVcdParamsOk

`func (o *VmwareConnectionParams) GetVcdParamsOk() (*VcdConnectionParams, bool)`

GetVcdParamsOk returns a tuple with the VcdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdParams

`func (o *VmwareConnectionParams) SetVcdParams(v VcdConnectionParams)`

SetVcdParams sets VcdParams field to given value.

### HasVcdParams

`func (o *VmwareConnectionParams) HasVcdParams() bool`

HasVcdParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


