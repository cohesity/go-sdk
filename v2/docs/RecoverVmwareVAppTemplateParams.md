# RecoverVmwareVAppTemplateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VmwareTargetParams** | Pointer to [**NullableRecoverVmwareVAppTemplateParamsVmwareTargetParams**](RecoverVmwareVAppTemplateParamsVmwareTargetParams.md) |  | [optional] 

## Methods

### NewRecoverVmwareVAppTemplateParams

`func NewRecoverVmwareVAppTemplateParams(targetEnvironment string, ) *RecoverVmwareVAppTemplateParams`

NewRecoverVmwareVAppTemplateParams instantiates a new RecoverVmwareVAppTemplateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppTemplateParamsWithDefaults

`func NewRecoverVmwareVAppTemplateParamsWithDefaults() *RecoverVmwareVAppTemplateParams`

NewRecoverVmwareVAppTemplateParamsWithDefaults instantiates a new RecoverVmwareVAppTemplateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *RecoverVmwareVAppTemplateParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverVmwareVAppTemplateParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverVmwareVAppTemplateParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVmwareTargetParams

`func (o *RecoverVmwareVAppTemplateParams) GetVmwareTargetParams() RecoverVmwareVAppTemplateParamsVmwareTargetParams`

GetVmwareTargetParams returns the VmwareTargetParams field if non-nil, zero value otherwise.

### GetVmwareTargetParamsOk

`func (o *RecoverVmwareVAppTemplateParams) GetVmwareTargetParamsOk() (*RecoverVmwareVAppTemplateParamsVmwareTargetParams, bool)`

GetVmwareTargetParamsOk returns a tuple with the VmwareTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareTargetParams

`func (o *RecoverVmwareVAppTemplateParams) SetVmwareTargetParams(v RecoverVmwareVAppTemplateParamsVmwareTargetParams)`

SetVmwareTargetParams sets VmwareTargetParams field to given value.

### HasVmwareTargetParams

`func (o *RecoverVmwareVAppTemplateParams) HasVmwareTargetParams() bool`

HasVmwareTargetParams returns a boolean if a field has been set.

### SetVmwareTargetParamsNil

`func (o *RecoverVmwareVAppTemplateParams) SetVmwareTargetParamsNil(b bool)`

 SetVmwareTargetParamsNil sets the value for VmwareTargetParams to be an explicit nil

### UnsetVmwareTargetParams
`func (o *RecoverVmwareVAppTemplateParams) UnsetVmwareTargetParams()`

UnsetVmwareTargetParams ensures that no value is present for VmwareTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


