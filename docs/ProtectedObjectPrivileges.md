# ProtectedObjectPrivileges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectedObjectsprivilegesType** | Pointer to **NullableString** | Specifies if all, none or specific protected objects are allowed to be accessed. Specifies if all, none or specific protected objects are allowed to be accessed. kNone - None of the protected objects have access. kAll - All the protected objects have access. kSpecific - Only specific protected objects have access. | [optional] 
**ProtectionSourceIds** | Pointer to **[]int64** | Specifies the ids of the protection sources which are allowed to be accessed in case the privilege type is kSpecific. | [optional] 

## Methods

### NewProtectedObjectPrivileges

`func NewProtectedObjectPrivileges() *ProtectedObjectPrivileges`

NewProtectedObjectPrivileges instantiates a new ProtectedObjectPrivileges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectPrivilegesWithDefaults

`func NewProtectedObjectPrivilegesWithDefaults() *ProtectedObjectPrivileges`

NewProtectedObjectPrivilegesWithDefaults instantiates a new ProtectedObjectPrivileges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectedObjectsprivilegesType

`func (o *ProtectedObjectPrivileges) GetProtectedObjectsprivilegesType() string`

GetProtectedObjectsprivilegesType returns the ProtectedObjectsprivilegesType field if non-nil, zero value otherwise.

### GetProtectedObjectsprivilegesTypeOk

`func (o *ProtectedObjectPrivileges) GetProtectedObjectsprivilegesTypeOk() (*string, bool)`

GetProtectedObjectsprivilegesTypeOk returns a tuple with the ProtectedObjectsprivilegesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedObjectsprivilegesType

`func (o *ProtectedObjectPrivileges) SetProtectedObjectsprivilegesType(v string)`

SetProtectedObjectsprivilegesType sets ProtectedObjectsprivilegesType field to given value.

### HasProtectedObjectsprivilegesType

`func (o *ProtectedObjectPrivileges) HasProtectedObjectsprivilegesType() bool`

HasProtectedObjectsprivilegesType returns a boolean if a field has been set.

### SetProtectedObjectsprivilegesTypeNil

`func (o *ProtectedObjectPrivileges) SetProtectedObjectsprivilegesTypeNil(b bool)`

 SetProtectedObjectsprivilegesTypeNil sets the value for ProtectedObjectsprivilegesType to be an explicit nil

### UnsetProtectedObjectsprivilegesType
`func (o *ProtectedObjectPrivileges) UnsetProtectedObjectsprivilegesType()`

UnsetProtectedObjectsprivilegesType ensures that no value is present for ProtectedObjectsprivilegesType, not even an explicit nil
### GetProtectionSourceIds

`func (o *ProtectedObjectPrivileges) GetProtectionSourceIds() []int64`

GetProtectionSourceIds returns the ProtectionSourceIds field if non-nil, zero value otherwise.

### GetProtectionSourceIdsOk

`func (o *ProtectedObjectPrivileges) GetProtectionSourceIdsOk() (*[]int64, bool)`

GetProtectionSourceIdsOk returns a tuple with the ProtectionSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceIds

`func (o *ProtectedObjectPrivileges) SetProtectionSourceIds(v []int64)`

SetProtectionSourceIds sets ProtectionSourceIds field to given value.

### HasProtectionSourceIds

`func (o *ProtectedObjectPrivileges) HasProtectionSourceIds() bool`

HasProtectionSourceIds returns a boolean if a field has been set.

### SetProtectionSourceIdsNil

`func (o *ProtectedObjectPrivileges) SetProtectionSourceIdsNil(b bool)`

 SetProtectionSourceIdsNil sets the value for ProtectionSourceIds to be an explicit nil

### UnsetProtectionSourceIds
`func (o *ProtectedObjectPrivileges) UnsetProtectionSourceIds()`

UnsetProtectionSourceIds ensures that no value is present for ProtectionSourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


