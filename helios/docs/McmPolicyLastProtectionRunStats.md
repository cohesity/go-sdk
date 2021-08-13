# McmPolicyLastProtectionRunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster Id. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster Incarnation Id. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region Id. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the policy Id. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the policy name. | [optional] 
**LastRunStats** | Pointer to [**McmPolicyLastProtectionRunStatsByPolicy**](McmPolicyLastProtectionRunStatsByPolicy.md) |  | [optional] 

## Methods

### NewMcmPolicyLastProtectionRunStats

`func NewMcmPolicyLastProtectionRunStats() *McmPolicyLastProtectionRunStats`

NewMcmPolicyLastProtectionRunStats instantiates a new McmPolicyLastProtectionRunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmPolicyLastProtectionRunStatsWithDefaults

`func NewMcmPolicyLastProtectionRunStatsWithDefaults() *McmPolicyLastProtectionRunStats`

NewMcmPolicyLastProtectionRunStatsWithDefaults instantiates a new McmPolicyLastProtectionRunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *McmPolicyLastProtectionRunStats) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmPolicyLastProtectionRunStats) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmPolicyLastProtectionRunStats) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmPolicyLastProtectionRunStats) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmPolicyLastProtectionRunStats) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmPolicyLastProtectionRunStats) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmPolicyLastProtectionRunStats) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmPolicyLastProtectionRunStats) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmPolicyLastProtectionRunStats) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmPolicyLastProtectionRunStats) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmPolicyLastProtectionRunStats) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmPolicyLastProtectionRunStats) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRegionId

`func (o *McmPolicyLastProtectionRunStats) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *McmPolicyLastProtectionRunStats) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *McmPolicyLastProtectionRunStats) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *McmPolicyLastProtectionRunStats) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *McmPolicyLastProtectionRunStats) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *McmPolicyLastProtectionRunStats) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetPolicyId

`func (o *McmPolicyLastProtectionRunStats) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *McmPolicyLastProtectionRunStats) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *McmPolicyLastProtectionRunStats) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *McmPolicyLastProtectionRunStats) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *McmPolicyLastProtectionRunStats) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *McmPolicyLastProtectionRunStats) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *McmPolicyLastProtectionRunStats) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *McmPolicyLastProtectionRunStats) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *McmPolicyLastProtectionRunStats) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *McmPolicyLastProtectionRunStats) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *McmPolicyLastProtectionRunStats) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *McmPolicyLastProtectionRunStats) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetLastRunStats

`func (o *McmPolicyLastProtectionRunStats) GetLastRunStats() McmPolicyLastProtectionRunStatsByPolicy`

GetLastRunStats returns the LastRunStats field if non-nil, zero value otherwise.

### GetLastRunStatsOk

`func (o *McmPolicyLastProtectionRunStats) GetLastRunStatsOk() (*McmPolicyLastProtectionRunStatsByPolicy, bool)`

GetLastRunStatsOk returns a tuple with the LastRunStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStats

`func (o *McmPolicyLastProtectionRunStats) SetLastRunStats(v McmPolicyLastProtectionRunStatsByPolicy)`

SetLastRunStats sets LastRunStats field to given value.

### HasLastRunStats

`func (o *McmPolicyLastProtectionRunStats) HasLastRunStats() bool`

HasLastRunStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


