# HeliosTenantAssignmentsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | Pointer to **NullableString** | Specifies the the cluster on which properties are to be assigned. The format is clusterId:clusterIncarnationId. | [optional] 
**PolicyIds** | Pointer to **[]string** | List of policies on the cluster, to be assigned to the Tenant. | [optional] 

## Methods

### NewHeliosTenantAssignmentsParams

`func NewHeliosTenantAssignmentsParams() *HeliosTenantAssignmentsParams`

NewHeliosTenantAssignmentsParams instantiates a new HeliosTenantAssignmentsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTenantAssignmentsParamsWithDefaults

`func NewHeliosTenantAssignmentsParamsWithDefaults() *HeliosTenantAssignmentsParams`

NewHeliosTenantAssignmentsParamsWithDefaults instantiates a new HeliosTenantAssignmentsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *HeliosTenantAssignmentsParams) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosTenantAssignmentsParams) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosTenantAssignmentsParams) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosTenantAssignmentsParams) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosTenantAssignmentsParams) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosTenantAssignmentsParams) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetPolicyIds

`func (o *HeliosTenantAssignmentsParams) GetPolicyIds() []string`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *HeliosTenantAssignmentsParams) GetPolicyIdsOk() (*[]string, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *HeliosTenantAssignmentsParams) SetPolicyIds(v []string)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *HeliosTenantAssignmentsParams) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### SetPolicyIdsNil

`func (o *HeliosTenantAssignmentsParams) SetPolicyIdsNil(b bool)`

 SetPolicyIdsNil sets the value for PolicyIds to be an explicit nil

### UnsetPolicyIds
`func (o *HeliosTenantAssignmentsParams) UnsetPolicyIds()`

UnsetPolicyIds ensures that no value is present for PolicyIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


