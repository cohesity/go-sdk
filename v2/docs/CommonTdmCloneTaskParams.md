# CommonTdmCloneTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment of the TDM Clone task. | 
**PolicyId** | Pointer to **NullableString** | Specifies the ID of the policy, which should be used to protect this clone. This is useful for automatic snapshots. This must be specified if either of protectionGroupId and protectionGroupName is specified. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the ID of an existing protection group, which should start protecting this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupName should not be specified. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the name of a new protection group, which should be created to protect this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupId should not be specified. | [optional] 

## Methods

### NewCommonTdmCloneTaskParams

`func NewCommonTdmCloneTaskParams(environment NullableString, ) *CommonTdmCloneTaskParams`

NewCommonTdmCloneTaskParams instantiates a new CommonTdmCloneTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmCloneTaskParamsWithDefaults

`func NewCommonTdmCloneTaskParamsWithDefaults() *CommonTdmCloneTaskParams`

NewCommonTdmCloneTaskParamsWithDefaults instantiates a new CommonTdmCloneTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CommonTdmCloneTaskParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonTdmCloneTaskParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonTdmCloneTaskParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *CommonTdmCloneTaskParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonTdmCloneTaskParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetPolicyId

`func (o *CommonTdmCloneTaskParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CommonTdmCloneTaskParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CommonTdmCloneTaskParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CommonTdmCloneTaskParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CommonTdmCloneTaskParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CommonTdmCloneTaskParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetProtectionGroupId

`func (o *CommonTdmCloneTaskParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *CommonTdmCloneTaskParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *CommonTdmCloneTaskParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *CommonTdmCloneTaskParams) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *CommonTdmCloneTaskParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *CommonTdmCloneTaskParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *CommonTdmCloneTaskParams) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *CommonTdmCloneTaskParams) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *CommonTdmCloneTaskParams) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *CommonTdmCloneTaskParams) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *CommonTdmCloneTaskParams) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *CommonTdmCloneTaskParams) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


