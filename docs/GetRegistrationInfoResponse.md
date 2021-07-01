# GetRegistrationInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootNodes** | Pointer to [**[]ProtectionSourceTreeInfo**](ProtectionSourceTreeInfo.md) | Specifies the registration, protection and permission information of either all or a subset of registered Protection Sources matching the filter parameters. overrideDescription: true | [optional] 
**Stats** | Pointer to [**NullableProtectionSummary**](ProtectionSummary.md) | Specifies the sum of all the stats of protection of Protection Sources and views selected by the query parameters. | [optional] 
**StatsByEnv** | Pointer to [**[]ProtectionSummaryByEnv**](ProtectionSummaryByEnv.md) | Specifies the breakdown of the stats by environment overrideDescription: true | [optional] 

## Methods

### NewGetRegistrationInfoResponse

`func NewGetRegistrationInfoResponse() *GetRegistrationInfoResponse`

NewGetRegistrationInfoResponse instantiates a new GetRegistrationInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegistrationInfoResponseWithDefaults

`func NewGetRegistrationInfoResponseWithDefaults() *GetRegistrationInfoResponse`

NewGetRegistrationInfoResponseWithDefaults instantiates a new GetRegistrationInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootNodes

`func (o *GetRegistrationInfoResponse) GetRootNodes() []ProtectionSourceTreeInfo`

GetRootNodes returns the RootNodes field if non-nil, zero value otherwise.

### GetRootNodesOk

`func (o *GetRegistrationInfoResponse) GetRootNodesOk() (*[]ProtectionSourceTreeInfo, bool)`

GetRootNodesOk returns a tuple with the RootNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNodes

`func (o *GetRegistrationInfoResponse) SetRootNodes(v []ProtectionSourceTreeInfo)`

SetRootNodes sets RootNodes field to given value.

### HasRootNodes

`func (o *GetRegistrationInfoResponse) HasRootNodes() bool`

HasRootNodes returns a boolean if a field has been set.

### SetRootNodesNil

`func (o *GetRegistrationInfoResponse) SetRootNodesNil(b bool)`

 SetRootNodesNil sets the value for RootNodes to be an explicit nil

### UnsetRootNodes
`func (o *GetRegistrationInfoResponse) UnsetRootNodes()`

UnsetRootNodes ensures that no value is present for RootNodes, not even an explicit nil
### GetStats

`func (o *GetRegistrationInfoResponse) GetStats() ProtectionSummary`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetRegistrationInfoResponse) GetStatsOk() (*ProtectionSummary, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetRegistrationInfoResponse) SetStats(v ProtectionSummary)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetRegistrationInfoResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *GetRegistrationInfoResponse) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *GetRegistrationInfoResponse) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetStatsByEnv

`func (o *GetRegistrationInfoResponse) GetStatsByEnv() []ProtectionSummaryByEnv`

GetStatsByEnv returns the StatsByEnv field if non-nil, zero value otherwise.

### GetStatsByEnvOk

`func (o *GetRegistrationInfoResponse) GetStatsByEnvOk() (*[]ProtectionSummaryByEnv, bool)`

GetStatsByEnvOk returns a tuple with the StatsByEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByEnv

`func (o *GetRegistrationInfoResponse) SetStatsByEnv(v []ProtectionSummaryByEnv)`

SetStatsByEnv sets StatsByEnv field to given value.

### HasStatsByEnv

`func (o *GetRegistrationInfoResponse) HasStatsByEnv() bool`

HasStatsByEnv returns a boolean if a field has been set.

### SetStatsByEnvNil

`func (o *GetRegistrationInfoResponse) SetStatsByEnvNil(b bool)`

 SetStatsByEnvNil sets the value for StatsByEnv to be an explicit nil

### UnsetStatsByEnv
`func (o *GetRegistrationInfoResponse) UnsetStatsByEnv()`

UnsetStatsByEnv ensures that no value is present for StatsByEnv, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


