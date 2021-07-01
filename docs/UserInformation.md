# UserInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeSubtenantObjects** | Pointer to **NullableBool** | Whether objects owned by subtenants should be returned. This would require a prefix search with the passed tenant_id. All tenants are considered sub-tenants of the admin. For GET requests, if tenant id is empty(admin user is querying) and if this flag is false, we will only return untagged objects. If it is true, we will return everything. | [optional] 
**PulseAttributeVec** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the KeyValuePair that client (eg. Iris) wants to persist along with the corresponding (soon-to-be-created) Pulse task for the current action. Eg. pulse_attribute_vec can drive user notifications by associating a Pulse Task with user SID and later Pulse can be searched by client specified Sid to get all finished tasks for the logged in user. | [optional] 
**SidVec** | Pointer to [**[]ClusterConfigProtoSID**](ClusterConfigProtoSID.md) | If specified, only the objects associated with these SIDs should be returned. | [optional] 
**TenantIdVec** | Pointer to **[]string** | If specified, only the objects associated with this tenant should be returned. A given tenant ID is always a prefix of the ids of its subtenants. Eg. if tenant_id of cluster admin is empty string then it will be a prefix match for all the tenants on the cluster. | [optional] 

## Methods

### NewUserInformation

`func NewUserInformation() *UserInformation`

NewUserInformation instantiates a new UserInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInformationWithDefaults

`func NewUserInformationWithDefaults() *UserInformation`

NewUserInformationWithDefaults instantiates a new UserInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeSubtenantObjects

`func (o *UserInformation) GetIncludeSubtenantObjects() bool`

GetIncludeSubtenantObjects returns the IncludeSubtenantObjects field if non-nil, zero value otherwise.

### GetIncludeSubtenantObjectsOk

`func (o *UserInformation) GetIncludeSubtenantObjectsOk() (*bool, bool)`

GetIncludeSubtenantObjectsOk returns a tuple with the IncludeSubtenantObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubtenantObjects

`func (o *UserInformation) SetIncludeSubtenantObjects(v bool)`

SetIncludeSubtenantObjects sets IncludeSubtenantObjects field to given value.

### HasIncludeSubtenantObjects

`func (o *UserInformation) HasIncludeSubtenantObjects() bool`

HasIncludeSubtenantObjects returns a boolean if a field has been set.

### SetIncludeSubtenantObjectsNil

`func (o *UserInformation) SetIncludeSubtenantObjectsNil(b bool)`

 SetIncludeSubtenantObjectsNil sets the value for IncludeSubtenantObjects to be an explicit nil

### UnsetIncludeSubtenantObjects
`func (o *UserInformation) UnsetIncludeSubtenantObjects()`

UnsetIncludeSubtenantObjects ensures that no value is present for IncludeSubtenantObjects, not even an explicit nil
### GetPulseAttributeVec

`func (o *UserInformation) GetPulseAttributeVec() []KeyValuePair`

GetPulseAttributeVec returns the PulseAttributeVec field if non-nil, zero value otherwise.

### GetPulseAttributeVecOk

`func (o *UserInformation) GetPulseAttributeVecOk() (*[]KeyValuePair, bool)`

GetPulseAttributeVecOk returns a tuple with the PulseAttributeVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulseAttributeVec

`func (o *UserInformation) SetPulseAttributeVec(v []KeyValuePair)`

SetPulseAttributeVec sets PulseAttributeVec field to given value.

### HasPulseAttributeVec

`func (o *UserInformation) HasPulseAttributeVec() bool`

HasPulseAttributeVec returns a boolean if a field has been set.

### SetPulseAttributeVecNil

`func (o *UserInformation) SetPulseAttributeVecNil(b bool)`

 SetPulseAttributeVecNil sets the value for PulseAttributeVec to be an explicit nil

### UnsetPulseAttributeVec
`func (o *UserInformation) UnsetPulseAttributeVec()`

UnsetPulseAttributeVec ensures that no value is present for PulseAttributeVec, not even an explicit nil
### GetSidVec

`func (o *UserInformation) GetSidVec() []ClusterConfigProtoSID`

GetSidVec returns the SidVec field if non-nil, zero value otherwise.

### GetSidVecOk

`func (o *UserInformation) GetSidVecOk() (*[]ClusterConfigProtoSID, bool)`

GetSidVecOk returns a tuple with the SidVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidVec

`func (o *UserInformation) SetSidVec(v []ClusterConfigProtoSID)`

SetSidVec sets SidVec field to given value.

### HasSidVec

`func (o *UserInformation) HasSidVec() bool`

HasSidVec returns a boolean if a field has been set.

### SetSidVecNil

`func (o *UserInformation) SetSidVecNil(b bool)`

 SetSidVecNil sets the value for SidVec to be an explicit nil

### UnsetSidVec
`func (o *UserInformation) UnsetSidVec()`

UnsetSidVec ensures that no value is present for SidVec, not even an explicit nil
### GetTenantIdVec

`func (o *UserInformation) GetTenantIdVec() []string`

GetTenantIdVec returns the TenantIdVec field if non-nil, zero value otherwise.

### GetTenantIdVecOk

`func (o *UserInformation) GetTenantIdVecOk() (*[]string, bool)`

GetTenantIdVecOk returns a tuple with the TenantIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdVec

`func (o *UserInformation) SetTenantIdVec(v []string)`

SetTenantIdVec sets TenantIdVec field to given value.

### HasTenantIdVec

`func (o *UserInformation) HasTenantIdVec() bool`

HasTenantIdVec returns a boolean if a field has been set.

### SetTenantIdVecNil

`func (o *UserInformation) SetTenantIdVecNil(b bool)`

 SetTenantIdVecNil sets the value for TenantIdVec to be an explicit nil

### UnsetTenantIdVec
`func (o *UserInformation) UnsetTenantIdVec()`

UnsetTenantIdVec ensures that no value is present for TenantIdVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


