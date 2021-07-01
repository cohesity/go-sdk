# TenantConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies name of the tenant. | [optional] 
**Restricted** | Pointer to **NullableBool** | Whether the user is a restricted user. A restricted user can only view the objects he has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Array of Roles.  Specifies the Cohesity roles to associate with the user such as such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for this user. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantConfig

`func NewTenantConfig() *TenantConfig`

NewTenantConfig instantiates a new TenantConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantConfigWithDefaults

`func NewTenantConfigWithDefaults() *TenantConfig`

NewTenantConfigWithDefaults instantiates a new TenantConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TenantConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRestricted

`func (o *TenantConfig) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *TenantConfig) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *TenantConfig) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *TenantConfig) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *TenantConfig) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *TenantConfig) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *TenantConfig) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TenantConfig) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TenantConfig) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *TenantConfig) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *TenantConfig) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *TenantConfig) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetTenantId

`func (o *TenantConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantConfig) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantConfig) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


