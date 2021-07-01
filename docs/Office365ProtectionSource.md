# Office365ProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the Office 365 entity. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the office 365 entity. | [optional] 
**PrimarySMTPAddress** | Pointer to **NullableString** | Specifies the SMTP address for the Outlook source. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the Office 365 entity. Specifies the type of Office 365 entity &#39;kDomain&#39; indicates the O365 domain through which authentication occurs. &#39;kOutlook&#39; indicates the Exchange online entities. &#39;kMailbox&#39; indicates the user&#39;s mailbox account. &#39;kUsers&#39; indicates the container for User entities. &#39;kGroups&#39; indicates the container for Group entities. &#39;kSites&#39; indicates the container for Site entities. &#39;kUser&#39; indicates an Office365 User entity. &#39;kGroup&#39; indicates an Office365 Group entity. &#39;kSite&#39; indicates an Office365 SharePoint Site entity. | [optional] 
**UserInfo** | Pointer to [**Office365UserInfo**](Office365UserInfo.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID of the Office 365 entity. | [optional] 

## Methods

### NewOffice365ProtectionSource

`func NewOffice365ProtectionSource() *Office365ProtectionSource`

NewOffice365ProtectionSource instantiates a new Office365ProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ProtectionSourceWithDefaults

`func NewOffice365ProtectionSourceWithDefaults() *Office365ProtectionSource`

NewOffice365ProtectionSourceWithDefaults instantiates a new Office365ProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Office365ProtectionSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Office365ProtectionSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Office365ProtectionSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Office365ProtectionSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Office365ProtectionSource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Office365ProtectionSource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *Office365ProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Office365ProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Office365ProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Office365ProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Office365ProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Office365ProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrimarySMTPAddress

`func (o *Office365ProtectionSource) GetPrimarySMTPAddress() string`

GetPrimarySMTPAddress returns the PrimarySMTPAddress field if non-nil, zero value otherwise.

### GetPrimarySMTPAddressOk

`func (o *Office365ProtectionSource) GetPrimarySMTPAddressOk() (*string, bool)`

GetPrimarySMTPAddressOk returns a tuple with the PrimarySMTPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySMTPAddress

`func (o *Office365ProtectionSource) SetPrimarySMTPAddress(v string)`

SetPrimarySMTPAddress sets PrimarySMTPAddress field to given value.

### HasPrimarySMTPAddress

`func (o *Office365ProtectionSource) HasPrimarySMTPAddress() bool`

HasPrimarySMTPAddress returns a boolean if a field has been set.

### SetPrimarySMTPAddressNil

`func (o *Office365ProtectionSource) SetPrimarySMTPAddressNil(b bool)`

 SetPrimarySMTPAddressNil sets the value for PrimarySMTPAddress to be an explicit nil

### UnsetPrimarySMTPAddress
`func (o *Office365ProtectionSource) UnsetPrimarySMTPAddress()`

UnsetPrimarySMTPAddress ensures that no value is present for PrimarySMTPAddress, not even an explicit nil
### GetType

`func (o *Office365ProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Office365ProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Office365ProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Office365ProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Office365ProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Office365ProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserInfo

`func (o *Office365ProtectionSource) GetUserInfo() Office365UserInfo`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *Office365ProtectionSource) GetUserInfoOk() (*Office365UserInfo, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *Office365ProtectionSource) SetUserInfo(v Office365UserInfo)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *Office365ProtectionSource) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetUuid

`func (o *Office365ProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Office365ProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Office365ProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Office365ProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Office365ProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Office365ProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


