# ProtectedVmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionJobs** | Pointer to [**[]ProtectionJob**](ProtectionJob.md) | Specifies the list of Protection Jobs that protect the VM. | [optional] 
**ProtectionPolicies** | Pointer to [**[]ProtectionPolicy**](ProtectionPolicy.md) | Specifies the list of Policies that are used by the Protection Jobs. | [optional] 
**ProtectionSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**Stats** | Pointer to [**NullableProtectionSummary**](ProtectionSummary.md) | Specifies the protection stats of VM. | [optional] 

## Methods

### NewProtectedVmInfo

`func NewProtectedVmInfo() *ProtectedVmInfo`

NewProtectedVmInfo instantiates a new ProtectedVmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedVmInfoWithDefaults

`func NewProtectedVmInfoWithDefaults() *ProtectedVmInfo`

NewProtectedVmInfoWithDefaults instantiates a new ProtectedVmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionJobs

`func (o *ProtectedVmInfo) GetProtectionJobs() []ProtectionJob`

GetProtectionJobs returns the ProtectionJobs field if non-nil, zero value otherwise.

### GetProtectionJobsOk

`func (o *ProtectedVmInfo) GetProtectionJobsOk() (*[]ProtectionJob, bool)`

GetProtectionJobsOk returns a tuple with the ProtectionJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobs

`func (o *ProtectedVmInfo) SetProtectionJobs(v []ProtectionJob)`

SetProtectionJobs sets ProtectionJobs field to given value.

### HasProtectionJobs

`func (o *ProtectedVmInfo) HasProtectionJobs() bool`

HasProtectionJobs returns a boolean if a field has been set.

### SetProtectionJobsNil

`func (o *ProtectedVmInfo) SetProtectionJobsNil(b bool)`

 SetProtectionJobsNil sets the value for ProtectionJobs to be an explicit nil

### UnsetProtectionJobs
`func (o *ProtectedVmInfo) UnsetProtectionJobs()`

UnsetProtectionJobs ensures that no value is present for ProtectionJobs, not even an explicit nil
### GetProtectionPolicies

`func (o *ProtectedVmInfo) GetProtectionPolicies() []ProtectionPolicy`

GetProtectionPolicies returns the ProtectionPolicies field if non-nil, zero value otherwise.

### GetProtectionPoliciesOk

`func (o *ProtectedVmInfo) GetProtectionPoliciesOk() (*[]ProtectionPolicy, bool)`

GetProtectionPoliciesOk returns a tuple with the ProtectionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionPolicies

`func (o *ProtectedVmInfo) SetProtectionPolicies(v []ProtectionPolicy)`

SetProtectionPolicies sets ProtectionPolicies field to given value.

### HasProtectionPolicies

`func (o *ProtectedVmInfo) HasProtectionPolicies() bool`

HasProtectionPolicies returns a boolean if a field has been set.

### SetProtectionPoliciesNil

`func (o *ProtectedVmInfo) SetProtectionPoliciesNil(b bool)`

 SetProtectionPoliciesNil sets the value for ProtectionPolicies to be an explicit nil

### UnsetProtectionPolicies
`func (o *ProtectedVmInfo) UnsetProtectionPolicies()`

UnsetProtectionPolicies ensures that no value is present for ProtectionPolicies, not even an explicit nil
### GetProtectionSource

`func (o *ProtectedVmInfo) GetProtectionSource() ProtectionSource`

GetProtectionSource returns the ProtectionSource field if non-nil, zero value otherwise.

### GetProtectionSourceOk

`func (o *ProtectedVmInfo) GetProtectionSourceOk() (*ProtectionSource, bool)`

GetProtectionSourceOk returns a tuple with the ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSource

`func (o *ProtectedVmInfo) SetProtectionSource(v ProtectionSource)`

SetProtectionSource sets ProtectionSource field to given value.

### HasProtectionSource

`func (o *ProtectedVmInfo) HasProtectionSource() bool`

HasProtectionSource returns a boolean if a field has been set.

### GetStats

`func (o *ProtectedVmInfo) GetStats() ProtectionSummary`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ProtectedVmInfo) GetStatsOk() (*ProtectionSummary, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ProtectedVmInfo) SetStats(v ProtectionSummary)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ProtectedVmInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *ProtectedVmInfo) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *ProtectedVmInfo) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


