# HeliosBaseTenantAssignmentsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | Pointer to **NullableString** | Specifies the the cluster on which properties are to be assigned. The format is clusterId:clusterIncarnationId. | [optional] 
**Policies** | Pointer to **[]string** | A list of Ids of properties assigned to the tenant. | [optional] 

## Methods

### NewHeliosBaseTenantAssignmentsResult

`func NewHeliosBaseTenantAssignmentsResult() *HeliosBaseTenantAssignmentsResult`

NewHeliosBaseTenantAssignmentsResult instantiates a new HeliosBaseTenantAssignmentsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosBaseTenantAssignmentsResultWithDefaults

`func NewHeliosBaseTenantAssignmentsResultWithDefaults() *HeliosBaseTenantAssignmentsResult`

NewHeliosBaseTenantAssignmentsResultWithDefaults instantiates a new HeliosBaseTenantAssignmentsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *HeliosBaseTenantAssignmentsResult) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosBaseTenantAssignmentsResult) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosBaseTenantAssignmentsResult) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosBaseTenantAssignmentsResult) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosBaseTenantAssignmentsResult) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosBaseTenantAssignmentsResult) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetPolicies

`func (o *HeliosBaseTenantAssignmentsResult) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *HeliosBaseTenantAssignmentsResult) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *HeliosBaseTenantAssignmentsResult) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *HeliosBaseTenantAssignmentsResult) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


