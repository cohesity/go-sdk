# RecoverVmwareVAppTemplateNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **NullableString** | Specifies the type of VMware source to which the vApp templatess are being restored. | 
**VCloudDirectorParams** | Pointer to [**RecoverVmwareVAppTemplateVCDSourceConfig**](RecoverVmwareVAppTemplateVCDSourceConfig.md) |  | [optional] 

## Methods

### NewRecoverVmwareVAppTemplateNewSourceConfig

`func NewRecoverVmwareVAppTemplateNewSourceConfig(sourceType NullableString, ) *RecoverVmwareVAppTemplateNewSourceConfig`

NewRecoverVmwareVAppTemplateNewSourceConfig instantiates a new RecoverVmwareVAppTemplateNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppTemplateNewSourceConfigWithDefaults

`func NewRecoverVmwareVAppTemplateNewSourceConfigWithDefaults() *RecoverVmwareVAppTemplateNewSourceConfig`

NewRecoverVmwareVAppTemplateNewSourceConfigWithDefaults instantiates a new RecoverVmwareVAppTemplateNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *RecoverVmwareVAppTemplateNewSourceConfig) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetVCloudDirectorParams

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) GetVCloudDirectorParams() RecoverVmwareVAppTemplateVCDSourceConfig`

GetVCloudDirectorParams returns the VCloudDirectorParams field if non-nil, zero value otherwise.

### GetVCloudDirectorParamsOk

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) GetVCloudDirectorParamsOk() (*RecoverVmwareVAppTemplateVCDSourceConfig, bool)`

GetVCloudDirectorParamsOk returns a tuple with the VCloudDirectorParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCloudDirectorParams

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) SetVCloudDirectorParams(v RecoverVmwareVAppTemplateVCDSourceConfig)`

SetVCloudDirectorParams sets VCloudDirectorParams field to given value.

### HasVCloudDirectorParams

`func (o *RecoverVmwareVAppTemplateNewSourceConfig) HasVCloudDirectorParams() bool`

HasVCloudDirectorParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


