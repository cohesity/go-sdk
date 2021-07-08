# AppsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowExternalTraffic** | Pointer to **NullableBool** | Whether to allow pod external traffic. | [optional] 
**AllowUnrestictedViewAccess** | Pointer to **NullableBool** | Whether to allow apps unrestricted view access. | [optional] 
**AppsMode** | Pointer to **NullableString** | Specifies the various modes for running apps. &#39;kDisabled&#39; specifies that apps are disabled. &#39;kBareMetal&#39; specifies that apps could only run in containers on the node (no VM). &#39;kVmOnly&#39; specifies that apps could only run in containers on a VM hosted by the node. | [optional] 
**AppsSubnet** | Pointer to [**Subnet**](Subnet.md) |  | [optional] 
**OvercommitMemoryPct** | Pointer to **NullableInt32** | The system memory to overcommit for apps. | [optional] 
**ReservedCpuMillicores** | Pointer to **NullableInt32** | The CPU millicores to reserve for apps. | [optional] 
**ReservedMemoryPct** | Pointer to **NullableInt32** | The system memory to reserve for apps. | [optional] 

## Methods

### NewAppsConfig

`func NewAppsConfig() *AppsConfig`

NewAppsConfig instantiates a new AppsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsConfigWithDefaults

`func NewAppsConfigWithDefaults() *AppsConfig`

NewAppsConfigWithDefaults instantiates a new AppsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowExternalTraffic

`func (o *AppsConfig) GetAllowExternalTraffic() bool`

GetAllowExternalTraffic returns the AllowExternalTraffic field if non-nil, zero value otherwise.

### GetAllowExternalTrafficOk

`func (o *AppsConfig) GetAllowExternalTrafficOk() (*bool, bool)`

GetAllowExternalTrafficOk returns a tuple with the AllowExternalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExternalTraffic

`func (o *AppsConfig) SetAllowExternalTraffic(v bool)`

SetAllowExternalTraffic sets AllowExternalTraffic field to given value.

### HasAllowExternalTraffic

`func (o *AppsConfig) HasAllowExternalTraffic() bool`

HasAllowExternalTraffic returns a boolean if a field has been set.

### SetAllowExternalTrafficNil

`func (o *AppsConfig) SetAllowExternalTrafficNil(b bool)`

 SetAllowExternalTrafficNil sets the value for AllowExternalTraffic to be an explicit nil

### UnsetAllowExternalTraffic
`func (o *AppsConfig) UnsetAllowExternalTraffic()`

UnsetAllowExternalTraffic ensures that no value is present for AllowExternalTraffic, not even an explicit nil
### GetAllowUnrestictedViewAccess

`func (o *AppsConfig) GetAllowUnrestictedViewAccess() bool`

GetAllowUnrestictedViewAccess returns the AllowUnrestictedViewAccess field if non-nil, zero value otherwise.

### GetAllowUnrestictedViewAccessOk

`func (o *AppsConfig) GetAllowUnrestictedViewAccessOk() (*bool, bool)`

GetAllowUnrestictedViewAccessOk returns a tuple with the AllowUnrestictedViewAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnrestictedViewAccess

`func (o *AppsConfig) SetAllowUnrestictedViewAccess(v bool)`

SetAllowUnrestictedViewAccess sets AllowUnrestictedViewAccess field to given value.

### HasAllowUnrestictedViewAccess

`func (o *AppsConfig) HasAllowUnrestictedViewAccess() bool`

HasAllowUnrestictedViewAccess returns a boolean if a field has been set.

### SetAllowUnrestictedViewAccessNil

`func (o *AppsConfig) SetAllowUnrestictedViewAccessNil(b bool)`

 SetAllowUnrestictedViewAccessNil sets the value for AllowUnrestictedViewAccess to be an explicit nil

### UnsetAllowUnrestictedViewAccess
`func (o *AppsConfig) UnsetAllowUnrestictedViewAccess()`

UnsetAllowUnrestictedViewAccess ensures that no value is present for AllowUnrestictedViewAccess, not even an explicit nil
### GetAppsMode

`func (o *AppsConfig) GetAppsMode() string`

GetAppsMode returns the AppsMode field if non-nil, zero value otherwise.

### GetAppsModeOk

`func (o *AppsConfig) GetAppsModeOk() (*string, bool)`

GetAppsModeOk returns a tuple with the AppsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsMode

`func (o *AppsConfig) SetAppsMode(v string)`

SetAppsMode sets AppsMode field to given value.

### HasAppsMode

`func (o *AppsConfig) HasAppsMode() bool`

HasAppsMode returns a boolean if a field has been set.

### SetAppsModeNil

`func (o *AppsConfig) SetAppsModeNil(b bool)`

 SetAppsModeNil sets the value for AppsMode to be an explicit nil

### UnsetAppsMode
`func (o *AppsConfig) UnsetAppsMode()`

UnsetAppsMode ensures that no value is present for AppsMode, not even an explicit nil
### GetAppsSubnet

`func (o *AppsConfig) GetAppsSubnet() Subnet`

GetAppsSubnet returns the AppsSubnet field if non-nil, zero value otherwise.

### GetAppsSubnetOk

`func (o *AppsConfig) GetAppsSubnetOk() (*Subnet, bool)`

GetAppsSubnetOk returns a tuple with the AppsSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsSubnet

`func (o *AppsConfig) SetAppsSubnet(v Subnet)`

SetAppsSubnet sets AppsSubnet field to given value.

### HasAppsSubnet

`func (o *AppsConfig) HasAppsSubnet() bool`

HasAppsSubnet returns a boolean if a field has been set.

### GetOvercommitMemoryPct

`func (o *AppsConfig) GetOvercommitMemoryPct() int32`

GetOvercommitMemoryPct returns the OvercommitMemoryPct field if non-nil, zero value otherwise.

### GetOvercommitMemoryPctOk

`func (o *AppsConfig) GetOvercommitMemoryPctOk() (*int32, bool)`

GetOvercommitMemoryPctOk returns a tuple with the OvercommitMemoryPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommitMemoryPct

`func (o *AppsConfig) SetOvercommitMemoryPct(v int32)`

SetOvercommitMemoryPct sets OvercommitMemoryPct field to given value.

### HasOvercommitMemoryPct

`func (o *AppsConfig) HasOvercommitMemoryPct() bool`

HasOvercommitMemoryPct returns a boolean if a field has been set.

### SetOvercommitMemoryPctNil

`func (o *AppsConfig) SetOvercommitMemoryPctNil(b bool)`

 SetOvercommitMemoryPctNil sets the value for OvercommitMemoryPct to be an explicit nil

### UnsetOvercommitMemoryPct
`func (o *AppsConfig) UnsetOvercommitMemoryPct()`

UnsetOvercommitMemoryPct ensures that no value is present for OvercommitMemoryPct, not even an explicit nil
### GetReservedCpuMillicores

`func (o *AppsConfig) GetReservedCpuMillicores() int32`

GetReservedCpuMillicores returns the ReservedCpuMillicores field if non-nil, zero value otherwise.

### GetReservedCpuMillicoresOk

`func (o *AppsConfig) GetReservedCpuMillicoresOk() (*int32, bool)`

GetReservedCpuMillicoresOk returns a tuple with the ReservedCpuMillicores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCpuMillicores

`func (o *AppsConfig) SetReservedCpuMillicores(v int32)`

SetReservedCpuMillicores sets ReservedCpuMillicores field to given value.

### HasReservedCpuMillicores

`func (o *AppsConfig) HasReservedCpuMillicores() bool`

HasReservedCpuMillicores returns a boolean if a field has been set.

### SetReservedCpuMillicoresNil

`func (o *AppsConfig) SetReservedCpuMillicoresNil(b bool)`

 SetReservedCpuMillicoresNil sets the value for ReservedCpuMillicores to be an explicit nil

### UnsetReservedCpuMillicores
`func (o *AppsConfig) UnsetReservedCpuMillicores()`

UnsetReservedCpuMillicores ensures that no value is present for ReservedCpuMillicores, not even an explicit nil
### GetReservedMemoryPct

`func (o *AppsConfig) GetReservedMemoryPct() int32`

GetReservedMemoryPct returns the ReservedMemoryPct field if non-nil, zero value otherwise.

### GetReservedMemoryPctOk

`func (o *AppsConfig) GetReservedMemoryPctOk() (*int32, bool)`

GetReservedMemoryPctOk returns a tuple with the ReservedMemoryPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedMemoryPct

`func (o *AppsConfig) SetReservedMemoryPct(v int32)`

SetReservedMemoryPct sets ReservedMemoryPct field to given value.

### HasReservedMemoryPct

`func (o *AppsConfig) HasReservedMemoryPct() bool`

HasReservedMemoryPct returns a boolean if a field has been set.

### SetReservedMemoryPctNil

`func (o *AppsConfig) SetReservedMemoryPctNil(b bool)`

 SetReservedMemoryPctNil sets the value for ReservedMemoryPct to be an explicit nil

### UnsetReservedMemoryPct
`func (o *AppsConfig) UnsetReservedMemoryPct()`

UnsetReservedMemoryPct ensures that no value is present for ReservedMemoryPct, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


