# AddActiveDirectoryPrincipalsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies a description about the user or group. | [optional] 
**DomainName** | **NullableString** | Specifies the domain of the Active Directory where the referenced principal is stored. | 
**Name** | **NullableString** | Specifies the name of the Active Directory principal, that will be referenced by the group or user. The name of the Active Directory principal is used for naming the new group or user on the Cohesity Cluster. | 
**ObjectClass** | **NullableString** | Specifies the type of Active Directory principal.&lt;br&gt; &#39;User&#39; specifies a user object class.&lt;br&gt; &#39;Group&#39; specifies a group object class.&lt;br&gt; &#39;ServiceAccount&#39; specifies a service account object class. | 
**Restricted** | Pointer to **NullableBool** | Whether the principal is a restricted principal. A restricted principal can only view the objects he has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with this user or group such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for this group or user. For example if the &#39;joe&#39; user is added form the Active Directory and is associated with the Cohesity &#39;View&#39; role,&#39;joe&#39; can log in to the Cohesity Dashboard and has a read-only view of the data on the Cohesity Cluster. | [optional] 

## Methods

### NewAddActiveDirectoryPrincipalsParameters

`func NewAddActiveDirectoryPrincipalsParameters(domainName NullableString, name NullableString, objectClass NullableString, ) *AddActiveDirectoryPrincipalsParameters`

NewAddActiveDirectoryPrincipalsParameters instantiates a new AddActiveDirectoryPrincipalsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddActiveDirectoryPrincipalsParametersWithDefaults

`func NewAddActiveDirectoryPrincipalsParametersWithDefaults() *AddActiveDirectoryPrincipalsParameters`

NewAddActiveDirectoryPrincipalsParametersWithDefaults instantiates a new AddActiveDirectoryPrincipalsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AddActiveDirectoryPrincipalsParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddActiveDirectoryPrincipalsParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddActiveDirectoryPrincipalsParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddActiveDirectoryPrincipalsParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddActiveDirectoryPrincipalsParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddActiveDirectoryPrincipalsParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomainName

`func (o *AddActiveDirectoryPrincipalsParameters) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddActiveDirectoryPrincipalsParameters) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddActiveDirectoryPrincipalsParameters) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### SetDomainNameNil

`func (o *AddActiveDirectoryPrincipalsParameters) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *AddActiveDirectoryPrincipalsParameters) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetName

`func (o *AddActiveDirectoryPrincipalsParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddActiveDirectoryPrincipalsParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddActiveDirectoryPrincipalsParameters) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AddActiveDirectoryPrincipalsParameters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddActiveDirectoryPrincipalsParameters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectClass

`func (o *AddActiveDirectoryPrincipalsParameters) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AddActiveDirectoryPrincipalsParameters) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AddActiveDirectoryPrincipalsParameters) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.


### SetObjectClassNil

`func (o *AddActiveDirectoryPrincipalsParameters) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *AddActiveDirectoryPrincipalsParameters) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetRestricted

`func (o *AddActiveDirectoryPrincipalsParameters) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *AddActiveDirectoryPrincipalsParameters) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *AddActiveDirectoryPrincipalsParameters) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *AddActiveDirectoryPrincipalsParameters) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *AddActiveDirectoryPrincipalsParameters) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *AddActiveDirectoryPrincipalsParameters) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *AddActiveDirectoryPrincipalsParameters) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AddActiveDirectoryPrincipalsParameters) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AddActiveDirectoryPrincipalsParameters) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AddActiveDirectoryPrincipalsParameters) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *AddActiveDirectoryPrincipalsParameters) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *AddActiveDirectoryPrincipalsParameters) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


