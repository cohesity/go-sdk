# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description for the attachment. | [optional] 
**InterfaceGroups** | Pointer to **[]string** | Specifies the network interface groups. | [optional] 
**Interfaces** | Pointer to **[]string** | Specifies the network interfaces | [optional] 
**IpsetNames** | Pointer to **[]string** | Specifies the ip sets. | [optional] 
**IsImplicit** | Pointer to **NullableBool** |  | [optional] [readonly] 
**Profile** | Pointer to **NullableString** | Specifies the firewall profile. | [optional] 
**Subnets** | Pointer to **[]string** | Specifies the subnets. | [optional] 

## Methods

### NewAttachment

`func NewAttachment() *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Attachment) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Attachment) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Attachment) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Attachment) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *Attachment) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *Attachment) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDescription

`func (o *Attachment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Attachment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Attachment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Attachment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Attachment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Attachment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInterfaceGroups

`func (o *Attachment) GetInterfaceGroups() []string`

GetInterfaceGroups returns the InterfaceGroups field if non-nil, zero value otherwise.

### GetInterfaceGroupsOk

`func (o *Attachment) GetInterfaceGroupsOk() (*[]string, bool)`

GetInterfaceGroupsOk returns a tuple with the InterfaceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceGroups

`func (o *Attachment) SetInterfaceGroups(v []string)`

SetInterfaceGroups sets InterfaceGroups field to given value.

### HasInterfaceGroups

`func (o *Attachment) HasInterfaceGroups() bool`

HasInterfaceGroups returns a boolean if a field has been set.

### GetInterfaces

`func (o *Attachment) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Attachment) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Attachment) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Attachment) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIpsetNames

`func (o *Attachment) GetIpsetNames() []string`

GetIpsetNames returns the IpsetNames field if non-nil, zero value otherwise.

### GetIpsetNamesOk

`func (o *Attachment) GetIpsetNamesOk() (*[]string, bool)`

GetIpsetNamesOk returns a tuple with the IpsetNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsetNames

`func (o *Attachment) SetIpsetNames(v []string)`

SetIpsetNames sets IpsetNames field to given value.

### HasIpsetNames

`func (o *Attachment) HasIpsetNames() bool`

HasIpsetNames returns a boolean if a field has been set.

### GetIsImplicit

`func (o *Attachment) GetIsImplicit() bool`

GetIsImplicit returns the IsImplicit field if non-nil, zero value otherwise.

### GetIsImplicitOk

`func (o *Attachment) GetIsImplicitOk() (*bool, bool)`

GetIsImplicitOk returns a tuple with the IsImplicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImplicit

`func (o *Attachment) SetIsImplicit(v bool)`

SetIsImplicit sets IsImplicit field to given value.

### HasIsImplicit

`func (o *Attachment) HasIsImplicit() bool`

HasIsImplicit returns a boolean if a field has been set.

### SetIsImplicitNil

`func (o *Attachment) SetIsImplicitNil(b bool)`

 SetIsImplicitNil sets the value for IsImplicit to be an explicit nil

### UnsetIsImplicit
`func (o *Attachment) UnsetIsImplicit()`

UnsetIsImplicit ensures that no value is present for IsImplicit, not even an explicit nil
### GetProfile

`func (o *Attachment) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Attachment) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Attachment) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Attachment) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *Attachment) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *Attachment) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetSubnets

`func (o *Attachment) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *Attachment) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *Attachment) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *Attachment) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


