# FirewallProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directions** | Pointer to **[]string** | Specifies the packet direction settings. | [optional] 
**Name** | **NullableString** | Specifies the name of the profile. | 
**Ports** | Pointer to **[]string** | Specifies the port along with the protocol settings. For example 22/tcp, 68/udp. | [optional] 

## Methods

### NewFirewallProfile

`func NewFirewallProfile(name NullableString, ) *FirewallProfile`

NewFirewallProfile instantiates a new FirewallProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallProfileWithDefaults

`func NewFirewallProfileWithDefaults() *FirewallProfile`

NewFirewallProfileWithDefaults instantiates a new FirewallProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirections

`func (o *FirewallProfile) GetDirections() []string`

GetDirections returns the Directions field if non-nil, zero value otherwise.

### GetDirectionsOk

`func (o *FirewallProfile) GetDirectionsOk() (*[]string, bool)`

GetDirectionsOk returns a tuple with the Directions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirections

`func (o *FirewallProfile) SetDirections(v []string)`

SetDirections sets Directions field to given value.

### HasDirections

`func (o *FirewallProfile) HasDirections() bool`

HasDirections returns a boolean if a field has been set.

### GetName

`func (o *FirewallProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallProfile) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FirewallProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FirewallProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPorts

`func (o *FirewallProfile) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *FirewallProfile) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *FirewallProfile) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *FirewallProfile) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


