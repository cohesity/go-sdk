# VmwareStandbyObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableInt64** | Specifies the entity id of the corresponding backup object for which this standby is configured. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id to which this standby object belongs. | [optional] [readonly] 
**CdpStandbyStatus** | Pointer to **NullableString** | Specifies the current status of the standby object protected using continuous data protection policy. | [optional] 
**GuestId** | Pointer to **NullableString** | Specifies the guest ID(OS) of the standby VM for the corresponding backup object. | [optional] [readonly] 
**StandbyMOref** | Pointer to [**MOref**](MOref.md) |  | [optional] 
**StandbyTime** | Pointer to **NullableInt64** | Specifies the time till which the standby object has been hydrated for the corresponding backup object. | [optional] 

## Methods

### NewVmwareStandbyObject

`func NewVmwareStandbyObject() *VmwareStandbyObject`

NewVmwareStandbyObject instantiates a new VmwareStandbyObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareStandbyObjectWithDefaults

`func NewVmwareStandbyObjectWithDefaults() *VmwareStandbyObject`

NewVmwareStandbyObjectWithDefaults instantiates a new VmwareStandbyObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *VmwareStandbyObject) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *VmwareStandbyObject) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *VmwareStandbyObject) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *VmwareStandbyObject) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *VmwareStandbyObject) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *VmwareStandbyObject) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetProtectionGroupId

`func (o *VmwareStandbyObject) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *VmwareStandbyObject) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *VmwareStandbyObject) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *VmwareStandbyObject) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *VmwareStandbyObject) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *VmwareStandbyObject) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetCdpStandbyStatus

`func (o *VmwareStandbyObject) GetCdpStandbyStatus() string`

GetCdpStandbyStatus returns the CdpStandbyStatus field if non-nil, zero value otherwise.

### GetCdpStandbyStatusOk

`func (o *VmwareStandbyObject) GetCdpStandbyStatusOk() (*string, bool)`

GetCdpStandbyStatusOk returns a tuple with the CdpStandbyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpStandbyStatus

`func (o *VmwareStandbyObject) SetCdpStandbyStatus(v string)`

SetCdpStandbyStatus sets CdpStandbyStatus field to given value.

### HasCdpStandbyStatus

`func (o *VmwareStandbyObject) HasCdpStandbyStatus() bool`

HasCdpStandbyStatus returns a boolean if a field has been set.

### SetCdpStandbyStatusNil

`func (o *VmwareStandbyObject) SetCdpStandbyStatusNil(b bool)`

 SetCdpStandbyStatusNil sets the value for CdpStandbyStatus to be an explicit nil

### UnsetCdpStandbyStatus
`func (o *VmwareStandbyObject) UnsetCdpStandbyStatus()`

UnsetCdpStandbyStatus ensures that no value is present for CdpStandbyStatus, not even an explicit nil
### GetGuestId

`func (o *VmwareStandbyObject) GetGuestId() string`

GetGuestId returns the GuestId field if non-nil, zero value otherwise.

### GetGuestIdOk

`func (o *VmwareStandbyObject) GetGuestIdOk() (*string, bool)`

GetGuestIdOk returns a tuple with the GuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestId

`func (o *VmwareStandbyObject) SetGuestId(v string)`

SetGuestId sets GuestId field to given value.

### HasGuestId

`func (o *VmwareStandbyObject) HasGuestId() bool`

HasGuestId returns a boolean if a field has been set.

### SetGuestIdNil

`func (o *VmwareStandbyObject) SetGuestIdNil(b bool)`

 SetGuestIdNil sets the value for GuestId to be an explicit nil

### UnsetGuestId
`func (o *VmwareStandbyObject) UnsetGuestId()`

UnsetGuestId ensures that no value is present for GuestId, not even an explicit nil
### GetStandbyMOref

`func (o *VmwareStandbyObject) GetStandbyMOref() MOref`

GetStandbyMOref returns the StandbyMOref field if non-nil, zero value otherwise.

### GetStandbyMOrefOk

`func (o *VmwareStandbyObject) GetStandbyMOrefOk() (*MOref, bool)`

GetStandbyMOrefOk returns a tuple with the StandbyMOref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyMOref

`func (o *VmwareStandbyObject) SetStandbyMOref(v MOref)`

SetStandbyMOref sets StandbyMOref field to given value.

### HasStandbyMOref

`func (o *VmwareStandbyObject) HasStandbyMOref() bool`

HasStandbyMOref returns a boolean if a field has been set.

### GetStandbyTime

`func (o *VmwareStandbyObject) GetStandbyTime() int64`

GetStandbyTime returns the StandbyTime field if non-nil, zero value otherwise.

### GetStandbyTimeOk

`func (o *VmwareStandbyObject) GetStandbyTimeOk() (*int64, bool)`

GetStandbyTimeOk returns a tuple with the StandbyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyTime

`func (o *VmwareStandbyObject) SetStandbyTime(v int64)`

SetStandbyTime sets StandbyTime field to given value.

### HasStandbyTime

`func (o *VmwareStandbyObject) HasStandbyTime() bool`

HasStandbyTime returns a boolean if a field has been set.

### SetStandbyTimeNil

`func (o *VmwareStandbyObject) SetStandbyTimeNil(b bool)`

 SetStandbyTimeNil sets the value for StandbyTime to be an explicit nil

### UnsetStandbyTime
`func (o *VmwareStandbyObject) UnsetStandbyTime()`

UnsetStandbyTime ensures that no value is present for StandbyTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


