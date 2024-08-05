# SlaveData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedSessionId** | Pointer to **NullableInt64** | Expected gandalf session_id of the slave. | [optional] 
**IndexingCookie** | Pointer to **NullableString** | Information from where the slave can resume indexing work on the object. | [optional] 
**Stats** | Pointer to [**NullableGaiaIndexingStats**](GaiaIndexingStats.md) |  | [optional] 
**TaskHandle** | Pointer to **NullableInt64** | Task handle assigned by the master. | [optional] 

## Methods

### NewSlaveData

`func NewSlaveData() *SlaveData`

NewSlaveData instantiates a new SlaveData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaveDataWithDefaults

`func NewSlaveDataWithDefaults() *SlaveData`

NewSlaveDataWithDefaults instantiates a new SlaveData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedSessionId

`func (o *SlaveData) GetExpectedSessionId() int64`

GetExpectedSessionId returns the ExpectedSessionId field if non-nil, zero value otherwise.

### GetExpectedSessionIdOk

`func (o *SlaveData) GetExpectedSessionIdOk() (*int64, bool)`

GetExpectedSessionIdOk returns a tuple with the ExpectedSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSessionId

`func (o *SlaveData) SetExpectedSessionId(v int64)`

SetExpectedSessionId sets ExpectedSessionId field to given value.

### HasExpectedSessionId

`func (o *SlaveData) HasExpectedSessionId() bool`

HasExpectedSessionId returns a boolean if a field has been set.

### SetExpectedSessionIdNil

`func (o *SlaveData) SetExpectedSessionIdNil(b bool)`

 SetExpectedSessionIdNil sets the value for ExpectedSessionId to be an explicit nil

### UnsetExpectedSessionId
`func (o *SlaveData) UnsetExpectedSessionId()`

UnsetExpectedSessionId ensures that no value is present for ExpectedSessionId, not even an explicit nil
### GetIndexingCookie

`func (o *SlaveData) GetIndexingCookie() string`

GetIndexingCookie returns the IndexingCookie field if non-nil, zero value otherwise.

### GetIndexingCookieOk

`func (o *SlaveData) GetIndexingCookieOk() (*string, bool)`

GetIndexingCookieOk returns a tuple with the IndexingCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingCookie

`func (o *SlaveData) SetIndexingCookie(v string)`

SetIndexingCookie sets IndexingCookie field to given value.

### HasIndexingCookie

`func (o *SlaveData) HasIndexingCookie() bool`

HasIndexingCookie returns a boolean if a field has been set.

### SetIndexingCookieNil

`func (o *SlaveData) SetIndexingCookieNil(b bool)`

 SetIndexingCookieNil sets the value for IndexingCookie to be an explicit nil

### UnsetIndexingCookie
`func (o *SlaveData) UnsetIndexingCookie()`

UnsetIndexingCookie ensures that no value is present for IndexingCookie, not even an explicit nil
### GetStats

`func (o *SlaveData) GetStats() GaiaIndexingStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SlaveData) GetStatsOk() (*GaiaIndexingStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SlaveData) SetStats(v GaiaIndexingStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SlaveData) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *SlaveData) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *SlaveData) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetTaskHandle

`func (o *SlaveData) GetTaskHandle() int64`

GetTaskHandle returns the TaskHandle field if non-nil, zero value otherwise.

### GetTaskHandleOk

`func (o *SlaveData) GetTaskHandleOk() (*int64, bool)`

GetTaskHandleOk returns a tuple with the TaskHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskHandle

`func (o *SlaveData) SetTaskHandle(v int64)`

SetTaskHandle sets TaskHandle field to given value.

### HasTaskHandle

`func (o *SlaveData) HasTaskHandle() bool`

HasTaskHandle returns a boolean if a field has been set.

### SetTaskHandleNil

`func (o *SlaveData) SetTaskHandleNil(b bool)`

 SetTaskHandleNil sets the value for TaskHandle to be an explicit nil

### UnsetTaskHandle
`func (o *SlaveData) UnsetTaskHandle()`

UnsetTaskHandle ensures that no value is present for TaskHandle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


