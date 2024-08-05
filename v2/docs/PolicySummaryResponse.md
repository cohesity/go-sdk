# PolicySummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProtectionRunSummary** | Pointer to [**LastProtectionRunSummary**](LastProtectionRunSummary.md) |  | [optional] 
**PaginationCookie** | Pointer to **NullableString** | If there are more results to display, use this value to get the next set of results, by using this value in paginationCookie param for the next request to GetProtectionPolicySummary. | [optional] 
**Policy** | Pointer to [**ProtectionPolicy**](ProtectionPolicy.md) |  | [optional] 
**ProtectionGroupsSummary** | Pointer to [**[]ProtectionGroupRun**](ProtectionGroupRun.md) | Specifies the list of Protection Groups associated with the given Protection Policy. This is only populated if the type of the Protection Policy is kRegular. | [optional] 
**ProtectionRunsSummary** | Pointer to [**ProtectionRunsInPolicySummary**](ProtectionRunsInPolicySummary.md) |  | [optional] 
**ProtectionSourcesSummary** | Pointer to [**[]ProtectionSourceSummary**](ProtectionSourceSummary.md) | Specifies the list of Protection Sources which are protected under the given policy. This is only populated if the policy is of type kRPO. | [optional] 

## Methods

### NewPolicySummaryResponse

`func NewPolicySummaryResponse() *PolicySummaryResponse`

NewPolicySummaryResponse instantiates a new PolicySummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySummaryResponseWithDefaults

`func NewPolicySummaryResponseWithDefaults() *PolicySummaryResponse`

NewPolicySummaryResponseWithDefaults instantiates a new PolicySummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProtectionRunSummary

`func (o *PolicySummaryResponse) GetLastProtectionRunSummary() LastProtectionRunSummary`

GetLastProtectionRunSummary returns the LastProtectionRunSummary field if non-nil, zero value otherwise.

### GetLastProtectionRunSummaryOk

`func (o *PolicySummaryResponse) GetLastProtectionRunSummaryOk() (*LastProtectionRunSummary, bool)`

GetLastProtectionRunSummaryOk returns a tuple with the LastProtectionRunSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProtectionRunSummary

`func (o *PolicySummaryResponse) SetLastProtectionRunSummary(v LastProtectionRunSummary)`

SetLastProtectionRunSummary sets LastProtectionRunSummary field to given value.

### HasLastProtectionRunSummary

`func (o *PolicySummaryResponse) HasLastProtectionRunSummary() bool`

HasLastProtectionRunSummary returns a boolean if a field has been set.

### GetPaginationCookie

`func (o *PolicySummaryResponse) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *PolicySummaryResponse) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *PolicySummaryResponse) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *PolicySummaryResponse) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *PolicySummaryResponse) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *PolicySummaryResponse) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetPolicy

`func (o *PolicySummaryResponse) GetPolicy() ProtectionPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicySummaryResponse) GetPolicyOk() (*ProtectionPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicySummaryResponse) SetPolicy(v ProtectionPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicySummaryResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProtectionGroupsSummary

`func (o *PolicySummaryResponse) GetProtectionGroupsSummary() []ProtectionGroupRun`

GetProtectionGroupsSummary returns the ProtectionGroupsSummary field if non-nil, zero value otherwise.

### GetProtectionGroupsSummaryOk

`func (o *PolicySummaryResponse) GetProtectionGroupsSummaryOk() (*[]ProtectionGroupRun, bool)`

GetProtectionGroupsSummaryOk returns a tuple with the ProtectionGroupsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupsSummary

`func (o *PolicySummaryResponse) SetProtectionGroupsSummary(v []ProtectionGroupRun)`

SetProtectionGroupsSummary sets ProtectionGroupsSummary field to given value.

### HasProtectionGroupsSummary

`func (o *PolicySummaryResponse) HasProtectionGroupsSummary() bool`

HasProtectionGroupsSummary returns a boolean if a field has been set.

### SetProtectionGroupsSummaryNil

`func (o *PolicySummaryResponse) SetProtectionGroupsSummaryNil(b bool)`

 SetProtectionGroupsSummaryNil sets the value for ProtectionGroupsSummary to be an explicit nil

### UnsetProtectionGroupsSummary
`func (o *PolicySummaryResponse) UnsetProtectionGroupsSummary()`

UnsetProtectionGroupsSummary ensures that no value is present for ProtectionGroupsSummary, not even an explicit nil
### GetProtectionRunsSummary

`func (o *PolicySummaryResponse) GetProtectionRunsSummary() ProtectionRunsInPolicySummary`

GetProtectionRunsSummary returns the ProtectionRunsSummary field if non-nil, zero value otherwise.

### GetProtectionRunsSummaryOk

`func (o *PolicySummaryResponse) GetProtectionRunsSummaryOk() (*ProtectionRunsInPolicySummary, bool)`

GetProtectionRunsSummaryOk returns a tuple with the ProtectionRunsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunsSummary

`func (o *PolicySummaryResponse) SetProtectionRunsSummary(v ProtectionRunsInPolicySummary)`

SetProtectionRunsSummary sets ProtectionRunsSummary field to given value.

### HasProtectionRunsSummary

`func (o *PolicySummaryResponse) HasProtectionRunsSummary() bool`

HasProtectionRunsSummary returns a boolean if a field has been set.

### GetProtectionSourcesSummary

`func (o *PolicySummaryResponse) GetProtectionSourcesSummary() []ProtectionSourceSummary`

GetProtectionSourcesSummary returns the ProtectionSourcesSummary field if non-nil, zero value otherwise.

### GetProtectionSourcesSummaryOk

`func (o *PolicySummaryResponse) GetProtectionSourcesSummaryOk() (*[]ProtectionSourceSummary, bool)`

GetProtectionSourcesSummaryOk returns a tuple with the ProtectionSourcesSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourcesSummary

`func (o *PolicySummaryResponse) SetProtectionSourcesSummary(v []ProtectionSourceSummary)`

SetProtectionSourcesSummary sets ProtectionSourcesSummary field to given value.

### HasProtectionSourcesSummary

`func (o *PolicySummaryResponse) HasProtectionSourcesSummary() bool`

HasProtectionSourcesSummary returns a boolean if a field has been set.

### SetProtectionSourcesSummaryNil

`func (o *PolicySummaryResponse) SetProtectionSourcesSummaryNil(b bool)`

 SetProtectionSourcesSummaryNil sets the value for ProtectionSourcesSummary to be an explicit nil

### UnsetProtectionSourcesSummary
`func (o *PolicySummaryResponse) UnsetProtectionSourcesSummary()`

UnsetProtectionSourcesSummary ensures that no value is present for ProtectionSourcesSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


