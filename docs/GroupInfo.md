# GroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies domain name of the user. | [optional] 
**GroupName** | Pointer to **NullableString** | Specifies group name of the group. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies unique Security ID (SID) of the user. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants to which the group belongs to. | [optional] 

## Methods

### NewGroupInfo

`func NewGroupInfo() *GroupInfo`

NewGroupInfo instantiates a new GroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInfoWithDefaults

`func NewGroupInfoWithDefaults() *GroupInfo`

NewGroupInfoWithDefaults instantiates a new GroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *GroupInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GroupInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GroupInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GroupInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GroupInfo) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GroupInfo) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetGroupName

`func (o *GroupInfo) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupInfo) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupInfo) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupInfo) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *GroupInfo) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *GroupInfo) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetSid

`func (o *GroupInfo) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *GroupInfo) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *GroupInfo) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *GroupInfo) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *GroupInfo) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *GroupInfo) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTenantIds

`func (o *GroupInfo) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *GroupInfo) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *GroupInfo) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *GroupInfo) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *GroupInfo) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *GroupInfo) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


