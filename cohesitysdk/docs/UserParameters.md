# UserParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalGroupNames** | Pointer to **[]string** | Array of Additional Groups.  Specifies the names of additional groups this User may belong to. | [optional] 
**ClusterIdentifiers** | Pointer to [**[]ClusterIdentifier**](ClusterIdentifier.md) | Specifies the list of clusters this user has access to. If this is not specified, access will be granted to all clusters. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description about the user. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the fully qualified domain name (FQDN) of an Active Directory or LOCAL for the default LOCAL domain on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain. | [optional] 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user becomes effective. Until that time, the user cannot log in. | [optional] 
**EmailAddress** | Pointer to **NullableString** | Specifies the email address of the user. | [optional] 
**ExpiredTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user becomes expired. After that, the user cannot log in. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password of this user. | [optional] 
**PrimaryGroupName** | Pointer to **NullableString** | Specifies the name of the primary group of this User. | [optional] 
**PrivilegeIds** | Pointer to **[]string** | Array of Privileges.  Specifies the Cohesity privileges from the roles. This will be populated based on the union of all privileges in roles. Type for unique privilege Id values. All below enum values specify a value for all uniquely defined privileges in Cohesity. | [optional] 
**Restricted** | Pointer to **NullableBool** | Whether the user is a restricted user. A restricted user can only view the objects he has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Array of Roles.  Specifies the Cohesity roles to associate with the user such as such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for this user. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the login name of the user. | [optional] 

## Methods

### NewUserParameters

`func NewUserParameters() *UserParameters`

NewUserParameters instantiates a new UserParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserParametersWithDefaults

`func NewUserParametersWithDefaults() *UserParameters`

NewUserParametersWithDefaults instantiates a new UserParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalGroupNames

`func (o *UserParameters) GetAdditionalGroupNames() []string`

GetAdditionalGroupNames returns the AdditionalGroupNames field if non-nil, zero value otherwise.

### GetAdditionalGroupNamesOk

`func (o *UserParameters) GetAdditionalGroupNamesOk() (*[]string, bool)`

GetAdditionalGroupNamesOk returns a tuple with the AdditionalGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGroupNames

`func (o *UserParameters) SetAdditionalGroupNames(v []string)`

SetAdditionalGroupNames sets AdditionalGroupNames field to given value.

### HasAdditionalGroupNames

`func (o *UserParameters) HasAdditionalGroupNames() bool`

HasAdditionalGroupNames returns a boolean if a field has been set.

### SetAdditionalGroupNamesNil

`func (o *UserParameters) SetAdditionalGroupNamesNil(b bool)`

 SetAdditionalGroupNamesNil sets the value for AdditionalGroupNames to be an explicit nil

### UnsetAdditionalGroupNames
`func (o *UserParameters) UnsetAdditionalGroupNames()`

UnsetAdditionalGroupNames ensures that no value is present for AdditionalGroupNames, not even an explicit nil
### GetClusterIdentifiers

`func (o *UserParameters) GetClusterIdentifiers() []ClusterIdentifier`

GetClusterIdentifiers returns the ClusterIdentifiers field if non-nil, zero value otherwise.

### GetClusterIdentifiersOk

`func (o *UserParameters) GetClusterIdentifiersOk() (*[]ClusterIdentifier, bool)`

GetClusterIdentifiersOk returns a tuple with the ClusterIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifiers

`func (o *UserParameters) SetClusterIdentifiers(v []ClusterIdentifier)`

SetClusterIdentifiers sets ClusterIdentifiers field to given value.

### HasClusterIdentifiers

`func (o *UserParameters) HasClusterIdentifiers() bool`

HasClusterIdentifiers returns a boolean if a field has been set.

### SetClusterIdentifiersNil

`func (o *UserParameters) SetClusterIdentifiersNil(b bool)`

 SetClusterIdentifiersNil sets the value for ClusterIdentifiers to be an explicit nil

### UnsetClusterIdentifiers
`func (o *UserParameters) UnsetClusterIdentifiers()`

UnsetClusterIdentifiers ensures that no value is present for ClusterIdentifiers, not even an explicit nil
### GetDescription

`func (o *UserParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UserParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *UserParameters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserParameters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserParameters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserParameters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *UserParameters) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UserParameters) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEffectiveTimeMsecs

`func (o *UserParameters) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *UserParameters) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *UserParameters) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *UserParameters) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *UserParameters) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *UserParameters) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetEmailAddress

`func (o *UserParameters) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserParameters) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserParameters) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UserParameters) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *UserParameters) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *UserParameters) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetExpiredTimeMsecs

`func (o *UserParameters) GetExpiredTimeMsecs() int64`

GetExpiredTimeMsecs returns the ExpiredTimeMsecs field if non-nil, zero value otherwise.

### GetExpiredTimeMsecsOk

`func (o *UserParameters) GetExpiredTimeMsecsOk() (*int64, bool)`

GetExpiredTimeMsecsOk returns a tuple with the ExpiredTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimeMsecs

`func (o *UserParameters) SetExpiredTimeMsecs(v int64)`

SetExpiredTimeMsecs sets ExpiredTimeMsecs field to given value.

### HasExpiredTimeMsecs

`func (o *UserParameters) HasExpiredTimeMsecs() bool`

HasExpiredTimeMsecs returns a boolean if a field has been set.

### SetExpiredTimeMsecsNil

`func (o *UserParameters) SetExpiredTimeMsecsNil(b bool)`

 SetExpiredTimeMsecsNil sets the value for ExpiredTimeMsecs to be an explicit nil

### UnsetExpiredTimeMsecs
`func (o *UserParameters) UnsetExpiredTimeMsecs()`

UnsetExpiredTimeMsecs ensures that no value is present for ExpiredTimeMsecs, not even an explicit nil
### GetPassword

`func (o *UserParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UserParameters) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UserParameters) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPrimaryGroupName

`func (o *UserParameters) GetPrimaryGroupName() string`

GetPrimaryGroupName returns the PrimaryGroupName field if non-nil, zero value otherwise.

### GetPrimaryGroupNameOk

`func (o *UserParameters) GetPrimaryGroupNameOk() (*string, bool)`

GetPrimaryGroupNameOk returns a tuple with the PrimaryGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroupName

`func (o *UserParameters) SetPrimaryGroupName(v string)`

SetPrimaryGroupName sets PrimaryGroupName field to given value.

### HasPrimaryGroupName

`func (o *UserParameters) HasPrimaryGroupName() bool`

HasPrimaryGroupName returns a boolean if a field has been set.

### SetPrimaryGroupNameNil

`func (o *UserParameters) SetPrimaryGroupNameNil(b bool)`

 SetPrimaryGroupNameNil sets the value for PrimaryGroupName to be an explicit nil

### UnsetPrimaryGroupName
`func (o *UserParameters) UnsetPrimaryGroupName()`

UnsetPrimaryGroupName ensures that no value is present for PrimaryGroupName, not even an explicit nil
### GetPrivilegeIds

`func (o *UserParameters) GetPrivilegeIds() []string`

GetPrivilegeIds returns the PrivilegeIds field if non-nil, zero value otherwise.

### GetPrivilegeIdsOk

`func (o *UserParameters) GetPrivilegeIdsOk() (*[]string, bool)`

GetPrivilegeIdsOk returns a tuple with the PrivilegeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeIds

`func (o *UserParameters) SetPrivilegeIds(v []string)`

SetPrivilegeIds sets PrivilegeIds field to given value.

### HasPrivilegeIds

`func (o *UserParameters) HasPrivilegeIds() bool`

HasPrivilegeIds returns a boolean if a field has been set.

### SetPrivilegeIdsNil

`func (o *UserParameters) SetPrivilegeIdsNil(b bool)`

 SetPrivilegeIdsNil sets the value for PrivilegeIds to be an explicit nil

### UnsetPrivilegeIds
`func (o *UserParameters) UnsetPrivilegeIds()`

UnsetPrivilegeIds ensures that no value is present for PrivilegeIds, not even an explicit nil
### GetRestricted

`func (o *UserParameters) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *UserParameters) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *UserParameters) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *UserParameters) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *UserParameters) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *UserParameters) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *UserParameters) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserParameters) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserParameters) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserParameters) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UserParameters) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UserParameters) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetUsername

`func (o *UserParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserParameters) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserParameters) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


