# RetrieveArchiveTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskUid** | Pointer to **NullableString** | Specifies the globally unique id for this retrieval of an archive task. | [optional] 
**UptierExpiryTimes** | Pointer to **[]int64** | Specifies how much time the retrieved entity is present in the hot-tiers. | [optional] 

## Methods

### NewRetrieveArchiveTask

`func NewRetrieveArchiveTask() *RetrieveArchiveTask`

NewRetrieveArchiveTask instantiates a new RetrieveArchiveTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveArchiveTaskWithDefaults

`func NewRetrieveArchiveTaskWithDefaults() *RetrieveArchiveTask`

NewRetrieveArchiveTaskWithDefaults instantiates a new RetrieveArchiveTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskUid

`func (o *RetrieveArchiveTask) GetTaskUid() string`

GetTaskUid returns the TaskUid field if non-nil, zero value otherwise.

### GetTaskUidOk

`func (o *RetrieveArchiveTask) GetTaskUidOk() (*string, bool)`

GetTaskUidOk returns a tuple with the TaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUid

`func (o *RetrieveArchiveTask) SetTaskUid(v string)`

SetTaskUid sets TaskUid field to given value.

### HasTaskUid

`func (o *RetrieveArchiveTask) HasTaskUid() bool`

HasTaskUid returns a boolean if a field has been set.

### SetTaskUidNil

`func (o *RetrieveArchiveTask) SetTaskUidNil(b bool)`

 SetTaskUidNil sets the value for TaskUid to be an explicit nil

### UnsetTaskUid
`func (o *RetrieveArchiveTask) UnsetTaskUid()`

UnsetTaskUid ensures that no value is present for TaskUid, not even an explicit nil
### GetUptierExpiryTimes

`func (o *RetrieveArchiveTask) GetUptierExpiryTimes() []int64`

GetUptierExpiryTimes returns the UptierExpiryTimes field if non-nil, zero value otherwise.

### GetUptierExpiryTimesOk

`func (o *RetrieveArchiveTask) GetUptierExpiryTimesOk() (*[]int64, bool)`

GetUptierExpiryTimesOk returns a tuple with the UptierExpiryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptierExpiryTimes

`func (o *RetrieveArchiveTask) SetUptierExpiryTimes(v []int64)`

SetUptierExpiryTimes sets UptierExpiryTimes field to given value.

### HasUptierExpiryTimes

`func (o *RetrieveArchiveTask) HasUptierExpiryTimes() bool`

HasUptierExpiryTimes returns a boolean if a field has been set.

### SetUptierExpiryTimesNil

`func (o *RetrieveArchiveTask) SetUptierExpiryTimesNil(b bool)`

 SetUptierExpiryTimesNil sets the value for UptierExpiryTimes to be an explicit nil

### UnsetUptierExpiryTimes
`func (o *RetrieveArchiveTask) UnsetUptierExpiryTimes()`

UnsetUptierExpiryTimes ensures that no value is present for UptierExpiryTimes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


