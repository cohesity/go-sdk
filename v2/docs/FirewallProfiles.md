# FirewallProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]FirewallProfile**](FirewallProfile.md) |  | [optional] 

## Methods

### NewFirewallProfiles

`func NewFirewallProfiles() *FirewallProfiles`

NewFirewallProfiles instantiates a new FirewallProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallProfilesWithDefaults

`func NewFirewallProfilesWithDefaults() *FirewallProfiles`

NewFirewallProfilesWithDefaults instantiates a new FirewallProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *FirewallProfiles) GetProfiles() []FirewallProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FirewallProfiles) GetProfilesOk() (*[]FirewallProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FirewallProfiles) SetProfiles(v []FirewallProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FirewallProfiles) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


