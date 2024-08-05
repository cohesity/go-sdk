# UpdateGroupParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the group. | [optional] 
**LocalGroupParams** | Pointer to [**UpdateGroupParametersLocalGroupParams**](UpdateGroupParametersLocalGroupParams.md) |  | [optional] 
**Restricted** | Pointer to **NullableBool** | Specifies whether the Group is restricted. A restricted group can only view &amp; manage the objects it has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with the group. The Cohesity roles determine privileges on the Cohesity Cluster for this group. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies a list of tenant ids who can access this group. | [optional] 

## Methods

### NewUpdateGroupParameters

`func NewUpdateGroupParameters() *UpdateGroupParameters`

NewUpdateGroupParameters instantiates a new UpdateGroupParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupParametersWithDefaults

`func NewUpdateGroupParametersWithDefaults() *UpdateGroupParameters`

NewUpdateGroupParametersWithDefaults instantiates a new UpdateGroupParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateGroupParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroupParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroupParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroupParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateGroupParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateGroupParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalGroupParams

`func (o *UpdateGroupParameters) GetLocalGroupParams() UpdateGroupParametersLocalGroupParams`

GetLocalGroupParams returns the LocalGroupParams field if non-nil, zero value otherwise.

### GetLocalGroupParamsOk

`func (o *UpdateGroupParameters) GetLocalGroupParamsOk() (*UpdateGroupParametersLocalGroupParams, bool)`

GetLocalGroupParamsOk returns a tuple with the LocalGroupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroupParams

`func (o *UpdateGroupParameters) SetLocalGroupParams(v UpdateGroupParametersLocalGroupParams)`

SetLocalGroupParams sets LocalGroupParams field to given value.

### HasLocalGroupParams

`func (o *UpdateGroupParameters) HasLocalGroupParams() bool`

HasLocalGroupParams returns a boolean if a field has been set.

### GetRestricted

`func (o *UpdateGroupParameters) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *UpdateGroupParameters) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *UpdateGroupParameters) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *UpdateGroupParameters) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *UpdateGroupParameters) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *UpdateGroupParameters) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *UpdateGroupParameters) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateGroupParameters) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateGroupParameters) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateGroupParameters) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTenantIds

`func (o *UpdateGroupParameters) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *UpdateGroupParameters) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *UpdateGroupParameters) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *UpdateGroupParameters) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


