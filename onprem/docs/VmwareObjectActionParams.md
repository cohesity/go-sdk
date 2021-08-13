# VmwareObjectActionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the action on the Object. | 
**EnableAppProtectionParams** | Pointer to [**VmwareObjectEnableAppProtectionParams**](VmwareObjectEnableAppProtectionParams.md) |  | [optional] 

## Methods

### NewVmwareObjectActionParams

`func NewVmwareObjectActionParams(action NullableString, ) *VmwareObjectActionParams`

NewVmwareObjectActionParams instantiates a new VmwareObjectActionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectActionParamsWithDefaults

`func NewVmwareObjectActionParamsWithDefaults() *VmwareObjectActionParams`

NewVmwareObjectActionParamsWithDefaults instantiates a new VmwareObjectActionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VmwareObjectActionParams) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmwareObjectActionParams) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmwareObjectActionParams) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *VmwareObjectActionParams) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *VmwareObjectActionParams) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetEnableAppProtectionParams

`func (o *VmwareObjectActionParams) GetEnableAppProtectionParams() VmwareObjectEnableAppProtectionParams`

GetEnableAppProtectionParams returns the EnableAppProtectionParams field if non-nil, zero value otherwise.

### GetEnableAppProtectionParamsOk

`func (o *VmwareObjectActionParams) GetEnableAppProtectionParamsOk() (*VmwareObjectEnableAppProtectionParams, bool)`

GetEnableAppProtectionParamsOk returns a tuple with the EnableAppProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAppProtectionParams

`func (o *VmwareObjectActionParams) SetEnableAppProtectionParams(v VmwareObjectEnableAppProtectionParams)`

SetEnableAppProtectionParams sets EnableAppProtectionParams field to given value.

### HasEnableAppProtectionParams

`func (o *VmwareObjectActionParams) HasEnableAppProtectionParams() bool`

HasEnableAppProtectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


