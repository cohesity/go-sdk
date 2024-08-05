# FirewallProfileParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the action. | 
**Description** | Pointer to **NullableString** | Specifies a description for the profile attachments. | [optional] 
**Direction** | Pointer to **NullableString** | Specifies the packet direction settings. | [optional] 
**InterfaceGroups** | Pointer to **[]string** | Specifies the network interface groups. | [optional] 
**Name** | **NullableString** | Specifies the name of the profile. | 
**Ports** | Pointer to **[]string** | Specifies the port along with the protocol settings. | [optional] 
**Subnets** | Pointer to **[]string** | Specifies the subnets. | [optional] 

## Methods

### NewFirewallProfileParams

`func NewFirewallProfileParams(action NullableString, name NullableString, ) *FirewallProfileParams`

NewFirewallProfileParams instantiates a new FirewallProfileParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallProfileParamsWithDefaults

`func NewFirewallProfileParamsWithDefaults() *FirewallProfileParams`

NewFirewallProfileParamsWithDefaults instantiates a new FirewallProfileParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FirewallProfileParams) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallProfileParams) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallProfileParams) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *FirewallProfileParams) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *FirewallProfileParams) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDescription

`func (o *FirewallProfileParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallProfileParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallProfileParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallProfileParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FirewallProfileParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FirewallProfileParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDirection

`func (o *FirewallProfileParams) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallProfileParams) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallProfileParams) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FirewallProfileParams) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *FirewallProfileParams) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *FirewallProfileParams) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetInterfaceGroups

`func (o *FirewallProfileParams) GetInterfaceGroups() []string`

GetInterfaceGroups returns the InterfaceGroups field if non-nil, zero value otherwise.

### GetInterfaceGroupsOk

`func (o *FirewallProfileParams) GetInterfaceGroupsOk() (*[]string, bool)`

GetInterfaceGroupsOk returns a tuple with the InterfaceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceGroups

`func (o *FirewallProfileParams) SetInterfaceGroups(v []string)`

SetInterfaceGroups sets InterfaceGroups field to given value.

### HasInterfaceGroups

`func (o *FirewallProfileParams) HasInterfaceGroups() bool`

HasInterfaceGroups returns a boolean if a field has been set.

### GetName

`func (o *FirewallProfileParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallProfileParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallProfileParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FirewallProfileParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FirewallProfileParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPorts

`func (o *FirewallProfileParams) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *FirewallProfileParams) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *FirewallProfileParams) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *FirewallProfileParams) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *FirewallProfileParams) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *FirewallProfileParams) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetSubnets

`func (o *FirewallProfileParams) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *FirewallProfileParams) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *FirewallProfileParams) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *FirewallProfileParams) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


