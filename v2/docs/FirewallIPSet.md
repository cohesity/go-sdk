# FirewallIPSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the IP set. | [optional] 
**Subnets** | Pointer to **[]string** | Specifies the subnets in the IP set. | [optional] 

## Methods

### NewFirewallIPSet

`func NewFirewallIPSet() *FirewallIPSet`

NewFirewallIPSet instantiates a new FirewallIPSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallIPSetWithDefaults

`func NewFirewallIPSetWithDefaults() *FirewallIPSet`

NewFirewallIPSetWithDefaults instantiates a new FirewallIPSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallIPSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallIPSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallIPSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallIPSet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FirewallIPSet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FirewallIPSet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubnets

`func (o *FirewallIPSet) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *FirewallIPSet) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *FirewallIPSet) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *FirewallIPSet) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


