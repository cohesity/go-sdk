# TenantDeletionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **NullableInt32** | Specifies the category of objects whose deletion state is being captured. | [optional] 
**FinishedAtTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the process finished. | [optional] 
**ProcessedAtNode** | Pointer to **NullableString** | Specifies the node ip where the process ran. Typically this would be Iris Master. | [optional] 
**RetryCount** | Pointer to **NullableInt64** | Specifies the number of times this task has been retried. | [optional] 
**StartedAtTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the process started. | [optional] 
**State** | Pointer to **NullableInt32** | Specifies the deletion completion state of the object category. | [optional] 

## Methods

### NewTenantDeletionInfo

`func NewTenantDeletionInfo() *TenantDeletionInfo`

NewTenantDeletionInfo instantiates a new TenantDeletionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDeletionInfoWithDefaults

`func NewTenantDeletionInfoWithDefaults() *TenantDeletionInfo`

NewTenantDeletionInfoWithDefaults instantiates a new TenantDeletionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TenantDeletionInfo) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TenantDeletionInfo) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TenantDeletionInfo) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TenantDeletionInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *TenantDeletionInfo) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TenantDeletionInfo) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetFinishedAtTimeMsecs

`func (o *TenantDeletionInfo) GetFinishedAtTimeMsecs() int64`

GetFinishedAtTimeMsecs returns the FinishedAtTimeMsecs field if non-nil, zero value otherwise.

### GetFinishedAtTimeMsecsOk

`func (o *TenantDeletionInfo) GetFinishedAtTimeMsecsOk() (*int64, bool)`

GetFinishedAtTimeMsecsOk returns a tuple with the FinishedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAtTimeMsecs

`func (o *TenantDeletionInfo) SetFinishedAtTimeMsecs(v int64)`

SetFinishedAtTimeMsecs sets FinishedAtTimeMsecs field to given value.

### HasFinishedAtTimeMsecs

`func (o *TenantDeletionInfo) HasFinishedAtTimeMsecs() bool`

HasFinishedAtTimeMsecs returns a boolean if a field has been set.

### SetFinishedAtTimeMsecsNil

`func (o *TenantDeletionInfo) SetFinishedAtTimeMsecsNil(b bool)`

 SetFinishedAtTimeMsecsNil sets the value for FinishedAtTimeMsecs to be an explicit nil

### UnsetFinishedAtTimeMsecs
`func (o *TenantDeletionInfo) UnsetFinishedAtTimeMsecs()`

UnsetFinishedAtTimeMsecs ensures that no value is present for FinishedAtTimeMsecs, not even an explicit nil
### GetProcessedAtNode

`func (o *TenantDeletionInfo) GetProcessedAtNode() string`

GetProcessedAtNode returns the ProcessedAtNode field if non-nil, zero value otherwise.

### GetProcessedAtNodeOk

`func (o *TenantDeletionInfo) GetProcessedAtNodeOk() (*string, bool)`

GetProcessedAtNodeOk returns a tuple with the ProcessedAtNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAtNode

`func (o *TenantDeletionInfo) SetProcessedAtNode(v string)`

SetProcessedAtNode sets ProcessedAtNode field to given value.

### HasProcessedAtNode

`func (o *TenantDeletionInfo) HasProcessedAtNode() bool`

HasProcessedAtNode returns a boolean if a field has been set.

### SetProcessedAtNodeNil

`func (o *TenantDeletionInfo) SetProcessedAtNodeNil(b bool)`

 SetProcessedAtNodeNil sets the value for ProcessedAtNode to be an explicit nil

### UnsetProcessedAtNode
`func (o *TenantDeletionInfo) UnsetProcessedAtNode()`

UnsetProcessedAtNode ensures that no value is present for ProcessedAtNode, not even an explicit nil
### GetRetryCount

`func (o *TenantDeletionInfo) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *TenantDeletionInfo) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *TenantDeletionInfo) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *TenantDeletionInfo) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### SetRetryCountNil

`func (o *TenantDeletionInfo) SetRetryCountNil(b bool)`

 SetRetryCountNil sets the value for RetryCount to be an explicit nil

### UnsetRetryCount
`func (o *TenantDeletionInfo) UnsetRetryCount()`

UnsetRetryCount ensures that no value is present for RetryCount, not even an explicit nil
### GetStartedAtTimeMsecs

`func (o *TenantDeletionInfo) GetStartedAtTimeMsecs() int64`

GetStartedAtTimeMsecs returns the StartedAtTimeMsecs field if non-nil, zero value otherwise.

### GetStartedAtTimeMsecsOk

`func (o *TenantDeletionInfo) GetStartedAtTimeMsecsOk() (*int64, bool)`

GetStartedAtTimeMsecsOk returns a tuple with the StartedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtTimeMsecs

`func (o *TenantDeletionInfo) SetStartedAtTimeMsecs(v int64)`

SetStartedAtTimeMsecs sets StartedAtTimeMsecs field to given value.

### HasStartedAtTimeMsecs

`func (o *TenantDeletionInfo) HasStartedAtTimeMsecs() bool`

HasStartedAtTimeMsecs returns a boolean if a field has been set.

### SetStartedAtTimeMsecsNil

`func (o *TenantDeletionInfo) SetStartedAtTimeMsecsNil(b bool)`

 SetStartedAtTimeMsecsNil sets the value for StartedAtTimeMsecs to be an explicit nil

### UnsetStartedAtTimeMsecs
`func (o *TenantDeletionInfo) UnsetStartedAtTimeMsecs()`

UnsetStartedAtTimeMsecs ensures that no value is present for StartedAtTimeMsecs, not even an explicit nil
### GetState

`func (o *TenantDeletionInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TenantDeletionInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TenantDeletionInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *TenantDeletionInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *TenantDeletionInfo) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *TenantDeletionInfo) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


