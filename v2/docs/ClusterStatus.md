# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirgapConfigStats** | Pointer to [**AirgapConfig**](AirgapConfig.md) |  | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the id of the cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] 
**CurrentOperation** | Pointer to **NullableString** | Specifies the current operation of the cluster. | [optional] 
**Message** | Pointer to **NullableString** | Specifies an optional message describing details of the cluster status. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the cluster. | [optional] 
**NodeStatuses** | Pointer to [**[]NodeStatusResult**](NodeStatusResult.md) | Specifies the status of each node on the cluster. | [optional] 
**RemovalState** | Pointer to **NullableString** | RemovalState specifies the possible states for removal operations happening on the cluster. | [optional] 
**ServicesSynced** | Pointer to **NullableBool** | Specifies whether or not the services are synced with the list of stopped services. | [optional] 
**SoftwareVersion** | Pointer to **NullableString** | Specifies the software version of the cluster. | [optional] 
**StoppedServices** | Pointer to **[]string** | Specifies the service name. | [optional] 
**SystemAppStatus** | Pointer to [**[]SystemAppStatusParams**](SystemAppStatusParams.md) | Specifies the status of each system app on the cluster. | [optional] 

## Methods

### NewClusterStatus

`func NewClusterStatus() *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirgapConfigStats

`func (o *ClusterStatus) GetAirgapConfigStats() AirgapConfig`

GetAirgapConfigStats returns the AirgapConfigStats field if non-nil, zero value otherwise.

### GetAirgapConfigStatsOk

`func (o *ClusterStatus) GetAirgapConfigStatsOk() (*AirgapConfig, bool)`

GetAirgapConfigStatsOk returns a tuple with the AirgapConfigStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirgapConfigStats

`func (o *ClusterStatus) SetAirgapConfigStats(v AirgapConfig)`

SetAirgapConfigStats sets AirgapConfigStats field to given value.

### HasAirgapConfigStats

`func (o *ClusterStatus) HasAirgapConfigStats() bool`

HasAirgapConfigStats returns a boolean if a field has been set.

### GetClusterId

`func (o *ClusterStatus) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterStatus) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterStatus) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterStatus) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ClusterStatus) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ClusterStatus) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ClusterStatus) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ClusterStatus) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ClusterStatus) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ClusterStatus) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ClusterStatus) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ClusterStatus) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetCurrentOperation

`func (o *ClusterStatus) GetCurrentOperation() string`

GetCurrentOperation returns the CurrentOperation field if non-nil, zero value otherwise.

### GetCurrentOperationOk

`func (o *ClusterStatus) GetCurrentOperationOk() (*string, bool)`

GetCurrentOperationOk returns a tuple with the CurrentOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOperation

`func (o *ClusterStatus) SetCurrentOperation(v string)`

SetCurrentOperation sets CurrentOperation field to given value.

### HasCurrentOperation

`func (o *ClusterStatus) HasCurrentOperation() bool`

HasCurrentOperation returns a boolean if a field has been set.

### SetCurrentOperationNil

`func (o *ClusterStatus) SetCurrentOperationNil(b bool)`

 SetCurrentOperationNil sets the value for CurrentOperation to be an explicit nil

### UnsetCurrentOperation
`func (o *ClusterStatus) UnsetCurrentOperation()`

UnsetCurrentOperation ensures that no value is present for CurrentOperation, not even an explicit nil
### GetMessage

`func (o *ClusterStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ClusterStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ClusterStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetName

`func (o *ClusterStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClusterStatus) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClusterStatus) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeStatuses

`func (o *ClusterStatus) GetNodeStatuses() []NodeStatusResult`

GetNodeStatuses returns the NodeStatuses field if non-nil, zero value otherwise.

### GetNodeStatusesOk

`func (o *ClusterStatus) GetNodeStatusesOk() (*[]NodeStatusResult, bool)`

GetNodeStatusesOk returns a tuple with the NodeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatuses

`func (o *ClusterStatus) SetNodeStatuses(v []NodeStatusResult)`

SetNodeStatuses sets NodeStatuses field to given value.

### HasNodeStatuses

`func (o *ClusterStatus) HasNodeStatuses() bool`

HasNodeStatuses returns a boolean if a field has been set.

### SetNodeStatusesNil

`func (o *ClusterStatus) SetNodeStatusesNil(b bool)`

 SetNodeStatusesNil sets the value for NodeStatuses to be an explicit nil

### UnsetNodeStatuses
`func (o *ClusterStatus) UnsetNodeStatuses()`

UnsetNodeStatuses ensures that no value is present for NodeStatuses, not even an explicit nil
### GetRemovalState

`func (o *ClusterStatus) GetRemovalState() string`

GetRemovalState returns the RemovalState field if non-nil, zero value otherwise.

### GetRemovalStateOk

`func (o *ClusterStatus) GetRemovalStateOk() (*string, bool)`

GetRemovalStateOk returns a tuple with the RemovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalState

`func (o *ClusterStatus) SetRemovalState(v string)`

SetRemovalState sets RemovalState field to given value.

### HasRemovalState

`func (o *ClusterStatus) HasRemovalState() bool`

HasRemovalState returns a boolean if a field has been set.

### SetRemovalStateNil

`func (o *ClusterStatus) SetRemovalStateNil(b bool)`

 SetRemovalStateNil sets the value for RemovalState to be an explicit nil

### UnsetRemovalState
`func (o *ClusterStatus) UnsetRemovalState()`

UnsetRemovalState ensures that no value is present for RemovalState, not even an explicit nil
### GetServicesSynced

`func (o *ClusterStatus) GetServicesSynced() bool`

GetServicesSynced returns the ServicesSynced field if non-nil, zero value otherwise.

### GetServicesSyncedOk

`func (o *ClusterStatus) GetServicesSyncedOk() (*bool, bool)`

GetServicesSyncedOk returns a tuple with the ServicesSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesSynced

`func (o *ClusterStatus) SetServicesSynced(v bool)`

SetServicesSynced sets ServicesSynced field to given value.

### HasServicesSynced

`func (o *ClusterStatus) HasServicesSynced() bool`

HasServicesSynced returns a boolean if a field has been set.

### SetServicesSyncedNil

`func (o *ClusterStatus) SetServicesSyncedNil(b bool)`

 SetServicesSyncedNil sets the value for ServicesSynced to be an explicit nil

### UnsetServicesSynced
`func (o *ClusterStatus) UnsetServicesSynced()`

UnsetServicesSynced ensures that no value is present for ServicesSynced, not even an explicit nil
### GetSoftwareVersion

`func (o *ClusterStatus) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *ClusterStatus) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *ClusterStatus) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *ClusterStatus) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *ClusterStatus) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *ClusterStatus) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetStoppedServices

`func (o *ClusterStatus) GetStoppedServices() []string`

GetStoppedServices returns the StoppedServices field if non-nil, zero value otherwise.

### GetStoppedServicesOk

`func (o *ClusterStatus) GetStoppedServicesOk() (*[]string, bool)`

GetStoppedServicesOk returns a tuple with the StoppedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedServices

`func (o *ClusterStatus) SetStoppedServices(v []string)`

SetStoppedServices sets StoppedServices field to given value.

### HasStoppedServices

`func (o *ClusterStatus) HasStoppedServices() bool`

HasStoppedServices returns a boolean if a field has been set.

### SetStoppedServicesNil

`func (o *ClusterStatus) SetStoppedServicesNil(b bool)`

 SetStoppedServicesNil sets the value for StoppedServices to be an explicit nil

### UnsetStoppedServices
`func (o *ClusterStatus) UnsetStoppedServices()`

UnsetStoppedServices ensures that no value is present for StoppedServices, not even an explicit nil
### GetSystemAppStatus

`func (o *ClusterStatus) GetSystemAppStatus() []SystemAppStatusParams`

GetSystemAppStatus returns the SystemAppStatus field if non-nil, zero value otherwise.

### GetSystemAppStatusOk

`func (o *ClusterStatus) GetSystemAppStatusOk() (*[]SystemAppStatusParams, bool)`

GetSystemAppStatusOk returns a tuple with the SystemAppStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAppStatus

`func (o *ClusterStatus) SetSystemAppStatus(v []SystemAppStatusParams)`

SetSystemAppStatus sets SystemAppStatus field to given value.

### HasSystemAppStatus

`func (o *ClusterStatus) HasSystemAppStatus() bool`

HasSystemAppStatus returns a boolean if a field has been set.

### SetSystemAppStatusNil

`func (o *ClusterStatus) SetSystemAppStatusNil(b bool)`

 SetSystemAppStatusNil sets the value for SystemAppStatus to be an explicit nil

### UnsetSystemAppStatus
`func (o *ClusterStatus) UnsetSystemAppStatus()`

UnsetSystemAppStatus ensures that no value is present for SystemAppStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


