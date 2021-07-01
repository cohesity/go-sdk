# AppInstanceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalNetworkInfo** | Pointer to [**ExternalNetworkInfo**](ExternalNetworkInfo.md) |  | [optional] 
**InstanceSize** | Pointer to **NullableString** | Instance size specification (e.g. small/medium/large). Used to determine container resources. | [optional] 
**ProtectedObjectPrivileges** | Pointer to [**ProtectedObjectPrivileges**](ProtectedObjectPrivileges.md) |  | [optional] 
**QosTier** | Pointer to **NullableString** | Specifies QoSTier of the app instance. Specifies QoS Tier for an app instance. App instances are allocated resources such as memory, CPU and IO based on their QoS Tier. kLow - Low QoS Tier. kMedium - Medium QoS Tier. kHigh - High QoS Tier. kMax - Max QoS Tier. | [optional] 
**ReadViewPrivileges** | Pointer to [**ViewPrivileges**](ViewPrivileges.md) |  | [optional] 
**ReadWriteViewPrivileges** | Pointer to [**ViewPrivileges**](ViewPrivileges.md) |  | [optional] 
**VmNumReplicasList** | Pointer to [**[]VmNumReplicas**](VmNumReplicas.md) | List of vm-name, replica count pairs to be used at the time of app instance launch. | [optional] 

## Methods

### NewAppInstanceSettings

`func NewAppInstanceSettings() *AppInstanceSettings`

NewAppInstanceSettings instantiates a new AppInstanceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceSettingsWithDefaults

`func NewAppInstanceSettingsWithDefaults() *AppInstanceSettings`

NewAppInstanceSettingsWithDefaults instantiates a new AppInstanceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalNetworkInfo

`func (o *AppInstanceSettings) GetExternalNetworkInfo() ExternalNetworkInfo`

GetExternalNetworkInfo returns the ExternalNetworkInfo field if non-nil, zero value otherwise.

### GetExternalNetworkInfoOk

`func (o *AppInstanceSettings) GetExternalNetworkInfoOk() (*ExternalNetworkInfo, bool)`

GetExternalNetworkInfoOk returns a tuple with the ExternalNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetworkInfo

`func (o *AppInstanceSettings) SetExternalNetworkInfo(v ExternalNetworkInfo)`

SetExternalNetworkInfo sets ExternalNetworkInfo field to given value.

### HasExternalNetworkInfo

`func (o *AppInstanceSettings) HasExternalNetworkInfo() bool`

HasExternalNetworkInfo returns a boolean if a field has been set.

### GetInstanceSize

`func (o *AppInstanceSettings) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *AppInstanceSettings) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *AppInstanceSettings) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *AppInstanceSettings) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### SetInstanceSizeNil

`func (o *AppInstanceSettings) SetInstanceSizeNil(b bool)`

 SetInstanceSizeNil sets the value for InstanceSize to be an explicit nil

### UnsetInstanceSize
`func (o *AppInstanceSettings) UnsetInstanceSize()`

UnsetInstanceSize ensures that no value is present for InstanceSize, not even an explicit nil
### GetProtectedObjectPrivileges

`func (o *AppInstanceSettings) GetProtectedObjectPrivileges() ProtectedObjectPrivileges`

GetProtectedObjectPrivileges returns the ProtectedObjectPrivileges field if non-nil, zero value otherwise.

### GetProtectedObjectPrivilegesOk

`func (o *AppInstanceSettings) GetProtectedObjectPrivilegesOk() (*ProtectedObjectPrivileges, bool)`

GetProtectedObjectPrivilegesOk returns a tuple with the ProtectedObjectPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedObjectPrivileges

`func (o *AppInstanceSettings) SetProtectedObjectPrivileges(v ProtectedObjectPrivileges)`

SetProtectedObjectPrivileges sets ProtectedObjectPrivileges field to given value.

### HasProtectedObjectPrivileges

`func (o *AppInstanceSettings) HasProtectedObjectPrivileges() bool`

HasProtectedObjectPrivileges returns a boolean if a field has been set.

### GetQosTier

`func (o *AppInstanceSettings) GetQosTier() string`

GetQosTier returns the QosTier field if non-nil, zero value otherwise.

### GetQosTierOk

`func (o *AppInstanceSettings) GetQosTierOk() (*string, bool)`

GetQosTierOk returns a tuple with the QosTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosTier

`func (o *AppInstanceSettings) SetQosTier(v string)`

SetQosTier sets QosTier field to given value.

### HasQosTier

`func (o *AppInstanceSettings) HasQosTier() bool`

HasQosTier returns a boolean if a field has been set.

### SetQosTierNil

`func (o *AppInstanceSettings) SetQosTierNil(b bool)`

 SetQosTierNil sets the value for QosTier to be an explicit nil

### UnsetQosTier
`func (o *AppInstanceSettings) UnsetQosTier()`

UnsetQosTier ensures that no value is present for QosTier, not even an explicit nil
### GetReadViewPrivileges

`func (o *AppInstanceSettings) GetReadViewPrivileges() ViewPrivileges`

GetReadViewPrivileges returns the ReadViewPrivileges field if non-nil, zero value otherwise.

### GetReadViewPrivilegesOk

`func (o *AppInstanceSettings) GetReadViewPrivilegesOk() (*ViewPrivileges, bool)`

GetReadViewPrivilegesOk returns a tuple with the ReadViewPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadViewPrivileges

`func (o *AppInstanceSettings) SetReadViewPrivileges(v ViewPrivileges)`

SetReadViewPrivileges sets ReadViewPrivileges field to given value.

### HasReadViewPrivileges

`func (o *AppInstanceSettings) HasReadViewPrivileges() bool`

HasReadViewPrivileges returns a boolean if a field has been set.

### GetReadWriteViewPrivileges

`func (o *AppInstanceSettings) GetReadWriteViewPrivileges() ViewPrivileges`

GetReadWriteViewPrivileges returns the ReadWriteViewPrivileges field if non-nil, zero value otherwise.

### GetReadWriteViewPrivilegesOk

`func (o *AppInstanceSettings) GetReadWriteViewPrivilegesOk() (*ViewPrivileges, bool)`

GetReadWriteViewPrivilegesOk returns a tuple with the ReadWriteViewPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteViewPrivileges

`func (o *AppInstanceSettings) SetReadWriteViewPrivileges(v ViewPrivileges)`

SetReadWriteViewPrivileges sets ReadWriteViewPrivileges field to given value.

### HasReadWriteViewPrivileges

`func (o *AppInstanceSettings) HasReadWriteViewPrivileges() bool`

HasReadWriteViewPrivileges returns a boolean if a field has been set.

### GetVmNumReplicasList

`func (o *AppInstanceSettings) GetVmNumReplicasList() []VmNumReplicas`

GetVmNumReplicasList returns the VmNumReplicasList field if non-nil, zero value otherwise.

### GetVmNumReplicasListOk

`func (o *AppInstanceSettings) GetVmNumReplicasListOk() (*[]VmNumReplicas, bool)`

GetVmNumReplicasListOk returns a tuple with the VmNumReplicasList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNumReplicasList

`func (o *AppInstanceSettings) SetVmNumReplicasList(v []VmNumReplicas)`

SetVmNumReplicasList sets VmNumReplicasList field to given value.

### HasVmNumReplicasList

`func (o *AppInstanceSettings) HasVmNumReplicasList() bool`

HasVmNumReplicasList returns a boolean if a field has been set.

### SetVmNumReplicasListNil

`func (o *AppInstanceSettings) SetVmNumReplicasListNil(b bool)`

 SetVmNumReplicasListNil sets the value for VmNumReplicasList to be an explicit nil

### UnsetVmNumReplicasList
`func (o *AppInstanceSettings) UnsetVmNumReplicasList()`

UnsetVmNumReplicasList ensures that no value is present for VmNumReplicasList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


