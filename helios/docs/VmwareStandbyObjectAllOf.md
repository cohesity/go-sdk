# VmwareStandbyObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdpStandbyStatus** | Pointer to **NullableString** | Specifies the current status of the standby object protected using continuous data protection policy. | [optional] 
**GuestId** | Pointer to **NullableString** | Specifies the guest ID(OS) of the standby VM for the corresponding backup object. | [optional] [readonly] 
**StandbyMOref** | Pointer to [**MOref**](MOref.md) |  | [optional] 
**StandbyTime** | Pointer to **NullableInt64** | Specifies the time till which the standby object has been hydrated for the corresponding backup object. | [optional] 

## Methods

### NewVmwareStandbyObjectAllOf

`func NewVmwareStandbyObjectAllOf() *VmwareStandbyObjectAllOf`

NewVmwareStandbyObjectAllOf instantiates a new VmwareStandbyObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareStandbyObjectAllOfWithDefaults

`func NewVmwareStandbyObjectAllOfWithDefaults() *VmwareStandbyObjectAllOf`

NewVmwareStandbyObjectAllOfWithDefaults instantiates a new VmwareStandbyObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdpStandbyStatus

`func (o *VmwareStandbyObjectAllOf) GetCdpStandbyStatus() string`

GetCdpStandbyStatus returns the CdpStandbyStatus field if non-nil, zero value otherwise.

### GetCdpStandbyStatusOk

`func (o *VmwareStandbyObjectAllOf) GetCdpStandbyStatusOk() (*string, bool)`

GetCdpStandbyStatusOk returns a tuple with the CdpStandbyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpStandbyStatus

`func (o *VmwareStandbyObjectAllOf) SetCdpStandbyStatus(v string)`

SetCdpStandbyStatus sets CdpStandbyStatus field to given value.

### HasCdpStandbyStatus

`func (o *VmwareStandbyObjectAllOf) HasCdpStandbyStatus() bool`

HasCdpStandbyStatus returns a boolean if a field has been set.

### SetCdpStandbyStatusNil

`func (o *VmwareStandbyObjectAllOf) SetCdpStandbyStatusNil(b bool)`

 SetCdpStandbyStatusNil sets the value for CdpStandbyStatus to be an explicit nil

### UnsetCdpStandbyStatus
`func (o *VmwareStandbyObjectAllOf) UnsetCdpStandbyStatus()`

UnsetCdpStandbyStatus ensures that no value is present for CdpStandbyStatus, not even an explicit nil
### GetGuestId

`func (o *VmwareStandbyObjectAllOf) GetGuestId() string`

GetGuestId returns the GuestId field if non-nil, zero value otherwise.

### GetGuestIdOk

`func (o *VmwareStandbyObjectAllOf) GetGuestIdOk() (*string, bool)`

GetGuestIdOk returns a tuple with the GuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestId

`func (o *VmwareStandbyObjectAllOf) SetGuestId(v string)`

SetGuestId sets GuestId field to given value.

### HasGuestId

`func (o *VmwareStandbyObjectAllOf) HasGuestId() bool`

HasGuestId returns a boolean if a field has been set.

### SetGuestIdNil

`func (o *VmwareStandbyObjectAllOf) SetGuestIdNil(b bool)`

 SetGuestIdNil sets the value for GuestId to be an explicit nil

### UnsetGuestId
`func (o *VmwareStandbyObjectAllOf) UnsetGuestId()`

UnsetGuestId ensures that no value is present for GuestId, not even an explicit nil
### GetStandbyMOref

`func (o *VmwareStandbyObjectAllOf) GetStandbyMOref() MOref`

GetStandbyMOref returns the StandbyMOref field if non-nil, zero value otherwise.

### GetStandbyMOrefOk

`func (o *VmwareStandbyObjectAllOf) GetStandbyMOrefOk() (*MOref, bool)`

GetStandbyMOrefOk returns a tuple with the StandbyMOref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyMOref

`func (o *VmwareStandbyObjectAllOf) SetStandbyMOref(v MOref)`

SetStandbyMOref sets StandbyMOref field to given value.

### HasStandbyMOref

`func (o *VmwareStandbyObjectAllOf) HasStandbyMOref() bool`

HasStandbyMOref returns a boolean if a field has been set.

### GetStandbyTime

`func (o *VmwareStandbyObjectAllOf) GetStandbyTime() int64`

GetStandbyTime returns the StandbyTime field if non-nil, zero value otherwise.

### GetStandbyTimeOk

`func (o *VmwareStandbyObjectAllOf) GetStandbyTimeOk() (*int64, bool)`

GetStandbyTimeOk returns a tuple with the StandbyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyTime

`func (o *VmwareStandbyObjectAllOf) SetStandbyTime(v int64)`

SetStandbyTime sets StandbyTime field to given value.

### HasStandbyTime

`func (o *VmwareStandbyObjectAllOf) HasStandbyTime() bool`

HasStandbyTime returns a boolean if a field has been set.

### SetStandbyTimeNil

`func (o *VmwareStandbyObjectAllOf) SetStandbyTimeNil(b bool)`

 SetStandbyTimeNil sets the value for StandbyTime to be an explicit nil

### UnsetStandbyTime
`func (o *VmwareStandbyObjectAllOf) UnsetStandbyTime()`

UnsetStandbyTime ensures that no value is present for StandbyTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


