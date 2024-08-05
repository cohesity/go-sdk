# CreateGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the group. | [optional] 
**Domain** | **string** | Specifies the domain of the group. For active directories, this is the fully qualified domain name (FQDN). It is &#39;LOCAL&#39; for local groups on the Cohesity Cluster. A group is uniquely identified by combination of the name and the domain. | 
**LocalGroupParams** | Pointer to [**CreateGroupParamsLocalGroupParams**](CreateGroupParamsLocalGroupParams.md) |  | [optional] 
**Name** | **string** | Specifies the name of the group. | 
**Restricted** | Pointer to **NullableBool** | Specifies whether the Group is restricted. A restricted group can only view &amp; manage the objects it has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with the group. The Cohesity roles determine privileges on the Cohesity Cluster for this group. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies a list of tenant ids who can access this group. | [optional] 

## Methods

### NewCreateGroupParams

`func NewCreateGroupParams(domain string, name string, ) *CreateGroupParams`

NewCreateGroupParams instantiates a new CreateGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupParamsWithDefaults

`func NewCreateGroupParamsWithDefaults() *CreateGroupParams`

NewCreateGroupParamsWithDefaults instantiates a new CreateGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateGroupParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateGroupParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateGroupParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *CreateGroupParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateGroupParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateGroupParams) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetLocalGroupParams

`func (o *CreateGroupParams) GetLocalGroupParams() CreateGroupParamsLocalGroupParams`

GetLocalGroupParams returns the LocalGroupParams field if non-nil, zero value otherwise.

### GetLocalGroupParamsOk

`func (o *CreateGroupParams) GetLocalGroupParamsOk() (*CreateGroupParamsLocalGroupParams, bool)`

GetLocalGroupParamsOk returns a tuple with the LocalGroupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroupParams

`func (o *CreateGroupParams) SetLocalGroupParams(v CreateGroupParamsLocalGroupParams)`

SetLocalGroupParams sets LocalGroupParams field to given value.

### HasLocalGroupParams

`func (o *CreateGroupParams) HasLocalGroupParams() bool`

HasLocalGroupParams returns a boolean if a field has been set.

### GetName

`func (o *CreateGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupParams) SetName(v string)`

SetName sets Name field to given value.


### GetRestricted

`func (o *CreateGroupParams) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *CreateGroupParams) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *CreateGroupParams) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *CreateGroupParams) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *CreateGroupParams) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *CreateGroupParams) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *CreateGroupParams) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateGroupParams) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateGroupParams) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateGroupParams) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTenantIds

`func (o *CreateGroupParams) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *CreateGroupParams) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *CreateGroupParams) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *CreateGroupParams) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


