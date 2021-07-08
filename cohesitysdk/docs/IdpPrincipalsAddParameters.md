# IdpPrincipalsAddParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the name of the Idp where the referenced principal is stored. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the type of the referenced Idp principal. If &#39;kGroup&#39;, the referenced Idp principal is a group. If &#39;kUser&#39;, the referenced Idp principal is a user. &#39;kUser&#39; specifies a user object class. &#39;kGroup&#39; specifies a group object class. &#39;kComputer&#39; specifies a computer object class. &#39;kWellKnownPrincipal&#39; specifies a well known principal. | [optional] 
**PrincipalName** | Pointer to **NullableString** | Specifies the name of the Idp principal, that will be referenced by the group or user. The name of the Idp principal is used for naming the new group or user on the Cohesity Cluster. | [optional] 
**Restricted** | Pointer to **NullableBool** | Whether the principal is a restricted principal. A restricted principal can only view the objects he has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Array of Roles.  Specifies the Cohesity roles to associate with this user or group such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for this group or user. For example if the &#39;joe&#39; user is added for the Active Directory &#39;joe&#39; user principal and is associated with the Cohesity &#39;View&#39; role, &#39;joe&#39; can log in to the Cohesity Dashboard and has a read-only view of the data on the Cohesity Cluster. | [optional] 

## Methods

### NewIdpPrincipalsAddParameters

`func NewIdpPrincipalsAddParameters() *IdpPrincipalsAddParameters`

NewIdpPrincipalsAddParameters instantiates a new IdpPrincipalsAddParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPrincipalsAddParametersWithDefaults

`func NewIdpPrincipalsAddParametersWithDefaults() *IdpPrincipalsAddParameters`

NewIdpPrincipalsAddParametersWithDefaults instantiates a new IdpPrincipalsAddParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *IdpPrincipalsAddParameters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdpPrincipalsAddParameters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdpPrincipalsAddParameters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdpPrincipalsAddParameters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *IdpPrincipalsAddParameters) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdpPrincipalsAddParameters) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetObjectClass

`func (o *IdpPrincipalsAddParameters) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *IdpPrincipalsAddParameters) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *IdpPrincipalsAddParameters) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *IdpPrincipalsAddParameters) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *IdpPrincipalsAddParameters) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *IdpPrincipalsAddParameters) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetPrincipalName

`func (o *IdpPrincipalsAddParameters) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *IdpPrincipalsAddParameters) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *IdpPrincipalsAddParameters) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *IdpPrincipalsAddParameters) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *IdpPrincipalsAddParameters) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *IdpPrincipalsAddParameters) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetRestricted

`func (o *IdpPrincipalsAddParameters) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *IdpPrincipalsAddParameters) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *IdpPrincipalsAddParameters) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *IdpPrincipalsAddParameters) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *IdpPrincipalsAddParameters) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *IdpPrincipalsAddParameters) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *IdpPrincipalsAddParameters) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IdpPrincipalsAddParameters) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IdpPrincipalsAddParameters) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IdpPrincipalsAddParameters) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IdpPrincipalsAddParameters) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IdpPrincipalsAddParameters) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


