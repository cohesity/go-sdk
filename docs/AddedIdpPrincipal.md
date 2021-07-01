# AddedIdpPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the group or user was added to the Cohesity Cluster. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the name of the Idp where the referenced principal is stored. | [optional] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the group or user was last modified on the Cohesity Cluster. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the type of the referenced Idp principal. If &#39;kGroup&#39;, the referenced Idp principal is a group. If &#39;kUser&#39;, the referenced Idp principal is a user. &#39;kUser&#39; specifies a user object class. &#39;kGroup&#39; specifies a group object class. &#39;kComputer&#39; specifies a computer object class. &#39;kWellKnownPrincipal&#39; specifies a well known principal. | [optional] 
**PrincipalName** | Pointer to **NullableString** | Specifies the name of the Idp principal, that will be referenced by the group or user. The name of the Idp principal is used for naming the new group or user on the Cohesity Cluster. | [optional] 
**Restricted** | Pointer to **NullableBool** | Whether the principal is a restricted principal. A restricted principal can only view the objects he has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Array of Roles.  Specifies the Cohesity roles to associate with this user or group such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for this group or user. For example if the &#39;joe&#39; user is added for the Active Directory &#39;joe&#39; user principal and is associated with the Cohesity &#39;View&#39; role, &#39;joe&#39; can log in to the Cohesity Dashboard and has a read-only view of the data on the Cohesity Cluster. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique Security ID (SID) of the Idp principal associated with this group or user. | [optional] 

## Methods

### NewAddedIdpPrincipal

`func NewAddedIdpPrincipal() *AddedIdpPrincipal`

NewAddedIdpPrincipal instantiates a new AddedIdpPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedIdpPrincipalWithDefaults

`func NewAddedIdpPrincipalWithDefaults() *AddedIdpPrincipal`

NewAddedIdpPrincipalWithDefaults instantiates a new AddedIdpPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimeMsecs

`func (o *AddedIdpPrincipal) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *AddedIdpPrincipal) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *AddedIdpPrincipal) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *AddedIdpPrincipal) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *AddedIdpPrincipal) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *AddedIdpPrincipal) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetDomain

`func (o *AddedIdpPrincipal) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AddedIdpPrincipal) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AddedIdpPrincipal) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AddedIdpPrincipal) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *AddedIdpPrincipal) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *AddedIdpPrincipal) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *AddedIdpPrincipal) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *AddedIdpPrincipal) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *AddedIdpPrincipal) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *AddedIdpPrincipal) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *AddedIdpPrincipal) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *AddedIdpPrincipal) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetObjectClass

`func (o *AddedIdpPrincipal) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AddedIdpPrincipal) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AddedIdpPrincipal) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *AddedIdpPrincipal) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *AddedIdpPrincipal) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *AddedIdpPrincipal) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetPrincipalName

`func (o *AddedIdpPrincipal) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *AddedIdpPrincipal) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *AddedIdpPrincipal) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *AddedIdpPrincipal) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *AddedIdpPrincipal) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *AddedIdpPrincipal) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetRestricted

`func (o *AddedIdpPrincipal) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *AddedIdpPrincipal) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *AddedIdpPrincipal) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *AddedIdpPrincipal) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *AddedIdpPrincipal) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *AddedIdpPrincipal) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *AddedIdpPrincipal) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AddedIdpPrincipal) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AddedIdpPrincipal) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AddedIdpPrincipal) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *AddedIdpPrincipal) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *AddedIdpPrincipal) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSid

`func (o *AddedIdpPrincipal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AddedIdpPrincipal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AddedIdpPrincipal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AddedIdpPrincipal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AddedIdpPrincipal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AddedIdpPrincipal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


