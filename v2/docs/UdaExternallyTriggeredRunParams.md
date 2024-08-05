# UdaExternallyTriggeredRunParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupArgs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies a map of custom arguments to be supplied to the plugin. | [optional] 
**ControlNode** | Pointer to **NullableString** | Specifies the IP or FQDN of the source host where this backup will run. | [optional] 

## Methods

### NewUdaExternallyTriggeredRunParams

`func NewUdaExternallyTriggeredRunParams() *UdaExternallyTriggeredRunParams`

NewUdaExternallyTriggeredRunParams instantiates a new UdaExternallyTriggeredRunParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaExternallyTriggeredRunParamsWithDefaults

`func NewUdaExternallyTriggeredRunParamsWithDefaults() *UdaExternallyTriggeredRunParams`

NewUdaExternallyTriggeredRunParamsWithDefaults instantiates a new UdaExternallyTriggeredRunParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupArgs

`func (o *UdaExternallyTriggeredRunParams) GetBackupArgs() []KeyValuePair`

GetBackupArgs returns the BackupArgs field if non-nil, zero value otherwise.

### GetBackupArgsOk

`func (o *UdaExternallyTriggeredRunParams) GetBackupArgsOk() (*[]KeyValuePair, bool)`

GetBackupArgsOk returns a tuple with the BackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupArgs

`func (o *UdaExternallyTriggeredRunParams) SetBackupArgs(v []KeyValuePair)`

SetBackupArgs sets BackupArgs field to given value.

### HasBackupArgs

`func (o *UdaExternallyTriggeredRunParams) HasBackupArgs() bool`

HasBackupArgs returns a boolean if a field has been set.

### SetBackupArgsNil

`func (o *UdaExternallyTriggeredRunParams) SetBackupArgsNil(b bool)`

 SetBackupArgsNil sets the value for BackupArgs to be an explicit nil

### UnsetBackupArgs
`func (o *UdaExternallyTriggeredRunParams) UnsetBackupArgs()`

UnsetBackupArgs ensures that no value is present for BackupArgs, not even an explicit nil
### GetControlNode

`func (o *UdaExternallyTriggeredRunParams) GetControlNode() string`

GetControlNode returns the ControlNode field if non-nil, zero value otherwise.

### GetControlNodeOk

`func (o *UdaExternallyTriggeredRunParams) GetControlNodeOk() (*string, bool)`

GetControlNodeOk returns a tuple with the ControlNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlNode

`func (o *UdaExternallyTriggeredRunParams) SetControlNode(v string)`

SetControlNode sets ControlNode field to given value.

### HasControlNode

`func (o *UdaExternallyTriggeredRunParams) HasControlNode() bool`

HasControlNode returns a boolean if a field has been set.

### SetControlNodeNil

`func (o *UdaExternallyTriggeredRunParams) SetControlNodeNil(b bool)`

 SetControlNodeNil sets the value for ControlNode to be an explicit nil

### UnsetControlNode
`func (o *UdaExternallyTriggeredRunParams) UnsetControlNode()`

UnsetControlNode ensures that no value is present for ControlNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


