# AddedActiveDirectoryPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies a description about the user or group. | [optional] 
**DomainName** | **NullableString** | Specifies the domain of the Active Directory where the referenced principal is stored. | 
**Name** | **NullableString** | Specifies the name of the Active Directory principal, that will be referenced by the group or user. The name of the Active Directory principal is used for naming the new group or user on the Cohesity Cluster. | 
**ObjectClass** | **NullableString** | Specifies the type of Active Directory principal.&lt;br&gt; &#39;User&#39; specifies a user object class.&lt;br&gt; &#39;Group&#39; specifies a group object class.&lt;br&gt; &#39;ServiceAccount&#39; specifies a service account object class. | 
**Restricted** | Pointer to **NullableBool** | Whether the principal is a restricted principal. A restricted principal can only view the objects he has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with this user or group such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for this group or user. For example if the &#39;joe&#39; user is added form the Active Directory and is associated with the Cohesity &#39;View&#39; role,&#39;joe&#39; can log in to the Cohesity Dashboard and has a read-only view of the data on the Cohesity Cluster. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the group or user was added to the Cohesity Cluster. | [optional] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the group or user was last modified on the Cohesity Cluster. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique Security ID (SID) of the Active Directory principal associated with this group or user. | [optional] 

## Methods

### NewAddedActiveDirectoryPrincipal

`func NewAddedActiveDirectoryPrincipal(domainName NullableString, name NullableString, objectClass NullableString, ) *AddedActiveDirectoryPrincipal`

NewAddedActiveDirectoryPrincipal instantiates a new AddedActiveDirectoryPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedActiveDirectoryPrincipalWithDefaults

`func NewAddedActiveDirectoryPrincipalWithDefaults() *AddedActiveDirectoryPrincipal`

NewAddedActiveDirectoryPrincipalWithDefaults instantiates a new AddedActiveDirectoryPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AddedActiveDirectoryPrincipal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddedActiveDirectoryPrincipal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddedActiveDirectoryPrincipal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddedActiveDirectoryPrincipal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddedActiveDirectoryPrincipal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddedActiveDirectoryPrincipal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomainName

`func (o *AddedActiveDirectoryPrincipal) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddedActiveDirectoryPrincipal) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddedActiveDirectoryPrincipal) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### SetDomainNameNil

`func (o *AddedActiveDirectoryPrincipal) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *AddedActiveDirectoryPrincipal) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetName

`func (o *AddedActiveDirectoryPrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddedActiveDirectoryPrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddedActiveDirectoryPrincipal) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AddedActiveDirectoryPrincipal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddedActiveDirectoryPrincipal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectClass

`func (o *AddedActiveDirectoryPrincipal) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AddedActiveDirectoryPrincipal) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AddedActiveDirectoryPrincipal) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.


### SetObjectClassNil

`func (o *AddedActiveDirectoryPrincipal) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *AddedActiveDirectoryPrincipal) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetRestricted

`func (o *AddedActiveDirectoryPrincipal) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *AddedActiveDirectoryPrincipal) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *AddedActiveDirectoryPrincipal) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *AddedActiveDirectoryPrincipal) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *AddedActiveDirectoryPrincipal) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *AddedActiveDirectoryPrincipal) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *AddedActiveDirectoryPrincipal) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AddedActiveDirectoryPrincipal) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AddedActiveDirectoryPrincipal) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AddedActiveDirectoryPrincipal) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *AddedActiveDirectoryPrincipal) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *AddedActiveDirectoryPrincipal) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *AddedActiveDirectoryPrincipal) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *AddedActiveDirectoryPrincipal) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *AddedActiveDirectoryPrincipal) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *AddedActiveDirectoryPrincipal) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *AddedActiveDirectoryPrincipal) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *AddedActiveDirectoryPrincipal) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *AddedActiveDirectoryPrincipal) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *AddedActiveDirectoryPrincipal) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *AddedActiveDirectoryPrincipal) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *AddedActiveDirectoryPrincipal) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *AddedActiveDirectoryPrincipal) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *AddedActiveDirectoryPrincipal) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetSid

`func (o *AddedActiveDirectoryPrincipal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AddedActiveDirectoryPrincipal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AddedActiveDirectoryPrincipal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AddedActiveDirectoryPrincipal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AddedActiveDirectoryPrincipal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AddedActiveDirectoryPrincipal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


