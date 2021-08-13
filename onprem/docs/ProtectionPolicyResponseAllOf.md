# ProtectionPolicyResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies a unique Policy id assigned by the Cohesity Cluster. | [optional] 
**TemplateId** | Pointer to **NullableString** | Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template. | [optional] 
**IsUsable** | Pointer to **NullableBool** | This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated. | [optional] 
**IsReplicated** | Pointer to **NullableBool** | This field is set to true when policy is the replicated policy. | [optional] 
**NumProtectionGroups** | Pointer to **NullableInt64** | Specifies the number of protection groups using the protection policy. | [optional] 
**NumProtectedObjects** | Pointer to **NullableInt64** | Specifies the number of protected objects using the protection policy. | [optional] 

## Methods

### NewProtectionPolicyResponseAllOf

`func NewProtectionPolicyResponseAllOf() *ProtectionPolicyResponseAllOf`

NewProtectionPolicyResponseAllOf instantiates a new ProtectionPolicyResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionPolicyResponseAllOfWithDefaults

`func NewProtectionPolicyResponseAllOfWithDefaults() *ProtectionPolicyResponseAllOf`

NewProtectionPolicyResponseAllOfWithDefaults instantiates a new ProtectionPolicyResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtectionPolicyResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionPolicyResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionPolicyResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionPolicyResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionPolicyResponseAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionPolicyResponseAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTemplateId

`func (o *ProtectionPolicyResponseAllOf) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ProtectionPolicyResponseAllOf) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ProtectionPolicyResponseAllOf) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ProtectionPolicyResponseAllOf) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *ProtectionPolicyResponseAllOf) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *ProtectionPolicyResponseAllOf) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetIsUsable

`func (o *ProtectionPolicyResponseAllOf) GetIsUsable() bool`

GetIsUsable returns the IsUsable field if non-nil, zero value otherwise.

### GetIsUsableOk

`func (o *ProtectionPolicyResponseAllOf) GetIsUsableOk() (*bool, bool)`

GetIsUsableOk returns a tuple with the IsUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsable

`func (o *ProtectionPolicyResponseAllOf) SetIsUsable(v bool)`

SetIsUsable sets IsUsable field to given value.

### HasIsUsable

`func (o *ProtectionPolicyResponseAllOf) HasIsUsable() bool`

HasIsUsable returns a boolean if a field has been set.

### SetIsUsableNil

`func (o *ProtectionPolicyResponseAllOf) SetIsUsableNil(b bool)`

 SetIsUsableNil sets the value for IsUsable to be an explicit nil

### UnsetIsUsable
`func (o *ProtectionPolicyResponseAllOf) UnsetIsUsable()`

UnsetIsUsable ensures that no value is present for IsUsable, not even an explicit nil
### GetIsReplicated

`func (o *ProtectionPolicyResponseAllOf) GetIsReplicated() bool`

GetIsReplicated returns the IsReplicated field if non-nil, zero value otherwise.

### GetIsReplicatedOk

`func (o *ProtectionPolicyResponseAllOf) GetIsReplicatedOk() (*bool, bool)`

GetIsReplicatedOk returns a tuple with the IsReplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplicated

`func (o *ProtectionPolicyResponseAllOf) SetIsReplicated(v bool)`

SetIsReplicated sets IsReplicated field to given value.

### HasIsReplicated

`func (o *ProtectionPolicyResponseAllOf) HasIsReplicated() bool`

HasIsReplicated returns a boolean if a field has been set.

### SetIsReplicatedNil

`func (o *ProtectionPolicyResponseAllOf) SetIsReplicatedNil(b bool)`

 SetIsReplicatedNil sets the value for IsReplicated to be an explicit nil

### UnsetIsReplicated
`func (o *ProtectionPolicyResponseAllOf) UnsetIsReplicated()`

UnsetIsReplicated ensures that no value is present for IsReplicated, not even an explicit nil
### GetNumProtectionGroups

`func (o *ProtectionPolicyResponseAllOf) GetNumProtectionGroups() int64`

GetNumProtectionGroups returns the NumProtectionGroups field if non-nil, zero value otherwise.

### GetNumProtectionGroupsOk

`func (o *ProtectionPolicyResponseAllOf) GetNumProtectionGroupsOk() (*int64, bool)`

GetNumProtectionGroupsOk returns a tuple with the NumProtectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProtectionGroups

`func (o *ProtectionPolicyResponseAllOf) SetNumProtectionGroups(v int64)`

SetNumProtectionGroups sets NumProtectionGroups field to given value.

### HasNumProtectionGroups

`func (o *ProtectionPolicyResponseAllOf) HasNumProtectionGroups() bool`

HasNumProtectionGroups returns a boolean if a field has been set.

### SetNumProtectionGroupsNil

`func (o *ProtectionPolicyResponseAllOf) SetNumProtectionGroupsNil(b bool)`

 SetNumProtectionGroupsNil sets the value for NumProtectionGroups to be an explicit nil

### UnsetNumProtectionGroups
`func (o *ProtectionPolicyResponseAllOf) UnsetNumProtectionGroups()`

UnsetNumProtectionGroups ensures that no value is present for NumProtectionGroups, not even an explicit nil
### GetNumProtectedObjects

`func (o *ProtectionPolicyResponseAllOf) GetNumProtectedObjects() int64`

GetNumProtectedObjects returns the NumProtectedObjects field if non-nil, zero value otherwise.

### GetNumProtectedObjectsOk

`func (o *ProtectionPolicyResponseAllOf) GetNumProtectedObjectsOk() (*int64, bool)`

GetNumProtectedObjectsOk returns a tuple with the NumProtectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProtectedObjects

`func (o *ProtectionPolicyResponseAllOf) SetNumProtectedObjects(v int64)`

SetNumProtectedObjects sets NumProtectedObjects field to given value.

### HasNumProtectedObjects

`func (o *ProtectionPolicyResponseAllOf) HasNumProtectedObjects() bool`

HasNumProtectedObjects returns a boolean if a field has been set.

### SetNumProtectedObjectsNil

`func (o *ProtectionPolicyResponseAllOf) SetNumProtectedObjectsNil(b bool)`

 SetNumProtectedObjectsNil sets the value for NumProtectedObjects to be an explicit nil

### UnsetNumProtectedObjects
`func (o *ProtectionPolicyResponseAllOf) UnsetNumProtectedObjects()`

UnsetNumProtectedObjects ensures that no value is present for NumProtectedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


