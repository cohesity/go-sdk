# FirewallEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Specifies the firewall profile attachments. | [optional] 
**Ipsets** | Pointer to [**[]FirewallIPSet**](FirewallIPSet.md) | Specifies the firewall ipsets. | [optional] 
**Profiles** | Pointer to [**[]FirewallProfile**](FirewallProfile.md) | Specifies the firewall profiles. | [optional] 

## Methods

### NewFirewallEntry

`func NewFirewallEntry() *FirewallEntry`

NewFirewallEntry instantiates a new FirewallEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallEntryWithDefaults

`func NewFirewallEntryWithDefaults() *FirewallEntry`

NewFirewallEntryWithDefaults instantiates a new FirewallEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *FirewallEntry) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *FirewallEntry) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *FirewallEntry) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *FirewallEntry) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetIpsets

`func (o *FirewallEntry) GetIpsets() []FirewallIPSet`

GetIpsets returns the Ipsets field if non-nil, zero value otherwise.

### GetIpsetsOk

`func (o *FirewallEntry) GetIpsetsOk() (*[]FirewallIPSet, bool)`

GetIpsetsOk returns a tuple with the Ipsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsets

`func (o *FirewallEntry) SetIpsets(v []FirewallIPSet)`

SetIpsets sets Ipsets field to given value.

### HasIpsets

`func (o *FirewallEntry) HasIpsets() bool`

HasIpsets returns a boolean if a field has been set.

### GetProfiles

`func (o *FirewallEntry) GetProfiles() []FirewallProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FirewallEntry) GetProfilesOk() (*[]FirewallProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FirewallEntry) SetProfiles(v []FirewallProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FirewallEntry) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


