# McmGetPolicyLastRunStatsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatsByPolicy** | Pointer to [**[]McmPolicyLastProtectionRunStats**](McmPolicyLastProtectionRunStats.md) | Specifies the last Protection Run stats by policy. | [optional] 

## Methods

### NewMcmGetPolicyLastRunStatsResponseBody

`func NewMcmGetPolicyLastRunStatsResponseBody() *McmGetPolicyLastRunStatsResponseBody`

NewMcmGetPolicyLastRunStatsResponseBody instantiates a new McmGetPolicyLastRunStatsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmGetPolicyLastRunStatsResponseBodyWithDefaults

`func NewMcmGetPolicyLastRunStatsResponseBodyWithDefaults() *McmGetPolicyLastRunStatsResponseBody`

NewMcmGetPolicyLastRunStatsResponseBodyWithDefaults instantiates a new McmGetPolicyLastRunStatsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatsByPolicy

`func (o *McmGetPolicyLastRunStatsResponseBody) GetStatsByPolicy() []McmPolicyLastProtectionRunStats`

GetStatsByPolicy returns the StatsByPolicy field if non-nil, zero value otherwise.

### GetStatsByPolicyOk

`func (o *McmGetPolicyLastRunStatsResponseBody) GetStatsByPolicyOk() (*[]McmPolicyLastProtectionRunStats, bool)`

GetStatsByPolicyOk returns a tuple with the StatsByPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByPolicy

`func (o *McmGetPolicyLastRunStatsResponseBody) SetStatsByPolicy(v []McmPolicyLastProtectionRunStats)`

SetStatsByPolicy sets StatsByPolicy field to given value.

### HasStatsByPolicy

`func (o *McmGetPolicyLastRunStatsResponseBody) HasStatsByPolicy() bool`

HasStatsByPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


