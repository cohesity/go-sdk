# ProtectionPolicySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProtectionRunSummary** | Pointer to [**LastProtectionRunSummary**](LastProtectionRunSummary.md) |  | [optional] 
**PaginationCookie** | Pointer to **NullableString** | If there are more results to display, use this value to get the next set of results, by using this value in paginationCookie param for the next request to GetProtectionPolicySummary. | [optional] 
**ProtectedSourcesSummary** | Pointer to [**[]ProtectedSourceSummary**](ProtectedSourceSummary.md) | Specifies the list of Protection Sources which are protected under the given policy. This is only populated if the policy is of type kRPO. | [optional] 
**ProtectionJobsSummary** | Pointer to [**[]ProtectionJobSummaryForPolicies**](ProtectionJobSummaryForPolicies.md) | Specifies the list of Protection Jobs associated with the given Protection Policy. This is only populated if the type of the Protection Policy is kRegular. | [optional] 
**ProtectionPolicy** | Pointer to [**ProtectionPolicy**](ProtectionPolicy.md) |  | [optional] 
**ProtectionRunsSummary** | Pointer to [**ProtectionRunsSummary**](ProtectionRunsSummary.md) |  | [optional] 

## Methods

### NewProtectionPolicySummary

`func NewProtectionPolicySummary() *ProtectionPolicySummary`

NewProtectionPolicySummary instantiates a new ProtectionPolicySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionPolicySummaryWithDefaults

`func NewProtectionPolicySummaryWithDefaults() *ProtectionPolicySummary`

NewProtectionPolicySummaryWithDefaults instantiates a new ProtectionPolicySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProtectionRunSummary

`func (o *ProtectionPolicySummary) GetLastProtectionRunSummary() LastProtectionRunSummary`

GetLastProtectionRunSummary returns the LastProtectionRunSummary field if non-nil, zero value otherwise.

### GetLastProtectionRunSummaryOk

`func (o *ProtectionPolicySummary) GetLastProtectionRunSummaryOk() (*LastProtectionRunSummary, bool)`

GetLastProtectionRunSummaryOk returns a tuple with the LastProtectionRunSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProtectionRunSummary

`func (o *ProtectionPolicySummary) SetLastProtectionRunSummary(v LastProtectionRunSummary)`

SetLastProtectionRunSummary sets LastProtectionRunSummary field to given value.

### HasLastProtectionRunSummary

`func (o *ProtectionPolicySummary) HasLastProtectionRunSummary() bool`

HasLastProtectionRunSummary returns a boolean if a field has been set.

### GetPaginationCookie

`func (o *ProtectionPolicySummary) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *ProtectionPolicySummary) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *ProtectionPolicySummary) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *ProtectionPolicySummary) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *ProtectionPolicySummary) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *ProtectionPolicySummary) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetProtectedSourcesSummary

`func (o *ProtectionPolicySummary) GetProtectedSourcesSummary() []ProtectedSourceSummary`

GetProtectedSourcesSummary returns the ProtectedSourcesSummary field if non-nil, zero value otherwise.

### GetProtectedSourcesSummaryOk

`func (o *ProtectionPolicySummary) GetProtectedSourcesSummaryOk() (*[]ProtectedSourceSummary, bool)`

GetProtectedSourcesSummaryOk returns a tuple with the ProtectedSourcesSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSourcesSummary

`func (o *ProtectionPolicySummary) SetProtectedSourcesSummary(v []ProtectedSourceSummary)`

SetProtectedSourcesSummary sets ProtectedSourcesSummary field to given value.

### HasProtectedSourcesSummary

`func (o *ProtectionPolicySummary) HasProtectedSourcesSummary() bool`

HasProtectedSourcesSummary returns a boolean if a field has been set.

### SetProtectedSourcesSummaryNil

`func (o *ProtectionPolicySummary) SetProtectedSourcesSummaryNil(b bool)`

 SetProtectedSourcesSummaryNil sets the value for ProtectedSourcesSummary to be an explicit nil

### UnsetProtectedSourcesSummary
`func (o *ProtectionPolicySummary) UnsetProtectedSourcesSummary()`

UnsetProtectedSourcesSummary ensures that no value is present for ProtectedSourcesSummary, not even an explicit nil
### GetProtectionJobsSummary

`func (o *ProtectionPolicySummary) GetProtectionJobsSummary() []ProtectionJobSummaryForPolicies`

GetProtectionJobsSummary returns the ProtectionJobsSummary field if non-nil, zero value otherwise.

### GetProtectionJobsSummaryOk

`func (o *ProtectionPolicySummary) GetProtectionJobsSummaryOk() (*[]ProtectionJobSummaryForPolicies, bool)`

GetProtectionJobsSummaryOk returns a tuple with the ProtectionJobsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobsSummary

`func (o *ProtectionPolicySummary) SetProtectionJobsSummary(v []ProtectionJobSummaryForPolicies)`

SetProtectionJobsSummary sets ProtectionJobsSummary field to given value.

### HasProtectionJobsSummary

`func (o *ProtectionPolicySummary) HasProtectionJobsSummary() bool`

HasProtectionJobsSummary returns a boolean if a field has been set.

### SetProtectionJobsSummaryNil

`func (o *ProtectionPolicySummary) SetProtectionJobsSummaryNil(b bool)`

 SetProtectionJobsSummaryNil sets the value for ProtectionJobsSummary to be an explicit nil

### UnsetProtectionJobsSummary
`func (o *ProtectionPolicySummary) UnsetProtectionJobsSummary()`

UnsetProtectionJobsSummary ensures that no value is present for ProtectionJobsSummary, not even an explicit nil
### GetProtectionPolicy

`func (o *ProtectionPolicySummary) GetProtectionPolicy() ProtectionPolicy`

GetProtectionPolicy returns the ProtectionPolicy field if non-nil, zero value otherwise.

### GetProtectionPolicyOk

`func (o *ProtectionPolicySummary) GetProtectionPolicyOk() (*ProtectionPolicy, bool)`

GetProtectionPolicyOk returns a tuple with the ProtectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionPolicy

`func (o *ProtectionPolicySummary) SetProtectionPolicy(v ProtectionPolicy)`

SetProtectionPolicy sets ProtectionPolicy field to given value.

### HasProtectionPolicy

`func (o *ProtectionPolicySummary) HasProtectionPolicy() bool`

HasProtectionPolicy returns a boolean if a field has been set.

### GetProtectionRunsSummary

`func (o *ProtectionPolicySummary) GetProtectionRunsSummary() ProtectionRunsSummary`

GetProtectionRunsSummary returns the ProtectionRunsSummary field if non-nil, zero value otherwise.

### GetProtectionRunsSummaryOk

`func (o *ProtectionPolicySummary) GetProtectionRunsSummaryOk() (*ProtectionRunsSummary, bool)`

GetProtectionRunsSummaryOk returns a tuple with the ProtectionRunsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunsSummary

`func (o *ProtectionPolicySummary) SetProtectionRunsSummary(v ProtectionRunsSummary)`

SetProtectionRunsSummary sets ProtectionRunsSummary field to given value.

### HasProtectionRunsSummary

`func (o *ProtectionPolicySummary) HasProtectionRunsSummary() bool`

HasProtectionRunsSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


