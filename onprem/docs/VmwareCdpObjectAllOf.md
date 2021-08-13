# VmwareCdpObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IoFilterStatus** | Pointer to **NullableString** | Specifies the state of CDP IO filter. CDP IO filter is an agent which will be installed on the object for performing continuous backup. &lt;br&gt; 1. &#39;kNotInstalled&#39; specifies that CDP is enabled on this object but filter is not installed. &lt;br&gt; 2. &#39;kInstallFilterInProgress&#39; specifies that IO filter installation is triggered and in progress. &lt;br&gt; 3. &#39;kFilterInstalledIOInactive&#39; specifies that IO filter is installed but IO streaming is disabled due to missing backup or explicitly disabled by the user. &lt;br&gt; 4. &#39;kIOActivationInProgress&#39; specifies that IO filter is activated to start streaming. &lt;br&gt; 5. &#39;kIOActive&#39; specifies that filter is attached to the object and started streaming. &lt;br&gt; 6. &#39;kIODeactivationInProgress&#39; specifies that deactivation has been initiated to stop the IO streaming. &lt;br&gt; 7. &#39;kUninstallFilterInProgress&#39; specifies that uninstallation of IO filter is in progress. | [optional] 
**IoFilterErrorMessage** | Pointer to **NullableString** | Specifies the error message related to IO filter if there is any. | [optional] 

## Methods

### NewVmwareCdpObjectAllOf

`func NewVmwareCdpObjectAllOf() *VmwareCdpObjectAllOf`

NewVmwareCdpObjectAllOf instantiates a new VmwareCdpObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareCdpObjectAllOfWithDefaults

`func NewVmwareCdpObjectAllOfWithDefaults() *VmwareCdpObjectAllOf`

NewVmwareCdpObjectAllOfWithDefaults instantiates a new VmwareCdpObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIoFilterStatus

`func (o *VmwareCdpObjectAllOf) GetIoFilterStatus() string`

GetIoFilterStatus returns the IoFilterStatus field if non-nil, zero value otherwise.

### GetIoFilterStatusOk

`func (o *VmwareCdpObjectAllOf) GetIoFilterStatusOk() (*string, bool)`

GetIoFilterStatusOk returns a tuple with the IoFilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoFilterStatus

`func (o *VmwareCdpObjectAllOf) SetIoFilterStatus(v string)`

SetIoFilterStatus sets IoFilterStatus field to given value.

### HasIoFilterStatus

`func (o *VmwareCdpObjectAllOf) HasIoFilterStatus() bool`

HasIoFilterStatus returns a boolean if a field has been set.

### SetIoFilterStatusNil

`func (o *VmwareCdpObjectAllOf) SetIoFilterStatusNil(b bool)`

 SetIoFilterStatusNil sets the value for IoFilterStatus to be an explicit nil

### UnsetIoFilterStatus
`func (o *VmwareCdpObjectAllOf) UnsetIoFilterStatus()`

UnsetIoFilterStatus ensures that no value is present for IoFilterStatus, not even an explicit nil
### GetIoFilterErrorMessage

`func (o *VmwareCdpObjectAllOf) GetIoFilterErrorMessage() string`

GetIoFilterErrorMessage returns the IoFilterErrorMessage field if non-nil, zero value otherwise.

### GetIoFilterErrorMessageOk

`func (o *VmwareCdpObjectAllOf) GetIoFilterErrorMessageOk() (*string, bool)`

GetIoFilterErrorMessageOk returns a tuple with the IoFilterErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoFilterErrorMessage

`func (o *VmwareCdpObjectAllOf) SetIoFilterErrorMessage(v string)`

SetIoFilterErrorMessage sets IoFilterErrorMessage field to given value.

### HasIoFilterErrorMessage

`func (o *VmwareCdpObjectAllOf) HasIoFilterErrorMessage() bool`

HasIoFilterErrorMessage returns a boolean if a field has been set.

### SetIoFilterErrorMessageNil

`func (o *VmwareCdpObjectAllOf) SetIoFilterErrorMessageNil(b bool)`

 SetIoFilterErrorMessageNil sets the value for IoFilterErrorMessage to be an explicit nil

### UnsetIoFilterErrorMessage
`func (o *VmwareCdpObjectAllOf) UnsetIoFilterErrorMessage()`

UnsetIoFilterErrorMessage ensures that no value is present for IoFilterErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


