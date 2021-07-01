# RoleDeleteParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **[]string** | Array of Role Names.  Specifies the list of roles to delete which are specified by role names. | 

## Methods

### NewRoleDeleteParameters

`func NewRoleDeleteParameters(names []string, ) *RoleDeleteParameters`

NewRoleDeleteParameters instantiates a new RoleDeleteParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDeleteParametersWithDefaults

`func NewRoleDeleteParametersWithDefaults() *RoleDeleteParameters`

NewRoleDeleteParametersWithDefaults instantiates a new RoleDeleteParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *RoleDeleteParameters) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *RoleDeleteParameters) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *RoleDeleteParameters) SetNames(v []string)`

SetNames sets Names field to given value.


### SetNamesNil

`func (o *RoleDeleteParameters) SetNamesNil(b bool)`

 SetNamesNil sets the value for Names to be an explicit nil

### UnsetNames
`func (o *RoleDeleteParameters) UnsetNames()`

UnsetNames ensures that no value is present for Names, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


