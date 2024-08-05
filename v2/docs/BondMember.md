# BondMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSecondary** | Pointer to **NullableBool** | Specifies whether or not this is a active secondary. This is only valid in ActiveBackup bonding mode. | [optional] 
**LinkState** | Pointer to **NullableString** | Bond secondary link state. | [optional] 
**MacAddress** | Pointer to **NullableString** | MAC address of the bond secondary. | [optional] 
**Name** | Pointer to **NullableString** | Name of the bond secondary. | [optional] 
**Slot** | Pointer to **NullableString** | Slot information of the bond secondary. | [optional] 
**Speed** | Pointer to **NullableString** | Speed of the bond secondary. | [optional] 
**Stats** | Pointer to [**InterfaceStats**](InterfaceStats.md) |  | [optional] 
**UplinkSwitch** | Pointer to [**UplinkSwitch**](UplinkSwitch.md) |  | [optional] 

## Methods

### NewBondMember

`func NewBondMember() *BondMember`

NewBondMember instantiates a new BondMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondMemberWithDefaults

`func NewBondMemberWithDefaults() *BondMember`

NewBondMemberWithDefaults instantiates a new BondMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSecondary

`func (o *BondMember) GetActiveSecondary() bool`

GetActiveSecondary returns the ActiveSecondary field if non-nil, zero value otherwise.

### GetActiveSecondaryOk

`func (o *BondMember) GetActiveSecondaryOk() (*bool, bool)`

GetActiveSecondaryOk returns a tuple with the ActiveSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSecondary

`func (o *BondMember) SetActiveSecondary(v bool)`

SetActiveSecondary sets ActiveSecondary field to given value.

### HasActiveSecondary

`func (o *BondMember) HasActiveSecondary() bool`

HasActiveSecondary returns a boolean if a field has been set.

### SetActiveSecondaryNil

`func (o *BondMember) SetActiveSecondaryNil(b bool)`

 SetActiveSecondaryNil sets the value for ActiveSecondary to be an explicit nil

### UnsetActiveSecondary
`func (o *BondMember) UnsetActiveSecondary()`

UnsetActiveSecondary ensures that no value is present for ActiveSecondary, not even an explicit nil
### GetLinkState

`func (o *BondMember) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *BondMember) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *BondMember) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *BondMember) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### SetLinkStateNil

`func (o *BondMember) SetLinkStateNil(b bool)`

 SetLinkStateNil sets the value for LinkState to be an explicit nil

### UnsetLinkState
`func (o *BondMember) UnsetLinkState()`

UnsetLinkState ensures that no value is present for LinkState, not even an explicit nil
### GetMacAddress

`func (o *BondMember) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BondMember) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BondMember) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BondMember) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *BondMember) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *BondMember) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetName

`func (o *BondMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BondMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BondMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BondMember) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BondMember) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BondMember) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSlot

`func (o *BondMember) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *BondMember) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *BondMember) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *BondMember) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### SetSlotNil

`func (o *BondMember) SetSlotNil(b bool)`

 SetSlotNil sets the value for Slot to be an explicit nil

### UnsetSlot
`func (o *BondMember) UnsetSlot()`

UnsetSlot ensures that no value is present for Slot, not even an explicit nil
### GetSpeed

`func (o *BondMember) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *BondMember) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *BondMember) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *BondMember) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *BondMember) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *BondMember) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetStats

`func (o *BondMember) GetStats() InterfaceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *BondMember) GetStatsOk() (*InterfaceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *BondMember) SetStats(v InterfaceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *BondMember) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetUplinkSwitch

`func (o *BondMember) GetUplinkSwitch() UplinkSwitch`

GetUplinkSwitch returns the UplinkSwitch field if non-nil, zero value otherwise.

### GetUplinkSwitchOk

`func (o *BondMember) GetUplinkSwitchOk() (*UplinkSwitch, bool)`

GetUplinkSwitchOk returns a tuple with the UplinkSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSwitch

`func (o *BondMember) SetUplinkSwitch(v UplinkSwitch)`

SetUplinkSwitch sets UplinkSwitch field to given value.

### HasUplinkSwitch

`func (o *BondMember) HasUplinkSwitch() bool`

HasUplinkSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


