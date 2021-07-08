# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalGroupNames** | Pointer to **[]string** | Array of Additional Groups.  Specifies the names of additional groups this User may belong to. | [optional] 
**AuthenticationType** | Pointer to **NullableString** | Specifies the authentication type of the user. &#39;kAuthLocal&#39; implies authenticated user is a local user. &#39;kAuthAd&#39; implies authenticated user is an Active Directory user. &#39;kAuthSalesforce&#39; implies authenticated user is a Salesforce user. &#39;kAuthGoogle&#39; implies authenticated user is a Google user. &#39;kAuthSso&#39; implies authenticated user is an SSO user. | [optional] 
**ClusterIdentifiers** | Pointer to [**[]ClusterIdentifier**](ClusterIdentifier.md) | Specifies the list of clusters this user has access to. If this is not specified, access will be granted to all clusters. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was created on the Cohesity Cluster. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description about the user. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the fully qualified domain name (FQDN) of an Active Directory or LOCAL for the default LOCAL domain on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain. | [optional] 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user becomes effective. Until that time, the user cannot log in. | [optional] 
**EmailAddress** | Pointer to **NullableString** | Specifies the email address of the user. | [optional] 
**ExpiredTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user becomes expired. After that, the user cannot log in. | [optional] 
**ForcePasswordChange** | Pointer to **NullableBool** | Specifies whether to force user to change password. | [optional] 
**GoogleAccount** | Pointer to [**GoogleAccountInfo**](GoogleAccountInfo.md) |  | [optional] 
**GroupRoles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with the user&#39; group. These roles can only be edited from group. | [optional] [readonly] 
**IdpUserInfo** | Pointer to [**IdpUserInfo**](IdpUserInfo.md) |  | [optional] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was last modified on the Cohesity Cluster. | [optional] 
**OrgMembership** | Pointer to [**[]TenantConfig**](TenantConfig.md) | OrgMembership contains the list of all available tenantIds for this user to switch to. Only when creating the session user, this field is populated on the fly. We discover the tenantIds from various groups assigned to the users. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password of this user. | [optional] 
**Preferences** | Pointer to [**Preferences**](Preferences.md) |  | [optional] 
**PrimaryGroupName** | Pointer to **NullableString** | Specifies the name of the primary group of this User. | [optional] 
**PrivilegeIds** | Pointer to **[]string** | Array of Privileges.  Specifies the Cohesity privileges from the roles. This will be populated based on the union of all privileges in roles. Type for unique privilege Id values. All below enum values specify a value for all uniquely defined privileges in Cohesity. | [optional] 
**Restricted** | Pointer to **NullableBool** | Whether the user is a restricted user. A restricted user can only view the objects he has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Array of Roles.  Specifies the Cohesity roles to associate with the user such as such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for this user. | [optional] 
**S3AccessKeyId** | Pointer to **NullableString** | Specifies the S3 Account Access Key ID. | [optional] 
**S3AccountId** | Pointer to **NullableString** | Specifies the S3 Account Canonical User ID. | [optional] 
**S3SecretKey** | Pointer to **NullableString** | Specifies the S3 Account Secret Key. | [optional] 
**SalesforceAccount** | Pointer to [**SalesforceAccountInfo**](SalesforceAccountInfo.md) |  | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique Security ID (SID) of the user. This field is mandatory in modifying user. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the effective Tenant ID of the user. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the login name of the user. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalGroupNames

`func (o *User) GetAdditionalGroupNames() []string`

GetAdditionalGroupNames returns the AdditionalGroupNames field if non-nil, zero value otherwise.

### GetAdditionalGroupNamesOk

`func (o *User) GetAdditionalGroupNamesOk() (*[]string, bool)`

GetAdditionalGroupNamesOk returns a tuple with the AdditionalGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGroupNames

`func (o *User) SetAdditionalGroupNames(v []string)`

SetAdditionalGroupNames sets AdditionalGroupNames field to given value.

### HasAdditionalGroupNames

`func (o *User) HasAdditionalGroupNames() bool`

HasAdditionalGroupNames returns a boolean if a field has been set.

### SetAdditionalGroupNamesNil

`func (o *User) SetAdditionalGroupNamesNil(b bool)`

 SetAdditionalGroupNamesNil sets the value for AdditionalGroupNames to be an explicit nil

### UnsetAdditionalGroupNames
`func (o *User) UnsetAdditionalGroupNames()`

UnsetAdditionalGroupNames ensures that no value is present for AdditionalGroupNames, not even an explicit nil
### GetAuthenticationType

`func (o *User) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *User) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *User) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *User) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *User) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *User) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetClusterIdentifiers

`func (o *User) GetClusterIdentifiers() []ClusterIdentifier`

GetClusterIdentifiers returns the ClusterIdentifiers field if non-nil, zero value otherwise.

### GetClusterIdentifiersOk

`func (o *User) GetClusterIdentifiersOk() (*[]ClusterIdentifier, bool)`

GetClusterIdentifiersOk returns a tuple with the ClusterIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifiers

`func (o *User) SetClusterIdentifiers(v []ClusterIdentifier)`

SetClusterIdentifiers sets ClusterIdentifiers field to given value.

### HasClusterIdentifiers

`func (o *User) HasClusterIdentifiers() bool`

HasClusterIdentifiers returns a boolean if a field has been set.

### SetClusterIdentifiersNil

`func (o *User) SetClusterIdentifiersNil(b bool)`

 SetClusterIdentifiersNil sets the value for ClusterIdentifiers to be an explicit nil

### UnsetClusterIdentifiers
`func (o *User) UnsetClusterIdentifiers()`

UnsetClusterIdentifiers ensures that no value is present for ClusterIdentifiers, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *User) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *User) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *User) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *User) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *User) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *User) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *User) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *User) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *User) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *User) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *User) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *User) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *User) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *User) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEffectiveTimeMsecs

`func (o *User) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *User) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *User) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *User) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *User) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *User) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetEmailAddress

`func (o *User) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *User) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *User) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *User) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *User) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *User) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetExpiredTimeMsecs

`func (o *User) GetExpiredTimeMsecs() int64`

GetExpiredTimeMsecs returns the ExpiredTimeMsecs field if non-nil, zero value otherwise.

### GetExpiredTimeMsecsOk

`func (o *User) GetExpiredTimeMsecsOk() (*int64, bool)`

GetExpiredTimeMsecsOk returns a tuple with the ExpiredTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimeMsecs

`func (o *User) SetExpiredTimeMsecs(v int64)`

SetExpiredTimeMsecs sets ExpiredTimeMsecs field to given value.

### HasExpiredTimeMsecs

`func (o *User) HasExpiredTimeMsecs() bool`

HasExpiredTimeMsecs returns a boolean if a field has been set.

### SetExpiredTimeMsecsNil

`func (o *User) SetExpiredTimeMsecsNil(b bool)`

 SetExpiredTimeMsecsNil sets the value for ExpiredTimeMsecs to be an explicit nil

### UnsetExpiredTimeMsecs
`func (o *User) UnsetExpiredTimeMsecs()`

UnsetExpiredTimeMsecs ensures that no value is present for ExpiredTimeMsecs, not even an explicit nil
### GetForcePasswordChange

`func (o *User) GetForcePasswordChange() bool`

GetForcePasswordChange returns the ForcePasswordChange field if non-nil, zero value otherwise.

### GetForcePasswordChangeOk

`func (o *User) GetForcePasswordChangeOk() (*bool, bool)`

GetForcePasswordChangeOk returns a tuple with the ForcePasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePasswordChange

`func (o *User) SetForcePasswordChange(v bool)`

SetForcePasswordChange sets ForcePasswordChange field to given value.

### HasForcePasswordChange

`func (o *User) HasForcePasswordChange() bool`

HasForcePasswordChange returns a boolean if a field has been set.

### SetForcePasswordChangeNil

`func (o *User) SetForcePasswordChangeNil(b bool)`

 SetForcePasswordChangeNil sets the value for ForcePasswordChange to be an explicit nil

### UnsetForcePasswordChange
`func (o *User) UnsetForcePasswordChange()`

UnsetForcePasswordChange ensures that no value is present for ForcePasswordChange, not even an explicit nil
### GetGoogleAccount

`func (o *User) GetGoogleAccount() GoogleAccountInfo`

GetGoogleAccount returns the GoogleAccount field if non-nil, zero value otherwise.

### GetGoogleAccountOk

`func (o *User) GetGoogleAccountOk() (*GoogleAccountInfo, bool)`

GetGoogleAccountOk returns a tuple with the GoogleAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAccount

`func (o *User) SetGoogleAccount(v GoogleAccountInfo)`

SetGoogleAccount sets GoogleAccount field to given value.

### HasGoogleAccount

`func (o *User) HasGoogleAccount() bool`

HasGoogleAccount returns a boolean if a field has been set.

### GetGroupRoles

`func (o *User) GetGroupRoles() []string`

GetGroupRoles returns the GroupRoles field if non-nil, zero value otherwise.

### GetGroupRolesOk

`func (o *User) GetGroupRolesOk() (*[]string, bool)`

GetGroupRolesOk returns a tuple with the GroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoles

`func (o *User) SetGroupRoles(v []string)`

SetGroupRoles sets GroupRoles field to given value.

### HasGroupRoles

`func (o *User) HasGroupRoles() bool`

HasGroupRoles returns a boolean if a field has been set.

### SetGroupRolesNil

`func (o *User) SetGroupRolesNil(b bool)`

 SetGroupRolesNil sets the value for GroupRoles to be an explicit nil

### UnsetGroupRoles
`func (o *User) UnsetGroupRoles()`

UnsetGroupRoles ensures that no value is present for GroupRoles, not even an explicit nil
### GetIdpUserInfo

`func (o *User) GetIdpUserInfo() IdpUserInfo`

GetIdpUserInfo returns the IdpUserInfo field if non-nil, zero value otherwise.

### GetIdpUserInfoOk

`func (o *User) GetIdpUserInfoOk() (*IdpUserInfo, bool)`

GetIdpUserInfoOk returns a tuple with the IdpUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUserInfo

`func (o *User) SetIdpUserInfo(v IdpUserInfo)`

SetIdpUserInfo sets IdpUserInfo field to given value.

### HasIdpUserInfo

`func (o *User) HasIdpUserInfo() bool`

HasIdpUserInfo returns a boolean if a field has been set.

### GetLastUpdatedTimeMsecs

`func (o *User) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *User) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *User) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *User) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *User) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *User) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetOrgMembership

`func (o *User) GetOrgMembership() []TenantConfig`

GetOrgMembership returns the OrgMembership field if non-nil, zero value otherwise.

### GetOrgMembershipOk

`func (o *User) GetOrgMembershipOk() (*[]TenantConfig, bool)`

GetOrgMembershipOk returns a tuple with the OrgMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembership

`func (o *User) SetOrgMembership(v []TenantConfig)`

SetOrgMembership sets OrgMembership field to given value.

### HasOrgMembership

`func (o *User) HasOrgMembership() bool`

HasOrgMembership returns a boolean if a field has been set.

### SetOrgMembershipNil

`func (o *User) SetOrgMembershipNil(b bool)`

 SetOrgMembershipNil sets the value for OrgMembership to be an explicit nil

### UnsetOrgMembership
`func (o *User) UnsetOrgMembership()`

UnsetOrgMembership ensures that no value is present for OrgMembership, not even an explicit nil
### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *User) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *User) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPreferences

`func (o *User) GetPreferences() Preferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *User) GetPreferencesOk() (*Preferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *User) SetPreferences(v Preferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *User) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetPrimaryGroupName

`func (o *User) GetPrimaryGroupName() string`

GetPrimaryGroupName returns the PrimaryGroupName field if non-nil, zero value otherwise.

### GetPrimaryGroupNameOk

`func (o *User) GetPrimaryGroupNameOk() (*string, bool)`

GetPrimaryGroupNameOk returns a tuple with the PrimaryGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroupName

`func (o *User) SetPrimaryGroupName(v string)`

SetPrimaryGroupName sets PrimaryGroupName field to given value.

### HasPrimaryGroupName

`func (o *User) HasPrimaryGroupName() bool`

HasPrimaryGroupName returns a boolean if a field has been set.

### SetPrimaryGroupNameNil

`func (o *User) SetPrimaryGroupNameNil(b bool)`

 SetPrimaryGroupNameNil sets the value for PrimaryGroupName to be an explicit nil

### UnsetPrimaryGroupName
`func (o *User) UnsetPrimaryGroupName()`

UnsetPrimaryGroupName ensures that no value is present for PrimaryGroupName, not even an explicit nil
### GetPrivilegeIds

`func (o *User) GetPrivilegeIds() []string`

GetPrivilegeIds returns the PrivilegeIds field if non-nil, zero value otherwise.

### GetPrivilegeIdsOk

`func (o *User) GetPrivilegeIdsOk() (*[]string, bool)`

GetPrivilegeIdsOk returns a tuple with the PrivilegeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeIds

`func (o *User) SetPrivilegeIds(v []string)`

SetPrivilegeIds sets PrivilegeIds field to given value.

### HasPrivilegeIds

`func (o *User) HasPrivilegeIds() bool`

HasPrivilegeIds returns a boolean if a field has been set.

### SetPrivilegeIdsNil

`func (o *User) SetPrivilegeIdsNil(b bool)`

 SetPrivilegeIdsNil sets the value for PrivilegeIds to be an explicit nil

### UnsetPrivilegeIds
`func (o *User) UnsetPrivilegeIds()`

UnsetPrivilegeIds ensures that no value is present for PrivilegeIds, not even an explicit nil
### GetRestricted

`func (o *User) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *User) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *User) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *User) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *User) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *User) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *User) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *User) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetS3AccessKeyId

`func (o *User) GetS3AccessKeyId() string`

GetS3AccessKeyId returns the S3AccessKeyId field if non-nil, zero value otherwise.

### GetS3AccessKeyIdOk

`func (o *User) GetS3AccessKeyIdOk() (*string, bool)`

GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeyId

`func (o *User) SetS3AccessKeyId(v string)`

SetS3AccessKeyId sets S3AccessKeyId field to given value.

### HasS3AccessKeyId

`func (o *User) HasS3AccessKeyId() bool`

HasS3AccessKeyId returns a boolean if a field has been set.

### SetS3AccessKeyIdNil

`func (o *User) SetS3AccessKeyIdNil(b bool)`

 SetS3AccessKeyIdNil sets the value for S3AccessKeyId to be an explicit nil

### UnsetS3AccessKeyId
`func (o *User) UnsetS3AccessKeyId()`

UnsetS3AccessKeyId ensures that no value is present for S3AccessKeyId, not even an explicit nil
### GetS3AccountId

`func (o *User) GetS3AccountId() string`

GetS3AccountId returns the S3AccountId field if non-nil, zero value otherwise.

### GetS3AccountIdOk

`func (o *User) GetS3AccountIdOk() (*string, bool)`

GetS3AccountIdOk returns a tuple with the S3AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccountId

`func (o *User) SetS3AccountId(v string)`

SetS3AccountId sets S3AccountId field to given value.

### HasS3AccountId

`func (o *User) HasS3AccountId() bool`

HasS3AccountId returns a boolean if a field has been set.

### SetS3AccountIdNil

`func (o *User) SetS3AccountIdNil(b bool)`

 SetS3AccountIdNil sets the value for S3AccountId to be an explicit nil

### UnsetS3AccountId
`func (o *User) UnsetS3AccountId()`

UnsetS3AccountId ensures that no value is present for S3AccountId, not even an explicit nil
### GetS3SecretKey

`func (o *User) GetS3SecretKey() string`

GetS3SecretKey returns the S3SecretKey field if non-nil, zero value otherwise.

### GetS3SecretKeyOk

`func (o *User) GetS3SecretKeyOk() (*string, bool)`

GetS3SecretKeyOk returns a tuple with the S3SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3SecretKey

`func (o *User) SetS3SecretKey(v string)`

SetS3SecretKey sets S3SecretKey field to given value.

### HasS3SecretKey

`func (o *User) HasS3SecretKey() bool`

HasS3SecretKey returns a boolean if a field has been set.

### SetS3SecretKeyNil

`func (o *User) SetS3SecretKeyNil(b bool)`

 SetS3SecretKeyNil sets the value for S3SecretKey to be an explicit nil

### UnsetS3SecretKey
`func (o *User) UnsetS3SecretKey()`

UnsetS3SecretKey ensures that no value is present for S3SecretKey, not even an explicit nil
### GetSalesforceAccount

`func (o *User) GetSalesforceAccount() SalesforceAccountInfo`

GetSalesforceAccount returns the SalesforceAccount field if non-nil, zero value otherwise.

### GetSalesforceAccountOk

`func (o *User) GetSalesforceAccountOk() (*SalesforceAccountInfo, bool)`

GetSalesforceAccountOk returns a tuple with the SalesforceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforceAccount

`func (o *User) SetSalesforceAccount(v SalesforceAccountInfo)`

SetSalesforceAccount sets SalesforceAccount field to given value.

### HasSalesforceAccount

`func (o *User) HasSalesforceAccount() bool`

HasSalesforceAccount returns a boolean if a field has been set.

### GetSid

`func (o *User) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *User) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *User) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *User) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *User) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *User) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTenantId

`func (o *User) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *User) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *User) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *User) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *User) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *User) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *User) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *User) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


