# VmwareCdpObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowReEnableCdp** | Pointer to **NullableBool** | Specifies if re-enabling CDP is allowed or not through UI without any job or policy update through API. | [optional] 
**CdpEnabled** | Pointer to **NullableBool** | Specifies whether CDP is currently active or not. CDP might have been active on this object before, but it might not be anymore. | [optional] 
**LastRunInfo** | Pointer to [**CdpObjectLastRunInfo**](CdpObjectLastRunInfo.md) |  | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id to which this CDP object belongs. | [optional] [readonly] 
**GuardrailsErrorMessage** | Pointer to **NullableString** | Specifies the error message from the guardrails info from cdp state if any. | [optional] 
**IoFilterErrorMessage** | Pointer to **NullableString** | Specifies the error message related to IO filter if there is any. | [optional] 
**IoFilterStatus** | Pointer to **NullableString** | Specifies the state of CDP IO filter. CDP IO filter is an agent which will be installed on the object for performing continuous backup. &lt;br&gt; 1. &#39;kNotInstalled&#39; specifies that CDP is enabled on this object but filter is not installed. &lt;br&gt; 2. &#39;kInstallFilterInProgress&#39; specifies that IO filter installation is triggered and in progress. &lt;br&gt; 3. &#39;kFilterInstalledIOInactive&#39; specifies that IO filter is installed but IO streaming is disabled due to missing backup or explicitly disabled by the user. &lt;br&gt; 4. &#39;kIOActivationInProgress&#39; specifies that IO filter is activated to start streaming. &lt;br&gt; 5. &#39;kIOActive&#39; specifies that filter is attached to the object and started streaming. &lt;br&gt; 6. &#39;kIODeactivationInProgress&#39; specifies that deactivation has been initiated to stop the IO streaming. &lt;br&gt; 7. &#39;kUninstallFilterInProgress&#39; specifies that uninstallation of IO filter is in progress. | [optional] 
**PreProcessingErrorMessage** | Pointer to **NullableString** | Specifies the error message from the cdp pre-processing stage if any. | [optional] 

## Methods

### NewVmwareCdpObject

`func NewVmwareCdpObject() *VmwareCdpObject`

NewVmwareCdpObject instantiates a new VmwareCdpObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareCdpObjectWithDefaults

`func NewVmwareCdpObjectWithDefaults() *VmwareCdpObject`

NewVmwareCdpObjectWithDefaults instantiates a new VmwareCdpObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowReEnableCdp

`func (o *VmwareCdpObject) GetAllowReEnableCdp() bool`

GetAllowReEnableCdp returns the AllowReEnableCdp field if non-nil, zero value otherwise.

### GetAllowReEnableCdpOk

`func (o *VmwareCdpObject) GetAllowReEnableCdpOk() (*bool, bool)`

GetAllowReEnableCdpOk returns a tuple with the AllowReEnableCdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReEnableCdp

`func (o *VmwareCdpObject) SetAllowReEnableCdp(v bool)`

SetAllowReEnableCdp sets AllowReEnableCdp field to given value.

### HasAllowReEnableCdp

`func (o *VmwareCdpObject) HasAllowReEnableCdp() bool`

HasAllowReEnableCdp returns a boolean if a field has been set.

### SetAllowReEnableCdpNil

`func (o *VmwareCdpObject) SetAllowReEnableCdpNil(b bool)`

 SetAllowReEnableCdpNil sets the value for AllowReEnableCdp to be an explicit nil

### UnsetAllowReEnableCdp
`func (o *VmwareCdpObject) UnsetAllowReEnableCdp()`

UnsetAllowReEnableCdp ensures that no value is present for AllowReEnableCdp, not even an explicit nil
### GetCdpEnabled

`func (o *VmwareCdpObject) GetCdpEnabled() bool`

GetCdpEnabled returns the CdpEnabled field if non-nil, zero value otherwise.

### GetCdpEnabledOk

`func (o *VmwareCdpObject) GetCdpEnabledOk() (*bool, bool)`

GetCdpEnabledOk returns a tuple with the CdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpEnabled

`func (o *VmwareCdpObject) SetCdpEnabled(v bool)`

SetCdpEnabled sets CdpEnabled field to given value.

### HasCdpEnabled

`func (o *VmwareCdpObject) HasCdpEnabled() bool`

HasCdpEnabled returns a boolean if a field has been set.

### SetCdpEnabledNil

`func (o *VmwareCdpObject) SetCdpEnabledNil(b bool)`

 SetCdpEnabledNil sets the value for CdpEnabled to be an explicit nil

### UnsetCdpEnabled
`func (o *VmwareCdpObject) UnsetCdpEnabled()`

UnsetCdpEnabled ensures that no value is present for CdpEnabled, not even an explicit nil
### GetLastRunInfo

`func (o *VmwareCdpObject) GetLastRunInfo() CdpObjectLastRunInfo`

GetLastRunInfo returns the LastRunInfo field if non-nil, zero value otherwise.

### GetLastRunInfoOk

`func (o *VmwareCdpObject) GetLastRunInfoOk() (*CdpObjectLastRunInfo, bool)`

GetLastRunInfoOk returns a tuple with the LastRunInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunInfo

`func (o *VmwareCdpObject) SetLastRunInfo(v CdpObjectLastRunInfo)`

SetLastRunInfo sets LastRunInfo field to given value.

### HasLastRunInfo

`func (o *VmwareCdpObject) HasLastRunInfo() bool`

HasLastRunInfo returns a boolean if a field has been set.

### GetProtectionGroupId

`func (o *VmwareCdpObject) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *VmwareCdpObject) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *VmwareCdpObject) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *VmwareCdpObject) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *VmwareCdpObject) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *VmwareCdpObject) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetGuardrailsErrorMessage

`func (o *VmwareCdpObject) GetGuardrailsErrorMessage() string`

GetGuardrailsErrorMessage returns the GuardrailsErrorMessage field if non-nil, zero value otherwise.

### GetGuardrailsErrorMessageOk

`func (o *VmwareCdpObject) GetGuardrailsErrorMessageOk() (*string, bool)`

GetGuardrailsErrorMessageOk returns a tuple with the GuardrailsErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailsErrorMessage

`func (o *VmwareCdpObject) SetGuardrailsErrorMessage(v string)`

SetGuardrailsErrorMessage sets GuardrailsErrorMessage field to given value.

### HasGuardrailsErrorMessage

`func (o *VmwareCdpObject) HasGuardrailsErrorMessage() bool`

HasGuardrailsErrorMessage returns a boolean if a field has been set.

### SetGuardrailsErrorMessageNil

`func (o *VmwareCdpObject) SetGuardrailsErrorMessageNil(b bool)`

 SetGuardrailsErrorMessageNil sets the value for GuardrailsErrorMessage to be an explicit nil

### UnsetGuardrailsErrorMessage
`func (o *VmwareCdpObject) UnsetGuardrailsErrorMessage()`

UnsetGuardrailsErrorMessage ensures that no value is present for GuardrailsErrorMessage, not even an explicit nil
### GetIoFilterErrorMessage

`func (o *VmwareCdpObject) GetIoFilterErrorMessage() string`

GetIoFilterErrorMessage returns the IoFilterErrorMessage field if non-nil, zero value otherwise.

### GetIoFilterErrorMessageOk

`func (o *VmwareCdpObject) GetIoFilterErrorMessageOk() (*string, bool)`

GetIoFilterErrorMessageOk returns a tuple with the IoFilterErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoFilterErrorMessage

`func (o *VmwareCdpObject) SetIoFilterErrorMessage(v string)`

SetIoFilterErrorMessage sets IoFilterErrorMessage field to given value.

### HasIoFilterErrorMessage

`func (o *VmwareCdpObject) HasIoFilterErrorMessage() bool`

HasIoFilterErrorMessage returns a boolean if a field has been set.

### SetIoFilterErrorMessageNil

`func (o *VmwareCdpObject) SetIoFilterErrorMessageNil(b bool)`

 SetIoFilterErrorMessageNil sets the value for IoFilterErrorMessage to be an explicit nil

### UnsetIoFilterErrorMessage
`func (o *VmwareCdpObject) UnsetIoFilterErrorMessage()`

UnsetIoFilterErrorMessage ensures that no value is present for IoFilterErrorMessage, not even an explicit nil
### GetIoFilterStatus

`func (o *VmwareCdpObject) GetIoFilterStatus() string`

GetIoFilterStatus returns the IoFilterStatus field if non-nil, zero value otherwise.

### GetIoFilterStatusOk

`func (o *VmwareCdpObject) GetIoFilterStatusOk() (*string, bool)`

GetIoFilterStatusOk returns a tuple with the IoFilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoFilterStatus

`func (o *VmwareCdpObject) SetIoFilterStatus(v string)`

SetIoFilterStatus sets IoFilterStatus field to given value.

### HasIoFilterStatus

`func (o *VmwareCdpObject) HasIoFilterStatus() bool`

HasIoFilterStatus returns a boolean if a field has been set.

### SetIoFilterStatusNil

`func (o *VmwareCdpObject) SetIoFilterStatusNil(b bool)`

 SetIoFilterStatusNil sets the value for IoFilterStatus to be an explicit nil

### UnsetIoFilterStatus
`func (o *VmwareCdpObject) UnsetIoFilterStatus()`

UnsetIoFilterStatus ensures that no value is present for IoFilterStatus, not even an explicit nil
### GetPreProcessingErrorMessage

`func (o *VmwareCdpObject) GetPreProcessingErrorMessage() string`

GetPreProcessingErrorMessage returns the PreProcessingErrorMessage field if non-nil, zero value otherwise.

### GetPreProcessingErrorMessageOk

`func (o *VmwareCdpObject) GetPreProcessingErrorMessageOk() (*string, bool)`

GetPreProcessingErrorMessageOk returns a tuple with the PreProcessingErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProcessingErrorMessage

`func (o *VmwareCdpObject) SetPreProcessingErrorMessage(v string)`

SetPreProcessingErrorMessage sets PreProcessingErrorMessage field to given value.

### HasPreProcessingErrorMessage

`func (o *VmwareCdpObject) HasPreProcessingErrorMessage() bool`

HasPreProcessingErrorMessage returns a boolean if a field has been set.

### SetPreProcessingErrorMessageNil

`func (o *VmwareCdpObject) SetPreProcessingErrorMessageNil(b bool)`

 SetPreProcessingErrorMessageNil sets the value for PreProcessingErrorMessage to be an explicit nil

### UnsetPreProcessingErrorMessage
`func (o *VmwareCdpObject) UnsetPreProcessingErrorMessage()`

UnsetPreProcessingErrorMessage ensures that no value is present for PreProcessingErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


